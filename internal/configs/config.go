package configs

import (
	"io/ioutil"
	"path"

	yaml "gopkg.in/yaml.v2"
)

const (
	// BuildVersion application
	BuildVersion = "BUILD_VERSION"
)

// Option application
type Option struct {
	Name string `yaml:"name"`
	HTTP struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"http"`
	Database struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
	Github struct {
		ClientID     string `yaml:"client_id"`
		ClientSecret string `yaml:"client_secret"`
	} `yaml:"github"`
	System struct {
		Attachments struct {
			Storage string `yaml:"storage"`
			Path    string `yaml:"path"`
		} `yaml:"attachments"`
	} `yaml:"system"`
	Operators []string `yaml:"operators"`

	Environment string
	OperatorSet map[string]bool
}

var appConfig *Option

// Init application
func Init(dir, env string) error {
	data, err := ioutil.ReadFile(path.Join(dir, "./config.yaml"))
	if err != nil {
		return err
	}

	var options map[string]Option
	err = yaml.Unmarshal(data, &options)
	if err != nil {
		return err
	}
	opt := options[env]
	opt.Environment = env
	opt.OperatorSet = make(map[string]bool)
	for _, operator := range opt.Operators {
		opt.OperatorSet[operator] = true
	}
	appConfig = &opt
	return nil
}
