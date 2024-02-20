package managermock

import (
	"database/sql"

	"github.com/stretchr/testify/mock"
)

type InfraManagerMock struct {
	mock.Mock
}

func (i *InfraManagerMock) Conn() *sql.DB {
	args := i.Called()
	return args.Get(0).(*sql.DB)
}
