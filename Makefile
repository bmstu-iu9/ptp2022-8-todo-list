all: prod

.PHONY: dev
dev:
	@docker compose -f docker-compose.dev.yml up --build

.PHONY: prod
prod:
	@docker compose up --build

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
	@docker compose up frontend --build

.PHONY: publish
publish: publish-api publish-frontend

.PHONY: publish-api
publish-api:
	@docker build -t slavatidika/api backend/ && docker push slavatidika/api

.PHONY: publish-frontend
publish-frontend:
	@docker build --target=deploy -t slavatidika/frontend . && docker push slavatidika/frontend
