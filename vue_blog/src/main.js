import '@babel/polyfill'
import 'mutationobserver-shim'
import Vue from 'vue'
import './plugins/bootstrap-vue'
import App from './App.vue'
import router from './router'
import'./plugins/element.js'
import 'element-ui/lib/theme-chalk/index.css'
import 'element-ui/lib/theme-chalk/display.css'
import './assets/css/global.css'   //导入全局样式
import './plugins/echarts-vue.js'

import  mavonEditor from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'

import axios from "axios"
import "./plugins/common"

import 'highlight.js/styles/googlecode.css' //样式文件
import 'highlight.js/styles/a11y-dark.css'

import {bubbles} from 'vue-canvas-effect';


Vue.component('bubbles-effect', bubbles);

Vue.use(mavonEditor)


axios.defaults.baseURL = window.location.origin + "/api"


//配置请求拦截器，用于在访问后端服务器时携带token令牌
axios.interceptors.request.use(config =>{
  let requestUrl = config.url
  let start = requestUrl.indexOf("/admin")
  if (start == 0) {   // 如果访问后台，需要带上token
    config.headers.Authorization = window.sessionStorage.getItem("token")
  }

  return config           //必须return config
})

// 配置响应拦截器，该拦截器用于拦截后端数据的响应
axios.interceptors.response.use(config => {
  if(config.status == 500) {
    router.push("/internalError");
  } else if(config.status == 404) {
    router.push("/notFound");
  }

  if(config.data.status == 300) {    // 返回码为300时，说明需要登录，才能访问后台
    router.push("/login");
  }

  return config;
})

Vue.prototype.$axios = axios

Vue.config.productionTip = false



new Vue({
  router,
  render: h => h(App),
  beforeCreate() {
    Vue.prototype.$bus = this   // 安装全局事件总线
  }
}).$mount('#app')
