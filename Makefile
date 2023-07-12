build-app:
	@go build -o bin/app ./cmd/app/

run: build-app
	@./bin/app

clean: 
	@rm -rf bin