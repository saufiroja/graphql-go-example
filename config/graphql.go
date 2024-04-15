package config

import "os"

func initGraphQL(conf *AppConfig) {
	port := os.Getenv("GRAPHQL_PORT")

	conf.GraphQL.Port = port
}
