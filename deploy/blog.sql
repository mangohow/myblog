/*
 Navicat MySQL Data Transfer

 Source Server         : centos_mysql
 Source Server Type    : MySQL
 Source Server Version : 80027
 Source Host           : 192.168.44.100:3306
 Source Schema         : blog

 Target Server Type    : MySQL
 Target Server Version : 80027
 File Encoding         : 65001

 Date: 11/04/2022 21:17:01
*/

CREATE DATABASE IF NOT EXISTS blog;

USE blog;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for t_bgimage
-- ----------------------------
DROP TABLE IF EXISTS `t_bgimage`;
CREATE TABLE `t_bgimage`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `url` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '图片url',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 19 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of t_bgimage
-- ----------------------------
INSERT INTO `t_bgimage` VALUES (1, 'http://192.168.44.100/images/firstPic/wallhaven-v9vv3p_1647674711.png');
INSERT INTO `t_bgimage` VALUES (2, 'http://192.168.44.100/images/firstPic/wallhaven-l37w6q_1647674789.jpg');
INSERT INTO `t_bgimage` VALUES (3, 'http://192.168.44.100/images/firstPic/wallhaven-72ozj3_1647674771.png');
INSERT INTO `t_bgimage` VALUES (4, 'http://192.168.44.100/images/firstPic/bg24.jpg');
INSERT INTO `t_bgimage` VALUES (5, 'http://192.168.44.100/images/firstPic/bg23.jpg');
INSERT INTO `t_bgimage` VALUES (6, 'http://192.168.44.100/images/firstPic/bg21.jpg');
INSERT INTO `t_bgimage` VALUES (7, 'http://192.168.44.100/images/firstPic/bg20.jpg');
INSERT INTO `t_bgimage` VALUES (8, 'http://192.168.44.100/images/firstPic/bg19.jpg');
INSERT INTO `t_bgimage` VALUES (9, 'http://192.168.44.100/images/firstPic/bg17.png');
INSERT INTO `t_bgimage` VALUES (10, 'http://192.168.44.100/images/firstPic/bg14.jpg');
INSERT INTO `t_bgimage` VALUES (11, 'http://192.168.44.100/images/firstPic/bg13.png');
INSERT INTO `t_bgimage` VALUES (12, 'http://192.168.44.100/images/firstPic/bg12.png');
INSERT INTO `t_bgimage` VALUES (13, 'http://192.168.44.100/images/firstPic/bg9.jpg');
INSERT INTO `t_bgimage` VALUES (14, 'http://192.168.44.100/images/firstPic/bg8.jpg');
INSERT INTO `t_bgimage` VALUES (15, 'http://192.168.44.100/images/firstPic/bg5.jpg');
INSERT INTO `t_bgimage` VALUES (16, 'http://192.168.44.100/images/firstPic/bg3.jpg');
INSERT INTO `t_bgimage` VALUES (17, 'http://192.168.44.100/images/firstPic/bg2.png');
INSERT INTO `t_bgimage` VALUES (18, 'http://192.168.44.100/images/firstPic/bg1.png');
INSERT INTO `t_bgimage` VALUES (19, 'http://192.168.44.100/images/firstPic/200_1647675014.jpg');

