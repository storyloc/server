package main

//go:generate go run ../../pkg/schema/gen/gen.go ../../pkg/schema

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/storyloc/server/pkg/graphql"
	"github.com/urfave/cli/v2"
	"log"
	"net/http"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "storyLock",
		Usage: "cli",
		Commands: []*cli.Command{
			server(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func server() *cli.Command {
	return &cli.Command{
		Name:  "server",
		Usage: "start storyLock server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "port",
				Value: "3000",
			},
		},
		Action: func(c *cli.Context) error {
			r := chi.NewRouter()
			r.Use(middleware.Logger)

			graphql.Handle(r)

			return http.ListenAndServe(fmt.Sprintf(":%s", c.String("port")), r)
		},
	}
}
