package config

type CClos func(c *Configuration)

type Configuration struct {
	Server  Server
	Storage Storage
}

func New(cs ...CClos) *Configuration {
	configuration := &Configuration{
		Server: Server{
			Port:     "3000",
			GraphiQl: false,
		},
		Storage: Storage{
			Type: "disk",
			Ipfs: StorageIpfs{
				Url: "localhost:5001",
			},
		},
	}

	configuration.Apply(cs...)

	return configuration
}

func (c *Configuration) Apply(css ...CClos) *Configuration {
	for _, cs := range css {
		cs(c)
	}

	return c
}
