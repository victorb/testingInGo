package testingInGo

import (
  "testing"
)

var LoginUis

func TestLoginMethodFailsWithoutCredentials(t *testing.T) {
  user := User{}

  if LoginUser() {
    t.Error("User could be logged in without credentials")
  }
}
