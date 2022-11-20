package config

import "time"

const (
	AppName string = "friendly-spoon"

	DefaultRESTPort        uint = 8080
	DefaultHealthcheckPort uint = 8081

	DefaultExternalGRPCPort uint = 50051
	DefaultInternalGRPCPort uint = 50052

	DefaultDBSchemaURL string = "file://sql"

	SessionsSecret   string        = "sf7+lSLzlDktkUBII2LkNTf1J3Xt6UtvP7goWhkoLNV4=" //todo: move to env
	SessionsLifetime time.Duration = 60 * time.Minute
)
