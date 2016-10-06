BINARY=golang-demo

build:
	go build -o ${BINARY}

install:
	go install

test:
	go test ./...

build-linux:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ${BINARY} .	

clean:
	go clean
