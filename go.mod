module github.com/jackstockley89/golangwebpage

go 1.18

require github.com/prometheus/client_golang v1.13.0

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.2-0.20181231171920-c182affec369 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/prometheus/common v0.37.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)

// A replace directive is needed for github.com/prometheus/prometheus to ensure running against the latest version of prometheus.
replace github.com/prometheus/prometheus => github.com/prometheus/prometheus v0.38.0
