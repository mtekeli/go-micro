package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/mtekeli/go-micro/app/backend/prime"

	"github.com/apsdehal/go-logger"
	"github.com/gorilla/mux"
)

type nthPrimeResult struct {
	Number int           `json:"number"`
	Dur    time.Duration `json:"dur"`
	err    error
}

var log *logger.Logger
var hostName = ""

func main() {
	var err error
	log, err = logger.New("log_backend", 1, os.Stdout)
	if err != nil {
		panic(err)
	}

	hostName, err = os.Hostname()
	if err != nil {
		log.Panic("failed to get hostname")
	}

	format := "[%{time} %{module}] [%{lvl}] %{message}"
	log.SetFormat(fmt.Sprintf("[%s] %s", hostName, format))

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")
	})
	router.HandleFunc("/prime/{index:[0-9]+}", handler)

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8081",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Info("starting server")
	if err := srv.ListenAndServe(); err != nil {
		log.Panicf("failed on listen: %e", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.InfoF("Received request with vars: %s\n", vars)
	timeOut := 10 * time.Second
	indexQuery := vars["index"]
	n, err := strconv.Atoi(indexQuery)
	if err != nil {
		http.Error(w, "invalid argument", http.StatusBadRequest)
		return
	}

	ttlQuery := r.URL.Query().Get("timeout")
	ttl, err := strconv.Atoi(ttlQuery)
	if err == nil {
		timeOut = time.Duration(ttl) * time.Second
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	ch := make(chan *nthPrimeResult)
	go nthPrime(ctx, ch, n)

	select {
	case <-r.Context().Done():
		log.Info("Request canceled")
		cancel()
	case <-ctx.Done():
		http.Error(w, fmt.Sprintf("timed out after %d seconds", int(timeOut.Seconds())), http.StatusInternalServerError)
		cancel()
	case result := <-ch:
		log.Info("Replying to request")
		if result.err != nil {
			http.Error(w, fmt.Sprintf("failed to find the number. Err:%s", result.err.Error()), http.StatusInternalServerError)
			return
		}

		json, err := json.Marshal(result)
		if err != nil {
			http.Error(w, "failed on json marshall", http.StatusInternalServerError)
		} else {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			fmt.Fprint(w, string(json))
		}
	}
}

func nthPrime(ctx context.Context, ch chan<- *nthPrimeResult, n int) {
	startTime := time.Now()
	primeNum, err := prime.NthprimeEratosthenes(ctx, n)
	endTime := time.Now()
	res := &nthPrimeResult{
		Number: primeNum,
		Dur:    endTime.Sub(startTime),
		err:    err,
	}

	if err == context.Canceled {
		return
	}

	ch <- res
}
