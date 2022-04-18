
<template>

    <div class="blog_bg">
        <!-- 标题 -->
        <div class="title">
            <h1>{{ blog.title }}</h1>
        </div>

        <!-- 博客信息 -->
        <div class="blogInfo">
            <span class="flag">{{ blog.flag }}</span>
            <span>
                <b-icon icon="person" font-scale="0.9"></b-icon>
                {{ blog.nickname }}
            </span>
            <span>
                <b-icon icon="calendar2-check" font-scale="1"></b-icon>
                {{ (blog.createTime || '').split('T')[0] }}
            </span>
            <span>
                <b-icon icon="journals" font-scale="0.8"></b-icon>
                {{ blog.typename }}
            </span>
            <span>
                <b-icon icon="eye" font-scale="1"></b-icon>
                {{ blog.views }}
            </span>
        </div>

        <!-- 正文 -->
        <div class="content-outer">
            <div class="tags">
                <b-icon icon="tags" font-scale="1.2"></b-icon>
                <span :key="item.id" v-for="item in tags"> {{ item.name }} </span>
            </div>
            <div class="content-inner" v-html="blog.content">

            </div>

            <!-- 赞赏 -->
            <div align="center" v-if="blog.appreciation">
                <b-button id="popover-button-variant" pill variant="outline-info">赞赏</b-button>
                <b-popover target="popover-button-variant"  placement="bottom" variant="info" triggers="focus">
                    <img :src="require('../assets/alipay1.jpg')" width="120" height="140" >
                    <img :src="require('../assets/wechat1.png')" width="120" height="140" >
                </b-popover>
            </div>
        </div>

        <!-- 版权声明 -->
        <div class="rights">
            <ul class="list">
                <li>作者：{{blog.nickname}}</li>
                <li>更新时间：{{ blog.updateTime | fromatDate("yyyy-MM-dd hh:mm:ss") }} </li>
                <li>版权声明：自由转载-非商用-非衍生-保持署名</li>
                <li>转载声明：如果是转载栈主转载的文章，请附上原文链接</li>
            </ul>
        </div>

        <!-- 评论 -->
        <div class="comment-outer">
             <div class="comment-body">
                <div style="font-weight: bold">评论</div>
                <hr>

                <!-- 评论内容 -->
                <div class="comment" :key="item.id" v-for="item in commentList">

                    <b-avatar class="comment-avatar" :src="item.avatar"></b-avatar>
                    <!-- 昵称和日期 -->
                    <!-- <div class="comment-author"> -->
                        <span class="author">
                            {{ item.nickname }}
                        </span>
                        <span class="date">
                            {{ item.createTime | fromatDate("yyyy-MM-dd hh:mm:ss") }}
                        </span>
                    <!-- </div> -->
                    <!-- 回复内容 -->
                    <div class="comment-content">
                        {{ item.content }}
                    </div>
                </div>

            </div>

            <!-- 回复栏 -->
            <textarea class="replay" v-model="comment.content">
            </textarea>

            <b-container fluid>
                <b-row >
                    <b-col sm="3" style="padding-left: 0;">
                        <b-form-input  type="text" placeholder="昵称" v-model="comment.nickname">
                        </b-form-input>
                    </b-col>
                    <b-col sm="3">
                        <b-form-input  type="email" placeholder="邮箱" v-model="comment.email">
                        </b-form-input>
                    </b-col>
                    <b-col sm="3">

                        <b-button variant="outline-info" @click="publishComment">
                            <b-icon icon="pencil-square" font-scale="0.9"></b-icon>
                            发布
                        </b-button>
                    </b-col>
                </b-row>
            </b-container>
        </div>

        <div style="height: 200px;"></div>
    </div>

</template>



