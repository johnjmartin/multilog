package multilog

import (
	"regexp"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

type logLine struct {
	log       string
	severity  Level
	timestamp time.Time
	included  bool
}

// LogLines is a slice of logLine
type LogLines []logLine

func newLogLineWithFilters(line string, start time.Time, minSeverity Level) logLine {
	// all improperly formatted logs will be filtered
	re := regexp.MustCompile(`\[([^\[\]]*)\]`)
	included := true

	submatchall := re.FindAllString(line, -1)
	if len(submatchall) < 2 {
		log.Warnf("Unsupported log format, log: %s", line)
		return logLine{included: false}
	}
	sev := severityFromLog(submatchall[1])
	if minSeverity > sev {
		included = false
	}
	ts, err := tsFromLog(submatchall[0])
	if err != nil {
		included = false
	}
	if ts.Before(start) {
		included = false
	}

	return logLine{log: line, severity: severityFromLog(submatchall[1]), timestamp: ts, included: included}
}

func (l LogLines) toStringSlice() []string {
	s := []string{}
	for _, line := range l {
		s = append(s, line.log)
	}
	return s
}

func (l LogLines) toOutput(entries int) string {
	linesSlice := l.toStringSlice()
	if len(linesSlice) < entries {
		entries = len(linesSlice)
	}

	return strings.Join(l.toStringSlice()[:entries], "\n")
}

func (l LogLines) Len() int { return len(l) }
func (l LogLines) Less(i, j int) bool {
	logi := l[i]
	logj := l[j]

	return logi.timestamp.Before(logj.timestamp)
}
func (l LogLines) Swap(i, j int) { l[i], l[j] = l[j], l[i] }

func tsFromLog(tsBlock string) (time.Time, error) {
	tsBlock = strings.Trim(tsBlock, "[")
	tsBlock = strings.Trim(tsBlock, "]")
	// constant for representing our custom timestamp, see https://golang.org/pkg/time/#Time.Format for more details
	t, err := time.Parse("01/02/2006 15:04:05.00", tsBlock)
	if err != nil {
		log.Warnf("invalid timestamp in log: %s", tsBlock)
		return time.Time{}, err
	}
	return t, nil
}

func severityFromLog(severityBlock string) Level {
	var severity Level
	severityBlock = strings.Trim(severityBlock, "[")
	severityBlock = strings.Trim(severityBlock, "]")
	severityBlock = strings.ToUpper(severityBlock)

	switch severityBlock {
	case "ALL":
		severity = ALL
	case "INFO":
		severity = INFO
	case "DEBUG":
		severity = DEBUG
	case "WARN":
		severity = WARN
	case "ERROR":
		severity = ERROR
	case "FATAL":
		severity = FATAL
	case "TRACE":
		severity = TRACE
	default:
		log.Warnf("unsupported severity in log: %s", severityBlock)
		// ensures this log is not included
		severity = -1
	}

	return severity
}
