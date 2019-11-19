package graphqlType

import "github.com/graphql-go/graphql"

// SelectorNodeType: Form violin
var SelectorNodeType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Selector",
		Fields: graphql.Fields{
			"server_uuid": &graphql.Field{
				Type: graphql.String,
			},
			"cpu": &graphql.Field{ // quota CPU
				Type: graphql.Int,
			},
			"memory": &graphql.Field{ //quota mem
				Type: graphql.Int,
			},
			"number_of_nodes": &graphql.Field{ //Should Clustering of nodes
				Type: graphql.Int,
			},
			"list_node": &graphql.Field{ //Should Clustering of nodes
				Type: graphql.Int,
			},
		},
	},
)
