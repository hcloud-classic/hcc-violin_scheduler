package graphql

import (
	"errors"
	graphqlType "hcc/violin-scheduler/action/graphql/type"
	"hcc/violin-scheduler/lib/logger"

	"github.com/graphql-go/graphql"
)

var queryTypes = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			// server DB
			"server": &graphql.Field{
				Type:        graphqlType.ServerType,
				Description: "Get server by uuid",
				Args: graphql.FieldConfigArgument{
					"uuid": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					logger.Logger.Println("Resolving: server")
					// return dao.ReadServer(params.Args)
					return "Not Use This", errors.New("Not Used")
				},
			},
		},
	})
