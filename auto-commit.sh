#!/bin/bash

randNum=$((RANDOM % 20))

go run ./main.go -amount $randNum --with-push