package model

type Config struct {
	Provider string `yaml:"provider"`
}

type GoogleDNS struct {
	Config
	Dns []DNS `yaml:"dns"`
}

type DNS struct {
	Name     string `yaml:"name"`
	Ip       string `yaml:"ip"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
