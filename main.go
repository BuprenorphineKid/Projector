package main

import (
  "os"
  "prj/cli"
)

func main() {
  def := []string{"help", "init", "default"}

  a := os.Args

  args := make([]string, len(a[1:]))

  if len(args) != 0 {
    copy(def[len(args) - 1:], args[len(args) - 1:])
    cli.Args(args...)
  }

  cli.Args(def...)
}
