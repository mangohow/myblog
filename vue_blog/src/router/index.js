import Vue from 'vue'
import VueRouter from 'vue-router'


const Index = () => import( "../pages")
const Home = () => import( "../pages/Home.vue")
const BlogDetail = () => import( "../pages/BlogDetail.vue")
const Types = () => import( "../pages/Types.vue")
const Tags = () => import( "../pages/Tags.vue")
const TimeLine = () => import( "../pages/TimeLine.vue")
const About = () => import( "../pages/About.vue")
const Essay = () => import( "../pages/Essay")
const MsgBoard = () => import( "../pages/MsgBoard")
const PhotoWall = () => import( "../pages/PhotoWall")
const ResourceLib = () => import( "../pages/ResourceLib")
const NotFound = () => import( "../components/error/NOTFOUND.vue")
const InternalError = () => import( "../components/error/INTERNALERROR.vue")


const Login = () => import("../pages/admin/Login")
const AdminHome  = () => import('../pages/admin/Home')
const Welcome  = () => import( "../pages/admin/Welcome.vue")
const ListBlogs  = () => import( "../pages/admin/blogs/ListBlogs.vue")
const ListTypes  = () => import( "../pages/admin/types/ListTypes.vue")
const ListTags  = () => import( "../pages/admin/tags/ListTags.vue")
const AddBlog  = () => import( "../pages/admin/blogs/AddBlog.vue")
const ListImages  = () => import( "../pages/admin/images/ListImages")
const ManageMessage  = () => import( "../pages/admin/messages/ManageMessage")
const ListMottos  = () => import( "../pages/admin/blogs/listMottos")
const ManageResLink  = () => import( "../pages/admin/resourecLink/manageResLink")
const ManageEssay  = () => import( "../pages/admin/essay/ManageEssay")

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
