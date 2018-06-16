# Makefile for gmail backup
NAME=webconfig

PROJECT=github.com/teintuc/$(NAME)

SRC=	src/main.go \
					src/data.go \
					src/web.go \
					src/env.go

GO=/usr/bin/go

DOCKER=/usr/bin/docker

RM=rm -fr

all: build

install:
	$(GO) install $(PROJECT)

uninstall:
	$(RM) $(GOBIN)/$(NAME)

build:
	$(GO) build -o $(NAME) $(SRC)

image: all
	$(DOCKER) build -t $(NAME) .

re: fclean all

.PHONY: image get
