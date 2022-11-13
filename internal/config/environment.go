package config

import "github.com/h-celel/mapenv"

type Environment struct {
	Auth0ClientID     string `mpe:"AUTH0_CLIENT_ID"`
	Auth0Domain       string `mpe:"AUTH0_DOMAIN"`
	Auth0ClientSecret string `mpe:"AUTH0_CLIENT_SECRET"`
	Auth0CallbackURL  string `mpe:"AUTH0_CALLBACK_URL"`

	RESTPort         uint   `mpe:"REST_PORT"`
	HealthcheckPort  uint   `mpe:"HEALTHCHECK_PORT"`
	GRPCExternalPort uint   `mpe:"GRPC_EXTERNAL_PORT"`
	GRPCInternalPort uint   `mpe:"GRPC_INTERNAL_PORT"`
	DBSchemaURL      string `mpe:"DB_SCHEMA_URL"`
}

func NewEnvironment() Environment {
	env := Environment{
		RESTPort:         DefaultRESTPort,
		HealthcheckPort:  DefaultHealthcheckPort,
		GRPCExternalPort: DefaultExternalGRPCPort,
		GRPCInternalPort: DefaultInternalGRPCPort,
		DBSchemaURL:      DefaultDBSchemaURL,
	}
	err := mapenv.Decode(&env)
	if err != nil {
		panic(err)
	}
	return env
}
