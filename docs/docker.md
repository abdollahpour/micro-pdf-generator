You can simple use docker to spin up a new micro-pdf-generator. For example:

```bash
docker run -d 8080:8080 abdollahpour/micro-pdf-generator
```

Server is ready on port 8080 of localhost

You can specify more setting using environment variables:

```
docker run -d 8080:8080 \
    -e MPG_PORT=8080 \
    -e MPG_HOST=0.0.0.0 \
    -e MPG_TIMEOUT=15 \
    -e MPG_MAX_SIZE=6 \
    abdollahpour/micro-pdf-generator
```

for more information please check [configurations](configurations.md).