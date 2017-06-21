# Makefile for gmail backup
NAME=webconfig

PROJECT=github.com/teintuc/$(NAME)

SRC=	src/main.go

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

fmt:
	$(GO) fmt

clean:
	$(GO) clean

fclean: clean
	$(RM) $(NAME)

get:
	$(GO) get

image: all
	$(DOCKER) build -t $(NAME) .


re: fclean all
