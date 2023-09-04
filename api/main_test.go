package api

import (
	"os"
	"testing"
	"time"

	db "github.com/MohammadZeyaAhmad/Bank-App/db/sqlc"
	"github.com/MohammadZeyaAhmad/Bank-App/util"
	"github.com/MohammadZeyaAhmad/Bank-App/worker"
	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}
   redisOpt := asynq.RedisClientOpt{
		Addr: config.RedisAddress,
	}
   
	taskDistributor := worker.NewRedisTaskDistributor(redisOpt)
	server, err := NewServer(config, store,taskDistributor)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}