#!/bin/bash

randNum=$((RANDOM % 20))

echo randNum

go run ./main.go -amount $randNum --with-push