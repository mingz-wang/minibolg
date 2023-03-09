package main

import (
	"os"

	_ "go.uber.org/automaxprocs"
	"minibolg/internal/minibolg"
)

// Go 程序的默认入口函数(主函数).
func main() {
	command := miniblog.NewMiniBlogCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
