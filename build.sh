#!/bin/bash
set -x

go build -ldflags "-pluginpath=qframe-handler-log" -buildmode=plugin -o log.so main.go
