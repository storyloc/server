package gql

import "github.com/graphql-go/graphql"

func (s schema) queryStory() *graphql.Field {
	return &graphql.Field{
		Type:        typeStory,
		Description: "get single story",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			id, _ := params.Args["id"].(string)
			return s.storyService.GetStory(id)
		},
	}
}
