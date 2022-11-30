
.PHONY: install-golangci-lint
install-golangci-lint:
	$(call print-target)
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh >lint-install.sh
	chmod u+x ./lint-install.sh && ./lint-install.sh -b $(WORKSPACE)/bin $(golangci-lint-version)
	$(RM) ./lint-install.sh

.PHONY: install-vulncheck
install-vulncheck:
	go install golang.org/x/vuln/cmd/govulncheck@latest

.PHONY: install-gosec
install-gosec:
	go install github.com/securego/gosec/cmd/gosec@latest

.PHONY: lint
lint: install-golangci-lint install-vulncheck
	golangci-lint run -v
	govulncheck ./...
	gosec ./...

.PHONY: install-fmt
install-fmt:
	go install -v golang.org/x/tools/cmd/goimports@latest

.PHONY: fmt
fmt: install-fmt
	gofmt -l -w .
	goimports -l -w .
	go fix ./...

.PHONY: build
build:
	go build cmd/rest-memory-service

.PHONY: run-inmem
run-inmem:
	docker-compose -f docker-compose.yaml up --build

.PHONY: run-mongo
run-mongo:
	docker-compose -f docker-compose.yaml -f docker-compose.generator.yaml up --build
