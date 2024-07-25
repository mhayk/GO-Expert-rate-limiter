package limiter

import (
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"

	"github.com/stretchr/testify/suite"
)

type RedisStoreTestSuite struct {
	suite.Suite
	redisServer *miniredis.Miniredis
	redisStore  *RedisStore
}

func (suite *RedisStoreTestSuite) SetupTest() {
	var err error
	suite.redisServer, err = miniredis.Run()
	suite.Require().NoError(err)

	suite.redisStore = NewRedisStore(suite.redisServer.Addr(), "")
}

func (suite *RedisStoreTestSuite) TearDownTest() {
	suite.redisServer.Close()
}

func (suite *RedisStoreTestSuite) TestGet() {
	key := "test-key"
	value := "test-value"
	suite.redisServer.Set(key, value)

	result, err := suite.redisStore.Get(key)
	suite.NoError(err)
	suite.Equal(value, result)
}

func (suite *RedisStoreTestSuite) TestGetNonExistentKey() {
	key := "non-existent-key"

	result, err := suite.redisStore.Get(key)
	suite.NoError(err)
	suite.Equal("", result)
}

func (suite *RedisStoreTestSuite) TestSet() {
	key := "test-key"
	value := "test-value"
	expiration := 10 * time.Second

	err := suite.redisStore.Set(key, value, expiration)
	suite.NoError(err)

	result, err := suite.redisStore.Get(key)
	suite.NoError(err)
	suite.Equal(value, result)
}

func (suite *RedisStoreTestSuite) TestIncr() {
	key := "test-key"

	err := suite.redisStore.Incr(key)
	suite.NoError(err)

	result, err := suite.redisStore.Get(key)
	suite.NoError(err)
	suite.Equal("1", result)
}

func (suite *RedisStoreTestSuite) TestExpire() {
	key := "test-key"
	value := "test-value"
	expiration := 1 * time.Second

	err := suite.redisStore.Set(key, value, 0)
	suite.NoError(err)

	err = suite.redisStore.Expire(key, expiration)
	suite.NoError(err)

	time.Sleep(2 * time.Second)

	result, err := suite.redisStore.Get(key)
	suite.NoError(err)
	suite.Equal("test-value", result)
}

func TestRedisStoreTestSuite(t *testing.T) {
	suite.Run(t, new(RedisStoreTestSuite))
}
