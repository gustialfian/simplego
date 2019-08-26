run:
	go run cmd/simplego/main.go \
  		--db "user=postgres password=mysecretpassword host=localhost port=8080 dbname=sandbox sslmode=disable" \
  		--port 8000