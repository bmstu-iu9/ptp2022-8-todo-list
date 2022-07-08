all: dev

.PHONY: dev
dev:
	@docker compose up --build

.PHONY: frontend
frontend:
	@docker compose up frontend

.PHONY: lint
lint:
	@docker run -it --rm $$(docker build --target lint -q .)
