migration-create:
	migrate create -ext sql -dir database/migrations -seq $(name)