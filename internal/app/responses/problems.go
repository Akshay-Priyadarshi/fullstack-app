package responses

const (
	ProblemTypeValidation = "Validation"
)

type Problem struct {
	Type    string         `json:"type"`
	Details map[string]any `json:"details"`
}
