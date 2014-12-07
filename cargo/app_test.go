package cargo

import (
  "github.com/codegangsta/cli"
  . "gopkg.in/check.v1"
)

type AppTestSuite struct{}

var _ = Suite(&AppTestSuite{})

var defaultApp *cli.App = CreateApp()

func (suite *AppTestSuite) TestGivesTheAppAName (c *C) {
  c.Assert(defaultApp.Name, Equals, "cargo")
}

func (suite *AppTestSuite) TestGivesTheAppAnUsage (c *C) {
  c.Assert(defaultApp.Usage, Equals, "manage you cloud server instances")
}

func (suite *AppTestSuite) TestHasAFlag (c *C) {
  c.Assert(len(defaultApp.Flags), Equals, 1)
}

func (suite *AppTestSuite) TestSetupsTheConfigFileFlagName (c *C) {
  configFile := defaultApp.Flags[0].(cli.StringFlag)

  c.Assert(configFile.Name, Equals, "config_file, f")
}

func (suite *AppTestSuite) TestSetupsTheConfigFileFlagValue (c *C) {
  configFile := defaultApp.Flags[0].(cli.StringFlag)

  c.Assert(configFile.Value, Matches, ".*\\.cargo")
}

func (suite *AppTestSuite) TestSetupsTheConfigFileFlagUsage (c *C) {
  configFile := defaultApp.Flags[0].(cli.StringFlag)

  c.Assert(configFile.Usage, Equals, "configuration file")
}

func (suite *AppTestSuite) TestTheAppHasOneCommand (c *C) {
  c.Assert(len(defaultApp.Commands), Equals, 1)
}

func (suite *AppTestSuite) TestConfiguresTheAccountCommandName (c *C) {
  account := defaultApp.Commands[0]

  c.Assert(account.Name, Equals, "account")
}

func (suite *AppTestSuite) TestConfiguresTheAccountCommandShortName (c *C) {
  account := defaultApp.Commands[0]

  c.Assert(account.ShortName, Equals, "a")
}

func (suite *AppTestSuite) TestConfiguresTheAccountCommandUsage (c *C) {
  account := defaultApp.Commands[0]

  c.Assert(account.Usage, Equals, "retreive your account information")
}

func (suite *AppTestSuite) TestConfiguresTheAccountCommandAction (c *C) {
  account := defaultApp.Commands[0]

  c.Assert(account.Action, NotNil)
}
