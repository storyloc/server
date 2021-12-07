package schemaGraphql

import (
	"github.com/graphql-go/graphql"
	"github.com/storyloc/server/pkg/service"
)

type Schema struct {
	// services
	profileService service.ProfileService
	storyService   service.StoryService
	// objects
	profile *graphql.Object
	story   *graphql.Object
}

func NewSchema(profileService service.ProfileService, storyService service.StoryService) (graphql.Schema, error) {
	obj := func(name, desc string, ifaces ...*graphql.Interface) *graphql.Object {
		return graphql.NewObject(graphql.ObjectConfig{
			Name:        name,
			Description: desc,
			Fields:      graphql.Fields{},
			Interfaces:  ifaces,
		})
	}

	addFields := func(o *graphql.Object, f graphql.Fields) {
		for n, f := range f {
			o.AddFieldConfig(n, f)
		}
	}

	schema := Schema{
		storyService:   storyService,
		profileService: profileService,

		profile: obj("Profile", "A profile."),
		story:   obj("Story", "A story."),
	}

	addFields(schema.story, schema.storyFields())
	addFields(schema.profile, schema.profileFields())

	return graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"profile": schema.profileField(),
				"story":   schema.storyField(),
				"stories": schema.storiesField(),
			},
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				"createProfile": schema.createProfileField(),
				"createStory":   schema.createStoryField(),
			},
		}),
	})
}
