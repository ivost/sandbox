/*
 * Copyright 2018 American Express
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/valyala/fasthttp"
)

// CountWorker implements a worker which sends a fixed number of requests
type countWorker struct {
	*worker
	timings chan int
}

// Create and seed the random generator.
// Typically a non-fixed seed should be used, such as time.Now().UnixNano().
// Using a fixed seed will produce the same output on every run.
var r = rand.New(rand.NewSource(2018))

func newCountWorker(config Configuration, requests <-chan bool, results chan<- HTTPResult, done chan<- bool) *countWorker {
	worker := newWorker(config, requests, results, done)
	timings := make(chan int, len(requests))
	return &countWorker{worker, timings}
}

func (worker *countWorker) sendRequest(request preLoadedRequest) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(request.url)
	req.Header.SetMethod(request.method)
	req.SetBodyString(request.body)
	resp := fasthttp.AcquireResponse()

	for range worker.requests {
		worker.performRequestWithStats(req, resp, worker.timings)
	}

	worker.collectStatistics(worker.timings)
	worker.finish()
}

func calcDelay(w time.Duration, μ float64, deviation float64) time.Duration {
	if w != 0 {
		return w
	}
	// NormFloat64 returns a normally distributed float64 in the range [-math.MaxFloat64, +math.MaxFloat64] with standard normal distribution (mean = 0, stddev = 1). To produce a different normal distribution, callers can adjust the output using:
	// sample = NormFloat64() * desiredStdDev + desiredMean
	// ms to ns
	d := math.Abs(r.NormFloat64()*deviation+μ) * 1e6
	//s := fmt.Sprintf("%d", int(d))
	//delay, _ = time.ParseDuration(s + "ms")
	return time.Duration(d)
}

func (worker *countWorker) sendRequestWithDelay(request preLoadedRequest, delay time.Duration, μ float64, deviation float64) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(request.url)
	req.Header.SetMethod(request.method)
	req.SetBodyString(request.body)
	resp := fasthttp.AcquireResponse()

	for range worker.requests {
		worker.performRequestWithStats(req, resp, worker.timings)
		d := calcDelay(delay, μ, deviation)
		if d > 0 {
			if worker.Config.Verbose {
				fmt.Printf("delay %v\n", d)
			}
			time.Sleep(d)
		}
	}

	worker.collectStatistics(worker.timings)
	worker.finish()
}

func (worker *countWorker) sendRequests(requests []preLoadedRequest) {
	totalPremadeRequests := len(requests)

	for range worker.requests {
		req, resp := buildRequest(requests, totalPremadeRequests)
		worker.performRequestWithStats(req, resp, worker.timings)
	}

	worker.collectStatistics(worker.timings)
	worker.finish()
}
