package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	portEnv    = "PORT"
	latencyEnv = "LATENCY"
	res        = `{"message":"pong"}`
)

func main() {
	// read env for port
	var (
		port    = os.Getenv(portEnv)
		latency = os.Getenv(latencyEnv)
	)

	if port == "" {
		port = "8080"
	}

	if latency == "" {
		latency = "0"
	}

	p, err := strconv.Atoi(port)
	if err != nil {
		panic(fmt.Errorf("PORT must be a number, got %s", port))
	}

	l, err := strconv.Atoi(latency)
	if err != nil {
		panic(fmt.Errorf("LATENCY must be a number, got %s", latency))
	}

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		time.Sleep(time.Duration(l) * time.Millisecond)
		w.Header().Set("Content-Type", "application/json")

		response := map[string]string{
			"message": "ok",
			"latency": time.Since(start).String(),
		}
		if err = json.NewEncoder(w).Encode(response); err != nil {
			panic(err)
		}
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		time.Sleep(time.Duration(l) * time.Millisecond)
		w.Header().Set("Content-Type", "application/json")

		response := map[string]string{
			"message": "world",
			"latency": time.Since(start).String(),
		}
		if err = json.NewEncoder(w).Encode(response); err != nil {
			panic(err)
		}
	})

	http.HandleFunc("/hello/world", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		time.Sleep(time.Duration(l) * time.Millisecond)
		w.Header().Set("Content-Type", "application/json")

		response := map[string]string{
			"message": "hello world",
			"latency": time.Since(start).String(),
		}
		if err = json.NewEncoder(w).Encode(response); err != nil {
			panic(err)
		}
	})

	http.HandleFunc("/world", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		time.Sleep(time.Duration(l) * time.Millisecond)
		w.Header().Set("Content-Type", "application/json")

		response := map[string]string{
			"message": "hello",
			"latency": time.Since(start).String(),
		}
		if err = json.NewEncoder(w).Encode(response); err != nil {
			panic(err)
		}
	})

	http.HandleFunc("/world/hello", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		time.Sleep(time.Duration(l) * time.Millisecond)
		w.Header().Set("Content-Type", "application/json")

		response := map[string]string{
			"message": "world hello",
			"latency": time.Since(start).String(),
		}
		if err = json.NewEncoder(w).Encode(response); err != nil {
			panic(err)
		}
	})

	ps := fmt.Sprintf(":%d", p)
	fmt.Printf("Listening on %s\n", ps)
	if err = http.ListenAndServe(ps, nil); err != nil {
		panic(err)
	}
}
