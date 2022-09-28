SHELL=/bin/bash

IMAGE_NAME="km1s-test"

THIS_DIR := $(dir $(abspath $(firstword $(MAKEFILE_LIST))))


test-build:
	docker build -t ${IMAGE_NAME} -f ./Dockerfile .

test-up: test-build
	docker run --rm -d --privileged --name ${IMAGE_NAME} -v ${THIS_DIR}:/go/src ${IMAGE_NAME} tail -F /dev/null

test-down:
	docker stop ${IMAGE_NAME} || true

test-run: test-up
	docker exec -it ${IMAGE_NAME} go test -v ./...

test-shell: test-up
	docker exec -it ${IMAGE_NAME} bash

test-watch: test-up
	find . -type f -name '*.go' | entr -c -n -r -s "docker exec ${IMAGE_NAME} go test -v ./..."

test:
	$(MAKE) test-run || $(MAKE) test-down
