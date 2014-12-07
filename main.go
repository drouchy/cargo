package main

import (
  "./cargo"
  "os"
)

func main() {
  cargo.CreateApp().Run(os.Args)
}
