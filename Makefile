all: dev

.PHONY: dev
dev:
	@docker compose up --build

.PHONY: frontend
frontend:
	@docker compose up frontend

.PHONY: build-lint
build-lint:
	@docker build --target lint . 2>&1 >/dev/null | tee .build.log

.PHONY: lint
lint: build-lint
	@docker run --rm -it $(shell grep "writing image" .build.log | head -n 1 | cut -d ' ' -f 4)
