package harvey

import (
	"bytes"
	"io"

	"sevki.org/build/build"
)

type Group struct {
	Name         string   `group:"name"`
	Dependencies []string `group:"deps"`
}

func (g *Group) Hash() []byte {

	return []byte{}
}

func (g *Group) Build(c *build.Context) error {
	return nil
}

func (g *Group) GetName() string {
	return g.Name
}

func (g *Group) GetDependencies() []string {
	return g.Dependencies
}
func (g *Group) GetSource() string {
	return ""
}

func (g *Group) Reader() io.Reader {
	return bytes.NewBufferString("")
}
func (g *Group) Installs() map[string]string {
	return nil
}
