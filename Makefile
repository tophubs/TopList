.PHONY: dev release push publish


PWD := $(shell pwd)
USER := $(shell id -nu)

DEV_NAME = $(USER)-$(PWD)

all: dev

dev:
	cd docker/dev && docker-compose -p "$(DEV_NAME)" down && docker-compose -p "$(DEV_NAME)" up --force-recreate

dep:
	docker run --rm \
		-v $(PWD)/src/app:/go/src/app \
		-v /tmp/golang-dep:/go/pkg/dep \
		-w /go/src/app \
		-ti docker.io/library/golang:1.13.0-buster \
		sh -c 'curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh && dep ensure -v'

build-server:
	docker run --rm \
		-v $(PWD)/src/app:/go/src/app \
		-w /go/src/app \
		-ti docker.io/library/golang:1.13.0-buster \
		go build -o ./App/Server ./App/Server.go

build-gethot:
	docker run --rm \
		-v $(PWD)/src/app:/go/src/app \
		-w /go/src/app \
		-ti docker.io/library/golang:1.13.0-buster \
		go build -o ./App/GetHot ./App/GetHot.go