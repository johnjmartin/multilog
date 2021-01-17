package multilog

import "time"

// MultiLog stores all log files that can be queried
type MultiLog struct {
	Logs map[string]logFile
}

// NewMultiLog creates an instance of the multilog api
func NewMultiLog(serverMap map[string]string) (*MultiLog, error) {
	// todo: convert servermap to logFile structs
	return &MultiLog{}, nil
}

// Query queries the MultiLog log files
func (m *MultiLog) Query(start time.Time, entries int, keys []string, minSeverity Level) (string, error) {
	return "", nil
}
