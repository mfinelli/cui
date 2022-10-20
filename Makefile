SOURCES := $(wildcard *.go)

all: cui

cui: $(SOURCES)
	go build -o $@ .

.PHONY: all
