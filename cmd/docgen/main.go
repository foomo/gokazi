package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/foomo/gokazi/cmd"
	cmdx "github.com/foomo/gokazi/pkg/cmd"
	"github.com/spf13/cobra/doc"
)

const outDir = "docs/reference/cli"

func main() {
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		log.Fatal(err)
	}

	indexPath := filepath.Join(outDir, "index.md")
	indexContent, _ := os.ReadFile(indexPath) // ignore: missing on first run

	root := cmd.New(cmdx.NewLogger())
	root.DisableAutoGenTag = true
	// cobra adds the completion subcommand lazily during Execute(); force it
	// here so it shows up in the generated tree.
	root.InitDefaultCompletionCmd()

	if err := doc.GenMarkdownTree(root, outDir); err != nil {
		log.Fatal(err)
	}

	if len(indexContent) > 0 {
		if err := os.WriteFile(indexPath, indexContent, 0o600); err != nil { //nolint:gosec
			log.Fatal(err)
		}
	}
}
