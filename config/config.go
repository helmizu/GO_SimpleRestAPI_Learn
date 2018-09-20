package config

type Config struct {
	Server   string
	Database string
}

func (c *Config) Init() {
	c.Server = "localhost"
	c.Database = "GoLearn"
}

func (c Config) Read() Config {
	return c
}
