<template>
    <div class="panel" id="rp">
        <div class="wrapper">
            <div class="info">
                <input type="text" placeholder="尊姓大名" v-model="postInfo.nickname">
                <input type="text" placeholder="邮箱 (仅作回复使用)" v-model="postInfo.email">
                <input type="text" placeholder="个人网站 (选填, 需包含http或https)" v-model="postInfo.url">
            </div>
            <textarea name="" id="" cols="30" rows="10" v-model="postInfo.content" placeholder="请输入评论内容..."></textarea>
            <div class="btn-area">
                <button @click="commit">提交评论</button>
                <button @click="clearAll">清空</button>
                <button v-if="closeBtn" @click="close">关闭</button>
            </div>
        </div>
    </div>

</template>

<script>
export default {
    name: "LeaveMessagePanel",
    props: ["closeBtn", "msgInfo"],
    watch: {
        msgInfo: {
            handler(newVal) {
                this.postInfo.parentId = newVal.parentId
                this.postInfo.topParentId = newVal.topParentId
            },
            deep: true
        }
    },
    data() {
        return {
            postInfo: {
                nickname: "",
                email: "",
                url: "",
                content: "",
                parentId: 0,
                topParentId: 0
            },
            commented: false,
            repeatCommitTimes: 0
        }
    },
    methods: {
        async commit() {
            if(!this.validate()) {
                return
            }

            this.trimSpace()
            const {data: res} = await this.$axios.post("/myblog/leaveMsg", this.postInfo);
            if(res.status === 101) {
                this.$message.success("收到靓仔的评论！")
                this.postInfo.content = ""
                if(this.closeBtn) {
                    this.close()
                }
            } else {
                this.commented = false
                this.$message.success("网络出现了点小问题，请重试！")
            }
        },
        validate() {
            // 评论后，需要等一分钟才能评论，不能重复提交，重复提交次数>3，就清空所有输入的内容
            if(this.commented) {
                this.$message.error("您的提交太过频繁，请稍后再试！")
                this.repeatCommitTimes++
                if(this.repeatCommitTimes > 3) {
                    this.clearAll()
                }
                return false
            }

            // 验证昵称
            if(this.postInfo.nickname.trim() === "") {
                this.$message.error("请输入您的昵称！")
                return false
            } else if(this.postInfo.nickname.length > 28) {
                this.$message.error("昵称长度限制为28个字符！")
                return
            }
            // 验证邮箱格式
            const t = /^[A-Za-zd0-9]+([-_.][A-Za-zd]+)*@([A-Za-zd]+[-.])+[A-Za-zd]{2,5}$/;
            if(!t.test(this.postInfo.email)) {
                this.$message.error("邮箱格式错误！")
                this.postInfo.email = ""
                return false
            }
            // 验证网址格式
            if(this.postInfo.url !== "") {
                const reg = /(http|ftp|https):\/\/[\w\-_]+(\.[\w\-_]+)+([\w\-\.,@?^=%&amp;:/~\+#]*[\w\-\@?^=%&amp;/~\+#])?/;
                if(!reg.test(this.postInfo.url)) {
                    this.$message.error("网址格式错误！")
                    this.postInfo.url = ""
                    return false
                }
            }

            // 验证输入内容
            if(this.postInfo.content == "") {
                this.$message.error("请输入评论内容！")
                return false
            } else if(this.postInfo.content.length > 1500) {
                this.$message.error("评论内容字符限制最大1500字！")
                return false
            }

            this.commented = true
            this.repeatCommitTimes = 0
            setTimeout(() => {
                this.commented = false
            }, 60 * 1000)

            return true
        },
        clearAll() {
            this.postInfo = {
                nickname: "",
                email: "",
                url: "",
                content: ""
            }
        },
        trimSpace() {
            this.postInfo.nickname.trim()
            this.postInfo.email.trim()
            this.postInfo.content.trim()
            this.postInfo.url.trim()
        },
        close() {
            this.$bus.$emit("hideReplay")
        }
    }
}
</script>

<style scoped>

.panel {
    border-radius: 8px;
    box-shadow: 0 15px 35px rgb(50 50 93 / 18%), 0 5px 15px rgb(0 0 0 / 18%) !important;
    background-color: #fff;
    width: 1070px;
}

.wrapper {
    padding: 10px 10px 15px;
}

input {
    width: 350px;
    height: 66px;
    outline: none;
    border: none;
    border-bottom: 1px dashed #dedede;
}

input:focus {
    border-bottom: 1px dashed #e4933c;
    box-shadow: 0 1px 0 0 #def605;
    transition: all 0.5s ease;
}



.info {
    margin-bottom: 10px;
}

textarea {
    outline: none;
    resize: none;
    width: 100%;
    border: none;
    border-bottom: 1px dashed #dedede;
    background: url("../assets/images/comment_bg.webp") no-repeat top right;
}

.btn-area {
    overflow: hidden;
    padding-top: 10px;
    padding-bottom: 10px;
}

button {
    float: right;
    border: 1px solid rgb(88 80 236);
    background-color: rgb(240 245 255);
    font-size: .875rem;
    padding: 0.625rem 1.25rem;
    border-radius: 0.25rem;
    margin-right: 20px;
}

button:active {
    transform: translateY(2px);
}
</style>