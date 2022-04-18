
<template>

    <div>
        <!--上方背景页面-->
        <div id="bg1" :style="bgObj" >
            <bubbles-effect :options="options" class="bubble"></bubbles-effect>
            <div id="motto" v-if="showMotto">
                <transition-group appear
                name="animate__animated animate__bounce animate__slow"
                enter-active-class="animate__bounceIn"
                leave-active-class="animate__bounceOut"
                >
                    <h1 key="1">{{firstBGPageInfo.curMotto.ch}}</h1>
                    <p key="2">{{firstBGPageInfo.curMotto.en}}</p>
                </transition-group>
            </div>
            <div class="scroll-down">
                <a class="anchor-down" @click.prevent="anchorDown"></a>
            </div>
        </div>



        <!--博客展示区-->
        <div id="area-blow">
            <div class="blog-area">

                <!--左侧博客信息区-->
                <div class="blog-left">
                    <!--用户信息栏-->
                    <UserInfoCard :count="blogCount"></UserInfoCard>
                    <!--最新推荐栏-->
                    <RecommendList :blogList="newRecommend"></RecommendList>
                    <!--标签云-->
                    <TagCloud></TagCloud>
                    <!--热门推荐-->
                    <RecommendList :blogList="hotBlogs"></RecommendList>
                </div>

                <!--右侧博客展示区-->
                <div class="blog-right">
                    <BlogCard :key="item.id" v-for="(item, index) in blogDetails"
                               :item="item" :imgRight="index % 2 === 0"
                                @click.native="blogDetail(item.id)">
                    </BlogCard>
                    <!--分页导航区-->
                    <Pagination @jumpPage="jumpPage" :pageInfo="{pageNum:queryInfo.pageNum, pages:pages}"></Pagination>

                </div>
            </div>
        </div>
    </div>

</template>



<script>
import img1 from "../assets/1.png"
import img2 from "../assets/2.jpg"
import img3 from "../assets/3.jpg"

import BackToTop from "../components/BackToTop";
import BlogCard from "../components/BlogCard";
import UserInfoCard from "../components/UserInfoCard";
import Pagination from "../components/Pagination";
import RecommendList from "../components/RecommendList";
import TagCloud from "../components/TagCloud";

import "animate.css"
// import "../assets/js/ribbons"

