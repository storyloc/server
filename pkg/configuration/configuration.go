package config

type CClos func(c *Configuration)

type Configuration struct {
	Server Server
	Ipfs   Ipfs
}

func New(cs ...CClos) *Configuration {
	configuration := &Configuration{
		Server: Server{
			Port:     "3000",
			GraphiQl: false,
		},
		Ipfs: Ipfs{
			Url: "localhost:5001",
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
