package v0

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"langgo/app/models"
	"langgo/app/pkg/common"
	"langgo/app/pkg/web"
	"langgo/app/repo"
	"langgo/bootstrap/plugins"
)

// CreateUserHandler    创建用户
//	@Summary		创建用户接口
//	@Description	创建用户接口
//	@Tags			用户
//	@Accept			application/json
//	@Param			RequestBody	body	models.CreateUser	true	"创建用户请求体"
//	@Produce		application/json
//	@Success		200	{object}	web.Response
//	@Router			/api/langgo/v0/user [post]
func CreateUserHandler(c *gin.Context) {
	var user models.CreateUser
	if err := web.GetByJSONBody(c, &user); err != nil {
		web.ParamsError(c, fmt.Sprintf("参数解析有误，详情:%s", err))
		return
	}

	var lgDB = new(plugins.LangGoDB).Use("default").NewDB()
	userUid := common.GenerateID()
	userInfo := models.User{
		UUID:     userUid,
		Name:     user.Name,
		Mobile:   user.Mobile,
		Password: user.Password,
		SoftDeletes: models.SoftDeletes{
			Status: 1,
		},
	}
	if err := repo.UserRepo.Create(lgDB, &userInfo); err != nil {
		web.InternalError(c, err.Error())
	}
	web.Success(c, userUid)
}

// QueryUserHandler    翻页查询用户
//	@Summary		翻页查询用户
//	@Description	翻页查询用户
//	@Tags			用户
//	@Accept			application/json
//	@Param			page	query	string	true	"翻页"
//	@Param			limit	query	string	true	"偏移"
//	@Produce		application/json
//	@Success		200	{object}	web.Response{data=models.QueryUserPage}
//	@Router			/api/langgo/v0/user [get]
func QueryUserHandler(c *gin.Context) {
	params := web.NewReqParams(c)
	params.PageByQuery().Cols("name", "mobile", "password").Eq("status", 1)

	var lgDB = new(plugins.LangGoDB).Use("default").NewDB()
	res, page, err := repo.UserRepo.FindPage(lgDB, &params.Condition)
	if err != nil {
		lgLogger.WithContext(c).Error("查询用户信息表失败，详情：", zap.Any("err", err.Error()))
		web.NotFoundResource(c, "用户不存在")
		return
	}

	var user []models.QueryUser
	for _, r := range res {
		user = append(user, models.QueryUser{
			Name:     r.Name,
			Mobile:   r.Mobile,
			Password: r.Password,
		})
	}

	userPage := models.QueryUserPage{
		Data: user,
		Page: page,
	}
	web.Success(c, userPage)
}

// QueryUserByUUIDHandler    根据用户ID查询用户
//	@Summary		通过用户ID查询用户
//	@Description	通过用户ID查询用户
//	@Tags			用户
//	@Accept			application/json
//	@Param			userid	path	string	true	"用户ID"
//	@Produce		application/json
//	@Success		200	{object}	web.Response{data=models.QueryUser}
//	@Router			/api/langgo/v0/user/{userid} [get]
func QueryUserByUUIDHandler(c *gin.Context) {
	userID := web.GetByPath(c, "userid")
	var lgDB = new(plugins.LangGoDB).Use("default").NewDB()
	res, err := repo.UserRepo.GetByUUID(lgDB, userID)
	if err != nil {
		lgLogger.WithContext(c).Error("查询用户信息表失败，详情：", zap.Any("err", err.Error()))
		web.NotFoundResource(c, "用户不存在")
		return
	}

	user := models.QueryUser{
		Name:     res.Name,
		Mobile:   res.Mobile,
		Password: res.Password,
	}
	web.Success(c, user)
}