export default {
    name: "Home",
    components: {BackToTop, BlogCard, Pagination, UserInfoCard, RecommendList, TagCloud},
    data() {
        return {
            options: {
                color: 'random', //Bubble color
                radius: 15, //Bubble radius
                densety: 0.2, // The larger the bubble density, the greater the density (suggest no more than 1).
                clearOffset: 0.2 // The larger the bubble disappears [0-1], the longer it disappears.
            },
            firstBGPageInfo: {
                curBg: "",
                bgs: [],
                curMotto: {
                    ch: "奋斗从未停止, 前进永无止境!",
                    en: "Struggle never stops, and progress never ends!"
                },
                mottos:[],
            },
            showCat: false,
            showMotto: true,
            pages: 1,              // 页面数量
            blogCount: 66,
            queryInfo: {
              pageNum: 1,
              pageSize: 8
            },
            blogDetails: [{
              id: 1,
              title: "博客系统开发",
              description: "本章主要介绍博客系统的开发工作...",
              firstPicture: img1,
              createTime: "2021-03-29",
              views: 1024,
              nickname: "Jack Ma",
              typename: "框架底层原理"
              },
              {
                id: 2,
                title: "博客系统开发",
                description: "本章主要介绍博客系统的开发工作...",
                firstPicture: img2,
                createTime: "2021-03-29",
                views: 1024,
                nickname: "Jack Ma",
                typename: "框架底层原理"
              },
              {
                id: 3,
                title: "博客系统开发",
                description: "本章主要介绍博客系统的开发工作...",
                firstPicture: img3,
                createTime: "2021-03-29",
                views: 1024,
                nickname: "Jack Ma",
                typename: "框架底层原理"
              }],
            newRecommend: {
                title: "最新推荐",
                icon: "icon-zuixingengxin",
                list: [
                    {id: 1, title: "MarkDownit使用"},
                    {id: 2, title: "Vue日期格式化过滤器"},
                    {id: 3, title: "前后端分离登录验证"},
                    {id: 4, title: "MavonEditor上传图片"},
                    {id: 5, title: "Springboot中PageHelper 分页查询使用方法"},
                    {id: 6, title: "mybatis根据日期查询、查询日期并返回"},
                    {id: 7, title: "maven中静态资源的过滤"}
                ]
            },
            hotBlogs: {
                title: "热门推荐",
                icon: "icon-fire",
                list: [
                    {id: 1, title: "MarkDownit使用"},
                    {id: 2, title: "Vue日期格式化过滤器"},
                    {id: 3, title: "前后端分离登录验证"},
                    {id: 4, title: "MavonEditor上传图片"},
                    {id: 5, title: "Springboot中PageHelper 分页查询使用方法"},
                    {id: 6, title: "mybatis根据日期查询、查询日期并返回"},
                    {id: 7, title: "maven中静态资源的过滤"}
                ]
            },

        }
    },
    created() {
        this.getBlogLists();
        this.getBackgImage()
        this.getMotto()
        this.getNewBlogs()
        this.getHotBlogs()
    },
    deactivated() {
        this.showMotto = false
    },
    activated() {
        this.showMotto = true
        window.scrollTo(0, 0)
    },
    computed: {
        bgObj() {
            if (this.firstBGPageInfo.curBg === "") {
                return ""
            }
            return {backgroundImage: 'url(' + this.firstBGPageInfo.curBg + ')'}
        }
    },
    methods: {
        // 获取背景图片地址
        async getBackgImage() {
            const {data: res} = await this.$axios.get("/myblog/bgs")
            if(res.status === 1) { //查询成功
                if (res.data.length > 0) {
                    this.firstBGPageInfo.bgs = res.data[0]
                    const n = Math.round(Math.random() * (res.data[0].length - 1));
                    this.firstBGPageInfo.curBg = this.firstBGPageInfo.bgs[n]
                }

            }
        },
        // 获取座右铭
        async getMotto() {
            const {data: res} = await this.$axios.get("/myblog/mottos")
            if(res.status === 1) {
                if (res.data.length > 0) {
                    this.firstBGPageInfo.mottos = res.data[0]
                    const n = Math.round(Math.random() * (res.data[0].length - 1));
                    this.firstBGPageInfo.curMotto = this.firstBGPageInfo.mottos[n]
                }
            }
        },
        anchorDown() {
            const offsetTop = document.getElementById("area-blow").offsetTop;
            window.scrollTo({top: offsetTop, behavior: 'smooth'})
        },
        // 获取博客列表
        async getBlogLists() {
            var keyWord = window.sessionStorage.getItem("keyWord");
            if(keyWord) {
                window.sessionStorage.removeItem("keyWord");
                const {data:res} = await this.$axios.get("/myblog/search", {params: {pageNum: this.queryInfo.pageNum, pageSize: this.queryInfo.pageSize, keyWord: keyWord}});
                if(res.status === 1) {
                    this.blogDetails = res.data.length > 0 ? res.data[0] : this.blogDetails
                    this.blogCount = res.data.length > 1 ? res.data[1] : this.blogCount
                    this.pages = Math.ceil(this.blogCount / this.queryInfo.pageSize)
                    if (this.pages <= 0) {
                        this.pages = 1
                    }
                }else {
                    this.$message.error("获取博客失败，请重试")
                }
            } else {
                const {data:res} = await this.$axios.get("/myblog/blogLists", {params: this.queryInfo});
                if(res.status === 1) {
                    this.blogDetails = res.data.length > 0 ? res.data[0] : this.blogDetails
                    this.blogCount = res.data.length > 1 ? res.data[1] : this.blogCount
                    this.pages = Math.ceil(this.blogCount / this.queryInfo.pageSize)
                    if (this.pages <= 0) {
                        this.pages = 1
                    }
                }else {
                    this.$message.error("获取博客失败，请重试")
                }
            }
        },
        // 获取最新推荐博客
        async getNewBlogs() {
            const {data:res} = await this.$axios.get("/myblog/newBlogs", {params: {countLimit: 10}});
            if(res.status === 1) {
                this.newRecommend.list = res.data.length > 0 ? res.data[0] : this.newRecommend.list;
            } else {
                this.$message.warning("获取最新推荐失败！")
            }
        },
        // 获取热门推荐博客
        async getHotBlogs() {
            const {data:res} = await this.$axios.get("/myblog/hotBlogs", {params: {countLimit: 10}});
            if(res.status === 1) {
                this.hotBlogs.list = res.data.length > 0 ? res.data[0] : this.hotBlogs.list;
            } else {
                this.$message.warning("获取热门推荐失败！")
            }
        },
        // 进入博客内容页面
        blogDetail(id) {
            this.$router.push({
                path: "/blogDetail",
                query: {
                    id: id
                }
            });
        },
        jumpPage(pageNum) {
            this.queryInfo.pageNum = pageNum;
            this.getBlogLists();
        }
    }
}
</script>



