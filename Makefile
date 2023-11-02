createmigration:
	migrate create -ext=sql -dir=pkg/sql/migrations -seq init

migrate:
	migrate -path=pkg/sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose up

migratedown:
	migrate -path=pkg/sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose down

.PHONY: createmigration migrate migratedown