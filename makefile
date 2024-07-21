init:
	sh ./bin/init.sh

start:
	docker compose up --build

down:
	docker compose down

v-down:
	docker compose down --volumes
