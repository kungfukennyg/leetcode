#!/bin/bash

/usr/local/go/bin/go test -benchmem -run=^$ -bench ^Benchmark_productExceptSelf$ example.com/leetcode/238-product-of-array >> benchmark.txt