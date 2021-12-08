package config

type StorageIpfs struct {
	Url string
}

func StorageIpfsUrl(val string) CClos {
	return func(c *Configuration) {
		if val != "" {
			c.Storage.Ipfs.Url = val
		}
	}
}
