package cargo

import (
  "strings"
  "bytes"
  . "gopkg.in/check.v1"
  "github.com/drouchy/dropletkit"
)

type AccountTestSuite struct{
  writer *bytes.Buffer
  err_writer *bytes.Buffer
  command AccountCommand
}

var _ = Suite(&AccountTestSuite{})

func (suite *AccountTestSuite) SetUpTest(c *C) {
  suite.writer     = new(bytes.Buffer)
  suite.err_writer = new(bytes.Buffer)
  suite.command    = AccountCommand{config: defaultConfig(), writer: suite.writer, error_writer: suite.err_writer, fetcher: mockFetcher}
}

func (suite *AccountTestSuite) TearDownTest(c *C) {
  suite.writer.Reset()
}

func (suite *AccountTestSuite) TestWritesTheAccountEmail (c *C) {
  suite.command.execute()

  line := strings.Split(suite.writer.String(), "\n")[0]
  c.Assert(line, Matches, ".*Email.*")
  c.Assert(line, Matches, ".*foo@example\\.com.*")
}

func (suite *AccountTestSuite) TestWritesTheAccountEmailVerified (c *C) {
  suite.command.execute()

  line := strings.Split(suite.writer.String(), "\n")[1]
  c.Assert(line, Matches, ".*Email verified.*")
  c.Assert(line, Matches, ".*true.*")
}

func (suite *AccountTestSuite) TestWritesTheAccountDropletLimit (c *C) {
  suite.command.execute()

  line := strings.Split(suite.writer.String(), "\n")[2]
  c.Assert(line, Matches, ".*Droplet limit.*")
  c.Assert(line, Matches, ".*210.*")
}

func (suite *AccountTestSuite) TestReturnsNoError (c *C) {
  error := suite.command.execute()

  c.Assert(error, IsNil)
}

func (suite *AccountTestSuite) TestDoesNotWriteAnythingOnStdoutWhenIsNotAuthenticated (c *C) {
  suite.command.fetcher = mockUnauthenticatedFetcher

  suite.command.execute()

  c.Assert(suite.writer.String(), Equals, "")
}

func (suite *AccountTestSuite) TestReturnsTheErrorWhenIsNotAuthenticated (c *C) {
  suite.command.fetcher = mockUnauthenticatedFetcher

  error := suite.command.execute()

  c.Assert(error.code, Equals, 401)
}


func defaultConfig() Config{
  return Config{DigitalOcean: DigitalOcean{Token: "asdfghjkl;"}}
}

func mockFetcher(options dropletkit.Options) (dropletkit.Account, error) {
  return dropletkit.Account{Uuid: "uuid1", Email: "foo@example.com", EmailVerified: true, DropletLimit: 210}, nil
}

func mockUnauthenticatedFetcher(options dropletkit.Options) (dropletkit.Account, error) {
  return dropletkit.Account{}, dropletkit.UnauthenticatedError
}

