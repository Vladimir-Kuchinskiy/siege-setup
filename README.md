# siege-setup

setup DB:

```bash
docker compose up -d
```

run api:
```bash
go run main.go
```

run siege:
```bash
siege -f siege.txt -b -c50 -t15s
```
