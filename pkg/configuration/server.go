package config

type Server struct {
	Port     string
	GraphiQl bool
}

func ServerPort(val string) CClos {
	return func(c *Configuration) {
		if val != "" {
			c.Server.Port = val
		}
	}
}

func ServerGraphiQl(val bool) CClos {
	return func(c *Configuration) {
		c.Server.GraphiQl = val
	}
}
