run:
	go run cmd/main.go

gen-swag:
	swag init -g ./api/api.go -o ./api/docs
