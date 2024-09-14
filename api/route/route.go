package route

import (
	"time"

	"github.com/MaogeJing/gohome/bootstrap"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(
	env *bootstrap.Env,
	timeout time.Duration,
	db gorm.DB,
	gin *gin.Engine,
) {
	// public API
	// publicRouter := gin.Group("")
	// middleware
}
