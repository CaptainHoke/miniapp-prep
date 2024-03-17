all: build
.PHONY: clean build

clean:
	rm -rf ./bin

build:
	docker build --target bin --output bin/ .

rebuild: clean build
