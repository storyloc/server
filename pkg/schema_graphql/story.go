package schemaGraphql

import (
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/mitchellh/mapstructure"
	"github.com/storyloc/server/pkg/storage"
)

func (s Schema) storyFields() graphql.Fields {
	return graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"owner": &graphql.Field{
			Type: graphql.NewNonNull(s.profile),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				story, ok := p.Source.(*storage.Story)
				if !ok {
					return nil, fmt.Errorf("unknown source type %T", p.Source)
				}

				return s.profileService.GetProfile(story.OwnerId)
			},
		},
		"createdAt": &graphql.Field{
			Type: graphql.DateTime,
		},
		"updatedAt": &graphql.Field{
			Type: graphql.DateTime,
		},
	}
}

func (s Schema) storyField() *graphql.Field {
	return &graphql.Field{
		Type: s.story,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id, ok := p.Args["id"].(string)
			if !ok {
				return nil, errors.New(`missing required argument "id"`)
			}

			return s.storyService.GetStory(id)
		},
	}
}

func (s Schema) storiesField() *graphql.Field {
	return &graphql.Field{
		Name: "Stories",
		Type: graphql.NewList(s.story),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return s.storyService.AllStories()
		},
	}
}

func (s Schema) createStoryField() *graphql.Field {
	return &graphql.Field{
		Type: s.story,
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Description: "Create a story with provided fields.",
				Type: graphql.NewInputObject(graphql.InputObjectConfig{
					Name: "CreateStoryInput",
					Fields: graphql.InputObjectConfigFieldMap{
						"name": &graphql.InputObjectFieldConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"ownerId": &graphql.InputObjectFieldConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
					},
				}),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			input, ok := p.Args["input"].(map[string]interface{})
			if !ok {
				return nil, errors.New("invalid input type")
			}

			var ts storage.Story
			if err := mapstructure.Decode(input, &ts); err != nil {
				return nil, err
			}

			return s.storyService.CreateStory(ts)
		},
	}
}