<script>
export default {
    data() {
        return {
            blog: {
                id: 0,
                title: "上邪~~",
                content: `上邪,<br> 我欲与君相知，<br>长命无绝衰。<br>山无陵，<br>江水为竭。<br>冬雷震震，<br>夏雨雪。<br>天地合，<br>乃敢与君绝。`,
                flag: "原创",
                updateTime: "2021-03-30",
                createTime: "2021-03-30",
                views: 1024,
                appreciation: false,
                typename: "框架底层原理",
                nickname: "Jack Ma"
            },
            tags: [
                {id: 1, name: "Java"},
                {id: 2, name: "Spring"},
                {id: 3, name: "框架"}
            ],
            comment: {         //评论数据
                nickname: "",
                email: "",
                content: "",
                avatar: "http://localhost:8080/images/firstPic/avatar1.png",
                blogId: 0
            },
            commentList: [
                {
                    id: 1,
                    nickname: "小兔叽",
                    email: "rabbit@blog.com",
                    content: "very good!",
                    avatar: "http://localhost:8080/firstPic/avatar1.png",
                    createTime: "2021-03-31 22:35",
                    blogId: 0,
                },
                {
                    id: 2,
                    nickname: "小脑斧",
                    email: "tiger@blog.com",
                    content: "excellent!",
                    avatar: "http://localhost:8080/firstPic/avatar1.png",
                    createTime: "2021-03-31 22:35",
                    blogId: 0,
                }
            ]
        }
    },
    created() {
        window.scrollTo(0, 0);
        this.getBlogDetails();
    },
    filters: {
        fromatDate: function(val, fmt) {
            var date = new Date(val);
            let ret;
            const opt = {
                "y+": date.getFullYear().toString(),        // 年
                "M+": (date.getMonth() + 1).toString(),     // 月
                "d+": date.getDate().toString(),            // 日
                "h+": date.getHours().toString(),           // 时
                "m+": date.getMinutes().toString(),         // 分
                "s+": date.getSeconds().toString()          // 秒
                // 有其他格式化字符需求可以继续添加，必须转化成字符串
            };
            for (let k in opt) {
                ret = new RegExp("(" + k + ")").exec(fmt);
                if (ret) {
                    fmt = fmt.replace(ret[1], (ret[1].length == 1) ? (opt[k]) : (opt[k].padStart(ret[1].length, "0")))
                };
            };
            return fmt;
        }
    },
    methods: {
        getBlogDetails: async function() {
            let blogId = this.$route.query.id;
            blogId = parseInt(blogId);
            const {data: res} = await this.$axios.get("/myblog/detailedBlog", {params: {id: blogId}});
            if(res.status === 1) {
                this.blog = res.data.length > 0 ? res.data[0] : this.blog;
                this.tags = res.data.length > 1 ? res.data[1] : this.tags;
                var hljs = require('highlight.js');
                var md = require('markdown-it')({
                                    html: true,
                                    linkify: true,
                                    typographer: true,
                                    highlight: function (str, lang) {
                                        if (lang && hljs.getLanguage(lang)) {
                                            try {
                                                return '<pre class="hljs"><code>' +
                                                    hljs.highlight(lang, str, true).value +
                                                    '</code></pre>';
                                            } catch (__) {}
                                        }

                                        return '<pre class="hljs"><code>' + md.utils.escapeHtml(str) + '</code></pre>';
                                    }

                });

                // 让图片居中
                md.renderer.rules.image = (token, idx, options, env, self) => {
                    return "<div align='center' class='blog_image'>" + md.renderer.renderToken(token, idx, options, env, self) + "</div>"
                }
                this.blog.content = md.render(this.blog.content);

                this.getCommentList();
            }
        },
        publishComment: async function () {
            this.comment.blogId = this.blog.id;
            if(!this.comment.content) {
                this.$message.warning("请输入评论内容")
                return
            }
            if(!this.comment.nickname) {
                this.$message.warning("请输入您的称呼")
                return
            }
            if(!this.comment.email) {
                this.$message.warning("请输入邮箱")
                return
            }
            if(!this.checkEmail()) {
                this.$message.warning("邮箱不合法")
                return
            }

            await this.$axios.post("/myblog/publishComment", this.comment);
            this.comment.nickname = "";
            this.comment.email = "";
            this.comment.content = "";

            this.getCommentList();
        },
        getCommentList: async function() {
            const {data: res} = await this.$axios.get("/myblog/commentList", {params: {id: this.blog.id}});
            if(res.status === 1) {
                this.commentList = res.data.length > 0 ? res.data[0] : this.commentList;
            }
        },
        checkEmail: function() {
            var reg = new RegExp("^[a-z0-9]+([._\\-]*[a-z0-9])*@([a-z0-9]+[-a-z0-9]*[a-z0-9]+.){1,63}[a-z0-9]+$"); //正则表达式
            if(this.comment.email === ""){ //输入不能为空
             　　　　alert("输入不能为空!");
             　　　　return false;
             　　}else if(!reg.test(this.comment.email)){ //正则验证不通过，格式不对
             　　　　return false;
             　　}else{
             　　　　return true;
             　　}
            }

    }
}
</script>

<style>

/*改变博客中的图片的最大宽度*/
.blog_image img {
    max-width: 1000px !important;
    cursor: pointer;
}

</style>

<style lang="less" scoped>

