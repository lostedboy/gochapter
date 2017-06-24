INSTALLATION

- Install dependencies
```bash
glide up --quick -v
```
- Run webserver
```bash
go run ./web/index.go --key=<google API key>
```

USAGE
- Get places suggestions
```bash
curl "http://192.168.33.10:8080/cities-suggestions?q=london"
```
- Get cities info
```bash
curl -XPOST "http://192.168.33.10:8080/cities-info" -d '{"place_id" : ["b1a8b96daab5065cf4a08f953e577c34cdf769c0"]}'
```
- Get distances between cities
```bash
go run ./cmd/distance.go --key=<google API key> Berlin Moscow London
```