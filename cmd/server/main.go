package main

import (
	"github.com/storyloc/server/pkg/configuration"
	"github.com/storyloc/server/pkg/graphql"
	"github.com/storyloc/server/pkg/server"
	"github.com/storyloc/server/pkg/service"
	"github.com/storyloc/server/pkg/storage"
	diskStorage "github.com/storyloc/server/pkg/storage/disk"
	ipfsStorage "github.com/storyloc/server/pkg/storage/ipfs"
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
	configuration := *config.New()

	serverCli(configuration, app)

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func serverCli(configuration config.Configuration, app *cli.App) {
	app.Commands = append(app.Commands, &cli.Command{
		Name:  "server",
		Usage: "start storyLock server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				// toDo: pic random port later
				Name:        "server-port",
				EnvVars:     []string{"SL_SERVER_PORT"},
				Usage:       "server port",
				DefaultText: configuration.Server.Port,
			},
			&cli.BoolFlag{
				Name:        "server-graphiql",
				EnvVars:     []string{"SL_SERVER_GRAPHIQL"},
				Usage:       "enable server graphiql",
				DefaultText: strconv.FormatBool(configuration.Server.GraphiQl),
			},
			&cli.StringFlag{
				Name:        "storage",
				EnvVars:     []string{"SL_STORAGE"},
				Usage:       "storage type",
				DefaultText: configuration.Storage.Type,
			},
			&cli.StringFlag{
				Name:        "storage-ipfs-url",
				EnvVars:     []string{"SL_STORAGE_IPFS_URL"},
				Usage:       "ipfs storage url",
				DefaultText: configuration.Storage.Ipfs.Url,
			},
		},
		Action: func(c *cli.Context) error {
			configuration.Apply(
				config.ServerPort(c.String("server-port")),
				config.ServerGraphiQl(c.Bool("server-graphiql")),
				config.StorageType(c.String("storage")),
				config.StorageIpfsUrl(c.String("storage-ipfs-url")),
			)

			var profileRepository storage.ProfileRepository
			var storyRepository storage.StoryRepository

			switch configuration.Storage.Type {
			case "ipfs":
				storyRepository = ipfsStorage.NewStoryRepository(configuration)
				profileRepository = ipfsStorage.NewProfileRepository(configuration)
			default:
				storyRepository = diskStorage.NewStoryRepository()
				profileRepository = diskStorage.NewProfileRepository()
			}

			storyService := service.NewStoryService(storyRepository)
			profileService := service.NewProfileService(profileRepository)

			graphqlSchema := graphql.NewSchema(profileService, storyService)
			graphqlServer, err := server.NewGraphqlServer(configuration, graphqlSchema)
			if err != nil {
				return err
			}

			return server.Start(configuration, graphqlServer)
		},
	})
}
