import Vue from 'vue'
import VueRouter from 'vue-router'
import Index from "../pages"
import Home from "../pages/Home.vue"
import BlogDetail from "../pages/BlogDetail.vue"
import Types from "../pages/Types.vue"
import Tags from "../pages/Tags.vue"
import TimeLine from "../pages/TimeLine.vue"
import About from "../pages/About.vue"
import Essay from "../pages/Essay"
import MsgBoard from "../pages/MsgBoard"
import PhotoWall from "../pages/PhotoWall"
import ResourceLib from "../pages/ResourceLib"
import NotFound from "../components/error/NOTFOUND.vue"
import InternalError from "../components/error/INTERNALERROR.vue"


import Login from "../pages/admin/Login"
import AdminHome from '../pages/admin/Home'
import Welcome from "../pages/admin/Welcome.vue"
import ListBlogs from "../pages/admin/blogs/ListBlogs.vue"
import ListTypes from "../pages/admin/types/ListTypes.vue"
import ListTags from "../pages/admin/tags/ListTags.vue"
import AddBlog from "../pages/admin/blogs/AddBlog.vue"
import ListImages from "../pages/admin/images/ListImages"
import ManageMessage from "../pages/admin/messages/ManageMessage";
import ListMottos from "../pages/admin/blogs/listMottos";
import ManageResLink from "../pages/admin/resourecLink/manageResLink";
import ManageEssay from "../pages/admin/essay/ManageEssay";

Vue.use(VueRouter)


const routes = [
  { path: "/", redirect: "/index" },
  {
    path: "/index",
    component: Index,
    redirect: "/home",
    children: [
      { path: "/home", component: Home },
      { path: "/blogDetail", component: BlogDetail },
      { path: "/types", component: Types },
      { path: "/tags", component: Tags },
      { path: "/timeLine", component: TimeLine },
      { path: "/essay", component: Essay },
      { path: "/msgBoard", component: MsgBoard },
      { path: "/photoWall", component: PhotoWall },
      { path: "/resourceLib", component: ResourceLib },
      { path: "/about", component: About },
      { path: "/notFound", component: NotFound},
      { path: "/internalError", component: InternalError}
    ]
  },
  { path: "/manage/login", component: Login },
  {
    path: "/manage/home",
    component: AdminHome,
    redirect: "/welcome",
    children: [
      { path: "/welcome", component: Welcome},
      { path: "/listBlogs", component: ListBlogs},
      { path: "/addBlog", component: AddBlog},
      { path: "/listMottos", component: ListMottos},
      { path: "/listTypes", component: ListTypes},
      { path: "/listTags", component: ListTags},
      { path: "/listImages", component: ListImages},
      { path: "/manageMsg", component: ManageMessage},
      { path: "/manageLink", component: ManageResLink},
      { path: "/manageEssay", component: ManageEssay}
    ]
  }
]

const router = new VueRouter({
  routes
})

//挂载路由导航守卫,此守卫是用来拦截页面访问的，如果没有token，则不能会被重定向到登录页
// 访问博客页面不需要token，直接放行
// 在登陆时不需要token，直接放行
router.beforeEach((to, from, next) => {
  // to: 将要访问的路径
  // form: 从哪个路径跳转而来
  // next 是一个函数，表示放行
  // next() 放行 next("/login") 强制跳转
  // 如果访问的不是管理员页面，直接放行
  let requestUrl = to.path
  let start = requestUrl.indexOf("/manage")
  if (start != 0) {
    return next();
  }

  // 如果访问的是登录页，也放行
  if(to.path == "/manage/login") {
    return next();
  }

  // 访问的是管理员页面，查看有没有token，如果有，则放行，否则，跳转到登录页
  const tokenStr = window.sessionStorage.getItem("token");
  if(!tokenStr) return next("/manage/login");
  next();
})


export default router
