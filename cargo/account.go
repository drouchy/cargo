package cargo

import (
  "fmt"
  "io"
  term "github.com/buger/goterm"
  "github.com/drouchy/dropletkit"
)

type AccountCommand struct {
  writer io.Writer
  fetcher dropletkit.AccountFetcher
  config Config
}

func (command AccountCommand) new() AccountCommand{
  return AccountCommand{}
}

func (command AccountCommand) execute() {
  options := dropletkit.DefaultOptions()
  options.Token = command.config.DigitalOcean.Token
  info := dropletkit.AccountInfo(options, command.fetcher)

  table := term.NewTable(0, 10, 5, ' ', 0)
  fmt.Fprintf(table, "Email\t%s\n", info.Email)
  fmt.Fprint(table, "Email verified\t",info.EmailVerified, "\n")
  fmt.Fprintf(table, "Droplet limit\t%d\n", info.DropletLimit)
  fmt.Fprintf(command.writer, table.String())
}
