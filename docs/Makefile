
all: validate

.PHONY: validate
validate:
	@docker run --rm -t -v "$(PWD)":/tmp stoplight/spectral lint -vr "/tmp/.spectral.yml" "/tmp/openapi.yaml"
