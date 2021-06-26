migrations-status:
	goose -dir ./migrations postgres "user=admin password=admin dbname=postgres sslmode=disable" status

migrations-new:
	goose -dir migrations create $(NAME) sql

migrations-up:
	goose -dir ./migrations  postgres "user=admin password=admin dbname=postgres sslmode=disable" up

migrations-down:
	goose -dir ./migrations  postgres "user=admin password=admin dbname=postgres sslmode=disable" down