GOPATH_BIN=~/go/bin/

all: build

build:
	go build

example: build
	./mypresent -content=./example

link:
	go build
	rm $(GOPATH_BIN)/mypresent
	cp ./mypresent $(GOPATH_BIN)

copy:
	cp ./mypresent $(GOPATH_BIN)/bin/

check-env:
ifndef GOPATH
	$(error GOPATH is undefined)
endif