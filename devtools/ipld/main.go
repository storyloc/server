//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/ipld/go-ipld-prime/schema"
	gengo "github.com/ipld/go-ipld-prime/schema/gen/go"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Must specify destination directory")
		os.Exit(1)
	}

	ts := schema.TypeSystem{}
	ts.Init()
	adjCfg := &gengo.AdjunctCfg{
		CfgUnionMemlayout: map[schema.TypeName]string{
			"Any": "interface",
		},
	}

	/*
		## Question
		- do we have the dataModel drawn somewhere?
		## ToDo schemas
		- check optional and nullable booleans twice
		- check if use plural or list notation, example: Audio -> Audios or Audio_List, or even mor real world like Recordings?
		- ConnectionType
		- InterestLevel
		- World -> profiles
		- World -> communites
		- Community -> Connections -> [Connection]
		- Community -> Contents -> [Link]
		- Community -> Template -> Link
		- Profile -> Friends -> [Profile]
		- Interest -> Level -> InterestLevel
		- Location -> In -> Location
		- Connection
		- Group
		- DAO
		- Company
		- Project
		- Task
		- Decision
		- Page
		- Chat
		- Event
		- Activity
		- Audio -> Length -> Bytes
		##
	*/

	ts.Accumulate(schema.SpawnBytes("Bytes"))
	ts.Accumulate(schema.SpawnString("String"))
	ts.Accumulate(schema.SpawnFloat("Int"))
	ts.Accumulate(schema.SpawnFloat("Float"))
	ts.Accumulate(schema.SpawnLink("Link"))
	ts.Accumulate(schema.SpawnBool("Bool"))

	ts.Accumulate(schema.SpawnStruct("World", []schema.StructField{}, schema.SpawnStructRepresentationMap(map[string]string{})))

	ts.Accumulate(schema.SpawnStruct("Community", []schema.StructField{
		schema.SpawnStructField("Name", "String", false, false),
		schema.SpawnStructField("Members", "Profiles", true, false),
		schema.SpawnStructField("Tags", "Tags", true, false),
	}, schema.SpawnStructRepresentationMap(map[string]string{})))
	ts.Accumulate(schema.SpawnList("Communities", "Community", false))

	ts.Accumulate(schema.SpawnStruct("Profile", []schema.StructField{
		schema.SpawnStructField("Name", "String", false, false),
		schema.SpawnStructField("Address", "String", false, false),
		schema.SpawnStructField("Image", "Image", false, false),
		schema.SpawnStructField("Locations", "LocationsOfInterest", false, false),
		schema.SpawnStructField("Position", "GeoPosition", false, false),
		schema.SpawnStructField("Interests", "Interests", false, false),
	}, schema.SpawnStructRepresentationMap(map[string]string{})))
	ts.Accumulate(schema.SpawnList("Profiles", "Profile", false))

	ts.Accumulate(schema.SpawnStruct("Image", []schema.StructField{
		schema.SpawnStructField("Data", "Bytes", false, false),
		schema.SpawnStructField("Size", "Int", false, false),
	}, schema.SpawnStructRepresentationMap(map[string]string{})))

	ts.Accumulate(schema.SpawnStruct("GeoPosition", []schema.StructField{
		schema.SpawnStructField("Longitude", "Float", false, false),
		schema.SpawnStructField("Latitude", "Float", false, false),
	}, schema.SpawnStructRepresentationMap(map[string]string{})))

	ts.Accumulate(schema.SpawnStruct("Tag", []schema.StructField{
		schema.SpawnStructField("Name", "String", false, false),
	}, schema.SpawnStructRepresentationMap(map[string]string{})))
	ts.Accumulate(schema.SpawnList("Tags", "Tag", false))

	ts.Accumulate(schema.SpawnStruct("Interest", []schema.StructField{
		schema.SpawnStructField("Tag", "Tag", false, false),
		schema.SpawnStructField("Experience", "Int", false, false),
	}, schema.SpawnStructRepresentationMap(map[string]string{})))
	ts.Accumulate(schema.SpawnList("Interests", "Interest", false))

	ts.Accumulate(schema.SpawnStruct("Location", []schema.StructField{
		schema.SpawnStructField("Name", "String", false, false),
	}, schema.SpawnStructRepresentationMap(map[string]string{})))
	ts.Accumulate(schema.SpawnList("Locations", "Location", false))

	ts.Accumulate(schema.SpawnStruct("LocationOfInterest", []schema.StructField{
		schema.SpawnStructField("History", "Locations", false, false),
		schema.SpawnStructField("Current", "Location", false, false),
		schema.SpawnStructField("Planning", "Locations", false, false),
	}, schema.SpawnStructRepresentationMap(map[string]string{})))
	ts.Accumulate(schema.SpawnList("LocationsOfInterest", "LocationOfInterest", false))

	ts.Accumulate(schema.SpawnStruct("Connection", []schema.StructField{}, schema.SpawnStructRepresentationMap(map[string]string{})))

	ts.Accumulate(schema.SpawnStruct("Group", []schema.StructField{}, schema.SpawnStructRepresentationMap(map[string]string{})))

	ts.Accumulate(schema.SpawnStruct("DAO", []schema.StructField{}, schema.SpawnStructRepresentationMap(map[string]string{})))

	ts.Accumulate(schema.SpawnStruct("Company", []schema.StructField{}, schema.SpawnStructRepresentationMap(map[string]string{})))

	ts.Accumulate(schema.SpawnStruct("Project", []schema.StructField{}, schema.SpawnStructRepresentationMap(map[string]string{})))

	ts.Accumulate(schema.SpawnStruct("Template", []schema.StructField{
		schema.SpawnStructField("Group", "Group", false, false),
		schema.SpawnStructField("DAO", "DAO", false, false),
		schema.SpawnStructField("Company", "Company", false, false),
		schema.SpawnStructField("Project", "Project", false, false),
	}, schema.SpawnStructRepresentationMap(map[string]string{})))

	ts.Accumulate(schema.SpawnStruct("Task", []schema.StructField{}, schema.SpawnStructRepresentationMap(map[string]string{})))

	ts.Accumulate(schema.SpawnStruct("Decision", []schema.StructField{}, schema.SpawnStructRepresentationMap(map[string]string{})))

	ts.Accumulate(schema.SpawnStruct("Page", []schema.StructField{}, schema.SpawnStructRepresentationMap(map[string]string{})))

	ts.Accumulate(schema.SpawnStruct("Chat", []schema.StructField{}, schema.SpawnStructRepresentationMap(map[string]string{})))

	ts.Accumulate(schema.SpawnStruct("Event", []schema.StructField{}, schema.SpawnStructRepresentationMap(map[string]string{})))

	ts.Accumulate(schema.SpawnStruct("Activity", []schema.StructField{}, schema.SpawnStructRepresentationMap(map[string]string{})))

	ts.Accumulate(schema.SpawnStruct("Audio", []schema.StructField{
		schema.SpawnStructField("Data", "Bytes", false, false),
		schema.SpawnStructField("Date", "Int", false, false),
	}, schema.SpawnStructRepresentationMap(map[string]string{})))
	ts.Accumulate(schema.SpawnList("Audios", "Audio", false))

	ts.Accumulate(schema.SpawnStruct("Story", []schema.StructField{
		schema.SpawnStructField("Owner", "Profile", false, false),
		schema.SpawnStructField("Date", "Int", false, false),
		schema.SpawnStructField("Image", "Image", false, false),
		schema.SpawnStructField("GeoPosition", "GeoPosition", false, false),
		schema.SpawnStructField("Recordings", "Audios", false, false),
		schema.SpawnStructField("Tags", "Tags", false, false),
		schema.SpawnStructField("IsPublic", "Bool", false, false),
		schema.SpawnStructField("Retention", "Int", false, false),
	}, schema.SpawnStructRepresentationMap(map[string]string{})))
	ts.Accumulate(schema.SpawnList("Stories", "Story", false))

	ts.Accumulate(schema.SpawnStruct("StoryCollection", []schema.StructField{
		schema.SpawnStructField("Name", "String", false, false),
		schema.SpawnStructField("Stories", "Stories", false, false),
		schema.SpawnStructField("StartLocation", "Location", false, false),
	}, schema.SpawnStructRepresentationMap(map[string]string{})))

	ts.Accumulate(schema.SpawnStruct("Content", []schema.StructField{
		schema.SpawnStructField("Task", "Task", false, false),
		schema.SpawnStructField("Decision", "Decision", false, false),
		schema.SpawnStructField("Page", "Page", false, false),
		schema.SpawnStructField("Chat", "Chat", false, false),
		schema.SpawnStructField("Event", "Event", false, false),
		schema.SpawnStructField("Activity", "Activity", false, false),
		schema.SpawnStructField("Story", "Story", false, false),
		schema.SpawnStructField("StoryCollection", "StoryCollection", false, false),
	}, schema.SpawnStructRepresentationMap(map[string]string{})))

	if errs := ts.ValidateGraph(); errs != nil {
		for _, err := range errs {
			fmt.Printf("- %s\n", err)
		}
		panic("not happening")
	}

	gengo.Generate(os.Args[1], "ipld", ts, adjCfg)
}
