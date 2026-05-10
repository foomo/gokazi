// Package cmd builds the gokazi cobra command tree.
//
// [New] returns the fully-wired root command with every subcommand
// attached. It is consumed by two binaries: cmd/gokazi, which executes
// the tree, and cmd/docgen, which walks the same tree to regenerate
// the Markdown CLI reference under docs/.
//
// Each subcommand has its own constructor ([NewList], [NewStop],
// [NewConfig], [NewVersion]). The root command ([NewRoot]) holds the
// persistent --debug flag, which toggles pterm debug output.
package cmd
