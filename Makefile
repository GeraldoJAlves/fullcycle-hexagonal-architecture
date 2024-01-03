mock:
	@mockgen -destination=internal/application/mocks/application.go -source=internal/application/product.go application
test:
	@go test ./...
coverage:
	@go test ./... -cover -coverprofile cover.out
	@go tool cover -html=cover.out