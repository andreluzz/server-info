package server_info

import (
	"github.com/stretchr/testify/suite"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ServerTestSuite struct {
	suite.Suite
}

func (suite *ServerTestSuite) TestServerGetIPAddress() {
	s := New()

	ip, err := s.GetIPAddress("Wi-Fi")
	assert.Error(suite.T(), err, "Should have an error object")

	s.Flush()
	ip, _ = s.GetIPAddress(s.Adapters[0].Name)
	assert.Equal(suite.T(), s.Adapters[0].IP, ip)

	ip, err = s.GetIPAddress()
	assert.NoError(suite.T(), err, "Should not have any error object")
	assert.Equal(suite.T(), s.Adapters[0].IP, ip)

	_, err = s.GetIPAddress("WrongAdapterName")
	assert.Error(suite.T(), err, "Should have an error object")
}

func (suite *ServerTestSuite) TestServerToJSON() {
	s := New()
	s.Flush()
	json, err := s.ToJSON()
	assert.NoError(suite.T(), err, "Should not have any error object")
	assert.NotEmpty(suite.T(), json, "Should have the json from the server")
}

func TestServerTestSuite(t *testing.T) {
	suite.Run(t, new(ServerTestSuite))
}