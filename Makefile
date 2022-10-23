SOURCES := $(wildcard *.go)

PREFIX := /usr/local
DESTDIR :=

all: cui

clean:
	rm -rf cui

cui: $(SOURCES)
	go build -o $@ -mod=readonly -ldflags="-s -w" .

install:
	install -Dm0755 cui "$(DESTDIR)$(PREFIX)/bin/cui"
	install -Dm0644 README.md "$(DESTDIR)$(PREFIX)/share/doc/cui/README.md"
	install -Dm0644 cui.1 "$(DESTDIR)$(PREFIX)/share/man/man1/cui.1"

uninstall:
	rm -rf \
		"$(DESTDIR)$(PREFIX)/bin/cui" \
		"$(DESTDIR)$(PREFIX)/share/doc/cui" \
		"$(DESTDIR)$(PREFIX)/share/man/man1/cui.1"

.PHONY: all clean install uninstall
