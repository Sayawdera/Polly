package main

import (
	"PolygonCLI/Commands"
	"github.com/TwiN/go-color"
	"github.com/spf13/cobra"
	"log"
	"os"
)

/*
|=============================================================
|					main()
|=============================================================
|
|
|
|
|=============================================================
*/
func main() {
	command := NewCommand()
	if err := command.Execute(); err != nil {
		log.Fatalf(color.InRed("[ERROR]")+"%s", err.Error())
	}
}

/*
|=============================================================
|					NewCommand()
|=============================================================
|
|
|
|
|=============================================================
*/
func NewCommand() *cobra.Command {
	c := &cobra.Command{
		Use:   "POLYGON CLI LATENCY",
		Short: "POLYGON CLI is the command line interface to work with the ddosify latencies API",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(1)
		},
	}
	c.AddCommand(Commands.NewExecCommand())

	return c
}
