module github.com/jackstockley89/golangwebpage

go 1.18

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/google/uuid v1.3.0
	github.com/lib/pq v1.10.6
	github.com/prometheus/client_golang v1.13.0
	github.com/stretchr/testify v1.8.0
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.2-0.20181231171920-c182affec369 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	github.com/rogpeppe/go-internal v1.9.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/prometheus/common v0.37.0 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
)

// A replace directive is needed for github.com/prometheus/prometheus to ensure running against the latest version of prometheus.
replace github.com/prometheus/prometheus => github.com/prometheus/prometheus v0.38.0
