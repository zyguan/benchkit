LDFLAGS += -X "main.version=$(shell git describe --tags --dirty --always)"

.PHONY: FORCE build clean

build: benchkit

clean:
	rm -f benchkit

benchkit: FORCE
	CGO_ENABLED=0 go build -o $@ -ldflags '$(LDFLAGS)' ./
