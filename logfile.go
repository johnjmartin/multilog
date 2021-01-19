package multilog

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// LogFile is a helper struct for reading log files provided a filePath
type LogFile struct {
	filePath string
}

// NewLogFile creates a logfile struct for reading logs out of a file
func NewLogFile(filePath string) (LogFile, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return LogFile{}, fmt.Errorf("log file at path %s could not be found", filePath)
	}
	return LogFile{filePath: filePath}, nil
}

func (l LogFile) read(start time.Time, entries int, minSeverity Level) (LogLines, error) {
	ll := LogLines{}

	file, err := os.Open(l.filePath)
	if err != nil {
		return LogLines{}, err
	}
	defer file.Close()

	// avoids reading the whole file into memory, will read the whole file line-by-line in the worst case
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineStr := scanner.Text()
		line := newLogLineWithFilters(lineStr, start, minSeverity)
		if line.included {
			ll = append(ll, line)
		}
		// We've got more than enough lines for our entries limit, return the collected lines
		if len(ll) >= entries {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return LogLines{}, err
	}
	return ll, nil
}
