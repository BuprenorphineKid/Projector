package cli

import(
)

func Args(s ...string) {
  switch s[1] {
  case "-h", "--help", "help", "Help":
    Help()
  case "-r", "--ruby", "rb", "ruby", "Ruby":
    rbArgs(s[2:]...)
  case "-g", "--go", "go", "Go":
    goArgs(s[2:]...)
  case "-b", "-sh", "--bash", "--shell", "bash", "shell", "Bash", "Shell":
    shArgs(s[2:]...)
  default:
    Help()
  }
}

func rbArgs(a ...string) {
  switch a[0] {
  case "init", "new", "project":
    project("rb", a[1])
  case "class":
    rbClass(a[1])
  case "module":
    rbModule(a[1])
  case "file", "script":
    rbFile(a[1])
  }
}

func goArgs(a ...string) {
  switch a[0] {
  case "init", "new", "project":
    project("go", a[1])
  case "pkg", "package":
    goPkg(a[1])
  case "file", "script":
    goFile(a[1], a[2])
  }
}

func shArgs(a ...string) {
  switch a[0] {
  case "init", "new", "project":
    project("sh", a[1])
  case "script", "file":
    shFile(a[1])
  }
}

func Help() {
  println("prj - short for \"project\", just a quicker and easier way to")
  println("spawn in some prefilled barebones templates for your desired")
  println("language. !!!!EVERY FINAL OPTION SHOULD ALSO BE RAN WITH A NAME")
  println("OF YOUR CHOICE OR THE DEFAULT WILL BE GIVEN!!!!")
  print(")-)-) ")
  println("Current working options are:\n")
  println("-h  \\\\  (--)help :  [init|new|project]")   
  println("-b  \\\\  (--)bash : [init|new|project|script] : [file|script]")
  println("-g  \\\\  (--)go : [init|new|project|module] : [file|script|pkg|package]")
  println("-r  \\\\  (--)ruby : [init|new|project|app] : [file|script|class|module]\n")
  println("usage")
  print(")-)-) ")
  println("prj -r class blocks | outputs: blocks.rb(class template")
  println("prj go init hangman| outputs: hangman/(new project dir)")
}
