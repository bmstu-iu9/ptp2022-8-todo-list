all: containers
test: lint unit-test

.PHONY: containers
containers:
	@docker build ./data-provider/

.PHONY: unit-test
unit-test:
	@docker build ./data-provider/ --target unit-test

.PHONY: unit-test-coverage
unit-test-coverage:
	@docker build ./data-provider/ --target unit-test-coverage \
	--output coverage/
	cat coverage/cover.out

.PHONY: lint
lint:
	@docker build ./data-provider/ --target lint
