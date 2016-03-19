package main

import (
	. "gopkg.in/check.v1"
	"testing"
)

// gocheck hook
func Test(t *testing.T) { TestingT(t) }

// Arrange
type TestBuilder struct {
	countries    []Country
	badCountries []Country
}

var _ = Suite(&TestBuilder{})

func (self *TestBuilder) SetUpSuite(c *C) {
	self.countries = []Country{
		Country{
			Name:    "test",
			infoRaw: "{}",
		},
	}

	self.badCountries = []Country{Country{Name: "test"}}

}

func (self *TestBuilder) TestOneCountyInfoError(c *C) {
	// Act
	json, err := prettyJson(self.badCountries)
	c.Logf("e: %v // %v", err, json)
	// Assert
	c.Assert(err, ErrorMatches, "can't parse \"info\".*")
}

func (self *TestBuilder) TestOneCountyInfoOK(c *C) {
	// Act
	json, err := prettyJson(self.countries)
	// Assert
	c.Assert(err, IsNil)
	c.Assert(json, Matches, ".*\"name\":\"test\".*")
}
