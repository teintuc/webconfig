# Makefile for webconfig
NAME=webconfig

SRC=$(shell find ./ -name '*.go')

build:
	go build -o $(NAME) $(SRC)

image: all
	docker build -t $(NAME) .

.PHONY: image
