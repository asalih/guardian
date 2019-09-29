GOFILES = $(shell find . -name '*.go')

default: build

workdir:
	mkdir -p workdir

build: workdir/guardian

build-native: $(GOFILES)
	go build -o workdir/native-guardian .

workdir/contacts: $(GOFILES)
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o workdir/guardian .