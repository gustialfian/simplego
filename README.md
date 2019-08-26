# simplego

## Run
```bash
make
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
- validation