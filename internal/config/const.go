package config

const (
	AppName string = "friendly-spoon"

	DefaultHealthcheckPort uint = 8080

	DefaultExternalGRPCHost string = ":50051"
	DefaultInternalGRPCHost string = ":50052"

	DefaultDBSchemaURL string = "file://sql"
)
