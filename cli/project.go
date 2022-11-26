package cli

import (
  "fmt"
  "os"
  str "strings"
)

func project(s string, name string) {
  switch s {
  case "sh":
    shProj(name)
  case "rb":
    rbProj(name)
  case "go":
    goProj(name)
  }
}

func shProj(n string) {
  os.MkdirAll(n, 0777)
  
  main, _ := os.Create(n + "/" + n + ".sh")
  main.WriteString("#/usr/bin/env bash")

  r, _ := os.Create(n + "/" + "README.md")
  r.Close()
}

func shFile(n string) {
  file, _ := os.Create(n + ".sh")

  file.WriteString("#/usr/bin/env bash")  
}

func rbProj(n string) {
  os.MkdirAll(n, 0777)
  os.MkdirAll(n + "/" + "lib", 0777)
  os.MkdirAll(n + "/" + "bin", 0777)
  main, _ := os.Create(n + "/" + "main.rb")

  main.WriteString("#/usr/bin/env ruby")
  main.Close()

  r, _ := os.Create(n + "/" + "README.md")
  r.Close()

  g, _ := os.Create(n + "/" + "Gemfile")
  g.Close()
}

func rbClass(n string) {
  file, _ := os.Create(n + ".rb")


  let := str.Split(n, "")
  let[0] = str.ToUpper(let[0])
  capt := str.Join(let, "")
  
  cont := fmt.Sprintf("class %s\n\n\tdef initialize()\n\n\tend\n\nend\n", capt)
  file.WriteString(cont)
}

func rbModule(n string) {
  file, _ := os.Create(n + ".rb")


  let := str.Split(n, "")
  let[0] = str.ToUpper(let[0])
  capt := str.Join(let, "")
  
  cont := fmt.Sprintf("module %s\n\nend\n", capt)
  file.WriteString(cont)
}

func rbFile(n string) {
  file, _ := os.Create(n + ".rb")


  let := str.Split(n, "")
  let[0] = str.ToUpper(let[0])
  capt := str.Join(let, "")
  
  cont := fmt.Sprintf("module %s\n\nend\n", capt)
  file.WriteString(cont)
}

func goProj(n string) {
  os.MkdirAll(n, 0777)
  main, _ := os.Create(n + "/" + "main.go")

  main.WriteString("package main\n\nimport (\n\n)\n\nfunc main() {\n\n}\n")
  main.Close()
    
  r, _ := os.Create(n + "/" + "README.md")
  r.Close()
}

func goPkg(n string) {
  os.MkdirAll(n, 0777)  
  pkgf, _ := os.Create(n + "/" + "main.go")

  cont := fmt.Sprintf("package %s\n\nimport (\n\n)\n\nfunc main() {\n\n}\n", n)
  pkgf.WriteString(cont)
  pkgf.Close()
}
  
func goFile(n string, p string) {
  os.MkdirAll(n, 0777)  
  pkgf, _ := os.Create(n + "/" + "main.go")

  cont := fmt.Sprintf("package %s\n\nimport (\n\n)\n\nfunc main() {\n\n}\n", p)
  pkgf.WriteString(cont)
  pkgf.Close()
}