-- ----------------------------
-- Table structure for t_blog
-- ----------------------------
DROP TABLE IF EXISTS `t_blog`;
CREATE TABLE `t_blog`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `title` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '博客标题',
  `content` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '博客内容',
  `first_picture` varchar(512) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '博客首图',
  `flag` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '原创、转载、翻译等',
  `views` int UNSIGNED NOT NULL COMMENT '浏览次数',
  `appreciation` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否开启赞赏',
  `share_statement` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否可以转载',
  `commentabled` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否可以评论',
  `published` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否发布',
  `recommend` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否推荐',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `type_id` bigint NOT NULL COMMENT '博客类型id',
  `user_id` bigint NOT NULL COMMENT '用户id',
  `description` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '博客描述',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_type_id`(`type_id`) USING BTREE COMMENT '类型id索引'
) ENGINE = InnoDB AUTO_INCREMENT = 47 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of t_blog
-- ----------------------------
INSERT INTO `t_blog` VALUES (1, 'maven中静态资源的过滤', '# maven中静态资源的过滤\r\n\r\npom.xml文件中加入下面配置\r\n\r\n### 可以过滤java和resources文件夹里面所有的的.properties和.xml文件\r\n**directory：指定资源所在的目录，目录的路径是相对于pom.xml文件的\r\nincludes：指定要包含哪些文件**\r\n**filtering标签中：false表示不过滤，true表示过滤**\r\n\r\n```java\r\n    <build>\r\n        <resources>\r\n            <resource>\r\n                <directory>src/main/java</directory>\r\n                <includes>\r\n                    <include>**/*.properties</include>\r\n                    <include>**/*.xml</include>\r\n                </includes>\r\n                <filtering>false</filtering>\r\n            </resource>\r\n            <resource>\r\n                <directory>src/main/resources</directory>\r\n                <includes>\r\n                    <include>**/*.properties</include>\r\n                    <include>**/*.xml</include>\r\n                </includes>\r\n                <filtering>false</filtering>\r\n            </resource>\r\n        </resources>\r\n    </build>\r\n```\r\n', 'https://img-blog.csdnimg.cn/ad219a20e5ac401ba091a0ea3bfb9863.png?x-oss-process=image/watermark,type_d3F5LXplbmhlaQ,shadow_50,text_Q1NETiBAQ29yZUR1bXDkuLY=,size_20,color_FFFFFF,t_70,g_se,x_16#pic_center', '原创', 148, 1, 1, 1, 1, 1, '2020-02-13 08:25:47', '2022-04-07 11:34:03', 8, 1, 'pom.xml文件中加入下面配置\r\n可以过滤java和resources文件夹里面所有的的.properties和.xml文件');
INSERT INTO `t_blog` VALUES (2, 'Springboot中PageHelper 分页查询使用方法', '### 一：导入依赖\n\n```java\n<dependency>\n	<groupId>com.github.pagehelper</groupId>\n	<artifactId>pagehelper-spring-boot-starter</artifactId>\n	<version>1.2.13</version>\n</dependency>\n```\n### 二：配置yml文件中PageHelper的属性\n\n```java\npagehelper:                #分页插件\n  helper-dialect: mysql\n  reasonable: true\n  support-methods-arguments: true\n  params:\n```\n### 三：在controller类中使用，\n##### 1.在查询方法上调用PageHelper.startPage()方法，设置分页的页数和每页信息数，\n##### 2.将查询出来的结果集用PageInfo的构造函数初始化为一个分页结果对象\n##### 3.将分页结果对象存入model，返回前端页面\n```java\n@GetMapping(\"/types\")\npublic String types(@RequestParam(required = false,defaultValue = \"1\",value = \"pagenum\")int pagenum, Model model){\n\n    PageHelper.startPage(pagenum, 5);  //pagenum：页数，pagesize:每页的信息数\n    \n    List<Type> allType = typeService.getAllType(); //调用业务层查询方法\n    \n    PageInfo<Type> pageInfo = new PageInfo<>(allType); //得到分页结果对象\n    \n    model.addAttribute(\"pageInfo\", pageInfo);  //携带分页结果信息\n    \n    return \"admin/types\";  //回到前端展示页面\n}\n```\n### 四：前端展示分页（thymeleaf+semantic-ui）,这里ui用自己的就行，比如bootstrap或layui，主要是thymeleaf的使用。\n\n```java\n<table  class=\"ui compact celled teal table\">\n  <thead>\n  <tr>\n    <th></th>\n    <th>名称</th>\n    <th>操作</th>\n  </tr>\n  </thead>\n  <tbody>\n  <tr th:each=\"type, iterStat : ${pageInfo.list}\">\n    <td th:text=\"${iterStat.count}\">1</td>\n    <td th:text=\"${type.name}\">摸鱼方法</td>\n    <td>\n      <a href=\"#\" th:href=\"@{/admin/types/{id}/input(id=${type.id})}\" class=\"ui mini teal basic button\">编辑</a>\n      <a href=\"#\" th:href=\"@{/admin/types/{id}/delete(id=${type.id})}\" class=\"ui mini red basic button\">删除</a>\n    </td>\n  </tr>\n  </tbody>\n  <tfoot>\n  <tr>\n    <th colspan=\"7\">\n      <div class=\"ui mini pagination menu\"  >\n        <div class=\"item\"><a th:href=\"@{/admin/types}\">首页</a></div>\n        <div class=\"item\"><a th:href=\"@{/admin/types(pagenum=${pageInfo.hasPreviousPage}?${pageInfo.prePage}:1)}\">上一页</a></div>\n        <div class=\"item\"><a th:href=\"@{/admin/types(pagenum=${pageInfo.hasNextPage}?${pageInfo.nextPage}:${pageInfo.pages})}\">下一页</a></div>\n        <div class=\"item\"><a th:href=\"@{/admin/types(pagenum=${pageInfo.pages})}\">尾页</a></div>\n      </div>\n      <a href=\"#\" th:href=\"@{/admin/types/input}\" class=\"ui mini right floated teal basic button\">新增</a>\n    </th>\n  </tr>\n  </tfoot>\n</table>\n\n<div class=\"ui segment m-inline-block\">\n  <p >当前第<span th:text=\"${pageInfo.pageNum}\"></span>页，总<span th:text=\"${pageInfo.pages}\"></span>页，共<span th:text=\"${pageInfo.total}\"></span>条记录</p>\n</div>\n```\n\n', 'http://192.168.44.100/images/firstPic/34_1649229836.jpg', '原创', 27, 1, 0, 1, 1, 1, '2021-02-01 08:25:47', '2022-04-06 15:23:59', 7, 1, '1.在查询方法上调用PageHelper.startPage()方法，设置分页的页数和每页信息数，\r\n2.将查询出来的结果集用PageInfo的构造函数初始化为一个分页结果对象\r\n3.将分页结果对象存入model，返回前端页面');
INSERT INTO `t_blog` VALUES (3, 'mybatis根据日期查询、查询日期并返回', '### 方法：\r\n\r\n#### 1.查询日期，返回日期字符串\r\n```handlebars\r\n//查询所有博客对应的年份，返回一个集合\r\nList<String> findGroupYear();  \r\n```\r\n#### 2.根据日期查询\r\n```handlebars\r\n//根据年份查询博客信息\r\nList<Blog> findByYear(@Param(\"year\") String year);  \r\n```\r\n\r\n\r\n### mybatis映射:\r\n\r\n\r\n```java\r\n<select id=\"findGroupYear\" resultType=\"String\">\r\n    select DATE_FORMAT(b.update_time, \'%Y\') from t_blog b\r\n</select>\r\n\r\n\r\n<select id=\"findByYear\" resultType=\"Blog\">\r\n    select b.title, b.update_time, b.id, b.flag\r\n    from t_blog b\r\n    where DATE_FORMAT(b.update_time, \'%Y\') = #{year}\r\n</select>\r\n```\r\n\r\n## 总结：\r\n**DATE_FORMAT(date,format)：date表示日期，format表示显示的格式。**\r\n**该方法可以把date类型转换为String类型的字符串**\r\n', 'http://p.qpic.cn/dnfbbspic/0/dnfbbs_dnfbbs_dnf_gamebbs_qq_com_forum_202002_04_082156ifotspmtuzcffycn.jpg/0', '原创', 11, 1, 1, 1, 1, 1, '2021-02-01 08:25:47', '2021-02-05 11:56:19', 3, 1, 'mybatis根据日期查询、查询日期并返回\r\n方法：\r\n1.查询日期，返回日期字符串');
INSERT INTO `t_blog` VALUES (4, 'MavonEditor上传图片', '# mavonEditor上传图片	\n\n## 前端设置\n\n\n\n**在<mavon-editor/>标签中设置@imgAdd事件，添加ref=md 以在imgAdd中获取mavonEditor对象。**\n\n![image20210327225526012.png](http://192.168.44.100/images/test.png)\n\n**在imgAdd中上传图片到后端服务器，每次粘贴图片到markdown编辑器中时就会触发imgAdd事件**\n\n```javascript\n//图片上传\n        // 绑定@imgAdd event\n        $imgAdd(pos, $file){\n            // 第一步.将图片上传到服务器.\n           var formdata = new FormData();\n           formdata.append(\'image\', $file);\n           this.$http({\n               url: \'/saveImages\',\n               method: \'post\',\n               data: formdata,\n               headers: { \'Content-Type\': \'multipart/form-data\' },\n           }).then((url) => {\n               // 第二步.将返回的url替换到文本原位置![...](./0) -> ![...](url)\n               // $vm.$img2Url 详情见本页末尾\n               const serverUrl = \"http://192.168.44.100/\";\n               var imageUrl = serverUrl +url.data;\n               this.$refs.md.$img2Url(pos, imageUrl);    //this.$refs.md为mavonEditor对象\n           })\n        }\n```\n\n\n\n## 后端服务器处理\n\n**后端服务器接收到图片，将图片保存在/resources/static/images/中，并返回访问该文件的url**\n\n```java\nprivate final String IMAGE_PATH = \"src/main/resources/static/images/\";\n    private final String IMAGE_URL = \"images/\";\n\n    @RequestMapping(value = \"/saveImages\", method = RequestMethod.POST)\n    @ResponseBody\n    public String saveImages(@RequestParam(\"image\") MultipartFile image) throws IOException {\n        String timeStr = String.valueOf(System.currentTimeMillis());\n        String imagePath = IMAGE_PATH + timeStr + image.getOriginalFilename();\n        String imageUrl = IMAGE_URL + timeStr + image.getOriginalFilename();\n        InputStream inputStream = image.getInputStream();\n        File file = new File(imagePath);\n        OutputStream os = new FileOutputStream(file);\n        int bytesRead = 0;\n        byte[] buffer = new byte[8192];\n        while ((bytesRead = inputStream.read(buffer, 0, 8192)) != -1) {\n            os.write(buffer, 0, bytesRead);\n        }\n\n        os.close();\n        inputStream.close();\n\n        return imageUrl;\n    }\n```\n\n**需要配置静态资源处理器，否则无法访问静态资源**\n\n```java\n@Configuration\npublic class WebConfig implements WebMvcConfigurer {\n    @Override\n    public void addResourceHandlers(ResourceHandlerRegistry registry) {\n//        registry.addResourceHandler(\"/images/**\").addResourceLocations(\"classpath:/static/images/\");\n        registry.addResourceHandler(\"/images/**\").addResourceLocations(\"file:E:\\\\java_workspace\\\\projectes\\\\vue_springboot_blog\\\\src\\\\main\\\\resources\\\\static\\\\images\\\\\");\n\n    }\n}\n\n/* \n	使用相对路径进行配置时，上传的图片无法立即访问，需要重启服务器才能访问到，不知为何...\n	查询百度，发现很多博客都是使用的绝对路径，因此决定一试，使用了绝对路径就可以了，不知为何...\n*/\n```\n\n', 'http://192.168.44.100/images/firstPic/41_1649229816.png', '原创', 75, 1, 0, 1, 1, 0, '2021-03-27 23:19:51', '2022-04-06 15:23:39', 1, 1, 'mavonEditor上传图片');
INSERT INTO `t_blog` VALUES (5, '前后端分离登录验证', '\n## 一、前端处理	\n\n### 1.前端从服务器获取token\n\n**在用户第一次登录服务器时，服务器会返回该用户登录的token，客户端收到token后将其保存至sessionStorage中，token只应在当前网站打开期间生效，而sessionStorage 是会话期间的存储，所以将Token保存到客户端的sessionStorage中。**\n\n**在浏览器的Application中可以产看sessionStorage**\n\n![image20210325223755742.png](http://192.168.44.100/images/1616903611540image-20210325223755742.png)\n\n\n\n### 2.前端挂载路由导航守卫\n\n**路由导航守卫的作用是在访问页面时，先查看token是否存在，如果存在则让行，否则转至登录页面，完成客户端的token校验。**\n\n```javascript\nrouter/index.js\n\n//挂载路由导航守卫\nrouter.beforeEach((to, from, next) => {\n  // to: 将要访问的路径\n  // form: 从哪个路径跳转而来\n  // next 是一个函数，表示放行\n  // next() 放行 next(\"/login\") 强制跳转\n\n  if(to.path == \"/login\") {\n    return next();\n  }\n\n  // 获取token\n  const tokenStr = window.sessionStorage.getItem(\"token\");\n  if(!tokenStr) return next(\"/login\");\n  next();\n})\n\n```\n\n### 3. 前端配置拦截器\n\n**在访问除了登录页面的其他页面时，需要携带token到服务器进行验证。因此可以配置一个请求拦截器在每次发起请求前，在header中加入token。**\n\n\n\n```javascript\nmain.js\n\n// 配置请求拦截器，携带token令牌\naxios.interceptors.request.use(config =>{    \n  config.headers.Authorization = window.sessionStorage.getItem(\"token\")\n\n  return config           //必须return config\n})\n```\n\n**后端服务器验证token发生异常时，会返回状态码为-1的ResponseResult，因此需要配置响应拦截器，数据返回后先通过响应拦截器，在响应拦截器中判断状态码如果为-1，则跳转至登录页面。**\n\n```javascript\n// 配置响应拦截器\naxios.interceptors.response.use(config =>{\n  if(config.data.status === -1) {\n    router.push(\"/login\");\n  } \n\n  return config\n})\n```\n\n\n\n## 二、后端处理	\n\n### 1、登录处理\n\n**首先需要一个生成token的工具类。**\n\n```java\npublic class JwtUtils {\n\n    final static String base64EncodedSecretKey = \"jwt_secret\";//私钥\n\n    final static long TOKEN_EXP = 1000 * 60 * 10;//过期时间,测试使用十分钟\n\n    /**\n     * xxx.xxx.xxx\n     * 头部：放有签名算法和令牌类型（这个就是JWT）\n     * 载荷：在令牌上附带的信息：比如用户的id，用户的电话号码，\n     * 这样以后验证了令牌之后就可以直接从这里获取信息而不用再查数据库了\n     * 签名：用来加令牌的\n     */\n    public static String createToken(String userId, String userName) {\n\n        Date expiresDate = new Date(System.currentTimeMillis() + TOKEN_EXP);\n        Algorithm algorithm = Algorithm.HMAC256(base64EncodedSecretKey);\n\n        return JWT.create().withAudience(userId)   //签发对象\n                .withIssuedAt(new Date())          //发行时间\n                .withExpiresAt(expiresDate)        //有效时间\n                .withClaim(\"username\", userName)    //载荷，随便写几个都可以\n                .sign(algorithm);   //加密\n    }\n\n    /**\n     * 检验合法性，\n     * @param token\n     */\n    public static void verifyToken(String token) /*throws TokenUnavailable*/ {\n        DecodedJWT jwt = null;\n        try {\n            JWTVerifier verifier = JWT.require(Algorithm.HMAC256(base64EncodedSecretKey)).build();\n            jwt = verifier.verify(token);\n        } catch (Exception e) {\n            //效验失败\n            //这里抛出的异常是我自定义的一个异常，你也可以写成别的\n            throw new TokenUnavailable();\n        }\n    }\n\n    /**\n     * 获取签发对象\n     */\n    public static String getAudience(String token)/* throws TokenUnavailable */{\n        String audience = null;\n        try {\n            audience = JWT.decode(token).getAudience().get(0);\n        } catch (JWTDecodeException j) {\n            //这里是token解析失败\n            throw new TokenUnavailable();\n        }\n        return audience;\n    }\n\n\n    /**\n     * 通过载荷名字获取载荷的值\n     */\n    public static Claim getClaimByName(String token, String name){\n        return JWT.decode(token).getClaim(name);\n    }\n\n```\n\n\n\n**每次登录后生成token，并将token发送给前端验证**\n\n```java\n\n    @RequestMapping(value = \"/login\", method = RequestMethod.POST)\n    @ResponseBody           //返回json数据格式\n    public ResponseResult<Token> login(@RequestBody User loginUser) {\n        System.out.println(loginUser);\n\n        String codedPassword = MD5Utils.code(loginUser.getPassword());\n        User user = userService.checkUser(loginUser.getUsername(), codedPassword);\n        ResponseResult<Token> responseResult = new ResponseResult<Token>();\n        if (user == null) {\n            responseResult.setStatus(ResponseResult.FAIL);\n            responseResult.setMessage(\"登录失败,请重新登录!\");\n\n        } else {\n            // 生成token\n            String tokenStr = JwtUtils.createToken(user.getId() + \"\", user.getUsername());\n            Token token = new Token(tokenStr, user.getId());\n            responseResult.setData(token);\n            responseResult.setMessage(\"登录成功!\");\n        }\n\n        return responseResult;\n    }\n```\n\n\n\n### 2、配置拦截器\n\n```java\n@Configuration\npublic class WebConfig implements WebMvcConfigurer {\n\n    @Bean\n    public JwtAuthenticationInterceptor authenticationInterceptor() {\n        return new JwtAuthenticationInterceptor();\n    }\n\n    @Override\n    public void addInterceptors(InterceptorRegistry registry) {\n        registry.addInterceptor(authenticationInterceptor())   //如果直接new JwtAuthenticationInterceptor()\n                .addPathPatterns(\"/admin/**\")                  // 会导致JwtAuthenticationInterceptor中的\n                .excludePathPatterns(\"/admin/login\");          // UserService注入失败\n    }\n}\n```\n\n\n\n```java\npublic class JwtAuthenticationInterceptor implements HandlerInterceptor {\n\n    @Autowired\n    private UserService userService;\n\n    @Override\n    public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object handler) throws Exception {\n        String token = request.getHeader(\"Authorization\");\n\n        // 如果不是映射到方法直接通过\n        if (!(handler instanceof HandlerMethod)) {\n            return true;\n        }\n\n        // 验证token\n        if(token == null) {\n            throw new NeedToLogin();\n        }\n        JwtUtils.verifyToken(token);\n\n        // 获取token中的userId\n        String userId = JwtUtils.getAudience(token);\n        // 获取username\n        Claim claim = JwtUtils.getClaimByName(token, \"username\");\n        String username = claim.asString();\n\n        Integer id = Integer.parseInt(userId);\n        System.out.println(userService);\n        // 从数据库中查询\n        String findUsername = userService.findById(id);\n        if(findUsername == null || !findUsername.equals(username)) {\n            throw new UserNotExist();\n        }\n\n        return true;\n    }\n}\n```\n\n\n\n### 3、配置异常处理器\n\n**验证token时可能会抛出UserNotExist、TokenUnavailable以及NeedToLogin等异常，需要对异常进行处理，捕获到该类异常后返回状态码为-1的消息。**\n\n```java\n\n@ControllerAdvice   //拦截所有controller抛出的异常，对异常进行统一的处理\npublic class ControllerExceptionHandler {\n\n    private final Logger logger = LoggerFactory.getLogger(this.getClass());\n\n    @ExceptionHandler(Exception.class)  //表示该方法可以处理所有类型异常\n    @ResponseBody\n    public ResponseResult<Boolean> exceptionHandler(HttpServletRequest request, Exception e) throws Exception {\n\n        //日志打印异常信息\n        logger.error(\"Request url: {}, Exception: {}\", request.getRequestURI(), e);\n\n        //不处理带有ResponseStatus注解的异常\n        if (AnnotationUtils.findAnnotation(e.getClass(), ResponseStatus.class) != null) {\n            throw e;\n        }\n\n        ResponseResult<Boolean> res = new ResponseResult<>(null, ResponseResult.UNAUTHORIZED, \"未获得授权\");\n\n        return res;\n    }\n}\n\n```\n\n', 'http://192.168.44.100/images/firstPic/33_1649229800.jpg', '原创', 599, 1, 0, 1, 1, 0, '2021-03-28 00:47:35', '2022-04-06 15:23:25', 1, 1, 'Vue登录验证');
INSERT INTO `t_blog` VALUES (6, 'Vue日期格式化过滤器', '# Vue日期格式化过滤器\n\n\n\n```java\nfilters: {\n        fromatDate: function(val, fmt) {\n            var date = new Date(val);\n            let ret;\n            const opt = {\n                \"y+\": date.getFullYear().toString(),        // 年\n                \"M+\": (date.getMonth() + 1).toString(),     // 月\n                \"d+\": date.getDate().toString(),            // 日\n                \"h+\": date.getHours().toString(),           // 时\n                \"m+\": date.getMinutes().toString(),         // 分\n                \"s+\": date.getSeconds().toString()          // 秒\n                // 有其他格式化字符需求可以继续添加，必须转化成字符串\n            };\n            for (let k in opt) {\n                ret = new RegExp(\"(\" + k + \")\").exec(fmt);\n                if (ret) {\n                    fmt = fmt.replace(ret[1], (ret[1].length == 1) ? (opt[k]) : (opt[k].padStart(ret[1].length, \"0\")))\n                };\n            };\n            return fmt;\n        }\n    },\n```\n\n\n\n### 使用方式:\n\n```javascript\n{{ item.createTime | fromatDate(\"yyyy-MM-dd hh:mm:ss\") }}\n```\n\n', 'http://192.168.44.100/images/firstPic/27_1649229768.png', '原创', 7, 1, 0, 1, 1, 0, '2021-04-09 21:16:33', '2022-04-06 15:23:03', 1, 1, 'vue中需要对日期进行格式化，这时可以使用过滤器来进行');
INSERT INTO `t_blog` VALUES (7, 'MarkDownit使用', '# MarkDownit使用\n\n**使用markdown编辑器生成的文本为markdown格式的，在展示时需要将markdown格式的文本转换为Html格式的，因此需要使用markdownit插件进行转换。**\n\n## 安装\n\n```javascript\nnpm install markdown-it --save\n```\n\n\n\n## 使用\n\n```js\nvar hljs = require(\'highlight.js\');\nvar md = require(\'markdown-it\')({\n    html: true,\n    linkify: true,\n    typographer: true,\n    highlight: function (str, lang) {\n        if (lang && hljs.getLanguage(lang)) {\n            try {\n                return \'<pre class=\"hljs\"><code>\' +\n                    hljs.highlight(lang, str, true).value +\n                    \'</code></pre>\';\n            } catch (__) {}\n        }\n\n        return \'<pre class=\"hljs\"><code>\' + md.utils.escapeHtml(str) + \'</code></pre>\';\n    }\n\n});\n\nthis.blog.content = md.render(this.blog.content);\n```\n\n**若要使用语法高亮，必须在main.js中导入css样式文件，否则无法语法高亮**\n\n```js\nimport \'highlight.js/styles/googlecode.css\' //样式文件\nimport \'highlight.js/styles/a11y-dark.css\'\n```\n\n', 'https://img-blog.csdnimg.cn/3ca9392481454b499b41de4538d6c3a0.png?x-oss-process=image/watermark,type_d3F5LXplbmhlaQ,shadow_50,text_Q1NETiBAQ29yZUR1bXDkuLY=,size_20,color_FFFFFF,t_70,g_se,x_16#pic_center', '原创', 7, 1, 0, 1, 1, 0, '2021-04-09 21:18:01', '2021-04-09 21:18:47', 1, 1, '\n使用markdown编辑器生成的文本为markdown格式的，在展示时需要将markdown格式的文本转换为Html格式的，因此需要使用markdownit插件进行转换。......\n');
INSERT INTO `t_blog` VALUES (44, 'Docker学习笔记——Docker基础', '\n\n\n\n## 一、Dokcer简介\n\n**官方介绍：**\n\n​		&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Docker 是一个开源的应用容器引擎，基于Go 语言并遵从 Apache2.0 协议开源。\n\n​		&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Docker 可以让开发者打包他们的应用以及依赖包到一个轻量级、可移植的容器中，然后发布到任何流行的 Linux 机器上，也可以实现虚拟化。\n\n​		&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;容器是完全使用沙箱机制，相互之间不会有任何接口（类似 iPhone 的 app）,更重要的是容器性能开销极低。\n\n​		&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Docker的构想是要实现“Build, Ship and Run Any App, Anywhere”，即通过对应用的封装（Packaging）、分发（Distribution）、部署（Deployment）、运行生命周期（Runtime）进行管理，达到应用组件级别的“一次封装，到处运行”。\n\n**本人理解：**\n\n​		&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;我们在开发到上线一个程序的时候需要接触很多环境，比如：开发环境（一般是在自己电脑）、测试环境、生产环境，在不同的环境中我们开发出的程序可能出现一些问题。可能在自己电脑上可以跑的程序，到了其它机器上就会因为环境不同而出现一些问题。而且换到不同的机器上我们就需要重新配置环境。例如下图所示。所以我们就想，如果能把环境一起打包运行，那样不就可以解决不同运行环境的问题了。使用Docker就可以解决这样的问题，我们可以使用Docker将程序以及环境打包为一个镜像，并上传到远程仓库，当需要在其它机器上运行时，只需要将镜像从远程仓库拉取下来直接运行即可。\n\n![image.png](http://47.99.200.62/images/blogImages/image_1647009087_resized.png)\n\n\n​		&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Docker是一个C/S架构的软件，Docker会启动一个守护进程用来接收客户端的命令并执行相应的命令。比如：docker run 镜像名，可以根据一个镜像运行一个容器。\n\n### 1、Docker中的名词概念\n\n**镜像（Image）：**\n\n​		&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Docker中的镜像就像是一个只读的模板，比如操作系统镜像，我们可以使用操作系统镜像来安装一个操作系统；又像是面向对象中的类，通过类我们可以创建一个对象。镜像是可以复用的，我们可以通过Docker运行一个镜像实例，运行起来的镜像就是一个容器。\n\n**容器（Container）：**\n\n​		&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Docker容器类似一个轻量级的沙箱，Docker利用容器技术来运行和隔离应用。容器是从镜像创建的应用运行实例。它可以启动、开始、停止、删除，而这些容器都是彼此相互隔离、互不可见的。\n\n​		&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;可以把容器看作一个简易版的Linux系统环境（包括root用户权限、进程空间、用户空间和网络等）以及运行在其中的应用程序打包而成的盒子。\n\n**仓库（Repository）**\n\n​		&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Docker仓库类似代码仓库，是Docker集中存放镜像文件的仓库，仓库分为私有仓库和公有仓库。可以通过Docker从仓库中拉取一个镜像，也可以上传自己的镜像。目前最大的公开仓库是官方提供的Docker Hub，就像Github一样，Docker Hub存放了很多官方的或者非官方的镜像，例如：nginx镜像、redis镜像等。官网：https://hub.docker.com/\n\n&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;如下图，下图中包含了三部分，分别时Docker客户端、Docker守护进程以及远程仓库。Docker客户端可以接收命令并通知Docker守护进程来进行一系列的操作，比如拉取镜像、运行容器等。Docker守护进程将会管理镜像以及容器。远程仓库存放了很多种不同的镜像。\n\n![image.png](http://47.99.200.62/images/blogImages/image_1647009110_resized.png)\n\n\nDocker的安装：[Ubuntu20.04安装Docker_Peerless__的博客-CSDN博客](https://blog.csdn.net/Peerless__/article/details/120906200)\n\n## 二、Docker基本使用\n\n**在使用下面的命令时，可以加参数 --help 来查看帮助文档。例如dokcer run --help:**\n\n下面显示了docker run命令的使用方式以及其它的可选参数和其解释。\n\n![image.png](http://47.99.200.62/images/blogImages/image_1647009127_resized.png)\n\n\n### 1、镜像\n\n#### 1.1 查看镜像\n\n```shell\ndocker images \n```\n\n![image.png](http://47.99.200.62/images/blogImages/image_1647009142.png)\n\n使用docker images命令可以查看本机已经存在的镜像。分别显示了镜像名、标签（版本信息）、镜像ID、创建时间以及镜像的大小。\n\n除此之外，还有一些可选参数，可以通过docker images --help 命令来查看，下面两个是比较常用的参数：\n\n```shell\n# -a参数，显示所有的镜像 docker images -a\n-a, --all             Show all images (default hides intermediate images)\n# -q参数，只显示镜像ID\n-q, --quiet           Only show image IDs\n```\n\n#### 1.2 给镜像设置标签\n\n```shell\ndocker tag sourname:tag  desname:tag\n#例如：\ndocker tag ubuntu:18.04 myubuntu:lastest\n```\n\n#### 1.3 搜索镜像\n\n```shell\ndocker search 镜像名\n\n#参数：\n# 过滤输出内容\n  -f, --filter filter   Filter output based on conditions provided\n# 限制输出个数\n  --limit int       Max number of search results (default 25)\n  \n# 例如\ndocker search --filter=is-official=true mysql           # 搜索mysql官方镜像 \ndocker search --filter=stars=200 mysql                  # 搜索star大于200的mysql镜像\ndocker search --limit 5 mysql                           # 只显示5个输出\n```\n\n![image.png](http://47.99.200.62/images/blogImages/image_1647009154.png)\n\n使用docker search可以从远程仓库搜索镜像，上面显示了镜像名、描述、stars、是否是官方镜像等信息。\n\n#### 1.4 拉取镜像\n\n```shell\ndocker pull 镜像名\ndocker pull 镜像名:tag\n\n#例如 \ndocker pull mysql             # 不加tag，将会拉取最新版本\ndocker pull mysql:5.7         # 拉取5.7版本的mysql\n```\n\n在镜像名后可以指定拉取的版本，如果不指定就会拉取最新版本，关于镜像的版本可以到docker hup来搜索。网址：https://hub.docker.com/\n\n![image.png](http://47.99.200.62/images/blogImages/image_1647009216_resized.png)\n\n\n#### 1.5 删除镜像\n\n使用下面命令可以删除本机仓库中的一个镜像，rmi ==> remove image\n\n```shell\ndocker rmi 镜像名 或 镜像ID\n\n#例如 docker rmi centos   或 docker rmi 9f266d35e02c , 命令Docker images可以查看镜像ID\n\n参数：\n# 强制删除，类似linux命令 rm -f \n-f, --force      Force removal of the image\n\n# 如果要删除所有的镜像可以使用以下命令, docker images -aq 会显示所有镜像的ID\ndocker rmi `docker images -aq`    或\ndocker rmi $(docker images -aq)\n```\n\n#### 1.6 导出和载入镜像\n\nDocker镜像的save和load命令可以将镜像导出到本地以及加载本地镜像到镜像库。使用save命令可以将镜像导出到本地，这样就可以跟别人分享自己的镜像，当然也可以提交到远程仓库，将在后面介绍。通过load命令则可以加载别人分享的镜像。\n\n```shell\ndocker save -o 生成的镜像名 镜像\n# 例如\ndocker save -o mysql5_7.tar mysql:5.7\n```\n\n```shell\ndocker load -i 本地镜像名  或者\ndocker load < 本地镜像名\n```\n\n\n\n### 2、容器\n\n#### 2.1 运行容器\n\n```shell\ndocker run 镜像名\n\n参数：\n  -d, --detach               Run container in background and print container ID  后台运行容器\n  -e, --env list             Set environment variables                           设置环境变量\n  --expose list              Expose a port or a range of ports                   暴露端口\n  -i, --interactive          Keep STDIN open even if not attached                交互模式\n  --name string              Assign a name to the container                      给当前运行的容器指定名字\n  -p, --publish list         Publish a container\'s port(s) to the host           指定端口映射\n  -P, --publish-all          Publish all exposed ports to random ports           随机指定暴露的端口映射\n  -t, --tty                  Allocate a pseudo-TTY                               分配一个终端\n  -v, --volume list          Bind mount a volume                                 挂载数据卷\n  --volumes-from list        Mount volumes from the specified container(s)       从一个数据卷容器挂载目录\n  -w, --workdir string       Working directory inside the container              指定容器中的工作目录\n```\n\n我们可以使用交互模式来运行一个容器，例如下面命令：\n\n```shell\ndocker run -it --name c1 centos /bin/bash\n```\n\n![image.png](http://47.99.200.62/images/blogImages/image_1647009235.png)\n\n\n-t选项让Docker分配一个伪终端并绑定到容器的标准输入上，-i则让容器的标准输入保持打开，--name给容器指定名字，centos为创建容器使用的镜像，/bin/bash为容器中运行的终端。如果镜像不存在，Docker会从远程仓库拉取，然后再运行。如下图，运行容器后发现终端发生了变化，说明我们已经进入了容器，用户为root用户，@后的一串字符为当前容器的ID。可以使用ls命令来查看当前目录下的文件。\n\n![image.png](http://47.99.200.62/images/blogImages/image_1647009250.png)\n\n输入exit命令可以退出当前容器。\n\n使用docker ps -a来查看所有容器，发现刚刚运行的容器在退出后就停止运行了。对于所创建的bash容器，当用户使用exit命令退出bash进程后，容器也会自动退出。这是因为对于容器来说，当其中的应用退出后，容器的使命完成，也就没有运行的必要了。\n\n![image.png](http://47.99.200.62/images/blogImages/image_1647009263.png)\n\n\n使用 Ctrl + p + q来退出容器可以让容器在后台继续运行。\n\n我们可以在启动时让容器在后台运行，这样在退出容器后，容器会继续运行。\n\n```shell\ndocker run -id --name c2 centos \n```\n\n![image.png](http://47.99.200.62/images/blogImages/image_1647009281.png)\n\n\n运行后，输出了运行的容器的ID。使用docker ps -a命令来查看所有容器，发现刚刚启动的容器还在运行\n\n可以使用下面的命令进入容器：\n\n```shell\ndocker exec -it c2 /bin/bash 或\ndocker exec -it 29f35e8928c5 /bin/bash\n```\n\n![image.png](http://47.99.200.62/images/blogImages/image_1647009296.png)\n\n在进入容器并退出后发现，c2容器依然在运行。\n\n当利用docker run来创建并启动容器时，docker在后台运行的标准操作包括：\n\n> - 检查本地是否存在指定的镜像，不存在就从远程仓库下载；\n> - 利用镜像创建一个容器，并启动该容器；\n> - 分配一个文件系统给容器，并在只读的镜像层外面挂载一层可读写层；\n> - 从宿主主机配置的网桥接口中桥接一个虚拟接口到容器中去；\n> - 从网桥的地址池配置一个IP地址给容器；\n> - 执行用户指定的应用程序；\n> - 执行完毕后容器被自动终止。\n\n#### 2.2 查看容器\n\n```shell\ndocker ps\n\n参数：\n# 查看所有容器\n  -a, --all             Show all containers (default shows just running)\n# 只显示容器ID\n  -q, --quiet           Only display container IDs\n```\n\n使用docker ps可以查看容器，类似linux的ps命令查看进程。不加参数只会显示当前正在运行的容器，-a参数可以显示所有容器，-q只显示容器ID，如果要删除所有容器就可以使用该参数了，类似删除全部镜像。\n\n#### 2.3 停止容器\n\n1）暂停容器\n\n```shell\n# 暂停容器\ndocker pause 容器名 或 ID\n# 继续运行暂停的容器\ndocker unpause 容器名 或 ID\n```\n\n2）终止容器\n\n```shell\ndocker stop 容器名 或 ID\n```\n\n![在这里插入图片描述](https://img-blog.csdnimg.cn/e2ad04fc299e49f4bb4ebd5ac176ad9b.png#pic_center)\n\n\n该命令会首先向容器发送SIGTERM信号，等待一段超时时间后（默认为10s），再发送SIGKILL信号来终止容器。\n\n#### 2.4 启动容器\n\n```shell\ndocker start 容器名 或 ID\n```\n\n![image.png](http://47.99.200.62/images/blogImages/image_1647009310.png)\n\n\n```shell\n# 重启容器,先终止再启动\ndocker restart 容器名 或 ID\n```\n\n#### 2.5 删除容器\n\n```shell\ndocker rm 容器名 或 ID\n\n参数：\n#强制删除容器\n  -f, --force     Force the removal of a running container (uses SIGKILL)\n```\n\n容器在运行时使用docker rm无法删除，需要加-f参数才可。\n\n如果要删除所有容器，也可以使用类似删除所有镜像的命令:\n\n```shell\ndocker rm `docker ps -aq`   或\ndocker rm $(docker ps -aq)\n```\n\n#### 2.6 导入和导出容器\n\n某些时候，需要将容器从一个系统迁移到另外一个系统，此时可以使用Docker的导入和导出功能。\n\n1）导出容器\n\n导出容器是指，导出一个已经创建的容器到另一个文件，不管这个容器是否处于运行状态。\n\n```shell\ndocker export -o exportname 容器名   或\ndocker export 容器名 > exportname\n\n# 例如\ndocker export -o test_for_run.tar ce1\ndocker export ce1 > test_for_run\n```\n\n之后，可将导出的tar文件传输到其它机器上，然后通过导入命令导入到系统中，实现容器的迁移。\n\n2）导入容器\n\n```shell\ndocker import filename \n```\n\n#### 2.7 查看容器的相关信息\n\n使用docker inspect可以查看容器的具体信息，会以json格式返回容器ID、创建时间、路径、状态、镜像、配置在内的各项信息。\n\n```shell\ndocker inspect 容器名 或 ID\n# 同样使用docker inspect 镜像名 或 镜像ID 也可以查看镜像的相关信息\n```\n\n![image.png](http://47.99.200.62/images/blogImages/image_1647009329_resized.png)\n\n里面有很多参数，在这里只展示了一部分。\n\n使用下面的命令可以查看docker容器中的进程信息。\n\n```shell\ndocker top 容器名 或 ID\n```\n\n#### 2.8 运行容器时的其它参数\n\n##### 设置环境变量\n\n```shell\n-e, --env list             Set environment variables       设置环境变量\n```\n\n使用-e命令可以给运行的容器设置环境变量\n\n![image.png](http://47.99.200.62/images/blogImages/image_1647009341.png)\n\n\n##### 暴露端口\n\n```she	\n--expose list              Expose a port or a range of ports                   暴露端口\n```\n\n例如，运行nginx时可以将容器的80端口进行暴露。\n\n##### 端口映射\n\n我们在外部机器是无法直接通过网络与容器进行通信的，但是宿主机是可以和容器直接通信的，因此我们可以借助宿主机来和容器进行通信。\n\n例如下图，在容器中运行了一个端口为3306的mysql数据库，在外部机器中是无法直接访问容器中的mysql的。因此可以配置一个端口映射，容器中的3306端口映射到宿主机的3307端口（也可以是其它端口），这样我们可以直接通过宿主机的IP以及端口来访问容器中的数据库。\n\n![image.png](http://47.99.200.62/images/blogImages/image_1647009361_resized.png)\n\n\n```shell\n  -p, --publish list         Publish a container\'s port(s) to the host           指定端口映射\n  -P, --publish-all          Publish all exposed ports to random ports           随机指定暴露的端口映射\n```\n\n例如在运行nginx时，可以将容器内的80端口映射为宿主机的8080端口，我们就可以在浏览器中通过8080端口来访问容器中的nginx。-p自己指定映射的端口。		\n\n![image.png](http://47.99.200.62/images/blogImages/image_1647009798.png)\n\n\n![image.png](http://47.99.200.62/images/blogImages/image_1647009810_resized.png)\n\n\n-P让docker来随机指定映射的端口。\n\n![image.png](http://47.99.200.62/images/blogImages/image_1647009376.png)\n\n![image.png](http://47.99.200.62/images/blogImages/image_1647009387_resized.png)\n\n\n可以使用docker inspect来查看映射的端口：\n\n![image.png](http://47.99.200.62/images/blogImages/image_1647009398.png)\n\n\n##### 指定工作目录\n\n```shell\n-w, --workdir string       Working directory inside the container              指定容器中的工作目录\n```\n\n![image.png](http://47.99.200.62/images/blogImages/image_1647009410.png)\n\n\n可以通过-w命令来指定容器中的工作目录，运行容器后，容器就在指定的工作目录。\n\n#### 2.9 数据管理\n\n在生产环境中使用Docker，往往需要对数据进行持久化，或者需要在多个容器直接进行数据共享，这必然涉及容器的数据管理操作，容器中的管理数据主要有两种方式：\n\n- 数据卷（Data Volumes）：容器内数据直接映射到本地主机环境；\n- 数据卷容器（Data Volume Containers）：使用特定容器维护数据卷。\n\n##### 挂载数据卷\n\n> 数据卷是一个可供容器使用的特殊目录，它将主机操作系统目录直接映射进容器，类似Linux中的mount行为。\n>\n> 数据卷可以提供很多有用的特性：\n>\n> - 数据卷可以在容器间共享和重用，容器间传递数据将变得高效与方便；\n> - 对数据卷内数据的修改会立马生效，无论是容器内操作还是本地操作；\n> - 对数据卷的更新不会影响镜像，解耦开应用和数据；\n> - 卷会一直存在，直到没有容器使用，可以安全地卸载它（即使容器被删除了，本地数据卷也不会被删除）。\n\n```shell\n  -v, --volume list          Bind mount a volume                                 挂载数据卷\n  # 使用：   -v 宿主机目录:容器目录\n```\n\n当容器在删除后，容器中的数据也就随之被删除了。如果我们运行了一个mysql的容器，当容器被删除后，我们不希望mysql中的数据也被删除。因此我们可以将容器中的数据与宿主机的一个目录建立一个映射，这两个目录中的数据会被同步。\n\n![image.png](http://47.99.200.62/images/blogImages/image_1647009429.png)\n\n\n如上图所示，运行了一个名字为c6的centos的容器，并指定了宿主机的/home/mgh/data与容器/usr/local/data的映射（如果路径中的文件夹不存在，则会创建），在容器的/usr/local/data目录中创建了hello.txt并写入内容。然后退出容器，在宿主机的/home/mgh/data目录中发现也有这样一个文件。同时在宿主机中创建了world.txt文件，进入容器后也可以看到该文件。\n\n我们也可以将多个容器中的一个目录挂载在宿主机的同一个目录下，这样就可以在容器间共享数据了。\n\n![image.png](http://47.99.200.62/images/blogImages/image_1647009443.png)\n\n\n##### 数据卷容器\n\n多个容器共享数据，我们可以将它们挂载在同一个目录，如果创建比较多的容器的话，这种方式操作还是比较麻烦的。另一种方法就是创建一个数据卷容器，在运行这个容器时将目录挂载在宿主机上，在启动其它容器时我们可以使用--volumes-from并指定数据卷容器来共享数据。\n\n```shell\n  --volumes-from list              Mount volumes from the specified container(s)  从一个特殊的容器中挂载数据卷\n  # 例如 也可以指定多个目录\n  docker run -id --name volume-container -v /home/data:/usr/local/data centos\n  docker run -id --name volume-des1 --volumes-from volume-container centos\n```\n\n![image.png](http://47.99.200.62/images/blogImages/image_1647009458_resized.png)\n\n\n如上图所示，先使用命令创建了一个数据卷容器，将容器中/usr/local/data挂载到宿主机的/home/data中，然后创建了三个容器并指定了数据卷容器。可以看到在数据卷容器的/usr/local/data目录下创建文件，其它几个容器都是可以共享数据的。\n\n> 如果删除了挂载地容器，数据卷并不会被自动删除。如果要删除一个数据卷，必须在删除最后一个还挂载它的容器时显式使用docker rm -v命令来指定同时删除关联地容器。\n>\n\n## 三、应用部署\n\n### 1、部署mysql\n\n步骤如下：\n\n1. 搜索mysql镜像\n\n   ```shell\n    docker search mysql\n   ```\n\n2. 拉取mysql镜像\n\n   ```shell\n   docker pull mysql:5.7\n   ```\n\n3. 运行容器设置端口映射\n\n   ```shell\n   # 创建mysql目录用于存储mysql数据\n   mkdir ~/mysql\n   cd ~/mysql\n   ```\n\n   ```shell\n   docker run -id \\\n   -p 3307:3306  \\\n   --name c_mysql  \\\n   -v ~/mysql/conf:/etc/mysql/conf.d  \\\n   -v ~/mysql/logs:/logs  \\\n   -v ~/mysql/data:/var/lib/mysql  \\\n   -e MYSQL_ROOT_PASSWORD=123456  \\\n   mysql:5.7\n   ```\n\n   参数说明：\n\n   -p 3307:3306 将容器的3306端口映射到宿主机的3307端口。\n\n   -v 挂载目录\n\n   -e MYSQL_ROOT_PASSWORD=123456  初始化root密码\n\n运行容器后，可以进入容器使用mysql客户端连接mysql查看。\n\n```shell\ndocker exec -it 容器ID /bin/bash\nmysql -uroot -p\n```\n\n![image.png](http://47.99.200.62/images/blogImages/image_1647009474.png)\n\n\n也可以使用navicat等工具来连接mysql：\n\n![image.png](http://47.99.200.62/images/blogImages/image_1647009485.png)\n\n\n### 2、部署tomcat\n\n步骤如下：\n\n1. 搜索tomcat镜像\n\n   ```shell\n    docker search tomcat\n   ```\n\n2. 拉取mysql镜像\n\n   ```shell\n   docker pull tomcat         # 直接拉取最新版\n   ```\n\n3. 运行容器设置端口映射\n\n   ```shell\n   # 创建tomcat目录用于存储tomcat数据信息\n   mkdir ~/tomcat\n   cd ~/tomcat\n   ```\n   \n   ```shell\n   docker run -id --name c_tomcat  \\\n   -p 8080:8080  \\\n   -v ~/tomcat:/usr/local/tomcat/webapps \\\n   tomcat \n   ```\n\n### 3、部署nginx\n\n1. 拉取nginx镜像\n\n   ```shell\n   docker pull nginx         # 直接拉取最新版\n   ```\n\n2. 运行容器设置端口映射\n\n   ```shell\n   mkdir -p ~/nginx/conf\n   cd ~/nginx/conf\n   \n   \n   vim nginx.conf  \n   # 在~/nginx/conf下创建nginx.conf配置文件，将下面的内容黏贴到nginx.conf中\n   ----------------------------------------------------------------\n   \n   user  nginx;\n   worker_processes  auto;\n   \n   error_log  /var/log/nginx/error.log notice;\n   pid        /var/run/nginx.pid;\n   \n   \n   events {\n       worker_connections  1024;\n   }\n   \n   \n   http {\n       include       /etc/nginx/mime.types;\n       default_type  application/octet-stream;\n   \n       log_format  main  \'$remote_addr - $remote_user [$time_local] \"$request\" \'\n                         \'$status $body_bytes_sent \"$http_referer\" \'\n                         \'\"$http_user_agent\" \"$http_x_forwarded_for\"\';\n   \n       access_log  /var/log/nginx/access.log  main;\n   \n       sendfile        on;\n       #tcp_nopush     on;\n   \n       keepalive_timeout  65;\n   \n       #gzip  on;\n   \n       include /etc/nginx/conf.d/*.conf;\n   }\n   \n   ```\n\n   ```shell\n   docker run -id --name=c_nginx \\\n   -p 80:80  \\\n   -v ~/nginx/conf/nginx.conf:/etc/nginx/nginx.conf  \\\n   -v ~/nginx/conf/logs:/var/log/nginx  \\\n   -v ~/nginx/conf/html:/usr/share/nginx/html  \\\n   nginx\n   ```\n\n### 4、部署redis\n\n步骤如下：\n\n1. 拉取redis镜像\n\n   ```shell\n   docker pull redis         # 直接拉取最新版\n   ```\n\n2. 运行容器设置端口映射\n\n   ```shell\n   docker run -id --name c_redis -p 6379:6379 redis\n   ```\n\n3. 使用外部机器连接redis\n\n   ```shell\n   redis-cli -h 宿主机ip -p 6379\n   ```\n\n   \n\n## 四、构建自己的镜像\n\n常用地创建镜像的方式有两种，分别是：将容器提交为一个镜像、使用Dockerfile创建镜像。\n\n### 1、Docker镜像原理\n\n> 首先先思考几个问题：\n>\n> 1. Docker 镜像本质是什么？\n> 2. Docker 中一个centos镜像为什么只有200MB，而一个centos操作系统的iso文件要几个个G？\n> 3. Docker 中一个tomcat镜像为什么有500MB，而一个tomcat安装包只有70多MB\n\n操作系统的组成部分：\n\n- 进程调度子系统\n- 进程通信子系统\n- 内存管理子系统\n- 设备管理子系统\n- <span style=\"color:red;\">文件管理子系统</span>\n- 网络通信子系统\n- 作业控制子系统\n\nLinux文件系统由bootfs和rootfs两部分组成\n\n- bootfs：包含bootloader（引导加载程序）和 kernel（内核）\n- rootfs： root文件系统，包含的就是典型 Linux 系统中的/dev，/proc，/bin，/etc等标准目录和文件\n\n不同的linux发行版，bootfs基本一样，而rootfs不同，如ubuntu，centos等。\n\n![image.png](http://47.99.200.62/images/blogImages/image_1647009514.png)\n\n\nDocker镜像是由特殊的文件系统叠加而成，最底端是 bootfs，并使用宿主机的bootfs ，第二层是 root文件系统rootfs,称为base image，然后再往上可以叠加其他的镜像文件。\n\n联合文件系统（Union File System）技术能够将不同的层整合成一个文件系统，为这些层提供了一个统一的视角，这样就隐藏了多层的存在，在用户的角度看来，只存在一个文件系统。\n\n一个镜像可以放在另一个镜像的上面。位于下面的镜像称为父镜像，最底部的镜像成为基础镜像。当从一个镜像启动容器时，Docker会在最顶层加载一个读写文件系统作为容器\n\n![image.png](http://47.99.200.62/images/blogImages/image_1647009529.png)\n\n\n因此上面思考的答案是：\n\n> 1. Docker 镜像本质是什么？\n>\n>    ​		：是一个分层的文件系统\n>\n> 2. Docker 中一个centos镜像为什么只有200MB，而一个centos操作系统的iso文件要几个个G？\n>\n>    ​		： Centos的iso镜像文件包含bootfs和rootfs，而docker的centos镜像复用操作系统的bootfs，只有rootfs和其他镜像层\n>\n> 3. Docker 中一个tomcat镜像为什么有500MB，而一个tomcat安装包只有70多MB\n>\n>    ​		： 由于docker中镜像是分层的，tomcat虽然只有70多MB，但他需要依赖于父镜像和基础镜像，所有整个对外暴露的tomcat镜像大小500多MB\n\n### 2、提交容器为镜像\n\n我们可以将一个容器转为一个镜像。也可以将一个镜像生成一个本地的压缩文件，在1.6节已经介绍。\n\n![image.png](http://47.99.200.62/images/blogImages/image_1647009541.png)\n\n\n我们从远程仓库拉取的官方的ubuntu系统是很纯净的，里面去除了很多不必要的软件，甚至连sudo命令的没有。我们可以运行一个从官网拉取的镜像，并在容器中安装vim、GCC、make等工具，然后将该容器提交为一个镜像。那么之后根据这个镜像运行新的容器时，容器中就会有这些工具了。\n\n```shell\ndocker commit 容器名 或 ID 生成的镜像标签\n\n# 例如\ndocker commit ubuntu1 myubuntu:lastest \n```\n\n### 3、使用Dockerfile创建镜像\n\nDockerfile是一个文本格式的配置文件，用户可以使用Dockerfile来快速创建自定义的镜像。Dockerfile一般由一行行指令组成，并且支持以#开头的注释行；每一条指令构建一层，基于基础镜像，最终构建出一个新的镜像；对于开发人员：可以为开发团队提供一个完全一致的开发环境；对于测试人员：可以直接拿开发时所构建的镜像或者通过Dockerfile文件构建一个新的镜像开始工作了；对于运维人员：在部署时，可以实现应用的无缝移植；\n\n一般而言，Dockerfile主体内容分为四部分：基础镜像信息、维护者信息、镜像操作指令和容器启动时执行指令。\n\n![image.png](http://47.99.200.62/images/blogImages/image_1647009558.png)\n\n\nDockerfile相关命令：\n\n| 关键字      | 作用                     | 备注                                                         |\n| ----------- | ------------------------ | ------------------------------------------------------------ |\n| FROM        | 指定父镜像               | 指定dockerfile基于哪个image构建                              |\n| MAINTAINER  | 作者信息                 | 用来标明这个dockerfile谁写的                                 |\n| LABEL       | 标签                     | 用来标明dockerfile的标签 可以使用Label代替Maintainer 最终都是在docker image基本信息中可以查看 |\n| RUN         | 执行命令                 | 执行一段命令 默认是/bin/sh 格式: RUN command 或者 RUN [\"command\" , \"param1\",\"param2\"]。RUN命令是在镜像构建时执行的。 |\n| CMD         | 容器启动命令             | 提供启动容器时候的默认命令 和ENTRYPOINT配合使用.格式 CMD command param1 param2 或者 CMD [\"command\" , \"param1\",\"param2\"]。CMD命令是在容器启动时执行的。 |\n| ENTRYPOINT  | 入口                     | 指定镜像的默认入口命令                      |\n| COPY        | 复制文件                 | build的时候复制文件到image中                                 |\n| ADD         | 添加文件                 | build的时候添加文件到image中 不仅仅局限于当前build上下文 可以来源于远程服务，添加压缩文件时会自动解压。 |\n| ENV         | 环境变量                 | 指定build时候的环境变量 可以在启动的容器的时候 通过-e覆盖 格式ENV name=value |\n| ARG         | 构建参数                 | 构建参数 只在构建的时候使用的参数 如果有ENV 那么ENV的相同名字的值始终覆盖arg的参数 |\n| VOLUME      | 定义外部可以挂载的数据卷 | 指定build的image那些目录可以启动的时候挂载到文件系统中 启动容器的时候使用 -v 绑定 格式 VOLUME [\"目录\"] |\n| EXPOSE      | 暴露端口                 | 定义容器运行的时候监听的端口 启动容器的使用-p来绑定暴露端口 格式: EXPOSE 8080 或者 EXPOSE 8080/udp |\n| WORKDIR     | 工作目录                 | 指定容器内部的工作目录 如果没有创建则自动创建 如果指定/ 使用的是绝对地址 如果不是/开头那么是在上一条workdir的路径的相对路径 |\n| USER        | 指定执行用户             | 指定build或者启动的时候 用户 在RUN CMD ENTRYPONT执行的时候的用户 |\n| HEALTHCHECK | 健康检查                 | 指定监测当前容器的健康监测的命令 基本上没用 因为很多时候 应用本身有健康监测机制 |\n| ONBUILD     | 触发器                   | 当存在ONBUILD关键字的镜像作为基础镜像的时候 当执行FROM完成之后 会执行 ONBUILD的命令 但是不影响当前镜像 用处也不怎么大 |\n| STOPSIGNAL  | 发送信号到宿主机         | 该STOPSIGNAL指令设置将发送到容器的系统调用信号以退出。       |\n| SHELL       | 指定执行脚本的shell      | 指定RUN CMD ENTRYPOINT 执行命令的时候 使用的shell            |\n\nDockerfile构建命令：\n\n```shell\ndocker build -f ./Dockerfile -t myimage .\n\n# -f 参数用于指定dockerfile的位置，如果在当前目录下切名字为Dockerfile，也可以不指定\n# -t target 用于指定生成的镜像名\n# .为构建的目录\n```\n\n#### **例1：构建Springboot项目的镜像：**\n\n写了一个简单的springboot程序：\n\n![image.png](http://47.99.200.62/images/blogImages/image_1647009579_resized.png)\n\n\n可以在浏览器中访问的到：\n\n![image.png](http://47.99.200.62/images/blogImages/image_1647009592.png)\n\n\n构建步骤：\n\n1.将项目打包为jar文件，并上传到linux的一个目录下，例如/home/mgh/docker_springboot文件下：\n\n```shell\nmkdir ~/docker_springboot\n```\n\n2.创建Dokcerfile文件：\n\n```shell\nvim Dockerfile\n```\n\n3.写Dockerfile：\n\ndockerfile 内容如下\n\n```shell\n# 由于springboot项目依赖java jdk，因此需要有jdk的环境。也可以是ubuntu或centos，但是其中没有java环境，还需要安装\nFROM java:8\n\n# 作者信息,也可以不要\nMAINTAINER author<author@hdu.edu.cn>\n\n# 将当前目录下的jar包添加到镜像中，并命名为app.jar\nADD springboot_test-0.0.1-SNAPSHOT.jar app.jar\n\n# 运行springboot项目 也可以是CMD [\"java\", \"-jar\", \"app.jar\"]\nCMD java -jar app.jar\n```\n\n4.构建镜像：\n\n```shell\ndocker build -t springboot_test .\n```\n\n5.构建成功\n\n![image.png](http://47.99.200.62/images/blogImages/image_1647009608.png)\n\n\n6.运行容器\n\n```shell\n docker run -id -p 8080:8080 springboot_test\n```\n\n7.在浏览器中访问\n\n![image.png](http://47.99.200.62/images/blogImages/image_1647009620.png)\n\n\n\n\n#### 例2：构建C++程序\n\n1.创建目录、写一个helloworld程序\n\n```shell\nmkdir ~/cpp_docker\nvim hello.cpp\n\n```\n程序如下，每隔一秒打印一次hello，world\n```cpp\n#include <iostream>\n#include <unistd.h>\n\nusing namespace std;\n\nint main()\n{\n    while(1)\n    {\n        cout << \"hello,world!\" << endl;\n        sleep(1);\n    }\n\n    return 0;\n}\n```\n\n2.创建Dockerfile文件并写入\n\n```shell\nvim Dockerfile\n\n# 内容如下：\n\n# 基础镜像\nFROM ubuntu:18.04\n\nMAINTAINER author<author@hdu.edu.cn>\n\n# 设置环境变量\nENV MYPATH /usr/local/test\n\n# 设置工作目录，设置工作目录后，启动容器，容器所在目录就是工作目录\nWORKDIR $MYPATH\n\n# 将当前文件夹下的cpp文件拷贝到工作目录\nCOPY hello.cpp $MYPATH\n\n# 安装g++ 以及 编译cpp文件。也可以将这三个命令分开写成三条，但是会生成三层，所以不建议\n# 在安装软件的时候，系统可能会让你输入y或n来同意或拒绝安装，使用-y参数就是同意安装。不然可能会失败。\nRUN apt-get update &&  \\\n    apt-get -y install g++ && \\\n    g++ hello.cpp -o hello\n\n# 容器启动时执行的命令\nCMD [\"./hello\"]\n\n```\n\n3.构建镜像\n\n```shell\ndocker build -t cpp_test .\n```\n\n4.运行容器\n\n```shell\ndocker run -it cpp_test\n```\n\n![image.png](http://47.99.200.62/images/blogImages/image_1647009634.png)\n\n\n可以看到容器中在打印hello，world。但是退出不了了，可以使用docker stop将这个容停止。\n\n#### 例3：构建Golang程序\n\n将使用gin框架的一个web程序构建为docker镜像：\n\n1.golang代码：\n\n```go\npackage main\n\nimport (\n	\"github.com/gin-gonic/gin\"\n	\"net/http\"\n)\n\n\nfunc main()  {\n	engine := gin.Default()\n\n	engine.GET(\"/jige\", func(context *gin.Context) {\n		data := map[string]interface{}{\n			\"name\" : \"kunkun\",\n			\"age\" : 23,\n			\"hobby\" : \"sing jump rap\",\n		}\n		context.JSON(http.StatusOK, data)\n\n	})\n\n	engine.Run(\"0.0.0.0:8080\")\n}\n\n```\n\ngo.mod\n\n```shell\ngo mod init\ngo mod tidy\n```\n\n2.将程序打包上传到linux的一个文件中\n\n```shell\nmkdir ~/docker_golang\n#程序文件名 gin_json\n\n```\n\n3.创建Dockerfile\n\n```shell\nvim Dockerfile\n\n# dockerfile内容如下\n# 基础镜像\nFROM golang\n\nMAINTAINER author <author@hdu.edu.cn>\n\n# 设置环境变量\nENV GO111MODULE on\n# 设置golang代理否则下载gin框架很可能会失败\nENV GOPROXY https://goproxy.cn\n\n# 设置工作目录\nWORKDIR /go/src/gin_json/\n\n# 把当前目录下所以文件拷贝到容器的工作目录\nCOPY . .\n\n# 安装gin框架依赖\nRUN go mod tidy\n\n# 容器启动运行程序\nCMD go run main.go\n```\n\n4.构建镜像\n\n```shell\ndocker build -t gotest .\n```\n\n5.运行容器\n\n```shell\ndocker run -id -p 9090:8080 gotest\n```\n\n6.浏览器访问页面\n\n![image.png](http://47.99.200.62/images/blogImages/image_1647009656.png)\n\n\n### 4、配置指令详细介绍：\n\n1.FROM ：\n\n> 指定所创建镜像的基础镜像，比如JAVA程序依赖于java jdk，因此可以选择一个带有java jdk的镜像作为基础镜像。\n>\n> 格式：FROM  image  或 FROM image:tag\n>\n> 任何Dockerfile中第一条指令必须为FROM指令。并且，如果在同一个Dockerfile中创建多个镜像时，可以使用多个FROM指令。\n>\n> 为了保证镜像精简，可以选用体积较小的镜像如Alpine或Debian作为基础镜像。\n\n2.ARG  ：\n\n> 定义创建镜像过程中使用的变量。\n>\n> 格式：ARG key  val  或 ARG key =val\n>\n> 当镜像构建完成后，ARG定义的变量将不再存在（ENV指定的变量将在镜像中保留）。\n\n3.ENV ：\n\n> 指定环境变量，在镜像生成过程中会被RUN指令使用，在镜像启动的容器中也会存在。\n>\n> 格式：ENV key  val  或 ENV  key=val\n>\n> 指令指定的环境变量在运行时可以被覆盖掉，如：docker run --env key=val built_image\n\n4.LABEL ：\n\n> LABEL指令可以为生成的镜像添加元数据标签信息。\n>\n> 格式：LABEL key=value key=value ...\n\n5.EXPOSE ：\n\n> 声明镜像内服务监听的端口。\n>\n> 格式：EXPOSE port ...\n>\n> 例如：\n>\n> EXPOSE 22 80 8080\n>\n> 该指令只是起到声明作用，并不会自动完成端口映射。\n\n6.ENTRYPOINT ：\n\n> 指定镜像的默认入口命令，该入口命令会在启动容器时作为根命令执行，所有传入值作为该命令的参数。\n>\n> 支持两种格式：\n>\n> ENTRYPOINT [\"executable\", \"param1\", \"param2\"]                      exec调用执行\n>\n> ENTRYPOINT command param1 param2                                    shell中执行\n>\n> 此外，CMD指令指定值将作为根命令的参数。每个Dockerfile中只能有一个ENTRYPOINT，当指定多个时，只有最后一个生效。\n\n7.CMD ：\n\n> CMD指令用来指定启动容器时默认执行的命令。\n>\n> 支持三种格式：\n>\n> CMD [\"executable\", \"param1\", \"param2\"]         相当于执行executable param1 param2\n>\n> CMD command param1 param2                      在默认的shell中执行，提供给需要交互的应用\n>\n> CMD [\"param1\", \"param2\"]                               提供给ENTRYPOINT的默认参数\n>\n> 每个Dockerfile只能有一条CMD命令。如果指定了多条命令，只有最后一条会被执行。如果用户启动容器时手动指定了运行的命令，则会覆盖掉CMD指定的命令。\n>\n> 例如：\n>\n> 在上面的构建C++程序中，指定了CMD运行hello程序：CMD [\"./hello\"]\n>\n> 因此如果在运行容器时指定了/bin/bash，那么./hello将会被覆盖掉，hello程序也就不会执行 ：  docker run -it cpp_test /bin/bash\n\n9.VOLUME ：\n\n> 创建一个数据卷挂载点。\n>\n> 格式： VOLUME  [\"/data\"]\n>\n> 运行容器时可以从本地主机或其它容器挂载数据卷，一般用来存放数据库和需要保持的数据或一些配置文件等。\n\n10.WORKDIR ：\n\n> 为后续的RUN、CMD、ENTRYPOINT指令配置工作目录\n>\n> 格式： WORKDIR dir\n>\n> 在WORKDIR指令中最好使用绝对路径。\n\n11.RUN ：\n\n> 运行指定命令。\n>\n> 格式：\n>\n> RUN  command  或 RUN [\"executable\", \"param1\", \"param2\"]\n>\n> 注意后者指令将会被解析为JSON数组，因此必须用双引号。前者默认在shell终端中运行命令，即/bin/bash -c；后者使用exec执行，不会启动shell环境。\n>\n> 指定使用其它终端类型可以通过第二种方式实现，例如 RUN [\"/bin/bash\", \"-c\", \"echo hello\"]。\n>\n> 每条RUN命令将在当前基础镜像上执行指定命令，并提交为新的镜像层，如果有多条RUN命令，分开来写将会产生很多层镜像层，因此最后将他们写成一条命令，可以用 \\ 来换行。例如：\n>\n>   RUN apt-get update \\\n>\n> ​				&& apt-get install -y libsnappy-dev zliblg-dev libbz2-dev \\\n>\n> ​				&& rm -rf /var/cache/apt \\\n>\n> ​				&& rm -rf /var/lib/apt/lists/*\n\n12.ADD ：\n\n> 添加内容到镜像。\n>\n> 格式： ADD src  dest\n>\n> 该命令将复制指定的src路径下内容到容器中的dest路径下。\n>\n> 其中src可以是Dockerfile所在目录的一个相对路径（文件或目录）；也可以是一个URL；还可以是一个tar文件，tar文件将会被自动解压为目录。路径支持正则表达式，例如：\n>\n> ADD *.cpp  /code/\n\n13.COPY ：\n\n> 复制内容到镜像。\n>\n> 格式：COPY src dest\n>\n> 复制本地主机的src下内容到镜像中的dest。目录不存在时，会自动创建。\n>\n> 路径同样支持正则表达式。\n>\n> COPY与ADD命令功能相似，当使用本地目录为源目录时，推荐使用COPY。\n\n**注意**：在构建镜像时最好保证指定目录（命令为：docker build -t test .， 最后那个.就是指定当前目录为构建目录）下没有其它不需要的文件，因为在构建镜像时是由Docker守护进程来执行的，因此客户端将会把指定目录下的所有内容都发送给Docker守护进程，如果当前目录下存在不需要的文件，将会拖慢镜像的构建。\n\n### 5.最佳实践\n\n所谓最佳实践，就是从需求出发，来定制适合自己、高效的镜像。\n\n**精简镜像用途**：尽量让每个镜像的用途都比较单一，避免构造大而复杂。\n\n**选用合适的基础镜像**：容器的核心是应用。选择过大的父镜像（如Ubuntu系统镜像）会造成最终生成应用镜像的臃肿，推荐选用瘦身过的应用镜像（如node:slim），或者较为小巧的系统镜像（如alpine、busybos或debian）；\n\n**正确使用版本号**：使用明确的版本号信息，如1.0，2.0，而非依赖于默认的lastest。通过版本号可以避免环境不一导致的问题；\n\n**减少镜像层数**：尽量合并RUN、AD和COPY指令。\n\n**及时删除临时文件和缓存文件：**特别是在执行apt-get命令后，/var/cache/apt下面会缓存了一些安装包；\n\n**提高生成速度**：减少内容目录下不必要的文件。\n\n\n[Docker进阶（Docker网络、Docker Compose）](https://blog.csdn.net/Peerless__/article/details/121053065)\n', 'http://47.99.200.62/images/firstPic/44_1647008689.png', '原创', 10, 1, 0, 1, 1, 0, '2022-03-11 14:25:01', '2022-03-11 14:44:33', 17, 1, '\n\n......');
INSERT INTO `t_blog` VALUES (45, '管理员测试', '测试新增', 'http://192.168.44.100/images/firstPic/11_1649302474.png', '原创', 1, 1, 1, 1, 1, 1, '2022-04-07 11:34:47', '2022-04-07 11:34:47', 1, 1, '测试');
INSERT INTO `t_blog` VALUES (46, '测试截取博客描述', '测试截取博客描述测试截取博客描述测试截取博客描述测试截取博客描述测试截取博客描述测试截取博客描述测试截取博客描述测试截取博客描述测试截取博客描述测试截取博客描述测试截取博客描述测试截取博客描述测试截取博客描述', 'http://192.168.44.100/images/firstPic/20_1649303465.png', '原创', 0, 1, 1, 1, 1, 0, '2022-04-07 11:51:22', '2022-04-07 11:51:22', 1, 1, '测试截取博客描述');
INSERT INTO `t_blog` VALUES (47, '测试添加图片', '![34.jpg](http://192.168.44.100/images/blogImages/34_1649320918.jpg)', 'http://192.168.44.100/images/firstPic/47_1649320924.png', '原创', 3, 1, 1, 1, 1, 0, '2022-04-07 16:42:35', '2022-04-07 16:42:35', 1, 1, '测试添加图片');

