## Envs for the project

### Server
- `SERVER_PORT`: Port for the server to listen on (default: 9000)
- `SERVER_SHUTDOWN_TIMEOUT`: Timeout for the server to shutdown (default: 5s)

### Database
- `DATABASE_DSN`: Database DSN
- `DATABASE_NEED_RECREATE`: If set to `true`, the database will be recreated on start (default: true)

## Server API
### Run
```bash
make start
```