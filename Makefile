init:
	sh ./bin/init.sh

start:
	docker compose up --build

down:
	docker compose down

teardown:
	docker compose down --volumes
