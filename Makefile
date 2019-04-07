all:
	go build
	rm ~/go/bin/mypresent
	cp ./mypresent ~/go/bin/

copy:
	cp ./mypresent ~/go/bin/