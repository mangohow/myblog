package router

import (
	"blog_web/router/admin"
	"github.com/gin-gonic/gin"
)

func RegisterBlogRouters(engine *gin.Engine) {
	homeRouter := NewHomeRouter()
	blogGroup := engine.Group("/api/myblog")
	{
		blogGroup.GET("/blogLists", homeRouter.HomeListBlogs)
		blogGroup.GET("/userInfo", homeRouter.GetHomePageUInfo)
		blogGroup.GET("/bgs", homeRouter.GetBgImages)
		blogGroup.GET("/newBlogs", homeRouter.GetNewBlogs)
		blogGroup.GET("/hotBlogs", homeRouter.GetHotBlogs)
		blogGroup.GET("/mottos", homeRouter.GetMottos)
		blogGroup.GET("/detailedBlog", homeRouter.GetDetailedBlog)
		blogGroup.GET("/commentList", homeRouter.GetCommentList)
		blogGroup.POST("/publishComment", homeRouter.PublishComment)
		blogGroup.GET("/search", homeRouter.SearchBlog)
	}

	typeListRouter := NewTypeListRouter()
	{
		blogGroup.GET("/typeList", typeListRouter.GetTypeList)
		blogGroup.GET("/typeBlogList", typeListRouter.GetBlogListByTypeid)
	}

	tagListRouter := NewTagListRouter()
	{
		blogGroup.GET("/tagList", tagListRouter.GetTagList)
		blogGroup.GET("/tagBlogList", tagListRouter.GetBlogListByTagId)
	}

	timeLineRouter := NewTimeLineRouter()
	{
		blogGroup.GET("/timeLine", timeLineRouter.GetTimeLinedBlogs)
		blogGroup.GET("/staticsBlog", timeLineRouter.GetGroupedBlogs)
	}

	resourceLibRouter := NewResourceLibRouter()
	{
		blogGroup.GET("/links", resourceLibRouter.LinkList)
	}

	leaveMessageRouter := NewLeaveMessageRouter()
	{
		blogGroup.POST("/leaveMsg", leaveMessageRouter.LeaveMessage)
		blogGroup.GET("/displayMsg", leaveMessageRouter.DisplayMessage)
	}

	essayRouter := NewEssayRouter()
	{
		blogGroup.GET("/essayList", essayRouter.EssayList)
	}
}

func RegisterBlogManageRouter(engine *gin.Engine) {
	loginRouter := admin.NewLoginRouter()
	engine.POST("/api/admin/login", loginRouter.Login)

	adminGroup := engine.Group("/api/admin")
	adminGroup.Use(admin.LoginAuthenticationMiddleware())
	blogRouter := admin.NewBlogRouter()
	{
		adminGroup.GET("/searchBlogs", blogRouter.SearchBlogs)
		adminGroup.DELETE("/deleteBlog", blogRouter.DeleteBlog)
		adminGroup.GET("/getFullBlog", blogRouter.GetFullBlog)
		adminGroup.PUT("/updateBlog", blogRouter.AddBlog)
		adminGroup.GET("/mottoList", blogRouter.MottoList)
		adminGroup.POST("addMotto", blogRouter.AddMotto)
		adminGroup.PUT("/updateMotto", blogRouter.UpdateMotto)
		adminGroup.DELETE("/deleteMotto", blogRouter.DeleteMotto)
	}

	typeRouter := admin.NewTypeRouter()
	{
		adminGroup.GET("/getAllTypes", typeRouter.GetAllTypes)
		adminGroup.GET("/getPageTypes", typeRouter.GetOnePageTypes)
		adminGroup.GET("/checkTypeExist", typeRouter.CheckTypeExist)
		adminGroup.DELETE("/deleteType", typeRouter.DeleteType)
		adminGroup.PUT("/updateType", typeRouter.UpdateType)
		adminGroup.POST("/addType", typeRouter.AddType)
	}

	tagRouter := admin.NewTagRouter()
	{
		adminGroup.GET("/getAllTags", tagRouter.GetAllTags)
		adminGroup.GET("/getPageTags", tagRouter.GetOnePageTags)
		adminGroup.GET("/checkTagExist", tagRouter.CheckTagExist)
		adminGroup.DELETE("/deleteTag", tagRouter.DeleteTag)
		adminGroup.PUT("/updateTag", tagRouter.UpdateTag)
		adminGroup.POST("/addTag", tagRouter.AddTag)
	}

	imageUploadRouter := admin.NewImageUploadRouter()
	{
		adminGroup.POST("/saveImages", imageUploadRouter.UploadBlogImage)
	}
	engine.POST("/api/admin/uploadImages", imageUploadRouter.UploadImage)
	engine.POST("/api/admin/uploadIcon", imageUploadRouter.UploadIcon)

	linksRouter := admin.NewLinksRouter()
	{
		adminGroup.GET("/pageLinks", linksRouter.LinksList)
		adminGroup.DELETE("/deleteLink", linksRouter.DeleteLink)
		adminGroup.POST("/addLink", linksRouter.AddLink)
		adminGroup.PUT("/updateLink", linksRouter.UpdateLink)
		adminGroup.GET("/categories", linksRouter.Categories)
		adminGroup.DELETE("/deleteCategory", linksRouter.DeleteCategory)
		adminGroup.POST("/addCategory", linksRouter.AddCategory)
		adminGroup.PUT("/updateCategory", linksRouter.UpdateCategory)
	}

	essayRouter := admin.NewEssayRouter()
	{
		adminGroup.GET("/essayList", essayRouter.EssayList)
		adminGroup.POST("/addEssay", essayRouter.AddEssay)
		adminGroup.DELETE("/deleteEssay", essayRouter.DeleteEssay)
		adminGroup.PUT("/updateEssay", essayRouter.UpdateEssay)
	}

	messageRouter := admin.NewMessageRouter()
	{
		adminGroup.GET("/msgList", messageRouter.MessageList)
		adminGroup.PUT("/updateMsgStatus", messageRouter.UpdateStatus)
	}
}

