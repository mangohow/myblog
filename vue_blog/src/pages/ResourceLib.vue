<template>
    <div>
        <TitleArea class="title-area" :info="info"></TitleArea>
        <div class="main">
            <transition appear
                        name="animate__animated animate__bounce animate__slow"
                        enter-active-class="animate__fadeInUp"
            >
                <div class="center-area">
                    <div class="category" :key="cates.id" v-for="cates in categories">
                        <h4 class="title">
                            {{ cates.name }}
                        </h4>
                        <ul class="clear-fix">
                            <li :key="links.id" v-for="links in cates.links">
                                <ResourceLabel :links="links"></ResourceLabel>
                            </li>
                        </ul>
                    </div>
                </div>
            </transition>
        </div>
    </div>

</template>

<script>

import TitleArea from "../components/TitleArea";
import ResourceLabel from "../components/ResourceLabel";

import "animate.css"

export default {
    name: "ResourceLib",
    components: { TitleArea, ResourceLabel },
    data() {
        return {
            info: {
                title: "资源库",
                desc: "工欲善其事，必先利其器！"
            },
            categories: []
        }
    },
    methods: {
        async getLinks() {
            const {data: res} = await this.$axios.get("/myblog/links");
            if(res.status !== 1) {
                this.$message.warning("获取链接失败，请重试!")
                return
            }
            this.categories.shift()
            let data
            let categories
            if (res.data.length > 1) {
                data = res.data[0]
                categories = res.data[1]
            } else {
                return
            }

            categories.forEach((val) => {
                const arr = data.filter((d) => {
                    return d.categoryId === val.id
                })
                if (arr.length > 0) {
                    this.categories.push({
                        id: val.id,
                        name: val.name,
                        links: arr
                    })
                }
            })
        }
    },
    created() {
        window.scrollTo(0, 0)
        this.getLinks()
    }
}
</script>

<style scoped>

ul, li {
    margin: 0;
    padding: 0;
}

.title-area {
    background: url("../assets/images/back4_3840_800.png") 0 0 / cover  no-repeat;
}

.main {
    background: url("../assets/images/back5.png") 0 0 / cover  no-repeat;
    background-attachment: fixed;
    padding-top: 80px;
    padding-bottom: 150px;
    min-height: 800px;
}

.center-area {
    width: 66%;
    background-color: rgba(255, 255, 255, .6);
    margin: 0 auto;
    border-radius: 5px;
    padding-bottom: 80px;
}

.category {
    padding-left: 30px;
    padding-top: 30px;
}

.title {
    height: 40px;
    line-height: 40px;
    font-size: 18px;
    font-weight: 400;
    color: #555;
    position: relative;
    padding-left: 28px;
}

.title::before {
    content: '\e611';
    position: absolute;
    left: 5px;
    top: 0px;
    font-family: "iconfont";
    font-size: 20px;
}

ul {
    padding-left: 20px;
}

ul li {
    float: left;
    list-style: none;
    margin-right: 30px;
    margin-top: 10px;
}


/*清除浮动*/
.clear-fix::after {
    content: '';
    display: block;
    height: 0;
    clear: both;
    visibility: hidden;
}

</style>
