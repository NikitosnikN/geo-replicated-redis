# API Server

API server is written on FastAPI (python) as fast solution and will be rewritten to webserver based to Go to handle 
much more concurrent requests.


### Installation

```bash
poetry install
```

### Configuration

Go to [config/.env.template](config%2F.env.template) and copy it to `config/.env` file, edit variables.

### Running locally

```bash
uvicorn src.app:app
```

### Running in docker

```bash
docker-compose up
```


