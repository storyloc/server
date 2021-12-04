// Code generated by github.com/ipld/go-ipld-prime/schema/gen/graphql/server, DO NOT EDIT.

package graphql

import (
	"context"
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	ipld "github.com/ipld/go-ipld-prime"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
	"github.com/storyloc/server/pkg/schema"
)

type nodeLoader func(ctx context.Context, cid cidlink.Link, builder ipld.NodeBuilder) (ipld.Node, error)

const nodeLoaderCtxKey = "NodeLoader"

var errNotNode = fmt.Errorf("Not IPLD Node")
var errInvalidLoader = fmt.Errorf("Invalid Loader Provided")
var errInvalidLink = fmt.Errorf("Invalid link")
var errUnexpectedType = "Unexpected type %T. expected %s"

func resolve_map_at(p graphql.ResolveParams) (interface{}, error) {
	ts, ok := p.Source.(ipld.Node)
	if !ok {
		return nil, errNotNode
	}
	arg := p.Args["key"]

	switch ta := arg.(type) {
	case ipld.Node:
		return ts.LookupByNode(ta)
	case string:
		return ts.LookupByString(ta)
	default:
		return nil, fmt.Errorf("unknown key type: %T", arg)
	}
}
func Bytes__type__serialize(value interface{}) interface{} {
	switch value := value.(type) {
	case ipld.Node:

		b, err := value.AsBytes()
		if err != nil {
			return err
		}
		return b

	default:
		return nil
	}
}
func Bytes__type__parse(value interface{}) interface{} {
	builder := schema.Type.Bytes__Repr.NewBuilder()
	switch v2 := value.(type) {
	case string:
		builder.AssignString(v2)
	case *string:
		builder.AssignString(*v2)
	default:
		return nil
	}
	return builder.Build()
}
func Bytes__type__parseLiteral(valueAST ast.Value) interface{} {
	builder := schema.Type.Bytes__Repr.NewBuilder()
	switch valueAST := valueAST.(type) {
	case *ast.StringValue:
		builder.AssignString(valueAST.Value)
	default:
		return nil
	}
	return builder.Build()
}

var Bytes__type = graphql.NewScalar(graphql.ScalarConfig{
	Name:         "Bytes",
	Description:  "Bytes",
	Serialize:    Bytes__type__serialize,
	ParseValue:   Bytes__type__parse,
	ParseLiteral: Bytes__type__parseLiteral,
})
var Communites__type = graphql.NewObject(graphql.ObjectConfig{
	Name: "Communites",
	Fields: graphql.Fields{
		"At": &graphql.Field{
			Type: Community__type,
			Args: graphql.FieldConfigArgument{
				"key": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				ts, ok := p.Source.(schema.Communites)
				if !ok {
					return nil, errNotNode
				}

				arg := p.Args["key"]
				var out ipld.Node
				var err error
				switch ta := arg.(type) {
				case ipld.Node:
					out, err = ts.LookupByNode(ta)
				case int64:
					out, err = ts.LookupByIndex(ta)
				default:
					return nil, fmt.Errorf("unknown key type: %T", arg)
				}

				return out, err

			},
		},
		"All": &graphql.Field{
			Type: graphql.NewList(Community__type),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				ts, ok := p.Source.(schema.Communites)
				if !ok {
					return nil, errNotNode
				}
				it := ts.ListIterator()
				children := make([]ipld.Node, 0)
				for !it.Done() {
					_, node, err := it.Next()
					if err != nil {
						return nil, err
					}

					children = append(children, node)
				}
				return children, nil
			},
		},
		"Range": &graphql.Field{
			Type: graphql.NewList(Community__type),
			Args: graphql.FieldConfigArgument{
				"skip": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"take": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				ts, ok := p.Source.(schema.Communites)
				if !ok {
					return nil, errNotNode
				}
				it := ts.ListIterator()
				children := make([]ipld.Node, 0)

				for !it.Done() {
					_, node, err := it.Next()
					if err != nil {
						return nil, err
					}

					children = append(children, node)
				}
				return children, nil
			},
		},
		"Count": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				ts, ok := p.Source.(schema.Communites)
				if !ok {
					return nil, errNotNode
				}
				return ts.Length(), nil
			},
		},
	},
})

func Community__Name__resolve(p graphql.ResolveParams) (interface{}, error) {
	ts, ok := p.Source.(schema.Community)
	if !ok {
		return nil, fmt.Errorf(errUnexpectedType, p.Source, "schema.Community")
	}

	return ts.FieldName().AsString()

}
func Community__Image__resolve(p graphql.ResolveParams) (interface{}, error) {
	ts, ok := p.Source.(schema.Community)
	if !ok {
		return nil, fmt.Errorf(errUnexpectedType, p.Source, "schema.Community")
	}

	f := ts.FieldImage()
	if f.Exists() {

		return f.Must(), nil

	} else {
		return nil, nil
	}

}

var Community__type = graphql.NewObject(graphql.ObjectConfig{
	Name: "Community",
	Fields: graphql.Fields{
		"Name": &graphql.Field{

			Type: graphql.NewNonNull(graphql.String),

			Resolve: Community__Name__resolve,
		},
		"Image": &graphql.Field{

			Type: Image__type,

			Resolve: Community__Image__resolve,
		},
	},
})

func Image__Data__resolve(p graphql.ResolveParams) (interface{}, error) {
	ts, ok := p.Source.(schema.Image)
	if !ok {
		return nil, fmt.Errorf(errUnexpectedType, p.Source, "schema.Image")
	}

	return ts.FieldData(), nil

}
func Image__Size__resolve(p graphql.ResolveParams) (interface{}, error) {
	ts, ok := p.Source.(schema.Image)
	if !ok {
		return nil, fmt.Errorf(errUnexpectedType, p.Source, "schema.Image")
	}

	return ts.FieldSize().AsInt()

}

var Image__type = graphql.NewObject(graphql.ObjectConfig{
	Name: "Image",
	Fields: graphql.Fields{
		"Data": &graphql.Field{

			Type: graphql.NewNonNull(Bytes__type),

			Resolve: Image__Data__resolve,
		},
		"Size": &graphql.Field{

			Type: graphql.NewNonNull(graphql.Int),

			Resolve: Image__Size__resolve,
		},
	},
})

func World__Communites__resolve(p graphql.ResolveParams) (interface{}, error) {
	ts, ok := p.Source.(schema.World)
	if !ok {
		return nil, fmt.Errorf(errUnexpectedType, p.Source, "schema.World")
	}

	return ts.FieldCommunites(), nil

}

var World__type = graphql.NewObject(graphql.ObjectConfig{
	Name: "World",
	Fields: graphql.Fields{
		"Communites": &graphql.Field{

			Type: graphql.NewNonNull(Communites__type),

			Resolve: World__Communites__resolve,
		},
	},
})

func init() {

}
