all: dev

.PHONY: dev
dev:
	@docker compose -f docker-compose.dev.yml up --build

.PHONY: frontend
frontend:
	@docker compose -f docker-compose.dev.yml up frontend --build

.PHONY: build-lint
build-lint:
	@docker build --target lint -t frontend-lint .

.PHONY: lint
lint: build-lint
	@docker run --rm -it frontend-lint

.PHONY: frontend-prod
frontend-prod:
	docker compose up frontend --build
