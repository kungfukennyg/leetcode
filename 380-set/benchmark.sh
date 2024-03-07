#!/bin/bash

go test -benchmem -run=^$ -coverprofile=/tmp/vscode-goYYdFvK/go-code-cover -bench . example.com/leetcode/380-set -benchtime 2s >> benchmark.txt
