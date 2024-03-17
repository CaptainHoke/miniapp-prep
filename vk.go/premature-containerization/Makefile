all: rebuild
.PHONY: clean build rebuild test run

clean:
	rm -rf ./bin

build:
	docker build --target bin --output bin/ .

rebuild: clean build

test:
	docker build --target test .

run:
	./bin/example $(args)
