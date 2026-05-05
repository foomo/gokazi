## gokazi completion fish

Generate the autocompletion script for fish

### Synopsis

Generate the autocompletion script for the fish shell.

To load completions in your current shell session:

	gokazi completion fish | source

To load completions for every new session, execute once:

	gokazi completion fish > ~/.config/fish/completions/gokazi.fish

You will need to start a new shell for this setup to take effect.


```
gokazi completion fish [flags]
```

### Options

```
  -h, --help              help for fish
      --no-descriptions   disable completion descriptions
```

### Options inherited from parent commands

```
      --debug   enable debug logs
```

### SEE ALSO

* [gokazi completion](gokazi_completion.md)	 - Generate the autocompletion script for the specified shell

