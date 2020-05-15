package graphqlType

import (
	"github.com/graphql-go/graphql"
)

// SelectorNodeType: Form violin
var SelectorNodeType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Selector",
		Fields: graphql.Fields{
			"node_uuid": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
		},
	},
)