-- ----------------------------
-- Table structure for t_blog_tags
-- ----------------------------
DROP TABLE IF EXISTS `t_blog_tags`;
CREATE TABLE `t_blog_tags`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `tag_id` int UNSIGNED NOT NULL COMMENT '标签id',
  `blog_id` int UNSIGNED NOT NULL COMMENT '博客id',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_blog_id`(`blog_id`) USING BTREE COMMENT '博客id索引'
) ENGINE = InnoDB AUTO_INCREMENT = 211 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of t_blog_tags
-- ----------------------------
INSERT INTO `t_blog_tags` VALUES (111, 4, 3);
INSERT INTO `t_blog_tags` VALUES (112, 6, 3);
INSERT INTO `t_blog_tags` VALUES (123, 3, 7);
INSERT INTO `t_blog_tags` VALUES (124, 11, 7);
INSERT INTO `t_blog_tags` VALUES (194, 16, 44);
INSERT INTO `t_blog_tags` VALUES (195, 17, 44);
INSERT INTO `t_blog_tags` VALUES (196, 3, 6);
INSERT INTO `t_blog_tags` VALUES (197, 3, 5);
INSERT INTO `t_blog_tags` VALUES (198, 4, 5);
INSERT INTO `t_blog_tags` VALUES (199, 5, 5);
INSERT INTO `t_blog_tags` VALUES (200, 3, 4);
INSERT INTO `t_blog_tags` VALUES (201, 4, 4);
INSERT INTO `t_blog_tags` VALUES (202, 6, 4);
INSERT INTO `t_blog_tags` VALUES (203, 4, 2);
INSERT INTO `t_blog_tags` VALUES (204, 5, 2);
INSERT INTO `t_blog_tags` VALUES (205, 6, 2);
INSERT INTO `t_blog_tags` VALUES (206, 5, 1);
INSERT INTO `t_blog_tags` VALUES (207, 6, 1);
INSERT INTO `t_blog_tags` VALUES (208, 3, 45);
INSERT INTO `t_blog_tags` VALUES (209, 4, 45);
INSERT INTO `t_blog_tags` VALUES (210, 4, 46);
INSERT INTO `t_blog_tags` VALUES (211, 3, 47);

-- ----------------------------
-- Table structure for t_comment
-- ----------------------------
DROP TABLE IF EXISTS `t_comment`;
CREATE TABLE `t_comment`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `nickname` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '昵称',
  `email` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '电子邮件',
  `content` varchar(1024) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '内容',
  `avatar` varchar(512) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '图像地址',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `blog_id` int NOT NULL COMMENT '博客id',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_blog_id`(`blog_id`) USING BTREE COMMENT '博客id索引'
) ENGINE = InnoDB AUTO_INCREMENT = 18 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of t_comment
-- ----------------------------
INSERT INTO `t_comment` VALUES (1, '小白', 'bai@qq.com', '小白的评论', 'http://192.168.44.100/images/firstPic/avatar1.png', '2021-02-05 11:56:19', 4);
INSERT INTO `t_comment` VALUES (2, '小红', 'hong@qq.com', '小红的评论', 'http://192.168.44.100/images/firstPic/avatar1.png', '2021-02-05 11:56:19', 4);
INSERT INTO `t_comment` VALUES (5, '小蓝', 'lan@qq.com', '小蓝的评论', 'http://192.168.44.100/images/firstPic/avatar1.png', '2021-02-05 11:56:19', 4);
INSERT INTO `t_comment` VALUES (7, '小马', '691639910@qq.com', '博主的评论', 'http://5b0988e595225.cdn.sohucs.com/images/20181103/feaa7d14883047fb81bbaa16f681f583.jpeg', '2021-02-05 11:56:19', 2);
INSERT INTO `t_comment` VALUES (8, '安东尼', '2333@qq.com', '不论是我的世界车水马龙繁华盛世 还是它们都瞬间消失化为须臾 我都会坚定地走向你 不慌张 不犹豫', 'http://192.168.44.100/images/firstPic/avatar1.png', '2021-02-05 11:56:19', 6);
INSERT INTO `t_comment` VALUES (9, '亚瑟', 'yase@163.com', 'good!', 'http://192.168.44.100/images/firstPic/avatar1.png', '2021-03-31 22:24:17', 23);
INSERT INTO `t_comment` VALUES (10, '约翰', 'yuehan@163.com', 'excellent!', 'http://192.168.44.100/images/firstPic/avatar1.png', '2021-03-31 22:26:47', 23);
INSERT INTO `t_comment` VALUES (11, '达奇', 'daqi@163.com', 'I have a plan! A good plan!', 'http://192.168.44.100/images/firstPic/avatar1.png', '2021-03-31 22:47:36', 23);
INSERT INTO `t_comment` VALUES (15, 'admin', 'test', 'test', 'http://192.168.44.100/images/firstPic/avatar1.png', '2022-03-03 20:11:36', 37);
INSERT INTO `t_comment` VALUES (16, 'admin', 'TEST', 'TEST', 'http://192.168.44.100/firstPic/avatar1.png', '2022-03-03 20:12:52', 37);
INSERT INTO `t_comment` VALUES (17, 'test', 'test', 'admin', 'http://192.168.44.100/images/firstPic/avatar1.png', '2022-03-03 20:17:36', 37);
INSERT INTO `t_comment` VALUES (18, '小红', '1789033422@qq.com', '试一试 ', 'http://192.168.44.100/images/firstPic/avatar1.png', '2022-03-18 12:46:41', 4);

