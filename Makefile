run:
	go run cmd/simplego/main.go \
  		--db "host=localhost port=6543 dbname=sandbox user=sandbox password=sandbox sslmode=disable" \
  		--port 8000