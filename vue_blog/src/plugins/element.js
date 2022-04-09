import Vue from 'vue'
import { Button, Form, FormItem, Input, Message, Container, Header,
     Aside, Main, Menu, MenuItem, Submenu, Breadcrumb, BreadcrumbItem,
     Card, Autocomplete, Row, Col, Checkbox, Option, Select,
     Table, TableColumn, MessageBox, Pagination, CheckboxGroup,
     Upload, Dialog } from 'element-ui'

Vue.use(Form)
Vue.use(Button)      //注册为全局可用的组件
Vue.use(FormItem)
Vue.use(Input)
Vue.use(Container)
Vue.use(Header)
Vue.use(Aside)
Vue.use(Main)
Vue.use(Menu)
Vue.use(MenuItem)
Vue.use(Submenu)
Vue.use(Breadcrumb)
Vue.use(BreadcrumbItem)
Vue.use(Card)
Vue.use(Autocomplete)
Vue.use(Row)
Vue.use(Col)
Vue.use(Checkbox)
Vue.use(Option)
Vue.use(Select)
Vue.use(Table)
Vue.use(TableColumn)
Vue.use(Pagination)
Vue.use(CheckboxGroup)
Vue.use(Upload)
Vue.use(Dialog)


Vue.prototype.$message = Message          //每个Vue对象可用通过this访问Message组件
Vue.prototype.$messageBox = MessageBox
