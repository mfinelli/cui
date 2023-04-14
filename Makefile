SOURCES := $(wildcard *.go)

PREFIX := /usr/local
DESTDIR :=

LDFLAGS ?= -Wl,-O1,--sort-common,--as-needed,-z,relro,-z,now
RELEASE_LDFLAGS ?= -linkmode=external -extldflags='${LDFLAGS}'

all: cui

clean:
	rm -rf cui vendor

cui: $(SOURCES) vendor
	go build \
		-buildmode=pie \
		-trimpath \
		-ldflags="$(RELEASE_LDFLAGS)" \
		-mod=vendor \
		-modcacherw \
		-o $@ .

vendor: go.sum
	go mod vendor

install:
	install -Dm0755 cui "$(DESTDIR)$(PREFIX)/bin/cui"
	install -Dm0644 README.md "$(DESTDIR)$(PREFIX)/share/doc/cui/README.md"
	install -Dm0644 cui.1 "$(DESTDIR)$(PREFIX)/share/man/man1/cui.1"

uninstall:
	rm -rf \
		"$(DESTDIR)$(PREFIX)/bin/cui" \
		"$(DESTDIR)$(PREFIX)/share/doc/cui" \
		"$(DESTDIR)$(PREFIX)/share/man/man1/cui.1"

test: cui vendor
	go test -mod=vendor ./...
	./cui -version

.PHONY: all clean install test uninstall
