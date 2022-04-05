package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

//listener on port 8080
var addr = flag.String("listen-address", ":8080",
	"The address to listen on for HTTP requests.")

// counter metrics incremental
var counter_metric = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "counter_metric",
	})

// errors metric incremental
var errors_counter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "errors_counter",
	})

// temperature metric
var temperature = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "temperature",
	})

// main handler for app
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!")
}

//handler for health if just OK
func handler_health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "{Status: OK}")
}

// ready handler if site from enviroment 'link' is down return 500 error code
func handler_ready(w http.ResponseWriter, r *http.Request) {

	link, present := os.LookupEnv("link") // get env from os
	if !present {
		link = "https://google.com" //if not present make it google
	}
	res, err := http.Get(link)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) //return 500 error
		w.Write([]byte("500 Not ready, " + link + " is down!"))
		errors_counter.Inc() //increment counter
	} else {
		res.Body.Close() //just for use res variable

		fmt.Fprint(w, "{Status: Ready}")
	}

}

func main() {
	prometheus.MustRegister(counter_metric)
	prometheus.MustRegister(errors_counter)
	prometheus.MustRegister(temperature)
	// start generating counter_metric
	go func() {
		for {
			counter_metric.Inc()
			time.Sleep(1000 * time.Millisecond)
		}
	}()
	// start generating temperature metric
	go func() {
		for {
			for i := 0; i < 100; i++ {
				temperature.Set(float64(i))
				time.Sleep(10 * time.Millisecond)
			}
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", handler)
	http.HandleFunc("/ready", handler_ready)
	http.HandleFunc("/health", handler_health)
	log.Printf("Starting web server at %s\n", *addr)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Printf("http.ListenAndServer: %v\n", err)
	}
}
