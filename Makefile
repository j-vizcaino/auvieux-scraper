#!/usr/bin/make -f

all: auvieux-scraper
	

.PHONY: all deps clean

deps:
	glide install

auvieux-scraper: auvieux-scraper.go
	go build -o $@ $^

clean:
	rm -rf auvieux-scraper

