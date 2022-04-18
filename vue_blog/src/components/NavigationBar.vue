<!--顶部导航栏组件-->
<template>
    <div class="navi-bar">
        <!-- 小站名字-->
        <p>{{ siteName }}</p>

        <!--中间搜索、导航区域-->
        <div class="navi-bar-center">
            <div class="navi-bar-item-group">
                <i @click="fullscreen" class="iconfont icon-quanpingmu fullscreen"></i>
                <!--搜索框-->
                <div class="search-warp">
                    <el-input
                        class="search-input"
                        size="mini"
                        placeholder="请输入内容"
                        prefix-icon="el-icon-search"
                        v-model="keyWord"
                        @keyup.enter.native="search"
                    >
                    </el-input>
                    <div @mouseleave="showResult=false;keyWord=''" v-show="showResult" class="search-result">
                        <div class="search-content">
                            <ul>
                                <li>{{ total ? '共' + total + '条记录！' : '没有找到记录！' }}</li>
                                <li :key="item.id" v-for="item in searchResult">
                                    <a @click="blogDetail(item.id)">{{ item.title }}</a>
                                </li>
                            </ul>
                        </div>
                    </div>
                </div>
                <!--导航按键-->
                <ul>
                    <li><NaviItem :activeRoute="activeRoute" index-name="首页" ico-class="icon-shouye" rout="/home"></NaviItem></li>
                    <li><NaviItem :activeRoute="activeRoute" index-name="分类" ico-class="icon-fenlei" rout="/types"></NaviItem></li>
                    <li><NaviItem :activeRoute="activeRoute" index-name="标签" ico-class="icon-biaoqian" rout="/tags"></NaviItem></li>
                    <li><NaviItem :activeRoute="activeRoute" index-name="归档" ico-class="icon-24gf-hourglass" rout="/timeLine"></NaviItem></li>
                    <li><NaviItem :activeRoute="activeRoute" index-name="随笔" ico-class="icon-yongyan" rout="/essay"></NaviItem></li>
                    <li><NaviItem :activeRoute="activeRoute" index-name="留言板" ico-class="icon-liuyanban" rout="/msgBoard"></NaviItem></li>
                    <li v-if="false"><NaviItem :activeRoute="activeRoute" index-name="照片墙" ico-class="icon-zhaopianxuanzhong" rout="/photoWall"></NaviItem></li>
                    <li><NaviItem :activeRoute="activeRoute" index-name="资源库" ico-class="icon-ziyuanku" rout="/resourceLib"></NaviItem></li>
                    <li><NaviItem :activeRoute="activeRoute" index-name="关于我" ico-class="icon-guanyuwomen" rout="/about"></NaviItem></li>
                    <li><NaviItem :activeRoute="activeRoute" index-name="后台管理" ico-class="icon-houtaiguanli-jifenguanli" rout="/login"></NaviItem></li>
                </ul>
            </div>
        </div>

    </div>

</template>

<script>

import "../assets/icon/iconfont.css"

import NaviItem from "./NaviItem";

export default {
    name: "NavigationBar",
    components: { NaviItem },
    data() {
        return {
            keyWord: "",
            siteName: "mgh的小站",
            activeRoute: "/home",
            searchResult: [],
            total: 0,
            showResult: false
        }
    },
    methods: {
        fullscreen() {
            const de = document.documentElement;
            if(de.requestFullscreen) {
                de.requestFullscreen();
            }
            window.scrollTo(0, 1);
        },
        async search() {
            if (!this.keyWord) {
                return
            }
            const {data:res} = await this.$axios.get("/myblog/search", {params: {keyWord: this.keyWord}})
            if (res.status !== 1) {
                this.$message.error("搜索失败，请重试！")
                return
            }
            if (res.data.length > 0) {
                this.searchResult = res.data[0]
                if (res.data.length > 1) {
                    this.total = res.data.length
                }
            }
            this.showResult = true
        },
        blogDetail(id) {
            this.showResult = false
            this.$router.push({
                path: "/blogDetail",
                query: {
                    id: id
                }
            });
        }
    },
    created() {
        this.activeRoute = this.$router.currentRoute.path
    },
    watch: {
        // 监听路由改变，路由改变时给activeRoute赋值，从而让导航栏组件高亮
        $route(to) {
            this.activeRoute = to.path
        }
    }

}
</script>


<style lang="less">
    .search-input {
        width: 200px !important;
        border: None;

    input {
        width: 200px;
        border: 1px solid #fda085 !important;
        border-radius: 20px;
        background: rgba(210,249,253,0.3);
    }

}


</style>

<style lang="less" scoped >

ul {
    margin: 0;
    display: inline-block;
}

li {
    display: inline-block;
    list-style: none;
    margin: 0 15px;
}

.search-warp {
    position: relative;
    display: inline-block;
}

.search-result {
    width: 240px;
    position: absolute;
    background-color: #fff;
    top: 45px;
    left: -20px;
    border-radius: 8px;

    .search-content {
        overflow-y:auto;
        max-height: 300px;
        margin: 8px 0;
    }
}

.search-result::before {
    content: '';
    position: absolute;
    display: inline-block;
    width: 16px;
    height: 16px;
    border-left: 1px solid #fff;
    border-top: 1px solid #fff;
    transform: rotate(45deg);
    top: -8px;
    left: 50px;
    background-color: #fff;
    border-top-left-radius: 3px;
}

.search-result ul  {
    padding-left: 16px;
    font-size: 14px;

    li {
        display: block;
        cursor: pointer;
        margin: 2px 0;

        a {
            outline: none;
            text-decoration: none;
            color: #61666d;
        }
    }

    li:first-child {
        cursor: default;
    }

}

p {
    float: left;
    line-height: 1.5;
    text-align: center;
    font-size: 24px;
    align-items: center;
    margin: 10px 0 0 50px;
    height: 40px;
    background: -webkit-linear-gradient(45deg,#70f7fe,#fbd7c6,#fdefac,#bfb5dd,#bed5f5);
    -webkit-background-clip: text;
    font-weight: bolder;
    color: transparent;
    user-select: none;
}

.navi-bar {
    z-index: 999;
    width: 100%;
    height: 60px;
    position: fixed;
    background: rgba(210,249,253,0.3);
}

.navi-bar-center {
    margin: 10px 0;
    height: 40px;
    width: 100%;
}

.navi-bar-item-group {
    display: inline-block;
    align-items: center;
    height: 100%;
    float: right;
    margin-right: 100px;
}

.fullscreen {
    color: #fda085;
    font-size: 20px;
    cursor: pointer;
    margin-right: 20px;
}

</style>
