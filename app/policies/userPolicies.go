// Package policies 用户授权
package policies

import (
	"liu/app/models/user"

	"github.com/gin-gonic/gin"
)

func CanModifyProject(ctx *gin.Context, _user user.User) bool {
	return false
}

// func CanViewProject(c *gin.Context, projectModel project.Project) bool {}
// func CanCreateProject(c *gin.Context, projectModel project.Project) bool {}
// func CanUpdateProject(c *gin.Context, projectModel project.Project) bool {}
// func CanDeleteProject(c *gin.Context, projectModel project.Project) bool {}
