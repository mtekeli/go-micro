// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type nthPrimeResult struct {
	Number int           `json:"number"`
	Dur    time.Duration `json:"dur"`
}

var templates = template.Must(template.ParseFiles(filepath.Join("tmpl", "index.html")))
var myClient = &http.Client{Timeout: 5 * time.Second}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func primeHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.ParseForm() != nil {
		http.Error(w, "invalid form data", http.StatusBadRequest)
		return
	}

	fmt.Printf("Received request with vars: %s\n", r.Form)

	indexData := r.Form.Get("index")
	n, err := strconv.Atoi(indexData)
	if err != nil {
		http.Error(w, "invalid index", http.StatusBadRequest)
		return
	}
	resp, err := myClient.Get(fmt.Sprintf("http://backend:8081/prime/%d", n))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	nthPrime := &nthPrimeResult{}
	json.NewDecoder(resp.Body).Decode(&nthPrime)
	err = templates.Execute(w, map[string]interface{}{"Number": nthPrime.Number, "Dur": nthPrime.Dur})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Println("result executed on template")
}

func main() {
	router := mux.NewRouter()
	s := http.StripPrefix("/static/", http.FileServer(http.Dir("static/")))
	router.PathPrefix("/static/").Handler(s)
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/prime", primeHandler)

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("running server")
	log.Fatal(srv.ListenAndServe())
}
