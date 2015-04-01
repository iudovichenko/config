package config

import (
	"flag"
)

// ReadProperties reads the specified file for properties
func ReadProperties(filename string) (*Properties, error) {
	return readPropertiesFile(filename)
}

// GetProperties uses the properties file defined by --props
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
