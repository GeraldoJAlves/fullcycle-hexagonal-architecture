test:
	@go test ./...
coverage:
	@go test ./... -cover -coverprofile cover.out
	@go tool cover -html=cover.out