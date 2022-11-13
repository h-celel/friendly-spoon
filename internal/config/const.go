package config

import "time"

const (
	AppName string = "friendly-spoon"

	DefaultRESTPort        uint = 8080
	DefaultHealthcheckPort uint = 8081

	DefaultExternalGRPCPort uint = 50051
	DefaultInternalGRPCPort uint = 50052

	DefaultDBSchemaURL string = "file://sql"

	SessionsSecret   string        = "some-secret"
	SessionsLifetime time.Duration = 15 * time.Minute
)
