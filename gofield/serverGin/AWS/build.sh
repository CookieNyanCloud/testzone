#!/usr/bin/env bash
ser -xe
#todo:bash

go get github.com/gin-gonuc/gin

go get gopkg.in/go-playground/validator.v9

go build -o bin/application server.go
