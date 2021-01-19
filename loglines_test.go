package multilog

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTsFromLog(t *testing.T) {
	ts, err := tsFromLog("02/28/2020 5:20:56.45")
	require.NoError(t, err)
	require.Equal(t, ts.Day(), 28)
}

func TestSeverityFromLog(t *testing.T) {
	sev := severityFromLog("[warn]")
	require.Equal(t, sev, WARN)
}

func TestSeverityFromLogInvalid(t *testing.T) {
	sev := severityFromLog("[invalid]")
	require.Equal(t, sev, Level(-1))
}

func TestNewLogLineIncluded(t *testing.T) {
	ts, _ := time.Parse(time.RFC3339, "2020-02-28T05:20:15Z")
	ll := newLogLineWithFilters("[02/28/2020 5:20:57.15][info][database server] Request to create database my_db7", ts, INFO)

	require.Equal(t, ll.severity, INFO)
	require.True(t, ll.included)
}

func TestNewLogLineFiltered(t *testing.T) {
	ts, _ := time.Parse(time.RFC3339, "2020-02-28T05:20:15Z")

	// filtered on severity
	ll := newLogLineWithFilters("[02/28/2020 5:20:57.15][info] Request to create database my_db7", ts, WARN)
	require.False(t, ll.included)

	ll = newLogLineWithFilters("[02/28/1020 0:00:57.15][info] Request to create database my_db7", ts, INFO)
	require.False(t, ll.included)
}
