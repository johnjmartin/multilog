package multilog

import (
	"fmt"
	"sort"
	"time"
)

// MultiLog stores all log files that can be queried
type MultiLog struct {
	Logs map[string]LogFile
}

// NewMultiLog creates an instance of the multilog api
func NewMultiLog(serverMap map[string]string) (*MultiLog, error) {
	logs := make(map[string]LogFile)
	for serverName, filePath := range serverMap {
		if l, err := NewLogFile(filePath); err == nil {
			logs[serverName] = l
		} else {
			return nil, fmt.Errorf("invalid log file %s", err)
		}
	}

	return &MultiLog{Logs: logs}, nil
}

// Query queries the MultiLog log files
func (m *MultiLog) Query(start time.Time, entries int, keys []string, minSeverity Level) (string, error) {
	// assuming keys correspond only to the keys in the MultiLog map, not the keys of the files log lines
	var ll LogLines = LogLines{}
	for _, key := range keys {
		if l, ok := m.Logs[key]; ok {
			// read lines from file, only returning a max of num entries, matching the severity & timestamp filters
			toAppend, err := l.read(start, entries, minSeverity)
			if err != nil {
				return "", err
			}
			ll = append(ll, toAppend...)
		} else {
			return "", fmt.Errorf("no log entry for key %s", key)
		}
	}

	// LogLines implements the sort.Interface methods for sorting support
	sort.Sort(ll)
	return ll.toOutput(entries), nil
}
