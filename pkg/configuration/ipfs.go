package config

type Ipfs struct {
	Url string
}

func IpfsUrl(val string) CClos {
	return func(c *Configuration) {
		if val != "" {
			c.Ipfs.Url = val
		}

	}
}
