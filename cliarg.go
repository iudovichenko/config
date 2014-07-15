package config

import (
	"flag"
)

func GetProperties(required bool) (*Properties, error) {
	filename := flag.String("props", "", "properties filename")
	flag.Parse()
	if *filename == "" {
		if required {
			return nil, NO_PROPS_FILE
		} else {
			return &Properties{make(map[string]string)}, nil
		}
	}
	return readPropertiesFile(*filename)
}
