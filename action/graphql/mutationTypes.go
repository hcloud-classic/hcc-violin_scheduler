package graphql

import (
	graphqlType "hcc/violin-scheduler/action/graphql/type"
	"hcc/violin-scheduler/lib/logger"
	"hcc/violin/driver"

	"github.com/graphql-go/graphql"
)

var mutationTypes = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		// server DB
		"create_server": &graphql.Field{
			Type:        graphqlType.ServerType,
			Description: "Create new server",
			Args: graphql.FieldConfigArgument{
				"subnet_uuid": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"os": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"server_name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"server_desc": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"cpu": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"memory": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"disk_size": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"status": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"user_uuid": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				logger.Logger.Println("Resolving: create_server")
				return driver.CreateServer(params)
			},
		},
	},
})
