# GO short
## Description
This is the implementation of the secong exercise of the [Gophercises](https://gophercises.com/) course.  
The exercise is about creating a URL shortener.
The proyect has 3 binaris that can be used:
- `cmd/go-short/go-short.go`: This is the first exercice from the course a shorter handle from a YAML.
- `cmd/go-short-v2/go-short-v2.go`: This is the second exercice from the course a shorter handle from a JSON.
- `cmd/go-short-sql/go-short-sql.go`: This is the third exercice from the course a shorter handle from a Data base.
And has 2 packages:
- `pkg/url-short/url-short.go`: This package has the logic of the shortener from a static file.
- `pkg/sqlite-url-short/sqlite-url-short.go`: This package has the logic of the the shortener from a SQLitle.

## Usage
```bash
go run {{binary option}}
```

## Examples
```bash
go run cmd/go-short/go-short.go
```
```bash
go run cmd/go-short-v2/go-short-v2.go
```
```bash
go run cmd/go-short-sql/go-short-sql.go
```

