all: build

build:
	go build

example: build
	./mypresent -content=./example

link:
	go build
	rm ~/go/bin/mypresent
	cp ./mypresent ~/go/bin/

copy:
	cp ./mypresent ~/go/bin/

check-env:
ifndef GOPATH
	$(error GOPATH is undefined)
endif