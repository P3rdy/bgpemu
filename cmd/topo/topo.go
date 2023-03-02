package topo

import (
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Create topo on cluster",
		RunE:  createFn,
	}
	topoCmd := &cobra.Command{
		Use:   "topo",
		Short: "Topology commands.",
	}
	topoCmd.AddCommand(createCmd)
	return topoCmd
}

func createFn(cmd *cobra.Command, args []string) error {
	
	return nil
}
