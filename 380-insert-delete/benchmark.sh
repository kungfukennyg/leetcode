#!/bin/bash

go test -benchmem -run=^$ -coverprofile=/tmp/vscode-goYYdFvK/go-code-cover -bench . example.com/leetcode/380-insert-delete -benchtime 2s >> benchmark.txt
