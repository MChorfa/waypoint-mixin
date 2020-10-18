package waypoint

import (
	"get.porter.sh/porter/pkg/mixin"
	"get.porter.sh/porter/pkg/porter/version"
	"github.com/MChorfa/waypoint-mixin/pkg"
)

func (m *Mixin) PrintVersion(opts version.Options) error {
	metadata := mixin.Metadata{
		Name: "waypoint",
		VersionInfo: mixin.VersionInfo{
			Version: pkg.Version,
			Commit:  pkg.Commit,
			Author:  "Mohamed Chorfa",
		},
	}
	return version.PrintVersion(m.Context, opts, metadata)
}
