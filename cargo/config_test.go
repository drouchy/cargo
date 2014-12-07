package cargo

import (
  . "gopkg.in/check.v1"
  "testing"
)

func Test(t *testing.T) { TestingT(t) }

type ConfigTestSuite struct{}

var _ = Suite(&ConfigTestSuite{})

var CONFIG_FILE = "../fixtures/base_config.json"

func (suite *ConfigTestSuite) TestDecodeDigitalOceanToken (c *C) {
  config := LoadConfig(CONFIG_FILE)

  c.Assert(config.DigitalOcean.Token, Equals, "qwertyuiop")
}
