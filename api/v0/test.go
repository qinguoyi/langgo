package v0

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"langgo/app/pkg/web"
	"langgo/bootstrap"
	"langgo/bootstrap/plugins"
)

// 不能提前创建，变量的初始化在main之前，导致lgDB为nil
//var lgDB = new(plugins.LangGoDB).NewDB()

var lgLogger *bootstrap.LangGoLogger

// TestHandler 测试
//	@Summary		测试接口
//	@Description	LangGo的测试接口
//	@Tags			测试
//	@Accept			application/json
//	@Produce		application/json
//	@Success		200	{object}	web.Response
//	@Router			/api/storage/v0/ping [get]
func TestHandler(c *gin.Context) {
	var lgDB = new(plugins.LangGoDB).Use("default").NewDB()
	var lgRedis = new(plugins.LangGoRedis).NewRedis()

	lgDB.Exec("select now();")
	lgLogger.WithContext(c).Info("test router")

	// Redis Test
	err := lgRedis.Set(c, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
	val, err := lgRedis.Get(c, "key").Result()
	if err != nil {
		panic(err)
	}
	lgLogger.WithContext(c).Info(fmt.Sprintf("%v", val))

	web.Success(c, "test router")
}
