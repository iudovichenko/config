package config

import (
	"errors"
)

var (
	NO_PROPS_FILE = errors.New("No properties file specified with -props")
	BAD_PROPERTY  = errors.New("Bad property found in properties file")
)
