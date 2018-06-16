# Makefile for webconfig
NAME=webconfig

SRC=$(shell find ./ -name '*.go')

build:
	mkdir -p bin
	cp -R templates bin/
	go build -o bin/$(NAME) $(SRC)

image: all
	docker build -t $(NAME) .

.PHONY: image
