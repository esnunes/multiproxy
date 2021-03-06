.PHONY: help # List all available targets
help:
	@grep '^.PHONY: .* #' Makefile | sed 's/\.PHONY: \(.*\) # \(.*\)/\1	\2/' | expand -t15

.PHONY: run # Run the application
run:
	@fileb0x assets.yaml
	@go run cmd/multiproxy/main.go examples/config.json

.PHONY: test # Run all tests including code coverage
test:
	@go test -v -cover ./... -coverprofile=coverage.out

.PHONY: cov # Run coverage report
cov: test
	@go tool cover -html=coverage.out

.PHONY: dep # Force download of all Go dependencies
dep:
	@dep ensure -v
	@go get -u -v github.com/UnnoTed/fileb0x

.PHONY: docker # Build docker image
docker:
	@docker build -t esnunes/multiproxy .

.PHONY: web-build
web-build:
	@cd web; yarn install; yarn build; cd -

