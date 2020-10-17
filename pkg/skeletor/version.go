package skeletor

import (
	"get.porter.sh/mixin/skeletor/pkg"
	"get.porter.sh/porter/pkg/mixin"
	"get.porter.sh/porter/pkg/porter/version"
)

func (m *Mixin) PrintVersion(opts version.Options) error {
	metadata := mixin.Metadata{
		Name: "skeletor",
		VersionInfo: mixin.VersionInfo{
			Version: pkg.Version,
			Commit:  pkg.Commit,
			Author:  "YOURNAME",
		},
	}
	return version.PrintVersion(m.Context, opts, metadata)
}
