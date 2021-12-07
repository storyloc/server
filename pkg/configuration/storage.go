package config

type Storage struct {
	Type string
	Ipfs StorageIpfs
}

func StorageType(val string) CClos {
	return func(c *Configuration) {
		if val != "" {
			c.Storage.Type = val
		}
	}
}
