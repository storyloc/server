package schemaGraphql

import (
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/mitchellh/mapstructure"
	"github.com/storyloc/server/pkg/storage"
)

func (s Schema) profileFields() graphql.Fields {
	return graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"stories": &graphql.Field{
			Type: graphql.NewList(s.story),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				profile, ok := p.Source.(*storage.Profile)
				if !ok {
					return nil, fmt.Errorf("unknown source type %T", p.Source)
				}
				var profileStories []*storage.Story

				stories, err := s.storyService.AllStories()
				if err != nil {
					return nil, err
				}

				for _, story := range stories {
					if story.OwnerId != profile.Id {
						continue
					}

					profileStories = append(profileStories, story)
				}

				return profileStories, nil
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

func (s Schema) profileField() *graphql.Field {
	return &graphql.Field{
		Type: s.profile,
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

			return s.profileService.GetProfile(id)
		},
	}
}

func (s Schema) createProfileField() *graphql.Field {
	return &graphql.Field{
		Type: s.profile,
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Description: "Create a profile with provided fields.",
				Type: graphql.NewInputObject(graphql.InputObjectConfig{
					Name: "CreateProfileInput",
					Fields: graphql.InputObjectConfigFieldMap{
						"name": &graphql.InputObjectFieldConfig{
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

			var tp storage.Profile
			if err := mapstructure.Decode(input, &tp); err != nil {
				return nil, err
			}

			return s.profileService.CreateProfile(tp)
		},
	}
}
