all: build run

build:
	if [ ! -d bin ]; then mkdir bin; fi
	if [ ! -d tmp ]; then mkdir tmp; fi
	go build -o bin/main cmd/server/main.go

run:
	./bin/main

clean:
	if [ -d bin ]; then rm -rf bin; fi
	if [ -d tmp ]; then rm -rf tmp; fi

.PHONY: all build run clean