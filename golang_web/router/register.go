package router

import (
	"blog_web/controller"
	"blog_web/controller/admin"
	"github.com/gin-gonic/gin"
)

func Register(engine *gin.Engine) {
	e := NewEngine(engine)
	registerBlogRouters(e)
	registerBlogManageRouter(e)
}

func registerBlogRouters(engine *Engine) {
	homeRouter := controller.NewHomeRouter()
	blogGroup := NewRouterGroup(engine.Group("/api/myblog"))
	{
		blogGroup.Get("/blogLists", homeRouter.HomeListBlogs)
		blogGroup.Get("/userInfo", homeRouter.GetHomePageUInfo)
		blogGroup.Get("/bgs", homeRouter.GetBgImages)
		blogGroup.Get("/newBlogs", homeRouter.GetNewBlogs)
		blogGroup.Get("/hotBlogs", homeRouter.GetHotBlogs)
		blogGroup.Get("/mottos", homeRouter.GetMottos)
		blogGroup.Get("/detailedBlog", homeRouter.GetDetailedBlog)
		blogGroup.Get("/commentList", homeRouter.GetCommentList)
		blogGroup.Post("/publishComment", homeRouter.PublishComment)
		blogGroup.Get("/search", homeRouter.SearchBlog)
	}

	typeListRouter := controller.NewTypeListRouter()
	{
		blogGroup.Get("/typeList", typeListRouter.GetTypeList)
		blogGroup.Get("/typeBlogList", typeListRouter.GetBlogListByTypeid)
	}

	tagListRouter := controller.NewTagListRouter()
	{
		blogGroup.Get("/tagList", tagListRouter.GetTagList)
		blogGroup.Get("/tagBlogList", tagListRouter.GetBlogListByTagId)
	}

	timeLineRouter := controller.NewTimeLineRouter()
	{
		blogGroup.Get("/timeLine", timeLineRouter.GetTimeLinedBlogs)
		blogGroup.Get("/staticsBlog", timeLineRouter.GetGroupedBlogs)
	}

	resourceLibRouter := controller.NewResourceLibRouter()
	{
		blogGroup.Get("/links", resourceLibRouter.LinkList)
	}

	leaveMessageRouter := controller.NewLeaveMessageRouter()
	{
		blogGroup.Post("/leaveMsg", leaveMessageRouter.LeaveMessage)
		blogGroup.Get("/displayMsg", leaveMessageRouter.DisplayMessage)
	}

	essayRouter := controller.NewEssayRouter()
	{
		blogGroup.Get("/essayList", essayRouter.EssayList)
	}
}

func registerBlogManageRouter(engine *Engine) {
	loginRouter := admin.NewLoginRouter()
	engine.Post("/api/admin/login", loginRouter.Login)

	adminGroup := NewRouterGroup(engine.Group("/api/admin"))
	adminGroup.Use(admin.LoginAuthenticationMiddleware())
	blogRouter := admin.NewBlogRouter()
	{
		adminGroup.Get("/searchBlogs", blogRouter.SearchBlogs)
		adminGroup.Delete("/deleteBlog", blogRouter.DeleteBlog)
		adminGroup.Get("/getFullBlog", blogRouter.GetFullBlog)
		adminGroup.Put("/updateBlog", blogRouter.AddBlog)
		adminGroup.Get("/mottoList", blogRouter.MottoList)
		adminGroup.Post("addMotto", blogRouter.AddMotto)
		adminGroup.Put("/updateMotto", blogRouter.UpdateMotto)
		adminGroup.Delete("/deleteMotto", blogRouter.DeleteMotto)
	}

	typeRouter := admin.NewTypeRouter()
	{
		adminGroup.Get("/getAllTypes", typeRouter.GetAllTypes)
		adminGroup.Get("/getPageTypes", typeRouter.GetOnePageTypes)
		adminGroup.Get("/checkTypeExist", typeRouter.CheckTypeExist)
		adminGroup.Delete("/deleteType", typeRouter.DeleteType)
		adminGroup.Put("/updateType", typeRouter.UpdateType)
		adminGroup.Post("/addType", typeRouter.AddType)
	}

	tagRouter := admin.NewTagRouter()
	{
		adminGroup.Get("/getAllTags", tagRouter.GetAllTags)
		adminGroup.Get("/getPageTags", tagRouter.GetOnePageTags)
		adminGroup.Get("/checkTagExist", tagRouter.CheckTagExist)
		adminGroup.Delete("/deleteTag", tagRouter.DeleteTag)
		adminGroup.Put("/updateTag", tagRouter.UpdateTag)
		adminGroup.Post("/addTag", tagRouter.AddTag)
	}

	// 可以选择使用本地服务器的图片存储或者阿里云OSS对象存储服务
	//imageUploadRouter := admin.NewLocalImageUploadRouter()
	imageUploadRouter := admin.NewAliOSSImageUploadRouter()
	{
		adminGroup.POST("/saveImages", imageUploadRouter.UploadBlogImage)
	}
	engine.POST("/api/admin/uploadImages", imageUploadRouter.UploadImage)
	engine.POST("/api/admin/uploadIcon", imageUploadRouter.UploadIcon)

	linksRouter := admin.NewLinksRouter()
	{
		adminGroup.Get("/pageLinks", linksRouter.LinksList)
		adminGroup.Delete("/deleteLink", linksRouter.DeleteLink)
		adminGroup.Post("/addLink", linksRouter.AddLink)
		adminGroup.Put("/updateLink", linksRouter.UpdateLink)
		adminGroup.Get("/categories", linksRouter.Categories)
		adminGroup.Delete("/deleteCategory", linksRouter.DeleteCategory)
		adminGroup.Post("/addCategory", linksRouter.AddCategory)
		adminGroup.Put("/updateCategory", linksRouter.UpdateCategory)
	}

	essayRouter := admin.NewEssayRouter()
	{
		adminGroup.Get("/essayList", essayRouter.EssayList)
		adminGroup.Post("/addEssay", essayRouter.AddEssay)
		adminGroup.Delete("/deleteEssay", essayRouter.DeleteEssay)
		adminGroup.Put("/updateEssay", essayRouter.UpdateEssay)
	}

	messageRouter := admin.NewMessageRouter()
	{
		adminGroup.Get("/msgList", messageRouter.MessageList)
		adminGroup.Put("/updateMsgStatus", messageRouter.UpdateStatus)
	}
}

