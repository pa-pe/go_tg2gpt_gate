package internal

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gitlab.com/AngelX/common/config"
	"testing"
	mockCache "upserv/src/service/mocks"
	"upserv/src/storage/mocks"
	"upserv/src/storage/model"
)

type GetHelloWorldSuite struct {
	cache *mockCache.ICache
	suite.Suite
	firstService     *helloWorldImpl
	mockFirstStorage *mocks.IHelloWorld
}

func (s *GetHelloWorldSuite) SetupSuite() {
	s.mockFirstStorage = &mocks.IHelloWorld{}
	s.cache = &mockCache.ICache{}
	s.firstService = NewHelloWorldService(s.mockFirstStorage, s.cache)

	err := config.Init("../../../config.ini", "")
	if !s.Nil(err) {
		return
	}
}

func (s *GetHelloWorldSuite) TestHelloWorld_Success() {
	ctx := context.Background()
	title := "Hello world"
	returnModel := &model.HelloWorld{Title: title}
	s.mockFirstStorage.On("Find", ctx).Return(returnModel, nil)
	s.cache.On("Load", ctx, "helloWorld", "first", mock.AnythingOfType("*model.HelloWorld")).Return(errors.New("not found"))
	s.cache.On("Put", ctx, "helloWorld", "first", mock.AnythingOfType("model.HelloWorld"), mock.Anything)
	result, err := s.firstService.Find(ctx)
	if !s.Nil(err) {
		return
	}

	s.Equal(*returnModel, *result)
}

func TestFirstService(t *testing.T) {
	suite.Run(t, new(GetHelloWorldSuite))
}
