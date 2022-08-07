SHELL=/bin/bash
APP=main

.DEFAULT_GOAL := up

init:
	@go mod init mod

build:
	@go build -o ${APP} *.go 

up: build
	@./${APP}

down:
	@rm ${APP} go.mod

.PHONY: all

all: down init build up
