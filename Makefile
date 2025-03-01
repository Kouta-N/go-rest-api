# デフォルトのターゲットを追加
.PHONY: migrate

migrate: 
	POSTGRES_PASSWORD=postgrespassword GO_ENV=dev go run migrate/migrate.go