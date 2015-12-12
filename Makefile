SRC=travis

default: test

test:
	go test -v $(SRC)/*.go
