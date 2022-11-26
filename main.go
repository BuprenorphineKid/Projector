package main

import (
  "os"
  "prj/cli"
)

func main() {
  a := os.Args
  def := make([]string, 3, 3)
  def[0] = "help"
  def[1] = "init"
  def[2] = "default"

  if len(a) != len(def) {
    a = append(a, def[len(a) - 1:]...)
  }
  cli.Args(a...)
}
