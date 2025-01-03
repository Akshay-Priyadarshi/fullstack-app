package responses

type Problem struct {
	Type    string         `json:"type"`
	Details map[string]any `json:"details"`
}
