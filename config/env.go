package config

import "log"

func Env() (Config, error) {
	env, err := LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	return env, nil
}
