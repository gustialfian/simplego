# simplego

## Run
```bash
go run cmd/simplego/main.go \
  --db "user=postgres password=mysecretpassword host=localhost port=8080 dbname=sandbox sslmode=disable" \
  --port 8000
```

## TODO:
- backing service
  - postgresql
  - redis
  - socket.io
  - rpc
- abstract main to app package
  - registerConfig()
  - registerRoute()
  - registerDB()
- Midleware
  - JWT Auth
  - CORS
  - Midleware for specific route
- split domain to 
  - handler
  - repository
  - service if necessary
- Graceful shutdown