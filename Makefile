run:
	go run ./cmd/api

.PHONY: run

help:
	go run ./cmd/api -help

PG_CONTAINER = postgres_be_movies
PG_USER = postgres
PG_PASSWORD = postgres
PG_DB = go_movies_1


.PHONY: pg_up
pg_up:
	docker run --name $(PG_CONTAINER) -e POSTGRES_PASSWORD=$(PG_PASSWORD) -e POSTGRES_USER=$(PG_USER) -p 5432:5432 -d postgres

.PHONY: pg_down
pg_down:
	docker stop $(PG_CONTAINER)
	docker rm $(PG_CONTAINER)

.PHONY: pg_connect_root
pg_connect:
	docker exec -it $(PG_CONTAINER) psql -U $(PG_USER) 

.PHONY: pg_connect_user
pg_connect_user:
	docker exec -it $(PG_CONTAINER) psql --username=greenlight --dbname=greenlight


.PHONY: pg_init_db
pg_init_db:
	docker exec -i $(PG_CONTAINER) psql -U $(PG_USER) < ./internal/sql/init_db/init.sql
