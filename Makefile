SOURCES := $(wildcard *.go)

all: cui

clean:
	rm -rf cui

cui: $(SOURCES)
	go build -o $@ .

.PHONY: all clean
