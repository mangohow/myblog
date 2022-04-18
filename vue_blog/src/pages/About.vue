
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
                            到了关于我了，关于我，关于我，该怎么介绍呢，感觉也没什么可写的。现在是2022年4月13日17点32分，外面下着大雨，我还在写这个博客系统。目前还是一名研究生，说到这个博客系统，还是在大四入学研究生的时候老师给的一个作业。在本科的时候，主学的是C++，Web开发接触的比较少，在开学前就跟着一个学长学Java和前端，当时完全是用来应付老师的作业的，就随便做了一个。但是后来看到别人的博客，又加上自己还是比较闲的，因此就想把之前的博客好好改造一下。但是，后端并没有使用Java，用的是Go语言，刚接触Go感觉这语言很反人类，类型非要写在后面、没有三元表达式等，但是使用习惯了之后，感觉这门语言还是很棒的，不像Java，太过于重，而C++的语法也变得愈发复杂，甚至现代的C++代码都让人看不懂。到目前为止，这个系统做的也差不多了，最后省的也就是测试Bug、优化代码以及部署了。
                        </p>
                        <p class="about-state">
                            说了这么多，还是在介绍这个博客😂。有时候感觉挺迷茫的，不过年轻人有点迷茫还不是很正常，还是忙点好，不管是学习还是运动。总想出去看看，但是这该死的疫情还不消失。青春才几年啊，疫情都已经持续了三年了。读书的时候想去工作，想去赚钱，这样就可以自己支配资金了，但是身边也有不少工作了两三年又来读研的朋友，还是好好珍惜上学的时光吧，以后也就只能回想了。就说这些吧，梦想可以看遍祖国的大江大海大山，加油！
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
            if (res.data.length > 0) {
                const data = res.data[0]
                this.userInfo.avatar = data.avatar
                this.userInfo.nickname = data.nickname
                this.userInfo.github = data.github
                this.userInfo.csdn = data.csdn
                this.userInfo.email = data.email
            }

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
