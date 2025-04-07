package main

import (
	"fmt"
	"os"
	"runtime/debug"

	cowsay "github.com/Code-Hex/Neo-cowsay/v2"
	"github.com/foomo/gokazi/cmd"
	cmdx "github.com/foomo/gokazi/pkg/cmd"
)

func main() {
	l := cmdx.NewLogger()

	root := cmd.NewRoot(l)
	root.AddCommand(
		cmd.NewList(l),
		cmd.NewStop(l),
		cmd.NewConfig(l),
		cmd.NewVersion(l),
	)

	say := func(msg string) string {
		if say, cerr := cowsay.Say(msg, cowsay.BallonWidth(80)); cerr == nil {
			msg = say
		}
		return msg
	}

	code := 0
	defer func() {
		if r := recover(); r != nil {
			l.Error(say("It's time to panic"))
			l.Error(fmt.Sprintf("%v", r))
			l.Error(string(debug.Stack()))
			code = 1
		}
		os.Exit(code)
	}()

	if err := root.Execute(); err != nil {
		l.Error(say("Ups, something went wrong"))
		l.Error(err.Error())
		code = 1
	}
}
