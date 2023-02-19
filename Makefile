migration-create:
	migrate create -ext sql -dir database/migrations -seq $(name)

migration-up:
	migrate -path database/migrations -database "postgresql://postgres:postgres@localhost:5432/videos?sslmode=disable" -verbose up

migration-down:
	migrate -path database/migrations -database "postgresql://postgres:postgres@localhost:5432/videos?sslmode=disable" -verbose down