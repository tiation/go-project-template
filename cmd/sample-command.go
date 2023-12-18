/*
Copyright © 2023 Dataflows
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/thedataflows/go-commons/pkg/config"
	"github.com/thedataflows/go-commons/pkg/defaults"
	"github.com/thedataflows/go-commons/pkg/log"
)

type SampleCommand struct {
	cmd    *cobra.Command
	parent *Root
}

var (
	_ = NewSampleCommand(root)
)

func init() {

}

func NewSampleCommand(parent *Root) *SampleCommand {
	c := &SampleCommand{
		parent: parent,
	}

	c.cmd = &cobra.Command{
		Use:           "sample-command",
		Short:         "This is a sample command",
		Long:          ``,
		Aliases:       []string{"c"},
		RunE:          c.RunCommand,
		SilenceErrors: parent.Cmd().SilenceErrors,
		SilenceUsage:  parent.Cmd().SilenceUsage,
	}
	parent.Cmd().AddCommand(c.cmd)

	c.cmd.PersistentFlags().StringP(
		c.KeyStringFlag1(),
		"c",
		c.DefaultStringFlag1(),
		"[Required] A flag called Flag1",
	)

	c.cmd.PersistentFlags().StringP(
		c.KeySomePath(),
		"p",
		c.DefaultSomePath(),
		"Some path flag",
	)

	defaultTimeout, _ := time.ParseDuration("10m0s")
	c.cmd.PersistentFlags().DurationP(
		c.KeyTimeout(),
		"t",
		defaultTimeout,
		"Some timeout flag",
	)

	// Bind flags to config
	config.ViperBindPFlagSet(c.cmd, c.cmd.PersistentFlags())

	return c
}

func (c *SampleCommand) RunCommand(cmd *cobra.Command, args []string) error {
	if err := c.CheckRequiredFlags(); err != nil {
		return err
	}

	log.Warnf("Running %s with #%d arguments", cmd.Name(), len(args))
	log.Warnf("--%s=%s", c.KeyStringFlag1(), c.StringFlag1())
	log.Warnf("--%s=%s", c.KeySomePath(), c.SomePath())
	log.Warnf("--%s=%s", c.KeyTimeout(), c.Timeout())
	log.Warnf("Stdin: %s", stdInBytes)

	// Some more logic

	return nil
}

func (c *SampleCommand) Cmd() *cobra.Command {
	return c.cmd
}

func (c *SampleCommand) CheckRequiredFlags() error {
	return config.CheckRequiredFlags(c.cmd, []string{c.KeyStringFlag1()})
}

// Flags keys, defaults and value getters
func (c *SampleCommand) KeyStringFlag1() string {
	const f = "flag1"
	return f
}

func (c *SampleCommand) DefaultStringFlag1() string {
	return defaults.Undefined
}

func (c *SampleCommand) StringFlag1() string {
	return config.ViperGetString(c.cmd, c.KeyStringFlag1())
}

func (c *SampleCommand) KeySomePath() string {
	return "some-path"
}

func (c *SampleCommand) DefaultSomePath() string {
	return fmt.Sprintf("./samplecommand-%s", c.DefaultStringFlag1())
}

func (c *SampleCommand) SomePath() string {
	samplecommandSomePath := config.ViperGetString(c.cmd, c.KeySomePath())
	if samplecommandSomePath == c.DefaultSomePath() {
		samplecommandSomePath = fmt.Sprintf(
			"%s/samplecommand-%s",
			c.parent.ProjectRoot(),
			c.StringFlag1(),
		)
	}
	return samplecommandSomePath
}

func (c *SampleCommand) KeyTimeout() string {
	const t = "timeout"
	return t
}

func (c *SampleCommand) Timeout() string {
	return config.ViperGetDuration(c.cmd, c.KeyTimeout()).String()
}
