NAME:=$(shell basename `git rev-parse --show-toplevel`)

all: run

run: build
	./${NAME}

build:
	go build
