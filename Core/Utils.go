package Core

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/adhocore/gronx"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"sort"
	"strconv"
	"time"
)

/*
|=============================================================
|					GetEnvironment()
|=============================================================
|
|
|
|
|=============================================================
*/
func GetEnvironment(Key string, Value string) string {
	if Key, OK := os.LookupEnv(Value); OK {
		return Key
	}
	return Value
}

/*
|=============================================================
|					ValidateURL()
|=============================================================
|
|
|
|
|=============================================================
*/
func ValidateURL(InputURL string) bool {
	URL, err := url.Parse(InputURL)
	return err == nil && URL.Scheme != "" && URL.Host != ""
}

/*
|=============================================================
|					ValidateIntervalTime()
|=============================================================
|
|
|
|
|=============================================================
*/
func ValidateIntervalTime(Interval string) bool {
	R := regexp.MustCompile(`^(\d*)(s|m|h)`)
	Matched := R.MatchString(Interval)
	return Matched
}

/*
|=============================================================
|					ValidateCronTime()
|=============================================================
|
|
|
|
|=============================================================
*/
func ValidateCronTime(CronTime string) bool {
	log.Println("[ INFO ]: CronTime to be validated ", CronTime)
	gron := gronx.New()
	if CronTime != "" {
		return gron.IsValid(CronTime)
	}
	return false
}

/*
|=============================================================
|					GetNextTimeCronTime()
|=============================================================
|
|
|
|
|=============================================================
*/
func GetNextTimeCronTime(CronTime string) int64 {
	NextTime, err := gronx.NextTick(CronTime, false)
	if err != nil {
		log.Println("[ INFO ]: CronTime format is not valid or not scheduled cr passed ", err)
		return -1
	}
	Duration := NextTime.Unix() - time.Now().Unix()
	return Duration
}

/*
|=============================================================
|					IntervalTimeToSeconds()
|=============================================================
|
|
|
|
|=============================================================
*/
func IntervalTimeToSeconds(Interval string) int {
	R := regexp.MustCompile(`^(\d*)(s|m|h)`)
	CaptureGroups := R.FindStringSubmatch(Interval)

	if len(CaptureGroups) < 1 {
		return -1
	}
	TimeValue, err := strconv.Atoi(CaptureGroups[1])

	if err != nil {
		return -1
	}
	TimeUnit := CaptureGroups[2]

	switch TimeUnit {
	case "s":
		return TimeValue
	case "m":
		return TimeValue * 60
	case "h":
		return TimeValue * 3600
	default:
		return 1
	}
}

/*
|=============================================================
|					doGetTokenRequests()
|=============================================================
|
|
|
|
|=============================================================
*/
func (this *LatencyChecker) doGetTokenRequests() (int, error) {
	if this.GetAPiKey() == "NOT_SET" {
		return -1, errors.New("[ ERROR ]: POLYGON_X_API_KEY Environment Variable not Set")
	}
	Request, err := http.NewRequest(http.MethodGet, this.GetServiceAPITokenURL(), nil)

	if err != nil {
		panic(err)
	}
	Request.Header.Add("X_API_KEY", this.GetAPiKey())
	Response, _ := http.DefaultClient.Do(Request)
	BodyResponse := &TokenAPIResponse{}
	Derr := json.NewDecoder(Response.Body).Decode(BodyResponse)

	if Derr != nil {
		return -2, Derr
	}

	if Response.StatusCode != http.StatusOK {
		return -3, errors.New("[ ERROR ]: Status Code Received " + strconv.Itoa(Response.StatusCode) + " ....But Status Code Expected" + strconv.Itoa(http.StatusOK))
	}
	defer Response.Body.Close()

	return BodyResponse.RequestCount, nil
}

/*
|=============================================================
|					doPostLatencyCheckRequest()
|=============================================================
|
|
|
|
|=============================================================
*/
func (this *LatencyChecker) doPostLatencyCheckRequest() (map[string]interface{}, error) {
	var ResponseLatency map[string]interface{}

	RequestBody := LatencyAPIRequest{
		TargetURL: this.TargetURL,
		Locations: this.GetLocations(),
	}
	RequestBodyJson, _ := json.Marshal(RequestBody)
	Body := bytes.NewReader(RequestBodyJson)

	Request, err := http.NewRequest(http.MethodGet, this.GetServiceAPIURL(), Body)
	if err != nil {
		panic(err)
	}

	Request.Header.Add("Content-Type", POLYGON_CLI_CONTENT_TYPE_REQUEST)
	Request.Header.Add("X-API-KEY", this.GetAPiKey())
	Response, _ := http.DefaultClient.Do(Request)
	defer Response.Body.Close()

	Derr := json.NewDecoder(Response.Body).Decode(&ResponseLatency)
	if Derr != nil {
		log.Println(Derr.Error())
		return nil, Derr
	}

	if Response.StatusCode != http.StatusOK {
		return nil, errors.New("[ ERROR ]: Status Code Received { " + strconv.Itoa(Response.StatusCode) + " } ...But Status Code Expected { " + strconv.Itoa(http.StatusOK) + "}")
	}

	return ResponseLatency, nil
}

/*
|=============================================================
|					getMinimumLatencies()
|=============================================================
|
|
|
|
|=============================================================
*/
<<<<<<< HEAD
func (this *LatencyChecker) getMinimumLatencies(Latencies map[string]float64) ([]string, []float64) {
=======
func (this *LatencyChecker) GetMinimumLatencies(Latencies map[string]float64) ([]string, []float64) {
>>>>>>> 9855083 (Completing Writting Tests and adding polly manager)
	OutputKeys := make([]string, len(Latencies))
	OutputLatency := make([]float64, len(Latencies))
	Keys := make([]string, 0, len(Latencies))

	for Value := range Latencies {
		Keys = append(Keys, Value)
	}
	sort.SliceStable(Keys, func(KEY int, VALUE int) bool {
		return Latencies[Keys[KEY]] < Latencies[Keys[VALUE]]
	})

	if this.GetOutputLocationsNumber() > len(Latencies) {
		this.SetOutputLocationsNumber(len(Latencies))
	}
	for KV := 0; KV < this.GetOutputLocationsNumber(); KV++ {
		OutputKeys[KV] = Keys[KV]
		OutputLatency[KV] = Latencies[Keys[KV]]
	}
	return OutputKeys, OutputLatency
}
