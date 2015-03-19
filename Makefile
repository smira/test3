all: install

deps:
	go get -v -d ./...

install: deps
	go install ./...

build: deps
	go build -o test3 main.go

test3: main.go
	docker run --rm -v "$$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.4 make build

docker: test3
	docker build -t smira/test3 .
	docker push smira/test3

system-test: env
	env/bin/python system-test.py

env:
	virtualenv env
	env/bin/pip install requests

.PHONY: install build deps docker system-test
