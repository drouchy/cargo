package cargo

import (
  . "gopkg.in/check.v1"
  "time"
)

type ErrorsTestSuite struct{}

var _ = Suite(&ErrorsTestSuite{})

func (suite *ErrorsTestSuite) TestCreatesAnNewErrorSetsTheMessage (c *C) {
  error := NewError(200, "the message of the error")

  c.Assert(error.message, Equals, "the message of the error")
}

func (suite *ErrorsTestSuite) TestCreatesAnNewErrorSetsTheErrorCode (c *C) {
  error := NewError(200, "the message of the error")

  c.Assert(error.code, Equals, 200)
}

func (suite *ErrorsTestSuite) TestCreatesAnNewErrorSetsTheTimestamp (c *C) {
  error := NewError(200, "the message of the error")
  now := time.Now().UnixNano()

  c.Assert(error.timestamp.UnixNano() < now && error.timestamp.UnixNano() > now - 1000, Equals, true)
}

func (suite *ErrorsTestSuite) TestImplementsTheErrorInterface (c *C) {
  error := NewError(200, "the message of the error")

  c.Assert(error.Error(), Matches, ".* \\[200\\] the message of the error")
}

func (suite *ErrorsTestSuite) TestTheErrorCodeReturnsTheCode (c *C) {
  error := NewError(1, "the message of the error")

  c.Assert(error.ErrorCode(), Equals, 1)
}
