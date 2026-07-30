package shared

type Finding struct {
	Rule          string
	Severity      string
	Category      string
	Resource      string
	Namespace     string
	Message       string
	RelatedEvents []string
	Diagnosis     string
}