// QueryUserByNameHandler    根据用户名称查询用户
//	@Summary		根据用户名称查询用户
//	@Description	根据用户名称查询用户
//	@Tags			用户
//	@Accept			application/json
//	@Param			name	path	string	true	"用户名称"
//	@Param			page	query	string	true	"翻页"
//	@Param			limit	query	string	true	"限制"
//	@Produce		application/json
//	@Success		200	{object}	web.Response{data=models.QueryUserPage}
//	@Router			/api/langgo/v0/user/name/{name} [get]
func QueryUserByNameHandler(c *gin.Context) {
	params := web.NewReqParams(c)
	params.PageByQuery().EqByPath("name").Cols("name", "mobile", "password").
		Eq("status", 1).Desc("updated_at")

	var lgDB = new(plugins.LangGoDB).Use("default").NewDB()
	res, page, err := repo.UserRepo.FindPage(lgDB, &params.Condition)
	if err != nil {
		lgLogger.WithContext(c).Error("查询用户信息表失败，详情：", zap.Any("err", err.Error()))
		web.NotFoundResource(c, "用户不存在")
		return
	}

	var user []models.QueryUser
	for _, r := range res {
		user = append(user, models.QueryUser{
			Name:     r.Name,
			Mobile:   r.Mobile,
			Password: r.Password,
		})
	}

	userPage := models.QueryUserPage{
		Data: user,
		Page: page,
	}
	web.Success(c, userPage)
}

// UpdateUserByUUIDHandler    根据用户ID更新用户
//	@Summary		根据用户ID更新用户
//	@Description	根据用户ID更新用户
//	@Tags			用户
//	@Accept			application/json
//	@Param			userid	path	string	true	"用户ID"
//	@Param			RequestBody	body	models.UpdateUser	true	"更新用户请求体"
//	@Produce		application/json
//	@Success		200	{object}	web.Response{data=models.QueryUser}
//	@Router			/api/langgo/v0/user/{userid} [patch]
func UpdateUserByUUIDHandler(c *gin.Context) {
	userID := web.GetByPath(c, "userid")
	var user models.UpdateUser
	if err := web.GetByJSONBody(c, &user); err != nil {
		web.ParamsError(c, fmt.Sprintf("参数解析有误，详情:%s", err))
		return
	}
	var lgDB = new(plugins.LangGoDB).Use("default").NewDB()
	if _, err := repo.UserRepo.GetByUUID(lgDB, userID); err != nil {
		lgLogger.WithContext(c).Error("查询用户信息表失败，详情：", zap.Any("err", err.Error()))
		web.NotFoundResource(c, "用户不存在")
		return
	}
	// 这里的事务仅仅是示例，单表更新不需要事务
	err := lgDB.Transaction(func(tx *gorm.DB) error {
		if err := repo.UserRepo.Updates(tx, userID, map[string]interface{}{
			"Mobile":   user.Mobile,
			"Password": user.Password,
		}); err != nil {
			return err
		}
		return nil
	})

	if err == nil {
		web.Success(c, user)
	} else {
		web.InternalError(c, fmt.Sprintf("更新用户信息失败，详情：%s", err.Error()))
	}
}

// DeleteUserByUUIDHandler    根据用户ID删除用户
//	@Summary		删除用户
//	@Description	删除用户
//	@Tags			用户
//	@Accept			application/json
//	@Param			userid	path	string	true	"用户名称"
//	@Produce		application/json
//	@Success		200	{object}	web.Response
//	@Router			/api/langgo/v0/user/{userid} [delete]
func DeleteUserByUUIDHandler(c *gin.Context) {
	userID := c.Param("userid")

	var lgDB = new(plugins.LangGoDB).Use("default").NewDB()

	if _, err := repo.UserRepo.GetByUUID(lgDB, userID); err != nil {
		lgLogger.WithContext(c).Error("查询用户信息表失败，详情：", zap.Any("err", err.Error()))
		web.NotFoundResource(c, "用户不存在")
		return
	}

	if err := repo.UserRepo.Delete(lgDB, userID); err != nil {
		lgLogger.WithContext(c).Error("删除用户信息表失败，详情：", zap.Any("err", err.Error()))
		web.InternalError(c, "删除用户信息失败")
		return
	}
	web.Success(c, "删除成功")
}
