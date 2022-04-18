
<template>

    <div class="type-bg">
        <div class="title" align="center">文章分类</div>
        <!--博客分类列表-->
        <transition appear
            name="animate__animated animate__bounce animate__slow"
            enter-active-class="animate__zoomIn"
            leave-active-class="animate__zoomOut"
        >
            <div class="type-area">
                <ul>
                    <li class="li-item" :key="item.id" v-for="item in types">
                        <BlogTypeItem @click.native="getBlogList(item.id)" :typeinfo="item" :activeId="currentTypeId"></BlogTypeItem>
                    </li>
                </ul>
            </div>
        </transition>


        <transition appear
                    name="animate__animated animate__bounce animate__slow"
                    enter-active-class="animate__fadeInLeft"
        >
            <div class="wapper">
                <!--博客列表栏-->
                <div class="blog-area">
                    <BlogCard class="blog-item" :key="item.id" v-for="(item, index) in blogDetails"
                              :item="item" :imgRight="index % 2 === 0"
                              @click.native="blogDetail(item.id)">
                    </BlogCard>
                </div>
                <!--分页导航栏-->
                <Pagination @jumpPage="jumpPage" :pageInfo="{pageNum:queryInfo.pageNum, pages:pages}"></Pagination>
            </div>
        </transition>

    </div>

</template>



<script>

import img1 from "../assets/1.png"
import img2 from "../assets/2.jpg"
import img3 from "../assets/3.jpg"

import BlogCard from "../components/BlogCard";
import Pagination from "../components/Pagination";
import BlogTypeItem from "../components/BlogTypeItem";


export default {
    components: { BlogCard, Pagination, BlogTypeItem },
    data() {
        return {
            currentTypeId: 0,
            pages: 1,              // 页面数量
            queryInfo: {
                pageNum: 1,
                pageSize: 5,
                typeId: 0
            },
            types: [{
                    id: 1,
                    name: "我的故事",
                    count: 5
                },
                {
                    id: 2,
                    name: "前端",
                    count: 8
                },
                 {
                    id: 3,
                    name: "后端",
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
                },
            ]
        }
    },
    created() {
        window.scrollTo(0, 0)
        this.getTypeList();
    },
    methods: {
        getTypeList: async function() {
            const {data: res} = await this.$axios.get("/myblog/typeList");
            if(res.status === 1) {
                this.types = res.data.length > 0 ? res.data[0] : this.types;
                this.types.sort((prev, next) => {
                    return next.count - prev.count
                })
            }

            await this.getBlogList(this.types[0].id);
        },
        getBlogList: async function(id) {
            this.currentTypeId = id;
            this.queryInfo.typeId = id;
            const {data: res} = await this.$axios.get("/myblog/typeBlogList", {params: this.queryInfo});
            if(res.status === 1) {
                this.blogDetails = res.data.length > 0 ? res.data[0] : this.blogDetails;
            } else {
                this.$message.error("获取博客失败，请重试！")
                return
            }

            const count = res.data.length > 1 ? res.data[1] : 0
            this.pages = Math.ceil(count / this.queryInfo.pageSize);
            if (this.pages <= 0) {
                this.pages = 1
            }
        },
        blogDetail(id) {
            this.$router.push({
                path: "/blogDetail",
                query: {
                    id: id
                }
            });
        },
        jumpPage(pageNum) {
            window.scrollTo(0, 0)
            this.queryInfo.pageNum = pageNum;
            this.getBlogList(this.currentTypeId);
        }
    },
}

</script>



<style lang="less" scoped>

.li-item {
    display: inline-block;
}

.type-area {
    display: block;
    width: 840px;
    margin: 0 auto;
    vertical-align: middle;
    padding-bottom: 24px;
}

.type-bg {
    background: url('~@/assets/images/back2.png') 0px 0px / cover no-repeat;
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

.wapper {
    padding-bottom: 50px;
}

.blog-area {
    width: 840px;
    margin: 0 auto;
    padding: 0 24px;
}

</style>