-- ----------------------------
-- Table structure for t_essay
-- ----------------------------
DROP TABLE IF EXISTS `t_essay`;
CREATE TABLE `t_essay`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `content` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '内容',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_create_time`(`create_time`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_essay
-- ----------------------------
INSERT INTO `t_essay` VALUES (1, '2022-04-06 18:09:32', '自信人生两百年，会当水击三千里！');
INSERT INTO `t_essay` VALUES (2, '2022-04-07 18:09:50', '如果你累了，学会休息，而不是放弃!');
INSERT INTO `t_essay` VALUES (3, '2022-04-08 18:10:10', '无论生活多么的复杂，总要保持自己的那一份优雅!');
INSERT INTO `t_essay` VALUES (4, '2021-05-01 18:10:48', '自信人生两百年，会当水击三千里！');
INSERT INTO `t_essay` VALUES (5, '2021-05-01 20:10:48', '如果你累了，学会休息，而不是放弃!');
INSERT INTO `t_essay` VALUES (6, '2022-04-06 19:09:15', '学会享受孤独！');

-- ----------------------------
-- Table structure for t_leave_message
-- ----------------------------
DROP TABLE IF EXISTS `t_leave_message`;
CREATE TABLE `t_leave_message`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `nickname` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '昵称',
  `email` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '邮箱地址',
  `url` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '个人网站',
  `content` varchar(1600) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '评论内容',
  `avatar` tinyint UNSIGNED NOT NULL COMMENT '头像',
  `create_time` datetime NOT NULL COMMENT '评论时间',
  `parent_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '父评论，0表示没有',
  `top_parent_id` int NOT NULL DEFAULT 0 COMMENT '顶级评论',
  `status` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '评论状态，0：未审核，1：审核通过，2：审核不通过',
  `ip` char(21) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '评论ip',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_status`(`status`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_leave_message
-- ----------------------------
INSERT INTO `t_leave_message` VALUES (1, '管理员', '123456@qq.com', '', 'This is a test.', 10, '2022-03-29 13:28:02', 0, 0, 1, '127.0.0.1:8082');
INSERT INTO `t_leave_message` VALUES (2, '管理员', '123456@qq.com', '', 'Test2', 1, '2022-03-29 13:29:36', 0, 0, 1, '127.0.0.1:8082');
INSERT INTO `t_leave_message` VALUES (3, '管理员', '456789@qq.com', 'https://www.baidu.com', 'Sub comment.', 6, '2022-03-30 03:07:59', 1, 1, 1, '127.0.0.1:8082');
INSERT INTO `t_leave_message` VALUES (4, '管理员', '456789@qq.com', 'https://www.baidu.com', 'Sub  Comment2.', 2, '2022-03-30 03:09:32', 1, 1, 1, '127.0.0.1:8082');
INSERT INTO `t_leave_message` VALUES (5, '管理员', '456789@qq.com', 'https://www.baidu.com', 'Sub Sub Comment.', 3, '2022-03-30 03:09:32', 3, 1, 1, '127.0.0.1:8082');
INSERT INTO `t_leave_message` VALUES (6, '管理员', '456789@qq.com', 'https://www.baidu.com', 'Sub Sub Comment.', 4, '2022-03-30 03:09:32', 2, 2, 1, '127.0.0.1:8082');
INSERT INTO `t_leave_message` VALUES (7, '管理员', '456789@qq.com', 'https://www.baidu.com', 'Sub Sub Comment.', 5, '2022-03-30 03:09:32', 2, 2, 1, '127.0.0.1:8082');
INSERT INTO `t_leave_message` VALUES (8, '管理员', '456789@qq.com', 'https://www.baidu.com', 'Sub Sub Comment.', 7, '2022-03-30 03:09:32', 0, 0, 1, '127.0.0.1:8082');
INSERT INTO `t_leave_message` VALUES (9, '管理员', '456789@qq.com', 'https://www.baidu.com', 'Sub Sub Comment.', 8, '2022-03-30 03:09:32', 6, 2, 1, '127.0.0.1:8082');
INSERT INTO `t_leave_message` VALUES (10, '管理员测试', 'aaa@qq.com', '', '测试发布子级评论', 2, '2022-03-30 09:28:30', 3, 1, 0, '192.168.44.1:56811');
INSERT INTO `t_leave_message` VALUES (11, '管理员测试', 'aaa@qq.com', '', '测试子级评论发布', 26, '2022-03-30 09:31:48', 3, 1, 1, '192.168.44.1:56979');
INSERT INTO `t_leave_message` VALUES (12, '无名', 'ccdd@qq.com', '', '测试', 25, '2022-03-30 11:52:18', 8, 8, 1, '192.168.44.1:63436');
INSERT INTO `t_leave_message` VALUES (13, '大名', 'eeff@qq.com', '', '测试', 12, '2022-03-30 11:55:41', 8, 8, 1, '192.168.44.1:63532');
INSERT INTO `t_leave_message` VALUES (14, '测试人员', 'fff@qq.com', '', 'Test', 12, '2022-03-30 12:14:49', 0, 0, 1, '192.168.44.1:64602');
INSERT INTO `t_leave_message` VALUES (15, '管理猿', '6666@qq.com', '', 'Success!', 27, '2022-03-30 13:48:32', 0, 0, 1, '192.168.44.1:55829');

-- ----------------------------
-- Table structure for t_motto
-- ----------------------------
DROP TABLE IF EXISTS `t_motto`;
CREATE TABLE `t_motto`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `ch` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '中文',
  `en` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '英文',
  `create_time` datetime NOT NULL COMMENT '创建日期',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of t_motto
-- ----------------------------
INSERT INTO `t_motto` VALUES (1, '心有多大，舞台就有多大!', 'How large the heart is,how large the stage is!', '2022-04-05 20:45:20');
INSERT INTO `t_motto` VALUES (2, '人生最可悲的事莫过于胸怀大志，却虚度光阴!', 'The biggest shame for a man is to waste his time having fun with the great dream!', '2022-04-06 20:45:25');
INSERT INTO `t_motto` VALUES (3, '无论生活多么的复杂，总要保持自己的那一份优雅!', 'No matter how complicated the life is ，Please be graceful all the time!', '2022-04-04 20:45:29');
INSERT INTO `t_motto` VALUES (4, '不需要任何人满意，但求于无愧于心!', 'I don\'t need to please everyone ，I just want to do everythingwithout regret for myself!', '2022-04-03 20:45:33');
INSERT INTO `t_motto` VALUES (5, '越努力，越幸运!', 'The harder you are working, the luckier you will be!', '2022-04-02 20:45:37');
INSERT INTO `t_motto` VALUES (6, '当你的才华还撑不起你的野心时，就要静下心来学习；当你的能力还撑不起你的梦想时，就要努力的去工作!', 'Please study hard，if your talent is not enough to achieve your ambition。 Please work hard，if your ability is not enough to make your dream come true!', '2022-04-13 20:45:40');
INSERT INTO `t_motto` VALUES (7, '如果你累了，学会休息，而不是放弃!', 'If you get tired, learn to rest, not to quit!', '2022-04-11 20:45:45');

-- ----------------------------
-- Table structure for t_resource_category
-- ----------------------------
DROP TABLE IF EXISTS `t_resource_category`;
CREATE TABLE `t_resource_category`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '类别名',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 21 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_resource_category
-- ----------------------------
INSERT INTO `t_resource_category` VALUES (1, 'Golang学习资料');
INSERT INTO `t_resource_category` VALUES (2, '优质博客');
INSERT INTO `t_resource_category` VALUES (3, '工具网站');
INSERT INTO `t_resource_category` VALUES (21, '前端开发');

-- ----------------------------
-- Table structure for t_resource_link
-- ----------------------------
DROP TABLE IF EXISTS `t_resource_link`;
CREATE TABLE `t_resource_link`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '链接名称',
  `desc` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '链接描述',
  `url` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT 'url',
  `category_id` int UNSIGNED NOT NULL COMMENT '类别id',
  `icon` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'icon',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_category_id`(`category_id`) USING BTREE COMMENT '类别id索引'
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_resource_link
-- ----------------------------
INSERT INTO `t_resource_link` VALUES (1, 'Go语言高级编程', 'Go语言高级编程(Advanced Go Programming)', 'https://chai2010.cn/advanced-go-programming-book/', 1, 'http://192.168.44.100/images/icons/favicon_1.ico');
INSERT INTO `t_resource_link` VALUES (2, 'Go标准库', 'Golang标准库文档', 'http://doc.golang.ltd/', 1, ' ');
INSERT INTO `t_resource_link` VALUES (3, '李文周的博客', '李文周的博客，里面含有Golang教程', 'https://www.liwenzhou.com/archives/', 1, 'http://192.168.44.100/images/icons/favicon_2.ico');
INSERT INTO `t_resource_link` VALUES (4, '鸟窝', '一位大佬的博客', 'https://colobu.com/', 1, '');
INSERT INTO `t_resource_link` VALUES (5, 'Go语言爱好者周刊', 'Go语言周刊，每周日发布', 'https://github.com/polaris1119/golangweekly', 1, '');
INSERT INTO `t_resource_link` VALUES (6, 'Go语言设计与实现', '深入了解Go语言的实现', 'https://draveness.me/golang/', 1, 'http://192.168.44.100/images/icons/favicon_3.png');
INSERT INTO `t_resource_link` VALUES (7, '程序员的工具箱', '各种程序员使用的在线工具', 'https://tool.lu/', 3, 'http://192.168.44.100/images/icons/favicon_4.ico');
INSERT INTO `t_resource_link` VALUES (8, '编程自学之路', '编程相关知识', 'https://www.r2coding.com/', 2, 'http://192.168.44.100/images/icons/favicon_5.ico');
INSERT INTO `t_resource_link` VALUES (9, '九月网址导航', '各种工具，PPT、高清图库...', 'https://jiuyueppt.com/', 3, 'http://192.168.44.100/images/icons/favicon_6.ico');
INSERT INTO `t_resource_link` VALUES (13, '在线画图工具', '一个在线的画图工具', 'https://www.processon.com/', 3, 'http://192.168.44.100/images/icons/favicon_1649424714.ico');
INSERT INTO `t_resource_link` VALUES (14, 'Animate.css', '一个css动画库，非常好用', 'https://animate.style/', 21, 'http://192.168.44.100/images/icons/favicon_1649486228.ico');
INSERT INTO `t_resource_link` VALUES (15, 'jQuery插件库', '提供各种最新的jQuery插件，各种css特效等', 'https://www.jq22.com/code2732', 21, 'http://192.168.44.100/images/icons/favicon_1649486404.ico');

