package utilities

type ServerConfig struct {
	MyAddr   string   `yaml:"addr"`
	ExtAddrs []string `yaml:"addrs"`
}
