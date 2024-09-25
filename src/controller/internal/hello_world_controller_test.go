package internal

import (
	"context"
	"github.com/stretchr/testify/suite"
	"gitlab.com/AngelX/common/config"
	"strconv"
	"testing"
	"upserv/src/http/request"
	"upserv/src/http/response"
	"upserv/src/storage/mocks"
	"upserv/src/storage/model"
)

type HelloWorldSuite struct {
	suite.Suite
	HelloWorldController *HelloWorldController
	mockFirstService     *mocks.IHelloWorld
}

func (s *HelloWorldSuite) SetupSuite() {
	s.mockFirstService = &mocks.IHelloWorld{}
	s.HelloWorldController = NewHelloWorldController(s.mockFirstService)
}

func (s *HelloWorldSuite) TestHelloWorld_Success() {
	ctx := context.Background()

	err := config.Init("../../../config.ini", "")
	if !s.Nil(err) {
		return
	}

	title := "Hello world"
	returnModel := &model.HelloWorld{Title: title}
	s.mockFirstService.On("Find", ctx).Return(returnModel, nil)

	lr := &request.GetHelloWorld{}
	result, err := s.HelloWorldController.GetHelloWorld(ctx, lr)

	if !s.Nil(err) {
		return
	}

	res := response.HelloWorld{}.From(returnModel)
	res.Cookies.Set(&response.Cookie{
		Name:  "hello_world_id",
		Value: strconv.FormatUint(uint64(returnModel.ID), 10),
	})
	s.Equal(&res, result)
}

func TestFirstController(t *testing.T) {
	suite.Run(t, new(HelloWorldSuite))
}
