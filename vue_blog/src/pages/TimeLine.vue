
<template>

    <div>
        <TitleArea class="title-area" :info="info"></TitleArea>
        <div class="main">
            <div class="center-area">
                <!-- 左边时间线 -->
                <div class="timeline">
                    <div class="article-title">文章列表</div>
                    <div class="article-list">
                        <TimeAxis :key="index" v-for="(item, index) in blogs" :list="item"></TimeAxis>
                    </div>
                </div>
                <!-- 右侧边栏 -->
                <div class="right-area">
                    <BlogTypePie></BlogTypePie>
                    <TagCloud class="tag-cloud"></TagCloud>
                </div>
            </div>
        </div>
    </div>

</template>



<script>

import "../assets/icon/iconfont.css"

import TitleArea from "../components/TitleArea";
import TimeAxis from "../components/TimeAxis";
import BlogTypePie from "../components/BlogTypePie";
import TagCloud from "../components/TagCloud";

export default {
    components: { TitleArea, TimeAxis, BlogTypePie, TagCloud },
    data() {
        return {
            info: {
                title: "归档",
                desc: "恰同学少年，风华正茂！"
            },
            blogsTotal: 3,
            blogs: [
                    [{
                    id: 1,
                    title: "上邪~~",
                    flag: "原创",
                    createTime: "2021-04-30T13:20:34.000+00:00"
                }],
                [{
                    id: 2,
                    title: "上邪~~",
                    flag: "原创",
                    createTime: "2021-03-30T13:20:34.000+00:00"
                }],
                [{
                    id: 3,
                    title: "上邪~~",
                    flag: "原创",
                    createTime: "2020-03-30T13:20:34.000+00:00"
                }]
            ],
        }
    },
    created() {
        window.scrollTo(0, 0)
        this.getBlogList();
    },
    methods: {
        getBlogList: async function() {
            const {data: res} = await this.$axios.get("/myblog/timeLine");
            this.blogs = res.data.length > 0 ? res.data[0] : this.blogs;
            this.blogsTotal = 0;
            for(let i = 0; i < this.blogs.length; i++) {
                this.blogsTotal += this.blogs[i].length;
            }
        },
    },
}
</script>



<style lang="less" scoped>

ul, li {
    margin: 0;
    padding: 0;
}

//标题区域
.title-area {
    background: url("../assets/images/back5__2496_520.png") 0 0 / cover no-repeat;
}

//下方整个区域
.main {
    min-height: 1000px;
    width: 100%;
    background: url("../assets/images/back6.png") 0 0 / cover no-repeat;
    background-attachment: fixed;
    padding-top: 50px;
}

//下方中心区域
.center-area {
    width: 1200px;
    margin: 0 auto;
}

.timeline {
    width: 840px;
    background-color: #fff;
    padding: 30px 0;
    border-radius: 6px;
    float: left;
}

.article-title {
    font-size: 36px;
    font-weight: 500;
    font-family: '楷体';
    text-align: center;
}


.right-area {
    width: 320px;
    float: left;
    margin-left: 26px;
}

.tag-cloud {
    width: 300px;
    margin: 24px 0 0;
}

.article-list {
    padding: 30px 0 30px 90px;
}


</style>
