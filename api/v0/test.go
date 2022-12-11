package v0

import (
	"StorageProxy/bootstrap"
	"StorageProxy/bootstrap/plugins"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var lgLogger *bootstrap.LangGoLogger

// 不能提前创建，变量的初始化在main之前，导致lgDB为nil
//var lgDB = new(plugins.LangGoDB).NewDB()

func Test(c *gin.Context) {
	var lgDB = new(plugins.LangGoDB).NewDB()
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

	c.String(http.StatusOK, "test router")
}
