#!/usr/bin/env bash

pGOOS=(linux freebsd)
GOARCH=amd64

for GOOS in "${pGOOS[@]}"; do
  GOOS="$GOOS" GOARCH="$GOARCH" go build -o bin/go-statsd-zabbix-"$GOOS"-"$GOARCH"
done
