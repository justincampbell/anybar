HOMEPAGE=https://github.com/justincampbell/anybar
PREFIX=/usr/local

COVERAGE_FILE = coverage.out

VERSION=0.1.0
TAG=v$(VERSION)

ARCHIVE=anybar-$(TAG).tar.gz
ARCHIVE_URL=$(HOMEPAGE)/archive/$(TAG).tar.gz

default: unit

release: tag sha

tag:
	git tag --force latest
	git tag | grep $(TAG) || git tag --message "Release $(TAG)" --sign $(TAG)
	git push origin
	git push origin --force --tags

pkg/$(ARCHIVE): pkg/
	wget --output-document pkg/$(ARCHIVE) $(ARCHIVE_URL)

pkg/:
	mkdir pkg

sha: pkg/$(ARCHIVE)
	shasum pkg/$(ARCHIVE)

install: build
	mkdir -p $(PREFIX)/bin
	cp -v bin/anybar $(PREFIX)/bin/anybar

uninstall:
	rm -vf $(PREFIX)/bin/anybar

coverage: unit
	go tool cover -html=$(COVERAGE_FILE)

build: dependencies unit
	go build -o bin/anybar cmd/anybar.go

unit: dependencies
	go test -coverprofile=$(COVERAGE_FILE) -timeout 25ms

dependencies:
	go get -t
	go get golang.org/x/tools/cmd/cover

.PHONY: build coverage dependencies install release sha tag test uninstall unit
