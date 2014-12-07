package main

import (
  "github.com/drouchy/cargo/cargo"
  "os"
)

func main() {
  cargo.CreateApp().Run(os.Args)
}
