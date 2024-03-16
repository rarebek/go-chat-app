postgresinit:
	sudo docker run --name postgres -p 5431:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=postgres -d postgres:latest

postgres:
	sudo docker exec -it postgres psql

createdb:
	sudo sudo docker exec -it postgres createdb --username=root --owner=root go-chat

dropdb:
	sudo sudo docker -it postgres dropdb go-chat

migrateup:
	migrate -path server/db/migrations -database "postgres://root:password@localhost:5431/go-chat?sslmode=disable" -verbose up

migratedown:
	migrate -path server/db/migrations -database "postgres://root:password@localhost:5431/go-chat?sslmode=disable" -verbose down


.PHONY: postgresinit postgres createdb dropdb migrateup migratedown