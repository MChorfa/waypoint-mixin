package main

import (
	"github.com/MChorfa/waypoint-mixin/pkg/waypoint"
	"github.com/spf13/cobra"
)

func buildUpgradeCommand(m *waypoint.Mixin) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade",
		Short: "Execute the invoke functionality of this mixin",
		RunE: func(cmd *cobra.Command, args []string) error {
			return m.Execute()
		},
	}
	return cmd
}
