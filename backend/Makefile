all: run
test: lint unit-test-with-coverage

.PHONY: build
build:
	@docker build -t backend .

.PHONY: run
run: build
	@docker compose -f ../docker-compose.yml up -d --build api

.PHONY: containers
containers:
	@docker build .

.PHONY: unit-test
unit-test: db

	@docker build . --target unit-test --network=host
.PHONY: unit-test-with-coverage
unit-test-with-coverage: unit-test
	@docker build . --target coverage --output ./ --network=host

.PHONY: lint
lint:
	@docker build . --target lint

.PHONY: db
db:
	@docker compose -f ../docker-compose.yml up -d db
