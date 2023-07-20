build-app:
	@go build -o bin/app ./cmd/app/

build-cli:
	@go build -o bin/bs ./cmd/bs/

run: build-app
	@./bin/app

db-init: build-cli
	@./bin/bs db init

db-migrate: build-cli
	@./bin/bs db migrate

db-rollback: build-cli
	@./bin/bs db rollback

db-lock: build-cli
	@./bin/bs db lock

db-unlock: build-cli
	@./bin/bs db unlock

db-create-sql: build-cli
	@./bin/bs db create_sql $(name)

db-create-go: build-cli
	@./bin/bs db create_go $(name)

db-status: build-cli
	@./bin/bs db status

db-mark-applied: build-cli
	@./bin/bs db mark_applied

clean:
	@rm -rf bin

