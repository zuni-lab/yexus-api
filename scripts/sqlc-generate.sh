rm -rf pkg/db/*.sql.go
docker run --rm -v $(pwd):/src -w /src sqlc/sqlc generate
