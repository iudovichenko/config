package config

import (
	"testing"

//	. "github.com/shoenig/assert"
)

func Test_parseProperties_envvar(t *testing.T) {
	text := "term=$TERM"
	p, e := parseProperties(text)
	Tassert(e == nil, t, "envvar props error")
	term := p.GetString("term")
	Tassert(term == "xterm", t, "envvar term props empty")
}

func Test_parseProperties(t *testing.T) {
	text := "a=1\nb= beta\nc =3\nd = four\n"
	_, e := parseProperties(text)
	Tassert(e == nil, t, "ok props errored: %v", e)

	text = "a=\nb= beta\nc =3\nd = four\n"
	_, e = parseProperties(text)
	Tassert(e != nil, t, "bad props did not error")
}

func Test_GetString(t *testing.T) {
	text := "a=1\nb= beta\nc =3\nd = four\n"
	p, _ := parseProperties(text)
	v := p.GetString("b")
	Tassert(v == "beta", t, "props.b wrong value")
}

func Test_GetStringOr(t *testing.T) {
	text := "a=1\nb= beta\nc =3\nd = four\n"
	p, _ := parseProperties(text)
	v := p.GetStringOr("z", "zeta")
	Tassert(v == "zeta", t, "props.z wrong default value")
}

func Test_GetInt(t *testing.T) {
	text := "a=1\nb= beta\nc =3\nd = four\n"
	p, _ := parseProperties(text)
	v := p.GetInt("c")
	Tassert(v == 3, t, "props.c wrong value")
}

func Test_GetIntOr_absent(t *testing.T) {
	text := "a=1\nb= beta\nc =3\nd = four\n"
	p, _ := parseProperties(text)
	v := p.GetIntOr("z", 6)
	Tassert(v == 6, t, "props.c absent wrong value")
}

func Test_GetIntOr_nonint(t *testing.T) {
	text := "a=1\nb= beta\nc = gamma\nd = four\n"
	p, _ := parseProperties(text)
	v := p.GetIntOr("c", 7)
	Tassert(v == 7, t, "props.c nonint wrong value")
}
