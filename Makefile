run :
	@go run cmd/api/main.go
run-with-air:
	@air
build-image:
	@docker build -t go-rest-api:latest .