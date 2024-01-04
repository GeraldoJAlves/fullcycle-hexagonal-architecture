mock:
	@mockgen -destination=internal/application/mocks/application.go -source=internal/application/product.go application
test:
	@go test ./...
coverage:
	@go test ./... -cover -coverprofile cover.out
	@go tool cover -html=cover.out
db_recreate:
	@docker compose exec app bash -c "rm sqlite.db && sqlite3 sqlite.db < ./docker/init.sql"
db_del_products:
	@docker compose exec app bash -c "echo 'delete from products' | sqlite3 sqlite.db"
db_get_products:
	@docker compose exec app bash -c "echo 'select * from products' | sqlite3 sqlite.db"
