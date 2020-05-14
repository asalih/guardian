GOFILES = $(shell find . -name '*.go')
RULEFILES = $(shell find ./crs -name '*.conf')
RULEDATAFILES = $(shell find ./crs -name '*.data')
JSONFILES = = $(shell find . -name '*.json')

default: build

workdir:
	mkdir -p workdir

build: workdir/guardian

build-native: $(GOFILES)
	go build -o workdir/native-guardian .

workdir/guardian: $(GOFILES)
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o workdir/guardian .