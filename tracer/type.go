package tracer

type JenkinsResponse struct {
	Class   string `json:"_class"`
	Actions []struct {
		Class  string `json:"_class,omitempty"`
		Causes []struct {
			Class            string `json:"_class"`
			Shortdescription string `json:"shortDescription"`
			Userid           string `json:"userId"`
			Username         string `json:"userName"`
		} `json:"causes,omitempty"`
	} `json:"actions"`
	Artifacts         []interface{} `json:"artifacts"`
	Building          bool          `json:"building"`
	Description       interface{}   `json:"description"`
	Displayname       string        `json:"displayName"`
	Duration          int           `json:"duration"`
	Estimatedduration int           `json:"estimatedDuration"`
	Executor          interface{}   `json:"executor"`
	Fulldisplayname   string        `json:"fullDisplayName"`
	ID                string        `json:"id"`
	Keeplog           bool          `json:"keepLog"`
	Number            int           `json:"number"`
	Queueid           int           `json:"queueId"`
	Result            string        `json:"result"`
	Timestamp         int64         `json:"timestamp"`
	URL               string        `json:"url"`
	Changesets        []interface{} `json:"changeSets"`
	Culprits          []interface{} `json:"culprits"`
	Nextbuild         interface{}   `json:"nextBuild"`
	Previousbuild     struct {
		Number int    `json:"number"`
		URL    string `json:"url"`
	} `json:"previousBuild"`
}
