format:
	goimports -w ./

lint:
	golint ./

test:
	go test ./...

ci:
	@echo "formatting..."
	@goimports -w ./

	@echo "linting..."
	@golint ./

	@echo "testing..."
	@go test ./...
