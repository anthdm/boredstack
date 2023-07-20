build-app:
	@go build -o bin/app ./cmd/app/

build-cli:
	@go build -o bin/bs ./cmd/bs/

run: build-app
	@./bin/app

clean: 
	@rm -rf bin