<style lang="less" scoped>

.animate__animated{
    animation-duration: 3s !important;
}

.bubble {
    width: 100%;
}

//座右铭部分
#motto {
    user-select: none;
    font-weight: 560;
    line-height: 1.25;
    color: white;
    text-align: center;
    position: relative;
    top: 42%;
    h1 {
        padding-bottom: 20px;
    }

    p {
        font-size: 22px;
    }
}


.scroll-down {
    width: 100%;
    height: 60px;
    position: relative;
    top: 75%;
    text-align: center;

    animation: bounce-in 5s 1s infinite;
}

.anchor-down {
    display: block;
    width: 36px;
    height: 36px;
    transform: rotate(45deg);
    position: absolute;
    left: 50%;
    margin-left: -18px;
    cursor: pointer;
}

.anchor-down::before {
    content: '';
    width: 28px;
    height: 28px;
    position: absolute;
    right: 14px;
    bottom: 14px;
    border-right: 5px solid #fccb90;
    border-bottom: 5px solid #fccb90;
    display: block;
    border-bottom-right-radius: 2px;
}

.anchor-down::after {
    content: '';
    width: 28px;
    height: 28px;
    position: absolute;
    right: 0;
    bottom: 0;
    border-right: 5px solid #84fab0;
    border-bottom: 5px solid #84fab0;
    display: block;
    border-bottom-right-radius: 2px;
}

@keyframes bounce-in {
    0% { transform: translateY(0); }
    20% { transform: translateY(0); }
    50% { transform: translateY(-20px);  border-right: #81f5ac; border-top:#81f5ac;}
    80% { transform: translateY(0); }
    100% { transform: translateY(0); }
}

//上方背景图
#bg1{
    width: 100%;
    height: 100vh;
    background: url('~@/assets/images/back4.png') 0 0 / cover no-repeat;
    overflow: hidden;
}

//下面区域
#area-blow {
    //background: url('~@/assets/images/back3.jpg') 0 0 / cover no-repeat;
    background-image: linear-gradient(to left, rgba(255, 0, 149, 0.2), rgba(0, 247, 255, 0.2));
    background-attachment: fixed;
}

// 下面中心区域
.blog-area {
    width: 1165px;
    margin: 0 auto;
    padding-top: 36px;
    padding-bottom: 64px;
    overflow: hidden;
}

.card {
    height: 260px;
    position: relative;
    background: #FFF;
    box-shadow: 0 1px 2px 0 rgba(34,36,38,.15);
    margin: 1rem 0;
    padding: 0px 24px !important;
    border-radius: .28571429rem;
    border: 1px solid rgba(34,36,38,.15);
    //opacity: 0.9;
    padding: 1.5em;
    font-size: 1rem;
    img {
        width: 360px !important;
    }
}

.blog-right {
    float: left;
}

.blog-left {
    float: left;
    width: 300px;
    margin-right: 15px;
    padding: 0 15px;
    border-radius: 6px;
    transition: all 0.3s;
}


</style>
