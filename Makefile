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

.PHONY: build-deploy
build-deploy:
	docker build --target deploy -t frontend-deploy .

.PHONY: deploy
deploy: build-deploy
	docker run --rm -it -p "3000:3000" frontend-deploy 
