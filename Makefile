proto:
	protoc --go_out=. --go-grpc_out=. ./internal/orb/proto/orb.proto

serve:
	go run . serve

beat:
	go run . beat

tests:
	go test ./...

up:
	docker compose up server-dev

start:
	docker compose start server-dev

stop:
	docker compose stop

down:
	docker compose down --remove-orphans -v

logs:
	docker compose logs -f

bash:
	docker compose exec server-dev sh
