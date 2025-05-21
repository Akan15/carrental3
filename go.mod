module github.com/Akan15/carrental3

go 1.23.0

toolchain go1.24.1

replace github.com/Akan15/carrental3/user-service => ./user-service

replace github.com/Akan15/carrental3/car-service => ./car-service

replace github.com/Akan15/carrental3/rental-service => ./rental-service

replace github.com/Akan15/carrental3/api-gateway => ./api-gateway

require (
	github.com/klauspost/compress v1.18.0 // indirect
	github.com/nats-io/nats.go v1.42.0 // indirect
	github.com/nats-io/nkeys v0.4.11 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	golang.org/x/crypto v0.37.0 // indirect
	golang.org/x/net v0.35.0 // indirect
	golang.org/x/sys v0.32.0 // indirect
	golang.org/x/text v0.24.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250218202821-56aae31c358a // indirect
	google.golang.org/grpc v1.72.1 // indirect
	google.golang.org/protobuf v1.36.5 // indirect
)
