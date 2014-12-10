package cargo

import (
  "github.com/codegangsta/cli"
  "github.com/drouchy/dropletkit"
  "os"
  "os/user"
  "path"
)

type Command interface {
  execute()
}

func CreateApp() *cli.App {
  app := cli.NewApp()
  app.Name = "cargo"
  app.Usage = "manage you cloud server instances"

  user, _ := user.Current()
  app.Flags = []cli.Flag {
    cli.StringFlag{
      Name: "config_file, f",
      Value: path.Join(user.HomeDir, ".cargo"),
      Usage: "configuration file",
    },
  }

  app.Commands = []cli.Command{
    {
      Name: "account",
      ShortName: "a",
      Usage: "retreive your account information",
      Action: func (context *cli.Context) {
        config := LoadConfig(context.GlobalString("config_file"))
        AccountCommand{writer: os.Stdout, error_writer: os.Stderr, config: config, fetcher: dropletkit.AccountFetcherImpl}.execute()
      },
    },
  }
  return app
}
