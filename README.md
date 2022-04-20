## 1、项目介绍

该项目采用Vue+MySQL+Golang+Gin框架编写，目前不支持移动端。

**vue_blog**：该目录是前端相关代码

**golang_web**：该目录是后端相关代码。

**deploy**：该目录是部署文件夹，部署的方式是Docker-Compose，使用nginx来做静态资源服务器以及反向代理服务器。



## 2、项目截图

`首页`

![image-20220409195326189](https://github.com/mangohow/myblog/blob/master/README.assets/image-20220409195326189.png)

![image-20220409195336045](https://github.com/mangohow/myblog/blob/master/README.assets/image-20220409195336045.png)

`文章标签`

![image-20220409195405913](https://github.com/mangohow/myblog/blob/master/README.assets/image-20220409195405913.png)

`归档`

![image-20220409195420221](https://github.com/mangohow/myblog/blob/master/README.assets/image-20220409195420221.png)

`后台管理`

![image-20220409195441357](https://github.com/mangohow/myblog/blob/master/README.assets/image-20220409195441357.png)

![image-20220409195452989](https://github.com/mangohow/myblog/blob/master/README.assets/image-20220409195452989.png)



## 3、项目部署

该项目提供了Docker-Compose的部署方式，在deploy中已经添加了相关的脚本文件，支持少量配置即可一键启动。

步骤如下：

1、将博客后端放入deploy/backend下，如果使用阿里云OSS需要修改conf/application.yaml中的配置

目录如下：

conf/  
└──  application.yaml

Dockerfile

myblog/  
├── conf  
├── db  
│&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;├── dao  
│&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;└── service  
├── images  
│&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;├── avatar  
│&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;├── blogImages  
│&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;├── firstPic  
│&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;└── icons  
├── logs  
├── model  
├── router  
│&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;└── admin  
└── utils  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;└── logger  

2、修改前端main.js的defaultUrl为你服务器的ip，将编译好的前端文件放入deploy/frontend/blog下

3、执行deploy下的start脚本

./start.sh serve [serverip] 

例如： ./start.sh serve http://192.168.44.100

​                                                                                                                   
