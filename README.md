INSTALLATION

- Install dependencies
```bash
glide up --quick -v
```
- Create databse
```sql
 CREATE DATABASE IF NOT EXISTS <databasename>
```
- Set up credantials
```bash
nano parameters.ini
```
- Run migrations
```bash
go run ./web/migrtions.go
```
- Run webserver
```bash
go run ./web/index.go
```

USAGE
- Load fixtures
```bash
go run ./web/fixtures.go
```
- Get places suggestions
```bash
curl "http://192.168.33.10:8080/cities-suggestions?q=london"
```
- Get cities info
```bash
curl -XPOST "http://192.168.33.10:8080/cities-info" -d '{"place_id" : ["ChIJdd4hrwug2EcRmSrV3Vo6llI"]}'
```
- Get distances between cities
```bash
go run ./cmd/distance.go Berlin Moscow London
```