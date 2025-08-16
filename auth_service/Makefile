MIGRATION_FOLDER=db/migrations
DB_USER=root:mySQL_1211@tcp(127.0.0.1:3306)/auth_dev

migrate-create:
	goose -dir $(MIGRATION_FOLDER) create $(name) sql

migrate-up:
	goose -dir $(MIGRATION_FOLDER) mysql "$(DB_USER)" up

migrate-down:
	goose -dir $(MIGRATION_FOLDER) mysql "$(DB_USER)" down

