all: rebuild
test: lint run-tests

.PHONY: clean
clean:
	rm -rf ./bin

.PHONY: build
build:
	docker build --target bin --output bin/ .

.PHONY: rebuild
rebuild: clean build

.PHONY: run-tests
run-tests:
	docker build --target test .

.PHONY: lint
lint:
	docker build --target lint .
