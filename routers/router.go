package routers

import (
	"github.com/Henate/Bloggor/middleware/jwt"
	"github.com/Henate/Bloggor/pkg/export"
	"github.com/Henate/Bloggor/pkg/logging"
	"github.com/Henate/Bloggor/pkg/upload"
	"github.com/Henate/Bloggor/routers/api"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"

	"github.com/Henate/Bloggor/pkg/setting"
	"github.com/Henate/Bloggor/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))

	r.GET("/auth", api.GetAuth)	//获取token
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload", api.UploadImage)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())	//使用jwt中间件
	{
		logging.Info("haha")
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
		apiv1.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
		apiv1.POST("/tags/export", v1.ExportTag)
		apiv1.POST("/tags/import", v1.ImportTag)
	}

	return r
}