package config

import "github.com/joho/godotenv"

type AppConfig struct {
	App struct {
		Env string
	}
	GraphQL struct {
		Port string
	}
	Postgres struct {
		Name string
		User string
		Pass string
		Host string
		Port string
		SSL  string
	}
}

var appConfig *AppConfig

func NewAppConfig() *AppConfig {
	// add config file path in .env
	_ = godotenv.Load("../.env")

	if appConfig == nil {
		appConfig = &AppConfig{}

		initApp(appConfig)
		initPostgres(appConfig)
		initGraphQL(appConfig)
	}

	return appConfig
}
