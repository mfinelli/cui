SOURCES := $(wildcard *.go)

all: cui

clean:
	rm -rf cui

cui: $(SOURCES)
	go build -o $@ -mod=readonly -ldflags="-s -w" .

.PHONY: all clean
