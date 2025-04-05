# App name
# ~/30H master Go master !2 ?3 > go run cmd/server/main.go                                              4s 23:51:27
APP_NAME := server

run:
	go run cmd/server/main.go
swag:
	swag init -g ./cmd/server/main.go -o ./cmd/swag/docs