createmigration:
	migrate create -ext=internal/sql -dir=internal/sql/migrations -seq init

migrate:
	migrate -path=internal/sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose up

migratedown:
	migrate -path=internal/sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose down

.PHONY: createmigration migrate migratedown