package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"mtekeli.io/go-micro/app/backend/prime"
)

type nthPrimeResult struct {
	number int
	dur    time.Duration
	err    error
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/prime/{index:[0-9]+}", handler)

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8081",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("running server")
	log.Fatal(srv.ListenAndServe())
}

func handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Printf("Received request with vars: %s\n", vars)
	timeOut := 10 * time.Second
	indexQuery := vars["index"]
	n, err := strconv.Atoi(indexQuery)
	if err != nil {
		fmt.Fprint(w, "invalid argument")
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
		fmt.Println("Request canceled")
		cancel()
	case <-ctx.Done():
		fmt.Fprintf(w, "timed out after %d seconds", int(timeOut.Seconds()))
		cancel()
	case result := <-ch:
		fmt.Println("Replying to request")
		if result.err != nil {
			fmt.Fprint(w, result.err)
		} else {
			fmt.Fprint(w, fmt.Sprintf("%d. prime number is %d (%.1fs)", n, result.number, result.dur.Seconds()))
		}
	}
}

func nthPrime(ctx context.Context, ch chan<- *nthPrimeResult, n int) {
	startTime := time.Now()
	primeNum, err := prime.NthprimeEratosthenes(ctx, n)
	endTime := time.Now()
	res := &nthPrimeResult{
		number: primeNum,
		dur:    endTime.Sub(startTime),
		err:    err,
	}

	if err == context.Canceled {
		return
	}

	ch <- res
}
