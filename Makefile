migration-create:
	@migrate create -ext sql -dir migrate/migrations/ -seq $(filter-out $@,$(MAKECMDGOALS))

migration-up:
	@go run migrate/main.go -up

migration-down:
	@go run migrate/main.go -down
