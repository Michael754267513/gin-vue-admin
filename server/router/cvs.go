package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitCvsRouter(Router *gin.RouterGroup) {
	CvsRouter := Router.Group("csv").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		CvsRouter.POST("createCvs", v1.CreateCvs)             // 新建Cvs
		CvsRouter.DELETE("deleteCvs", v1.DeleteCvs)           // 删除Cvs
		CvsRouter.DELETE("deleteCvsByIds", v1.DeleteCvsByIds) // 批量删除Cvs
		CvsRouter.PUT("updateCvs", v1.UpdateCvs)              // 更新Cvs
		CvsRouter.GET("findCvs", v1.FindCvs)                  // 根据ID获取Cvs
		CvsRouter.GET("getCvsList", v1.GetCvsList)            // 获取Cvs列表
	}
}
