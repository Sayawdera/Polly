package Core

type IGetLatencyCheckerMutators interface {
	GetServiceAPITokenURL() string
	GetServiceAPIURL() string
	GetTargetURL() string
	GetRuns() int
	GetWaitInterval() int
	GetLocations() []string
	GetAPiKey() string
	GetOutputLocationsNumber() int
}

type ISetLatencyCheckerMutators interface {
	SetServiceAPITokenURL(APITokenURL string)
	SetServiceAPIURL(APIURL string)
	SetTargetURL(TargetURL string)
	SetRuns(Runs int)
	SetWaitInterval(Interval int)
	SetLocations(Location []string)
	SetAPiKey(APIKEY string)
	SetOutputLocationsNumber(OutputLocation int)
}
