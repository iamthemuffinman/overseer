package cmd

import (
	"strings"

	"github.com/iamthemuffinman/cli"
)

type ProvisionPhysicalCommand struct {
	UI         cli.Ui
	ShutdownCh <-chan struct{}
}

func (c *ProvisionPhysicalCommand) Run(args []string) int {
	if len(args) == 0 {
		return cli.RunResultHelp
	}

	for _, arg := range args {
		if arg == "-h" || arg == "-help" || arg == "--help" {
			return cli.RunResultHelp
		}
	}

	doneCh := make(chan struct{})
	go func() {
		defer close(doneCh)
		// do stuff here
	}()

	select {
	case <-c.ShutdownCh:
		c.UI.Output("Interrupt received. Gracefully shutting down...")

		select {
		case <-c.ShutdownCh:
			c.UI.Error("Two interrupts received. Exiting immediately. Data loss may have occurred.")
			return 1
		case <-doneCh:
		}
	case <-doneCh:
	}

	return 0
}

func (c *ProvisionPhysicalCommand) Help() string {
	return c.helpProvisionPhysical()
}

func (c *ProvisionPhysicalCommand) Synopsis() string {
	return "Provision physical infrastructure"
}

func (c *ProvisionPhysicalCommand) helpProvisionPhysical() string {
	helpText := `
Usage: overseer provision physical [OPTIONS] [HOSTS]
`
	return strings.TrimSpace(helpText)
}
