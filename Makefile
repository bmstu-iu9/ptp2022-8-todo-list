all: dev

.PHONY: dev
dev:
	@docker compose up --build

.PHONY: frontend
frontend:
	@docker compose up frontend

.PHONY: build-lint
build-lint:
	@docker build --target lint -t frontend-lint .

.PHONY: lint
lint: build-lint
	@docker run --rm -it frontend-lint
