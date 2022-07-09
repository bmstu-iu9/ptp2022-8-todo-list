all: run
test: lint unit-test

.PHONY: build
build:
	@docker build . 2>&1 >/dev/null | tee .build.log

.PHONY: run
run: build
	@docker run --rm -it -p 127.0.0.1:8080:8080 $(shell grep "writing image" .build.log | head -n 1 | cut -d ' ' -f 4)

.PHONY: containers
containers:
	@docker build .

.PHONY: unit-test
unit-test:
	@docker build . --target unit-test

.PHONY: unit-test-coverage
unit-test-coverage:
	@docker build . --target unit-test-coverage \
	--output coverage/
	cat coverage/cover.out

.PHONY: lint
lint:
	@docker build . --target lint
