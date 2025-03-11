# デフォルトのターゲットを追加
.PHONY: migrate

migrate: 
	POSTGRES_PASSWORD=postgrespassword GO_ENV=dev go run migrate/migrate.go

run: 
	GO_ENV=dev go run main.go

# Docker起動するのを忘れないこと