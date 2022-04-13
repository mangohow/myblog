
<template>
    <div>
        <TitleArea class="title-area" :info="info"></TitleArea>
        <div class="bottom">
            <transition appear
                        name="animate__animated animate__bounce animate__slow"
                        enter-active-class="animate__slideInUp"
            >
                <div class="wrap">
                    <img :src="userInfo.avatar">
                    <h4>{{ userInfo.nickname }}</h4>
                    <ul class="contact">
                        <li><a :href="userInfo.github" target="_blank" class="iconfont icon-github"></a></li>
                        <li><a :href="userInfo.csdn" target="_blank" class="iconfont icon-csdn"></a></li>
                    </ul>
                    <div class="statement">

                    </div>
                    <div class="box-wrap">
                        <span class="side-label">关于我</span>
                        <h5 class="title">关于博主</h5>
                        <p class="about-state">
                            关于我
                        </p>
                        <p class="about-state">
                            第二段
                        </p>
                    </div>

                    <div class="box-wrap">
                        <span class="side-label bgc-pink">关于本站</span>
                        <h5 class="title">本站采用的技术</h5>
                        <h6 class="site-intro"><span>前端：</span>Vue + Element Ui</h6>
                        <h6 class="site-intro"><span>后端：</span>Golang + Gin + Mysql</h6>
                        <h6 class="site-intro"><span>部署方式：</span>提供了Docker-Compose的部署方式，只需少量配置即可一键启动，使用Nginx作为静态资源和反向代理服务器。</h6>
                        <h6 class="site-intro">本站已经开源在Github，地址：https://github.com/mangohow/myblog。</h6>
                    </div>

                    <div class="box-wrap">
                        <span class="side-label bgc-orange">联系方式</span>
                        <h6 class="site-intro" style="margin-top: 12px;"><span>Github：</span><a :href="userInfo.github" target="_blank"> {{ userInfo.github }}</a></h6>
                        <h6 class="site-intro"><span>CSDN：</span><a :href="userInfo.csdn" target="_blank"> Peerless</a></h6>
                        <h6 class="site-intro"><span>Email：</span>{{ userInfo.email }}</h6>
                    </div>

                </div>
            </transition>
        </div>
    </div>
</template>



<script>

import "../assets/icon/iconfont.css"
import TitleArea from "../components/TitleArea";

import "animate.css"

export default {
    components: { TitleArea },
    data() {
        return {
            info: {
                title: "关于我",
                desc: "自信人生二百年，会当水击三千里！"
            },
            userInfo: {
                avatar: "",
                nickname: "",
                github: "",
                csdn: "",
                email: ""
            }
        }
    },
    created() {
        window.scrollTo(0,0);
        this.getUserInfo()
    },
    methods: {
        async getUserInfo() {
            const {data:res} = await this.$axios.get("/myblog/userInfo")
            if(res.status !== 1) {
                this.$message.warning("获取用户信息失败！")
            }
            this.userInfo.avatar = res.data.avatar
            this.userInfo.nickname = res.data.nickname
            this.userInfo.github = res.data.github
            this.userInfo.csdn = res.data.csdn
            this.userInfo.email = res.data.email
        },
    }
}
</script>



<style lang="less" scoped>

ul {
    list-style: none;
    margin: 0;
    padding: 0;
}

.title-area {
    background: url("../assets/images/back12.jpg") 0 0 / cover  no-repeat;
}

.bottom {
    padding-bottom: 80px;
    background-color: #ececec;
}

.wrap {
    width: 60%;
    border-radius: 10px;
    margin: 0 auto;
    transform: translateY(-20px);
    box-shadow: 0 15px 35px rgb(50 50 93 / 10%), 0 5px 15px rgb(0 0 0 / 7%);
    background: #fff url("../assets/images/flowerbg.png") fixed top;
    padding: 50px 0;

    img {
        display: block;
        width: 160px;
        height: 160px;
        border-radius: 50%;
        margin: 0 auto;
    }

    h4 {
        text-align: center;
        margin: 12px 0;
        color: #34495e;
        font-style: italic;
    }
    .contact {
        text-align: center;
        margin: 20px 0;
        padding-bottom: 6px;
        position: relative;

        li {
            display: inline-block;
            a {
                text-decoration: none;
                font-size: 26px;
                margin: 0 12px;
                color: rgb(63, 63, 64);
            }
        }
    }


    .box-wrap {
        position: relative;
        width: 80%;
        margin: 30px auto 60px;
        border-radius: 8px;
        background-color: rgba(255, 255, 255, .5);
        box-shadow: -12px -17px 38px rgb(0 0 0 / 10%), 7px 18px 25px rgb(0 0 0 / 8%);
        transition: all .3s ease 0s, transform .6s cubic-bezier(.6, .2, .1, 1) 0s, -webkit-transform .6s cubic-bezier(.6, .2, .1, 1) 0s;
        color: #34495e;
        padding-top: 60px;
        padding-bottom: 18px;

        .side-label {
            position: absolute;
            background-color: #b28fce;
            color: #fff;
            font-style: italic;
            font-size: 14px;
            padding-left: 32px;
            padding-right: 24px;
            left: -16px;
            top: 20px;
            height: 32px;
            line-height: 32px;
            border-radius: 0 3px 3px 0;
        }

        .side-label::before {
            content: '';
            position: absolute;
            top: 100%;
            left: 0;
            width: 0;
            height: 0;
            background-color: transparent;
            border-style: solid;
            border-width: 0 16px 16px 0;
            border-color: transparent;
            border-right-color: #b28fce;
            filter: brightness(120%);
        }
    }

    .box-wrap:hover {
        transform: translateY(-5px);
    }

    .bgc-pink {
        background-color: #f596aa !important;
    }

    .bgc-pink::before {
        border-right-color: #f596aa !important;
    }

    .bgc-orange {
        background-color: #fb966e !important;
    }

    .bgc-orange::before {
        border-right-color: #fb966e !important;
    }

    .title {
        padding: 10px 20px;
        font-size: 20px;
        font-weight: 600;

    }

    font-family: 'Architects Daughter', cursive;

    .about-state {
        padding: 0 60px 0 36px;
        text-indent: 2rem;
        font-style: italic;
        font-size: 18px;
        font-weight: 520;
        margin-bottom: 8px;
    }

    .site-intro {
        padding-left: 32px;
        padding-right: 60px;
        font-size: 18px;

        span {
            font-weight: 700;
        }

        a {
            text-decoration: none;
            color: #42b983;
        }
    }

}

</style>
