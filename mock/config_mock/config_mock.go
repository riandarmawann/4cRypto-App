package configmock

import (
	"github.com/stretchr/testify/mock"
	"4crypto/config" // Gantilah YourProjectPath dengan path aktual ke paket config
)

type ConfigMock struct {
	mock.Mock
}

func (c *ConfigMock) NewConfig() (*config.Config, error) {
	args := c.Called()
	return args.Get(0).(*config.Config), args.Error(1)
}

func (c *ConfigMock) readConfig() error {
    args := c.Called()
    return args.Error(0)
}

