# Worker

Worker is written on golang to handle many events from Message Queue and deployed as side container to Redis
with low resource (CPU, RAM) consumption.


### Configuration

Export all variables in [.env.template](.env.template)

### Running locally

```bash
go run cmd/worker/main.go
```