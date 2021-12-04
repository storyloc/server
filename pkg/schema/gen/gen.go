//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"os"
	"path"

	gengraphql "github.com/ipld/go-ipld-graphql/gen"
	"github.com/ipld/go-ipld-prime/schema"
	gengo "github.com/ipld/go-ipld-prime/schema/gen/go"
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

	ts.Accumulate(schema.SpawnBytes("Bytes"))
	ts.Accumulate(schema.SpawnString("String"))
	ts.Accumulate(schema.SpawnInt("Int"))

	ts.Accumulate(schema.SpawnStruct("World", []schema.StructField{
		schema.SpawnStructField("Communites", "Communites", false, false),
	}, schema.SpawnStructRepresentationMap(map[string]string{})))

	ts.Accumulate(schema.SpawnStruct("Community", []schema.StructField{
		schema.SpawnStructField("Name", "String", false, false),
		schema.SpawnStructField("Image", "Image", true, false),
	}, schema.SpawnStructRepresentationMap(map[string]string{})))
	ts.Accumulate(schema.SpawnList("Communites", "Community", false))

	ts.Accumulate(schema.SpawnStruct("Image", []schema.StructField{
		schema.SpawnStructField("Data", "Bytes", false, false),
		schema.SpawnStructField("Size", "Int", false, false),
	}, schema.SpawnStructRepresentationMap(map[string]string{})))

	if errs := ts.ValidateGraph(); errs != nil {
		for _, err := range errs {
			fmt.Printf("- %s\n", err)
		}
		panic("not happening")
	}

	gengo.Generate(os.Args[1], "schema", ts, adjCfg)
	gengraphql.Generate(path.Join(os.Args[1], "..", "graphql"), "graphql", ts, "schema", "github.com/storyloc/server/pkg/schema")
}
