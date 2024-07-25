package limiter

import (
	"testing"
	"time"

	"github.com/mhayk/GO-Expert-rate-limiter/internal/infra/limiter/mock"
	"github.com/stretchr/testify/suite"
)

type LimiterTestSuite struct {
	suite.Suite
	limiter        *Limiter
	mockIPStore    *mock.Store
	mockTokenStore *mock.Store
}

func (suite *LimiterTestSuite) SetupTest() {
	suite.mockIPStore = mock.NewStore(suite.T())
	suite.mockTokenStore = mock.NewStore(suite.T())
	suite.limiter = NewLimiter(5, 60, 10, 60, suite.mockIPStore, suite.mockTokenStore)
}

func (suite *LimiterTestSuite) TestAllowRequestByIP() {
	ip := "127.0.0.1"
	key := "ratelimit:ip:127.0.0.1"

	suite.mockIPStore.EXPECT().Get(key).Return("4", nil)
	suite.mockIPStore.EXPECT().Incr(key).Return(nil)
	suite.mockIPStore.EXPECT().Expire(key, 60*time.Second).Return(nil)

	result := suite.limiter.AllowRequest(ip, "")
	suite.True(result)

	suite.mockIPStore.AssertExpectations(suite.T())
}

func (suite *LimiterTestSuite) TestDenyRequestByIP() {
	ip := "127.0.0.1"
	key := "ratelimit:ip:127.0.0.1"

	suite.mockIPStore.EXPECT().Get(key).Return("5", nil)

	result := suite.limiter.AllowRequest(ip, "")
	suite.False(result)

	suite.mockIPStore.AssertExpectations(suite.T())
}

func (suite *LimiterTestSuite) TestAllowRequestByToken() {
	token := "test-token"
	key := "ratelimit:token:test-token"

	suite.mockTokenStore.EXPECT().Get(key).Return("9", nil)
	suite.mockTokenStore.EXPECT().Incr(key).Return(nil)
	suite.mockTokenStore.EXPECT().Expire(key, 60*time.Second).Return(nil)

	result := suite.limiter.AllowRequest("", token)
	suite.True(result)

	suite.mockTokenStore.AssertExpectations(suite.T())
}

func (suite *LimiterTestSuite) TestDenyRequestByToken() {
	token := "test-token"
	key := "ratelimit:token:test-token"

	suite.mockTokenStore.EXPECT().Get(key).Return("10", nil)

	result := suite.limiter.AllowRequest("", token)
	suite.False(result)

	suite.mockTokenStore.AssertExpectations(suite.T())
}

func TestLimiterTestSuite(t *testing.T) {
	suite.Run(t, new(LimiterTestSuite))
}
