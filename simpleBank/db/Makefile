postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=secret -d postgres

createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres simple_bank

dropdb:
	docker exec -it postgres psql -U postgres -d postgres -c "DROP DATABASE simple_bank";
	#docker exec -it postgres --username=postgres dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down


.PHONY: postgres createdb dropdb migrateup migratedown

#todo:check, pass