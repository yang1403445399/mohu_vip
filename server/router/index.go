package router

import (
	"example/server/controller"
	"example/server/middleware"

	"github.com/gofiber/fiber/v2"
)

func InitRouter(app *fiber.App) {
	admin := app.Group("/admin")

	admin.Use(middleware.JWTAuthMiddleware())

	admin.Post("/user/login/submit", controller.CommonController{}.UserLoginSubmit)
	admin.Get("/user/login/verify", controller.CommonController{}.UserLoginVerify)
	admin.Get("/menu/list", controller.CommonController{}.MenuList)
	admin.Get("/browse/count", controller.CommonController{}.BrowseCount)
	admin.Get("/browse/region", controller.CommonController{}.BrowseRegion)
	admin.Get("/article/count", controller.CommonController{}.ArticleCount)
	admin.Get("/log/list", controller.CommonController{}.LogList)
	admin.Get("/basic/info", controller.CommonController{}.BasicInfo)
	admin.Post("/basic/submit", controller.CommonController{}.BasicSubmit)
	admin.Get("/image/list", controller.CommonController{}.ImageList)
	admin.Post("/image/upload", controller.CommonController{}.ImageUpload)
	admin.Post("/image/delete", controller.CommonController{}.ImageDelete)
	admin.Get("/image/column/list", controller.CommonController{}.ImageColumnList)
	admin.Post("/image/column/submit", controller.CommonController{}.ImageColumnSubmit)
	admin.Post("/image/column/delete", controller.CommonController{}.ImageColumnDelete)
	admin.Get("/banner/list", controller.CommonController{}.BannerList)
	admin.Get("/banner/info", controller.CommonController{}.BannerInfo)
	admin.Post("/banner/submit", controller.CommonController{}.BannerSubmit)
	admin.Post("/banner/delete", controller.CommonController{}.BannerDelete)
	admin.Get("/banner/type/list", controller.CommonController{}.BannerTypeList)
}
