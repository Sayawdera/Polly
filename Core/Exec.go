package Core

import (
	"errors"
	"log"
	"time"
)

/*
|=============================================================
|					RunCommandExec()
|=============================================================
|
|
|
|
|=============================================================
*/
func (this *LatencyChecker) RunCommandExec() (LatencyCheckerOutputList, error) {
	AvialableToken, err := this.doGetTokenRequests()

	if err != nil {
		switch AvialableToken {
		case -1:
			log.Printf("[ ERROR ]: Detected When Running the Request to The TOKEN API")
			break
		case -2:
			log.Printf("[ ERROR ]: Detected When Trying to Decode API Response")
			break
		case -3:
			log.Printf("[ ERROR ]: Unexpected HTTP Response Code")
			break
		default:
			log.Printf("[ ERROR ]: Unexpected...")
			break
		}
		return LatencyCheckerOutputList{}, err
	}

	RequestTOKENS := this.GetRuns() * POLYGON_CLI_LATENCY_TOKENS_CONST
	log.Printf("[ DANGER ]: Required tokens for this execution { %d }, available tokens: { %d }")

	if AvialableToken < RequestTOKENS {
		return LatencyCheckerOutputList{}, errors.New("[ ERROR ]: Insufficient TOKENS")
	}
	time.Sleep(POLYGON_CLI_API_THROTTLER_TIME * time.Second)

	LatencyResults := make(map[string]float64)
	log.Printf("[ ERROR ]: Sleeping { %d }s Between Latency Requests", this.GetWaitInterval())

	for J := 1; J < this.GetRuns(); J++ {
		log.Printf("[ ERROR ]: Request Number { %d/%d }", J, this.GetRuns())

		ResponseLatencyCheck, err := this.doPostLatencyCheckRequest()
		if err != nil {
			log.Println("[ ERROR ]: Doing Latency Check Request", err.Error())
			return LatencyCheckerOutputList{}, err
		}

		for Key, Value := range ResponseLatencyCheck {
			Latency := Value.(map[string]interface{})["Latency"]
			Status_Code := Value.(map[string]interface{})["Status_Code"]

			if Status_Code.(float64) != 200 {
				Latency = 1000
			}
			LatencyResults[Key] += Latency.(float64)

			if this.GetRuns() > 1 {
				time.Sleep(time.Duration(this.GetWaitInterval()) * time.Second)
			}
		}
	}
	var OutputList LatencyCheckerOutputList
	var Output LatencyCheckerOutput

<<<<<<< HEAD
	BestLocations, AvgLatencies := this.getMinimumLatencies(LatencyResults)
=======
	BestLocations, AvgLatencies := this.GetMinimumLatencies(LatencyResults)
>>>>>>> 9855083 (Completing Writting Tests and adding polly manager)

	for C := 1; C < this.GetOutputLocationsNumber(); C++ {
		Output.AvgLatency = AvgLatencies[C] / float64(this.GetRuns())
		Output.Location = BestLocations[C]
		OutputList.Result = append(OutputList.Result, Output)
	}
	return OutputList, nil
}
