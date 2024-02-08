package Core

type LatencyChecker struct {
	TargetURL             string
	Runs                  int
	WaitInterval          int
	Locations             []string
	APIKey                string
	ContentType           string
	OutputLocationsNumber int
	ServiceTokenAPIURL    string
	ServiceAPIURL         string
}

type LatencyCheckerOutput struct {
	Location   string  `json:"location", yaml:"location"`
	AvgLatency float64 `json:"avgLatency", yaml:"avgLatency"`
}

type LatencyCheckerOutputList struct {
	Result []LatencyCheckerOutput `json:"result", yaml:"result"`
}

type TokenAPIResponse struct {
	RequestCount int `json:"request_count", yaml:"request_count"`
	Duration     int `json:"duration", yaml:"duration"`
}

type LatencyAPIRequest struct {
	TargetURL string   `json:"target", yaml:"target"`
	Locations []string `json:"locations", yaml:"locations"`
}

const (
	POLYGON_CLI_CONTENT_TYPE_REQUEST = "application/json"
	POLYGON_CLI_TOKEN_API_URL        = "HTTPS"
	POLYGON_CLI_LATENCY_API_URL      = "HTTPS"
	POLYGON_CLI_LATENCY_TOKENS_CONST = 5000
	POLYGON_CLI_API_THROTTLER_TIME   = 1
)
