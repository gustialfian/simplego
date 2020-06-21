# simplego

## Run
```bash
# shell 1
docker run --name db --rm \
  -e POSTGRES_PASSWORD=sandbox \
  -e POSTGRES_USER=sandbox \
  -e POSTGRES_DB=sandbox \
  -p 6543:5432 \
  postgres:13-alpine

# shell 2
make
```

## TODO:
- backing service
  - [x] postgresql
  - [ ] redis
  - [ ] nats
  - [ ] rpc
- abstract main to app package
  - [x] registerConfig()
  - [x] registerRoute()
  - [x] registerDB()
- Midleware
  - [x] CORS
  - [ ] JWT Auth
  - [ ] Midleware for specific route
- split domain to 
  - handler
  - repository
  - service if necessary
- Graceful shutdown
- validation