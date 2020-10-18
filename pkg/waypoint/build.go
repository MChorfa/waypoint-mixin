package waypoint

import (
	"fmt"

	"get.porter.sh/porter/pkg/exec/builder"
	yaml "gopkg.in/yaml.v2"
)

// BuildInput represents stdin passed to the mixin for the build command.
type BuildInput struct {
	Config MixinConfig
}

// MixinConfig represents configuration that can be set on the waypoint mixin in porter.yaml
// mixins:
// - waypoint:
//	  clientVersion: "0.1.2"

type MixinConfig struct {
	ClientVersion string `yaml:"clientVersion,omitempty"`
}

const dockerfileLines = `ENV WAYPOINT_VERSION=%s
RUN apt-get update && apt-get install -y wget unzip && \
wget https://releases.hashicorp.com/waypoint/${WAYPOINT_VERSION}/waypoint_${WAYPOINT_VERSION}_linux_amd64.zip && \
unzip waypoint_${WAYPOINT_VERSION}_linux_amd64.zip -d /usr/local/bin
`

// Build will generate the necessary Dockerfile lines
// for an invocation image using this mixin
func (m *Mixin) Build() error {

	// Create new Builder.
	var input BuildInput

	err := builder.LoadAction(m.Context, "", func(contents []byte) (interface{}, error) {
		err := yaml.Unmarshal(contents, &input)
		return &input, err
	})
	if err != nil {
		return err
	}

	suppliedClientVersion := input.Config.ClientVersion

	if suppliedClientVersion != "" {
		m.ClientVersion = suppliedClientVersion
	}

	// Install waypoint
	fmt.Fprintf(m.Out, dockerfileLines, m.ClientVersion)
	return nil
}
