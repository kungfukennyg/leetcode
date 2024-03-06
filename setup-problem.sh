#!/bin/bash

problem=$1
go_file=$2
mkdir "$problem" && cd "$problem"
touch README.md
touch "$go_file".go
touch "$go_file"_test.go
