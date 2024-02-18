package delivery

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ServerTestSuite struct {
	suite.Suite
}

func (suite *ServerTestSuite) SetupTest() {
}

func (suite *ServerTestSuite) TestSetupControllers() {

}

func (suite *ServerTestSuite) TestRun() {

}

func TestServerMockTestSuite(t *testing.T) {
	suite.Run(t, new(ServerTestSuite))
}
