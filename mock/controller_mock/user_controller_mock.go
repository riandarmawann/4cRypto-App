package controllermock

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type UserControllermock struct {
	mock.Mock
}

func (u *UserControllermock) route() {
	u.Called()
}

func (u *UserControllermock) Create(ctx *gin.Context) {
	u.Called(ctx)
}

func (u *UserControllermock) FindById(ctx *gin.Context) {
	u.Called(ctx)
}

func (u *UserControllermock) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	u.Called(w, r)
}

func (u *UserControllermock) UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	u.Called(w, r)
}
