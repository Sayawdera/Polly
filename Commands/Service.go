package Commands

import (
	"PolygonCLI/Core"
	"encoding/json"
	"fmt"
	color "github.com/TwiN/go-color"
	"github.com/jedib0t/go-pretty/v6/table"
	"gopkg.in/yaml.v3"
	"os"
)

/*
|=============================================================
|					WriteOutputTable()
|=============================================================
|
|
|
|
|=============================================================
*/
func WriteOutputTable(lcol Core.LatencyCheckerOutputList) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{color.InWhite("Location"), color.InWhite("Average Latency")})

	if len(lcol.Result) > 0 {
		for I := range lcol.Result {
			if I == 0 {
				t.AppendRow([]interface{}{color.InGreen(lcol.Result[I].Location), color.InGreen(fmt.Sprintf("[ ERROR: { %d }]", lcol.Result[I].AvgLatency))})
				continue
			}
			t.AppendRow([]interface{}{color.InWhite(lcol.Result[I].Location), color.InWhite(fmt.Sprintf("[ ERROR ]: { %f }", lcol.Result[I].AvgLatency))})
		}
	}
	t.SetStyle(table.StyleLight)
	t.Render()
}

/*
|=============================================================
|					WriteOutputJson()
|=============================================================
|
|
|
|
|=============================================================
*/
func WriteOutputJson(lcol Core.LatencyCheckerOutputList) {
	JSON, _ := json.MarshalIndent(lcol, "", "        ")
	fmt.Println(string(JSON))
}

/*
|=============================================================
|					WriteOutputYaml()
|=============================================================
|
|
|
|
|=============================================================
*/
func WriteOutputYaml(lcol Core.LatencyCheckerOutputList) {
	YAML, _ := yaml.Marshal(lcol)
	fmt.Println(string(YAML))
}
