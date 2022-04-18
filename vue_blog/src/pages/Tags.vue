
<template>

    <div class="tag-bg">
        <div class="title" align="center">文章标签</div>

        <!-- 标签栏 -->
        <transition appear
                    name="animate__animated animate__bounce animate__slow"
                    enter-active-class="animate__fadeInDown"
                    leave-active-class="animate__fadeOutUp"
        >
            <ul class="tag-area">
                <li :key="item.id" v-for="item in tags">
                    <BlogTagItem @click.native="getBlogList(item.id)" :taginfo="item" :activeId="currentTagId"></BlogTagItem>
                </li>
            </ul>
        </transition>

        <transition appear
                    name="animate__animated animate__bounce animate__slow"
                    enter-active-class="animate__fadeInUp"
        >
            <div>
                <!-- 博客列表栏 -->
                <div class="blog-area">
                    <BlogCard class="blog-item" :key="item.id" v-for="(item, index) in blogDetails"
                              :item="item" :imgRight="index % 2 === 0"
                              @click.native="blogDetail(item.id)">
                    </BlogCard>
                </div>
                <Pagination class="pagebar" @jumpPage="jumpPage" :pageInfo="{pageNum:queryInfo.pageNum, pages:pages}"></Pagination>
            </div>
    </transition>

    </div>
</template>



<script>

import img1 from "../assets/1.png"

import BlogCard from "../components/BlogCard";
import Pagination from "../components/Pagination";
import BlogTagItem from "../components/BlogTagItem";

export default {
    components: { BlogCard, Pagination, BlogTagItem },
    data() {
        return {
            currentTagId: 0,
            pages: 1,              // 页面数量
            queryInfo: {
                pageNum: 1,
                pageSize: 5,
                tagId: 0
            },
            tags: [{
                    id: 1,
                    name: "前端",
                    count: 5
                },
                {
                    id: 2,
                    name: "后端",
                    count: 8
                },
                 {
                    id: 3,
                    name: "SpringBoot",
                    count: 5
                }
            ],
            blogDetails: [{
                id: 1,
                title: "博客系统开发",
                description: "本章主要介绍博客系统的开发工作...",
                firstPicture: img1,
                nickname: "Jack Ma",
                avatar: "https://placekitten.com/300/300",
                typename: "框架底层原理",
                createTime: "2021-03-29",
                views: 1024
            },
            {
                id: 2,
                title: "博客系统开发",
                description: "本章主要介绍博客系统的开发工作...",
                firstPicture: img1,
                nickname: "Jack Ma",
                avatar: "https://placekitten.com/300/300",
                typename: "框架底层原理",
                createTime: "2021-03-29",
                views: 1024
            },
            {
                id: 3,
                title: "博客系统开发",
                description: "本章主要介绍博客系统的开发工作...",
                firstPicture: img1,
                nickname: "Jack Ma",
                avatar: "https://placekitten.com/300/300",
                typename: "框架底层原理",
                createTime: "2021-03-29",
                views: 1024
            }],
        }
    },
    created() {
        window.scrollTo(0, 0)
        if (typeof(this.$route.query.id) !== "undefined") {
            this.getTagList(true)
        } else {
            this.getTagList(false);
        }
    },
    methods: {
        async getTagList(flag) {
            const {data: res} = await this.$axios.get("/myblog/tagList");
            if(res.status === 1) {
                this.tags = res.data.length > 0 ? res.data[0] : this.tags;
            }

            if (flag) {
                await this.getBlogList(this.$route.query.id);
            } else {
                await this.getBlogList(this.tags[0].id)
            }
        },
        async getBlogList(id) {
            this.currentTagId = id;
            this.queryInfo.tagId = id;
            const {data: res} = await this.$axios.get("/myblog/tagBlogList", {params: this.queryInfo});
            if(res.status === 1) {
                this.blogDetails = res.data.length > 0 ? res.data[0] : this.blogDetails;
            } else {
                this.$message.error("获取博客失败，请重试")
                return
            }

            const count = res.data.length > 1 ? res.data[1] : 0
            this.pages = Math.ceil(count / this.queryInfo.pageSize);
            if (this.pages <= 0) {
                this.pages = 1
            }
        },
        jumpPage(pageNum) {
            window.scrollTo(0, 0)
            this.queryInfo.pageNum = pageNum;
            this.getBlogList(this.currentTagId);
        },
        blogDetail(id) {
            this.$router.push({
                path: "/blogDetail",
                query: {
                    id: id
                }
            });
        }
    }
}

</script>



<style lang="less" scoped>

ul, li {
    margin: 0;
    padding: 0;
}

.tag-bg {
    background: url('~@/assets/images/bg11.jpg') 0 0 / cover no-repeat;
    background-attachment: fixed;
    min-height: 1000px;
}

.title {
    font-size: 450%;
    color: #ffffff;
    margin-bottom: 50px;
    bottom: 0 !important;
    right: 0 !important;
    font-family:'STXingkai';
    opacity: 0.5;
    padding-top: 6%;
}

.tag-area {
    width: 840px;
    margin: 0 auto;
    overflow: hidden;
}

.tag-area li {
    list-style: none;
    float: left;
}

.blog-area {
    width: 850px;
    margin: 40px auto 50px;
}

.clearfix::after {
    content: '';
    display: block;
    height: 0;
    clear: both;
    visibility: hidden;
}

.pagebar {
    padding-bottom: 50px;
}

</style>
