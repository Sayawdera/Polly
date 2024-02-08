package Commands

import (
	"PolygonCLI/Core"
	"errors"
	"github.com/spf13/cobra"
)

/*
|=============================================================
|					AddExecFlags()
|=============================================================
|
|
|
|
|=============================================================
*/
func AddExecFlags(cmd *cobra.Command) {
	Flags := cmd.Flags()
	Flags.StringVarP(&targetURL, "target-url", "t", "", "The target url. e.g: https://google.com")
	Flags.IntVarP(&numberOfRuns, "runs", "r", 1, "The number of executions.")
	Flags.StringVarP(&waitInterval, "interval", "i", "1m", "The amount of waiting time between runs.")
	Flags.StringArrayVarP(&locations, "locations", "l", []string{"EU.ES.*"}, "The array of locations to be requested. e.g: NA.US.*,NA.EU.*")
	Flags.IntVar(&outputLocationsNumber, "output-locations", 1, "The number of best locations to output.")
	Flags.StringVarP(&outputFormat, "output-format", "o", "table", "Output in an specific format. Usage: '-o [ table | yaml | json ]'")
	cmd.MarkFlagRequired("target-url")
}

/*
|=============================================================
|					NewExecCommand()
|=============================================================
|
|
|
|
|=============================================================
*/
func NewExecCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Exec the run command",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Validate command Args
			err := ValidateCommandArgs()
			if err != nil {
				return err
			}
			// Get waitIntervalInSeconds
			waitIntervalSeconds := Core.IntervalTimeToSeconds(waitInterval)
			LatencyChecker := Core.NewLatencyChecker(Core.GetEnvironment("POLYGON_X_API_KEY", "NOT_SET"), targetURL, numberOfRuns, waitIntervalSeconds, locations, outputLocationsNumber)
			res, err := LatencyChecker.RunCommandExec()
			switch {
			case outputFormat == "yaml":
				WriteOutputYaml(res)
			case outputFormat == "json":
				WriteOutputJson(res)
			default:
				WriteOutputTable(res)
			}
			return err
		},
	}
	AddExecFlags(cmd)
	return cmd
}

/*
|=============================================================
|					ValidateCommandArgs()
|=============================================================
|
|
|
|
|=============================================================
*/
func ValidateCommandArgs() error {
	validInterval := Core.ValidateIntervalTime(waitInterval)

	if !validInterval {
		return errors.New(" not valid interval")
	}
	validURL := Core.ValidateURL(targetURL)

	if !validURL {
		return errors.New(" not valid url")
	}
	return nil
}
