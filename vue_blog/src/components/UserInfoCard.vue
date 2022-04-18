<!--用户信息卡片组件-->
<template>
    <div class="box">
        <img @click="imgClicked" :src="userInfo.avatar" class="avatar">
        <h3 class="nickname">{{ userInfo.nickname }}</h3>
        <div class="blog-info">
            <div>
                <h3>{{ count }}</h3>
                <h6>文章</h6>
            </div>
            <div>
                <h3>{{ userInfo.tagCount }}</h3>
                <h6>标签</h6>
            </div>
        </div>
        <ul>
            <li><a :href="userInfo.github" target="_blank" class="iconfont icon-github"></a></li>
            <li><a :href="userInfo.csdn" target="_blank" class="iconfont icon-csdn"></a></li>
        </ul>
    </div>
</template>

<script>

import "../assets/icon/iconfont.css"

export default {
    name: "UserInfoCard",
    props: ["count"],
    data() {
        return {
            userInfo: {
                avatar: "https://placekitten.com/300/300",
                nickname: "",
                tagCount: 20,
                github: "https://www.baidu.com",
                csdn: "https://www.baidu.com "
            },
        }
    },
    methods: {
        imgClicked() {
            this.$router.push("/about")
        },
        // 获取用户信息
        async getUserInfo() {
            const {data:res} = await this.$axios.get("/myblog/userInfo")
            if(res.status !== 1) {
                this.$message.warning("获取用户信息失败！")
            }
            if (res.data.length > 0) {
                const data = res.data[0]
                this.userInfo.avatar = data.avatar
                this.userInfo.nickname = data.nickname
                this.userInfo.github = data.github
                this.userInfo.csdn = data.csdn
                if (res.data.length > 1){
                    this.userInfo.tagCount = res.data[1]
                }
            }
        },
    },
    created() {
        this.getUserInfo()
    }
}
</script>

<style lang="less" scoped>

.box {
    width: 270px;
    height: 300px;
    border-radius: 12px;
    text-align: center;
    background-image: url("~@/assets/images/beijing.jpg");
    user-select: none;

    img {
        width: 108px;
        height: 108px;
        border-radius: 50%;
        padding: 3px;
        margin-top: 16px;
        border: 3px dashed darkorange;
        cursor: pointer;
    }

}

.box:hover {
    transform: scale(1.03);
    box-shadow: 0 5px 10px rgb(0 0 0 / 25%), 0 10px 10px rgb(0 0 0 / 22%);
}

.box:hover img {
    transition:0.5s ease-in;
    transform:rotate(360deg);
}


.nickname {
    font-size: 1rem;
    margin: 16px 0;
}

.blog-info>div {
    display: inline-block;
    width: 108px;
    height: 48px;

    h3 {
        font-size: 1.2rem;
    }

    h6 {
        font-size: 0.7em;
    }
}

.blog-info div:first-child {
    border-right: 1px solid black;
}

ul {
    text-align: center;
    padding: 10px;
    margin: 15px 0 0 0;
    li {
        list-style: none;
        display: inline-block;
        cursor: pointer;

        a {
            font-size: 24px;
            text-decoration: none;
            outline: none;
        }
    }

    .icon-github {
        color: rgb(132, 155, 135);
        margin-right: 25px;
    }

    .icon-csdn {
        color: rgb(225, 91, 100);
        margin-left: 25px;
    }
}

</style>
