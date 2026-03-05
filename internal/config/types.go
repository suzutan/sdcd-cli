package config

type Config struct {
	CurrentContext string      `yaml:"current-context"`
	Contexts       []Context   `yaml:"contexts"`
	Preferences    Preferences `yaml:"preferences"`
}

type Context struct {
	Name   string `yaml:"name"`
	APIURL string `yaml:"api-url"`
	Token  string `yaml:"token"`
}

type Preferences struct {
	Output   string `yaml:"output"`
	NoColor  bool   `yaml:"no-color"`
	PageSize int    `yaml:"page-size"`
}
