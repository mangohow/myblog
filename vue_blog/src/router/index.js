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
  {
    path: "/",
    component: Index,
    meta: {
        auth: false
    },
    redirect: "/home",
    children: [
      { path: "/home", component: Home, meta: {auth: false} },
      { path: "/blogDetail", component: BlogDetail, meta: {auth: false} },
      { path: "/types", component: Types, meta: {auth: false} },
      { path: "/tags", component: Tags, meta: {auth: false} },
      { path: "/timeLine", component: TimeLine, meta: {auth: false} },
      { path: "/essay", component: Essay, meta: {auth: false} },
      { path: "/msgBoard", component: MsgBoard, meta: {auth: false} },
      { path: "/photoWall", component: PhotoWall, meta: {auth: false} },
      { path: "/resourceLib", component: ResourceLib, meta: {auth: false} },
      { path: "/about", component: About, meta: {auth: false} },
      { path: "/notFound", component: NotFound, meta: {auth: false} },
      { path: "/internalError", component: InternalError, meta: {auth: false} }
    ]
  },
  { path: "/login", component: Login },
  {
    path: "/manageHome",
    component: AdminHome,
    meta: {
        auth: false
    },
    redirect: "/welcome",
    children: [
      { path: "/welcome", component: Welcome, meta: {auth: true} },
      { path: "/listBlogs", component: ListBlogs, meta: {auth: true} },
      { path: "/addBlog", component: AddBlog, meta: {auth: true} },
      { path: "/listMottos", component: ListMottos, meta: {auth: true} },
      { path: "/listTypes", component: ListTypes, meta: {auth: true} },
      { path: "/listTags", component: ListTags, meta: {auth: true} },
      { path: "/listImages", component: ListImages, meta: {auth: true} },
      { path: "/manageMsg", component: ManageMessage, meta: {auth: true} },
      { path: "/manageLink", component: ManageResLink, meta: {auth: true} },
      { path: "/manageEssay", component: ManageEssay, meta: {auth: true} }
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
  // 如果访问的不是管理员或登录页面，直接放行
  if (!to.meta.auth) {
      return next()
  }

  // 访问的是管理员页面，查看有没有token，如果有，则放行，否则，跳转到登录页
  const tokenStr = window.sessionStorage.getItem("token");
  if(!tokenStr) return next("/login");
  next();
})


export default router
