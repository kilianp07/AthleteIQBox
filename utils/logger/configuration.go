package utils

// LoggerConfiguration holds the logger settings.
type LoggerConfiguration struct {
	OutputsConfig []map[string]interface{} `json:"outputs"` // Raw JSON maps for outputs
	Timestamp     bool                     `json:"timestamp"`
	Caller        bool                     `json:"caller"`
}
