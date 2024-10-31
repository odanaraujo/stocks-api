package config

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"gopkg.in/yaml.v2"
)

const localPath = "productapi/configs/%s.yml"
const dockerPath = "configs/%s.yml"

type Config struct {
	Env   string `yaml:url`
	MySQL struct {
		Url      string `yaml:url`
		DB       string `yaml:db`
		User     string `yaml:user`
		Password string `password`
	} `yaml:mysql`
}

func (c *Config) Print() {
	fmt.Printf("env=%s, mysql.url=%s, mysql.db=%s, mysql.user=%s, mysql.password=%s\n", c.Env, c.MySQL.Url, c.MySQL.DB, c.MySQL.User, c.MySQL.Password)
}

func Load(args []string) (*Config, error) {
	env, err := GetEnv(args)
	if err != nil {
		return nil, err
	}

	path := localPath
	if env == "docker" {
		path = dockerPath
	}

	f, err := os.Open(fmt.Sprintf(path, env))

	if err != nil {
		return nil, err
	}

	defer f.Close()

	cfg := &Config{}
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}

	cfg.Env = env
	cfg.Print()

	return cfg, nil
}

func GetEnv(args []string) (string, error) {
	envRegex, err := regexp.Compile(`env=\w+`)
	if err != nil {
		return "", err
	}

	for _, arg := range args {
		env := envRegex.FindString(arg)
		if env != "" {
			return strings.Replace(env, "env=", "", -1), nil
		}
	}

	return "local", nil
}
