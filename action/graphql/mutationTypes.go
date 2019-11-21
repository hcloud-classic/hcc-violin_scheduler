package graphql

import (
	graphqlType "hcc/violin-scheduler/action/graphql/type"
	"hcc/violin-scheduler/driver"
	"hcc/violin-scheduler/lib/logger"

	"github.com/graphql-go/graphql"
)

var mutationTypes = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		// server DB
		"schedule_nodes": &graphql.Field{
			Type:        graphqlType.SelectorNodeType,
			Description: "Scheduling Nodes",
			Args: graphql.FieldConfigArgument{
				"server_uuid": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"cpu": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"memory": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"nr_node": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				logger.Logger.Println("Resolving: schedule_nodes")
				return driver.ScheduleNodes(params.Args)
			},
		},

		"selected_nodes": &graphql.Field{
			Type:        graphql.NewList(graphql.String),
			Description: "Scheduling Nodes",
			Args: graphql.FieldConfigArgument{
				"server_uuid": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"cpu": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"memory": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"nr_node": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				logger.Logger.Println("Resolving: schedule_nodes")
				return driver.TestSchedule(params.Args)
			},
		},
	},
})
