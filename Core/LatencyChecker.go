package Core

/*
|=============================================================
|		              DeepCopyInfo()
|=============================================================
|
|
|
|
|=============================================================
*/
func (this *LatencyCheckerOutputList) DeepCopyInfo(self *LatencyCheckerOutputList) {
	*self = *this

	if this.Result != nil {

		this, self := &this.Result, &self.Result
		*self = make([]LatencyCheckerOutput, len(*this))

		for I := range *this {
			(*self)[I] = (*this)[I]
		}
	}
}

/*
|=============================================================
|		              DeepCopy()
|=============================================================
|
|
|
|
|=============================================================
*/
func (this *LatencyCheckerOutputList) DeepCopy() *LatencyCheckerOutputList {
	if this == nil {
		return nil
	}
	self := new(LatencyCheckerOutputList)
	this.DeepCopyInfo(self)
	return self
}

/*
|=============================================================
|		              NewLatencyChecker()
|=============================================================
|
|
|
|
|=============================================================
*/
func NewLatencyChecker(APIKey string, TargetULR string, Runs int, WaitInterval int, Locations []string, OutputLocationsNumber int) *LatencyChecker {
	return &LatencyChecker{
		TargetURL:             TargetULR,
		Runs:                  Runs,
		WaitInterval:          WaitInterval,
		Locations:             Locations,
		APIKey:                APIKey,
		ContentType:           POLYGON_CLI_CONTENT_TYPE_REQUEST,
		OutputLocationsNumber: OutputLocationsNumber,
		ServiceTokenAPIURL:    POLYGON_CLI_TOKEN_API_URL,
		ServiceAPIURL:         POLYGON_CLI_TOKEN_API_URL,
	}
}

/* ====================== GET MUTATORS ====================== */

/*
|=============================================================
|		              GetServiceAPITokenURL()
|=============================================================
|
|
|
|
|=============================================================
*/
func (this *LatencyChecker) GetServiceAPITokenURL() string {
	return this.ServiceTokenAPIURL
}

/*
|=============================================================
|		              GetServiceAPIURL()
|=============================================================
|
|
|
|
|=============================================================
*/
func (this *LatencyChecker) GetServiceAPIURL() string {
	return this.ServiceAPIURL
}

/*
|=============================================================
|		              GetTargetURL()
|=============================================================
|
|
|
|
|=============================================================
*/
func (this *LatencyChecker) GetTargetURL() string {
	return this.TargetURL
}

/*
|=============================================================
|		              GetRuns()
|=============================================================
|
|
|
|
|=============================================================
*/
func (this *LatencyChecker) GetRuns() int {
	return this.Runs
}

/*
|=============================================================
|		              GetWaitInterval()
|=============================================================
|
|
|
|
|=============================================================
*/
func (this *LatencyChecker) GetWaitInterval() int {
	return this.WaitInterval
}

/*
|=============================================================
|		              GetLocations()
|=============================================================
|
|
|
|
|=============================================================
*/
func (this *LatencyChecker) GetLocations() []string {
	return this.Locations
}

/*
|=============================================================
|		              GetAPiKey()
|=============================================================
|
|
|
|
|=============================================================
*/
func (this *LatencyChecker) GetAPiKey() string {
	return this.APIKey
}

/*
|=============================================================
|		              GetOutputLocationsNumber()
|=============================================================
|
|
|
|
|=============================================================
*/
func (this *LatencyChecker) GetOutputLocationsNumber() int {
	return this.OutputLocationsNumber
}

/* ====================== SET MUTATORS ====================== */

/*
|=============================================================
|		              SetServiceAPITokenURL()
|=============================================================
|
|
|
|
|=============================================================
*/
func (self *LatencyChecker) SetServiceAPITokenURL(APITokenURL string) {
	self.ServiceTokenAPIURL = APITokenURL
}

/*
|=============================================================
|		              SetServiceAPIURL()
|=============================================================
|
|
|
|
|=============================================================
*/
func (self *LatencyChecker) SetServiceAPIURL(APIURL string) {
	self.ServiceAPIURL = APIURL
}

/*
|=============================================================
|		              SetTargetURL()
|=============================================================
|
|
|
|
|=============================================================
*/
func (self *LatencyChecker) SetTargetURL(TargetURL string) {
	self.TargetURL = TargetURL
}

/*
|=============================================================
|		              SetRuns()
|=============================================================
|
|
|
|
|=============================================================
*/
func (self *LatencyChecker) SetRuns(Runs int) {
	self.Runs = Runs
}

/*
|=============================================================
|		              SetWaitInterval()
|=============================================================
|
|
|
|
|=============================================================
*/
func (self *LatencyChecker) SetWaitInterval(Interval int) {
	self.WaitInterval = Interval
}

/*
|=============================================================
|		              SetLocations()
|=============================================================
|
|
|
|
|=============================================================
*/
func (self *LatencyChecker) SetLocations(Location []string) {
	self.Locations = Location
}

/*
|=============================================================
|		              SetAPiKey()
|=============================================================
|
|
|
|
|=============================================================
*/
func (self *LatencyChecker) SetAPiKey(APIKEY string) {
	self.APIKey = APIKEY
}

/*
|=============================================================
|		              SetOutputLocationsNumber()
|=============================================================
|
|
|
|
|=============================================================
*/
func (self *LatencyChecker) SetOutputLocationsNumber(OutputLocation int) {
	self.OutputLocationsNumber = OutputLocation
}
