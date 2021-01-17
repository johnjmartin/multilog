package multilog

// Level dictates the logging level
type Level int

const (
	// ALL All levels including custom levels.
	ALL Level = iota

	// DEBUG Designates fine-grained informational events that are most useful to debug an application.
	DEBUG

	// INFO Designates informational messages that highlight the progress of the application at coarse-grained level
	INFO

	// WARN Designates potentially harmful situations.
	WARN

	// ERROR Designates error events that might still allow the application to continue running.
	ERROR

	// FATAL Designates very severe error events that will presumably lead the application to abort.
	FATAL

	// TRACE Designates finer-grained informational events than the DEBUG.
	TRACE
)
