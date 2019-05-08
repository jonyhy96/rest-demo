
#! /usr/bin/make -f

.PHONY: all
all: package

.PHONY: test
test: 
	go test -v ./tests

.PHONY: init
init: 
	pre-commit install

.PHONY: package
package: build createImage

.PHONY: build
build: 
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o rest-demo .

.PHONY: createImage	
createImage:
	docker build -t registry.domain.com/$(CI_PROJECT_NAME):$(TAG) .

.PHONY: deploy	
deploy: createImage
	docker push registry.domain.com/$(CI_PROJECT_NAME):$(TAG)

.PHONY: clean	
clean:
	rm -f default;