# Worker

Worker is written on golang to handle many events from Message Queue and deployed as side container to Redis
with low resource (CPU, RAM) consumption.

### Configuration

Copy [.env.template](.env.template) it to `.env` file, edit variables.

To run without docker export variables to env (example - `export $(cat .env | egrep -v "(^#.*|^$)" | xargs)`)

To run with docker-compose do nothing

### Running locally

```bash
go run cmd/worker/main.go
```

### Running in docker

```bash
docker-compose up
```