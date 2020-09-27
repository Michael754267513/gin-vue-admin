package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitProjectRouter(Router *gin.RouterGroup) {
	ProjectRouter := Router.Group("project").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		ProjectRouter.POST("createProject", v1.CreateProject)             // 新建Project
		ProjectRouter.DELETE("deleteProject", v1.DeleteProject)           // 删除Project
		ProjectRouter.DELETE("deleteProjectByIds", v1.DeleteProjectByIds) // 批量删除Project
		ProjectRouter.PUT("updateProject", v1.UpdateProject)              // 更新Project
		ProjectRouter.GET("findProject", v1.FindProject)                  // 根据ID获取Project
		ProjectRouter.GET("getProjectList", v1.GetProjectList)            // 获取Project列表
	}
}
