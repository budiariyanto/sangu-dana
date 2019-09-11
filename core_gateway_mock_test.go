package dana

import (
	"io"

	"github.com/stretchr/testify/mock"
)

type CoreGatewayMock struct {
	mock.Mock
	Client Client
}

func (m *CoreGatewayMock) Call(method, path string, header map[string]string, body io.Reader, v interface{}) error {
	args := m.Called(method, path, header, body, v)
	return args.Error(0)
}

func (m *CoreGatewayMock) Order() {

}
