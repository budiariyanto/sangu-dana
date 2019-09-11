package dana

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type CoreTestSuite struct {
	suite.Suite
	Gateway *CoreGatewayMock
}

func (c *CoreTestSuite) SetupTest() {
	client := NewClient()

	c.Gateway = &CoreGatewayMock{
		Client: client,
	}
}

func (c *CoreTestSuite) TestOrder(t *testing.T) {
	c.Gateway.Order()
	c.Gateway.On("Call").Return(nil)

}

func TestCoreTestSuite(t *testing.T) {
	suite.Run(t, new(CoreTestSuite))
}