.blog_bg {
    background: url('~@/assets/images/bg12.png') 0px 0px / cover no-repeat;
    background-attachment: fixed;
    min-height: 1000px;
    padding-top: 8%;
}

.title {
    color: #FFFFFF;
    font-size: 40px !important;
    font-family: STSong;
    text-align: center;
    margin-top: 20px;
}

.blogInfo {
    color: #FFFFFF;
    text-align: center;
    margin-top: 30px;
    font-size: 18px;
    padding-top: 5px;
    padding-bottom: 5px;

    span {
        margin-left: 12px;
    }

    .flag {
        background-color: #FFF;
        color: #F2711C;
        border-color: #F2711C;
        border: 1px solid;
        box-shadow: none;
        padding: 2px;
        border-radius: .28571429rem;
        font-size: 14px;
    }

}


.content-outer {
    width: 76%;
    min-height: 800px;
    margin-top: 60px;
    margin-left: 12%;
    margin-right: 12%;

    top: 0;
    bottom: 0;
    border-radius: .28571429rem;
    box-shadow: none;
    border: 1px solid #D4D4D5;
    padding: 1.5em;


    background: #FFF;
    opacity: 0.9;

}

.content-inner {
    padding: 4em 40px 2em;
    box-sizing: border-box;
    font: 1em/1.5 Tahoma,Helvetica,Arial,'宋体',sans-serif !important;
    color: rgba(0,0,0,.87);
    height: auto;

}

.tags {
    float: right;
    padding-top: 3px;
    padding-bottom: 3px;
    margin-right: 10px;

    span {
        margin-left: 12px;
        background-color: #FFF;
        color: #1b93a8;
        border-color: #F2711C;
        border: 1px solid;
        box-shadow: none;
        padding: 2px;
        border-radius: .28571429rem;
        font-size: 14px;
    }
}

.rights {
    width: 76%;
    margin-left: 12%;
    margin-right: 12%;
    border-radius: .28571429rem;
    box-shadow: 0 0 0 1px #a3c293 inset;
    border: 1px solid #a3c293;
    padding: 1.5em;


    background: #FCFFF5;
    color: #2C662D;
    opacity: 0.9;

    .list {
        text-align: left;
        padding: 0;
        opacity: .85;
        list-style-position: inside;
        margin-bottom: 0.8em;

    }
}

.comment-outer {
    width: 76%;
    margin-left: 12%;
    margin-right: 12%;
    border-radius: .28571429rem;
    box-shadow: none;
    border: 1px solid #D4D4D5;
    padding: 1.5em;

    min-height: 300px;
    background: #FFF;
    opacity: 0.9;


    //margin-bottom: 60px;
}

.comment-body {
    position: relative;
    background: #FFF;
    box-shadow: 0 1px 2px 0 rgb(34 36 38 / 15%);
    margin: 0.5rem 0;
    padding: 1em;
    border-radius: .28571429rem;
    border: 1px solid rgba(34,36,38,.15);
    border-top: 2px solid #00B5AD;
    min-height: 260px;

}

.comment {
    width: 90%;
    padding: 6px 18px;
    margin-left: 10px;
    margin-right: 20px;
    margin-top: 12px;
    display: block;

    .comment-avatar {
        float: left;
        margin-top: 6px;
        margin-right: 12px;
    }

    .author {
        font-size: 1.1em;
        color: rgba(0,0,0,.87);
        font-weight: 700;
        margin-left: 8px;
    }

    .date {
        display: inline-block;
        color: rgba(0,0,0,.4);
        font-size: .875em;
        margin-left: 8px;
    }


    .comment-content {
        margin-top: 4px;
        font-size: 15px;
        word-wrap: break-word;
        color: rgba(0,0,0,.87);
        line-height: 1.3;
        padding-left: 60px;
    }
}

.replay {
    width: 100%;
    min-height: 180px;
    background: #FFF;
    margin: 0.5rem 0;
    padding: 1em;
    border-radius: .28571429rem;

    padding: .78571429em 1em;
    border: 1px solid rgba(34,36,38,.15);
    outline: 0;
    color: rgba(0,0,0,.87);
    box-shadow: 0 0 0 0 transparent inset;
    transition: color .1s ease,border-color .1s ease;
    font-size: 1em;
    line-height: 1.2857;
    resize: vertical;
}

//// 多行文本溢出显示省略号
//.text-ellipsis {
//    //弹性伸缩盒子模型
//    display: -webkit-box;
//    //限制在一个块元素显示的文本行数
//    -webkit-line-clamp: 3;
//    //设置或检索伸缩盒子对象的子元素的排列方式
//    -webkit-box-orient: vertical;
//}

</style>




