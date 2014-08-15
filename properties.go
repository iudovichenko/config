package config

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Properties struct {
	props map[string]string
}

func readPropertiesFile(filename string) (*Properties, error) {
	if b, e := ioutil.ReadFile(filename); e != nil {
		return nil, e
	} else {
		return parseProperties(string(b))
	}
}

func parseProperties(content string) (*Properties, error) {
	p := &Properties{
		make(map[string]string),
	}
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		tokens := strings.SplitN(line, "=", 2)
		if len(tokens) != 2 {
			return nil, BAD_PROPERTY
		}
		key := strings.TrimSpace(tokens[0])
		val := strings.TrimSpace(tokens[1])
		if len(key) == 0 || len(val) == 0 {
			return nil, BAD_PROPERTY
		}

		// check if we should substitute an environment variable
		if val[0] == '$' {
			if len(val) <= 1 {
				return nil, BAD_PROPERTY
			}
			evar := val[1:]
			val = os.Getenv(evar)
		}

		p.props[key] = val
	}
	return p, nil
}

func (p *Properties) GetString(key string) string {
	return p.props[key]
}

func (p *Properties) GetStringOr(key, alt string) string {
	if v := p.props[key]; v == "" {
		return alt
	} else {
		return v
	}
}

func (p *Properties) GetInt(key string) int {
	i, _ := strconv.Atoi(p.props[key])
	return i
}

func (p *Properties) GetIntOr(key string, alt int) int {
	if i, e := strconv.Atoi(p.props[key]); e != nil {
		return alt
	} else {
		return i
	}
}
