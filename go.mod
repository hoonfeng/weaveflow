module ifaceconf

go 1.24.2

require (
	github.com/go-chi/chi/v5 v5.0.12
	github.com/go-chi/cors v1.2.1
	github.com/go-sql-driver/mysql v1.9.3
	github.com/golang-jwt/jwt/v5 v5.2.2
	github.com/hoonfeng/goproc v0.0.0
	github.com/lib/pq v1.10.9
	github.com/microsoft/go-mssqldb v1.9.4
	golang.org/x/crypto v0.38.0
	golang.org/x/sys v0.33.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/Microsoft/go-winio v0.6.2 // indirect
	github.com/golang-sql/civil v0.0.0-20220223132316-b832511892a9 // indirect
	github.com/golang-sql/sqlexp v0.1.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	golang.org/x/text v0.25.0 // indirect
)

replace github.com/hoonfeng/goproc => ./third_party/goproc
