
all: validate

.PHONY: validate
validate:
	@docker run --volume "$(PWD)":/data jamescooke/openapi-validator openapi.yml
