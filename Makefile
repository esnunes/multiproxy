.PHONY: help # List all available targets
help:
	@grep '^.PHONY: .* #' Makefile | sed 's/\.PHONY: \(.*\) # \(.*\)/\1	\2/' | expand -t15

.PHONY: run # Run the application
run:
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

.PHONY: docker # Build docker image
docker:
	@docker build -t esnunes/multiproxy .