-- ----------------------------
-- Table structure for t_tag
-- ----------------------------
DROP TABLE IF EXISTS `t_tag`;
CREATE TABLE `t_tag`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '标签名',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 17 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of t_tag
-- ----------------------------
INSERT INTO `t_tag` VALUES (3, '前端');
INSERT INTO `t_tag` VALUES (4, '后端');
INSERT INTO `t_tag` VALUES (5, 'springboot');
INSERT INTO `t_tag` VALUES (6, 'java');
INSERT INTO `t_tag` VALUES (16, '应用部署');
INSERT INTO `t_tag` VALUES (17, '容器');

-- ----------------------------
-- Table structure for t_type
-- ----------------------------
DROP TABLE IF EXISTS `t_type`;
CREATE TABLE `t_type`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '博客类型名',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 17 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of t_type
-- ----------------------------
INSERT INTO `t_type` VALUES (1, '博客系统');
INSERT INTO `t_type` VALUES (3, 'Mybatis');
INSERT INTO `t_type` VALUES (7, 'Springboot');
INSERT INTO `t_type` VALUES (8, 'Maven');
INSERT INTO `t_type` VALUES (17, 'Docker');

-- ----------------------------
-- Table structure for t_user
-- ----------------------------
DROP TABLE IF EXISTS `t_user`;
CREATE TABLE `t_user`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `nickname` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '昵称',
  `username` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '密码',
  `email` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '邮箱',
  `avatar` varchar(512) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '头像地址',
  `type` int UNSIGNED NOT NULL COMMENT '用户类型，管理员、普通用户',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `github` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'github地址',
  `csdn` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'csdn地址',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_username`(`username`) USING BTREE COMMENT '用户名索引'
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of t_user
-- ----------------------------
INSERT INTO `t_user` VALUES (1, 'xmg', 'mgh', 'e10adc3949ba59abbe56e057f20f883e', '691639910@qq.com', 'http://192.168.44.100/images/firstPic/avatar1.png', 1, '2021-02-01 18:25:26', '2022-03-02 11:52:19', 'https://www.baidu.com/', 'https://blog.csdn.net/Peerless__?type=blog');

SET FOREIGN_KEY_CHECKS = 1;
