package api

import (
	"git.zgwit.com/zgwit/iot-admin/internal/conf"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type paramFilter struct {
	Key   string   `form:"key"`
	Value []string `form:"value"`
}

type paramSearch struct {
	Offset    int           `form:"offset"`
	Length    int           `form:"length"`
	SortKey   string        `form:"sortKey"`
	SortOrder string        `form:"sortOrder"`
	Filters   []paramFilter `form:"filters"`
	Keyword   string        `form:"keyword"`
}

type paramId struct {
	Id int `uri:"id"`
}

type paramId2 struct {
	Id  int `uri:"id"`
	Id2 int `uri:"id2"`
}

func mustLogin(c *gin.Context) {
	session := sessions.Default(c)
	if user := session.Get("user"); user != nil {
		c.Next()
	} else {
		//TODO 检查OAuth2返回的code，进一步获取用户信息，放置到session中

		c.JSON(http.StatusUnauthorized, gin.H{"ok": false, "error": "Unauthorized"})
		c.Abort()
	}
}

func RegisterRoutes(app *gin.RouterGroup) {

	if conf.Config.SysAdmin.Enable {
		//检查 session，必须登录
		app.Use(mustLogin)
	} else if conf.Config.BaseAuth.Enable {
		//检查HTTP认证
		app.Use(gin.BasicAuth(gin.Accounts(conf.Config.BaseAuth.Users)))
	} else {
		//支持匿名访问
	}

	//TODO 转移至子目录，并使用中间件，检查session及权限
	app.POST("/channels", channels)
	app.POST("/channel", channelCreate)
	app.DELETE("/channel/:id", channelDelete)
	app.PUT("/channel/:id", channelModify)
	app.GET("/channel/:id", channelGet)
	app.GET("/channel/:id/start", channelStart)
	app.GET("/channel/:id/stop", channelStop)

	//app.POST("/channel/:id/links")

	//连接管理
	app.POST("/links", links)
	app.DELETE("/link/:id", linkDelete)
	app.PUT("/link/:id", linkModify)
	app.GET("/link/:id", linkGet)

	//插件管理
	app.POST("/plugins", plugins)
	app.POST("/plugin", pluginCreate)
	app.DELETE("/plugin/:id", pluginDelete)
	app.PUT("/plugin/:id", pluginModify)
	app.GET("/plugin/:id", pluginGet)

	//模型管理
	app.POST("/models", models)
	app.POST("/model", modelCreate)
	app.DELETE("/model/:id", modelDelete)
	app.PUT("/model/:id", modelModify)
	app.GET("/model/:id", modelGet)

	//app.GET("/model/:id/tunnels", nop)
	//app.GET("/model/:id/variables", nop)
	//app.GET("/model/:id/batches", nop)
	//app.GET("/model/:id/jobs", nop)
	//app.GET("/model/:id/strategies", nop)

	app.POST("/model/import", modelImport)
	app.GET("/model/:id/export", modelExport)

	app.POST("/tunnels", tunnels)
	app.POST("/tunnel", tunnelCreate)
	app.DELETE("/tunnel/:id", tunnelDelete)
	app.PUT("/tunnel/:id", tunnelModify)
	app.GET("/tunnel/:id", tunnelGet)

	app.POST("/variables", variables)
	app.POST("/variable", variableCreate)
	app.DELETE("/variable/:id", variableDelete)
	app.PUT("/variable/:id", variableModify)
	app.GET("/variable/:id", variableGet)

	app.POST("/batches", batches)
	app.POST("/batch", batchCreate)
	app.DELETE("/batch/:id", batchDelete)
	app.PUT("/batch/:id", batchModify)
	app.GET("/batch/:id", batchGet)

	app.POST("/jobs", jobs)
	app.POST("/job", jobCreate)
	app.DELETE("/job/:id", jobDelete)
	app.PUT("/job/:id", jobModify)
	app.GET("/job/:id", jobGet)

	app.POST("/strategies", strategies)
	app.POST("/strategy", strategyCreate)
	app.DELETE("/strategy/:id", strategyDelete)
	app.PUT("/strategy/:id", strategyModify)
	app.GET("/strategy/:id", strategyGet)

}

func replyOk(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"ok":   true,
		"data": data,
	})
}

func replyFail(ctx *gin.Context, err string) {
	ctx.JSON(http.StatusOK, gin.H{
		"ok":    false,
		"error": err,
	})
}

func replyError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusOK, gin.H{
		"ok":    false,
		"error": err.Error(),
	})
}

func nop(c *gin.Context) {
	c.String(http.StatusForbidden, "Unsupported")
}
