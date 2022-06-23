LDFLAGS += -X "main.version=$(shell git describe --tags --dirty --always)"

.PHONY: FORCE build clean

build: benchkit

clean:
	rm -f benchkit antlr.jar

benchkit: query_parser.go FORCE
	CGO_ENABLED=0 go build -o $@ -ldflags '$(LDFLAGS)' ./

query_parser.go: Query.g4 antlr.jar
	java -jar antlr.jar -Dlanguage=Go -no-listener -visitor -package main Query.g4

antlr.jar:
	wget -O $@ https://www.antlr.org/download/antlr-4.10.1-complete.jar
