package shared

type Finding struct {
	Rule      string
	Severity  string
	Category  string
	Resource  string
	Namespace string
	Message   string
	Diagnosis string
}
