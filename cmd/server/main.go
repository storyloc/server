package main

//go:generate go run ../../pkg/schema/gen.go ../../pkg/schema

import (
	"github.com/storyloc/server/pkg/configuration"
	"github.com/storyloc/server/pkg/server"
	"github.com/storyloc/server/pkg/server/gql"
	"github.com/storyloc/server/pkg/service"
	"github.com/storyloc/server/pkg/storage/file"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"strconv"
)

func main() {
	app := &cli.App{
		Name:  "storyLock",
		Usage: "cli",
	}
	configuration := config.New()

	serverCli(configuration, app)

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func serverCli(configuration *config.Configuration, app *cli.App) {
	app.Commands = append(app.Commands, &cli.Command{
		Name:  "server",
		Usage: "start storyLock server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				// toDo: pic random port later
				Name:        "port",
				EnvVars:     []string{"SL_PORT"},
				Usage:       "server port",
				DefaultText: configuration.Server.Port,
			},
			&cli.BoolFlag{
				Name:        "graphiql",
				EnvVars:     []string{"SL_GRAPHIQL"},
				Usage:       "enable graphiql",
				DefaultText: strconv.FormatBool(configuration.Server.GraphiQl),
			},
			&cli.StringFlag{
				Name:        "ipfs-url",
				EnvVars:     []string{"SL_IPFS_URL"},
				Usage:       "ipfs url",
				DefaultText: configuration.Ipfs.Url,
			},
		},
		Action: func(c *cli.Context) error {
			configuration.Apply(
				config.ServerPort(c.String("port")),
				config.ServerGraphiQl(c.Bool("graphiql")),
				config.IpfsUrl(c.String("ipfs-url")),
			)
			storyRepository := file.NewStoryRepository()
			storyService := service.NewStoryService(storyRepository)
			gqlServer, err := gql.NewServer(configuration, storyService)
			if err != nil {
				return err
			}

			servers := []server.Server{gqlServer}

			return server.Start(configuration, servers)
		},
	})
}
