INSTALLATION

- Install dependencies
```bash
cd ./src/
glide up --quick -v
```
- Run webserver
```bash
go run ./src/web/index.go
```

USAGE
- Get places suggestions
```bash
curl "http://192.168.33.10:8080/cities-suggestions?q=london"
```
- Get cities info
```bash
curl -XPOST "http://192.168.33.10:8080/places/" -d '{"place_id" : ["b1a8b96daab5065cf4a08f953e577c34cdf769c0"]}'
```
