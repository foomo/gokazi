// Package pterm provides a [log/slog.Handler] that routes records
// through the pterm package for styled terminal output.
//
// The handler is intended for command-line tools that already use
// pterm for status messages and want a single rendering pipeline for
// both pterm calls and slog records. Debug records are emitted only
// when pterm.PrintDebugMessages is enabled; all other levels are
// always rendered.
//
// Grouping ([slog.Handler.WithGroup]) is not supported and is silently
// ignored.
package pterm
