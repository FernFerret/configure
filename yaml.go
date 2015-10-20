package configure

import (
	"gopkg.in/yaml.v2"

	"github.com/gravitational/trace"
)

// ParseYAML parses yaml-encoded byte string into the struct
// passed to the function.
func ParseYAML(data []byte, cfg interface{}) error {
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return trace.Wrap(err)
	}
	return nil
}
