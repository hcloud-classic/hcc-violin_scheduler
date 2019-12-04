package graphqlType

import (
	"github.com/graphql-go/graphql"
)

// SelectorNodeType: Form violin
var SelectorNodeType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Selector",
		Fields: graphql.Fields{
			"node_uuid": &graphql.Field{ //Should Clustering of nodes
				Type: graphql.NewList(graphql.String),
			},
		},
	},
)
