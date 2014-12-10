package cargo

import (
  "fmt"
  "time"
)

type CargoError struct {
 message string
 code int
 timestamp time.Time
}

func NewError(code int, message string) CargoError {
  return CargoError{message: message, code: code, timestamp: time.Now()}
}

func (error CargoError) Error() string {
 return fmt.Sprintf("%v [%d] %s", error.timestamp, error.code, error.message)
}

func (error CargoError) ErrorCode() int {
  return error.code
}

func UnauthenticatedError(provider string) CargoError {
  return NewError(401, "Failed to authenticate to " + provider)
}
