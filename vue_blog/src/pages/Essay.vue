<template>
    <div>
        <TitleArea class="title-area" :info="info"></TitleArea>
        <div class="main">
            <div class="time-line">
                <ul :key="index" v-for="(item, index) in list">
                    <li>
                        <div class="year">{{ item.year }}</div>
                    </li>
                    <transition appear
                                name="animate__animated animate__bounce animate__slow"
                                :key="essay.id" v-for="(essay, index) in item.essays"
                                :enter-active-class="index % 2 === 0 ? 'animate__fadeInLeft' : 'animate__fadeInRight'"
                    >
                        <li  class="clear-fix">
                            <div :class="[index % 2 === 0 ? 'item-left' : 'item-right']">
                                <h4 class="time">{{ dateFormat(essay.createTime) }}</h4>
                                <p class="content">{{ essay.content }}</p>
                            </div>
                        </li>
                    </transition>
                </ul>
            </div>
        </div>
    </div>
</template>

<script>

import TitleArea from "../components/TitleArea";

import dayjs from 'dayjs'

export default {
    name: "Essay",
    components: { TitleArea },
    data() {
        return {
            info: {
                title:'日常随笔',
                desc:'记录生活的琐碎!',
            },
            list: []
        }
    },
    methods: {
        async getEssayList() {
            const {data:res} = await this.$axios.get("/myblog/essayList");
            if (res.status !== 1) {
                this.$message.error("网络出了点小问题，获取数据失败！")
                return
            }

            let data
            if (res.data.length > 0) {
                data = res.data[0]
            } else {
                return
            }

            let prev = dayjs(data[0].createTime).format("YYYY")
            let years = []
            data.forEach((val, index) => {
                if(index === 0) {
                    years.push({
                        count: 1,
                        year: prev
                    })
                    return
                }
                let cur = dayjs(val.createTime).format("YYYY")
                if (prev === cur) {
                    years[years.length - 1].count++
                } else {
                    years.push({
                        count: 1,
                        year: cur
                    })
                    prev = cur
                }

            })

            years.forEach((val) => {
                this.list.push({
                    year: val.year,
                    essays: data.splice(0, val.count)
                })
            })
        },
        dateFormat(t) {
            return dayjs(t).format("YYYY-MM-DD HH:mm:ss")
        },
    },
    created() {
        this.getEssayList()
    }
}
</script>

<style scoped>

ul, li {
    padding: 0;
    margin: 0;
}

.title-area {
    background: url("../assets/images/blog_bg_1280_300.png") 0 0 / cover  no-repeat;
}

.main {
    /*background: url("../assets/images/back7.jpg") 0 0 / cover  no-repeat;*/
    background: url("../assets/images/theme_bg.png") 0 0 / cover  no-repeat;
    background-attachment: fixed;
    width: 100%;
    min-height: 1200px;
    padding-top: 50px;
}

.time-line {
    width: 60%;
    margin: 0 auto;
    padding-top: 64px;
    position: relative;
    padding-bottom: 80px;
}

.time-line::before {
    content: '';
    height: 90%;
    width: 0;
    border: 2px solid rgba(143, 139, 236, .3);
    border-radius: 3px;
    position: absolute;
    left: 50%;
    top: 5%;
    transform: translateX(-50%);
}

.year {
    width: 80px;
    height: 48px;
    line-height: 42px;
    border: 2px solid #aba9e8;
    border-radius: 8px;
    text-align: center;
    font-size: 20px;
    color: #aba9e8;
    font-weight: 600;
    margin: 0 auto;
}


ul li {
    list-style: none;
    margin-top: 16px;
}

ul li:first-child {
    margin-top: 80px;
}

ul:first-child li:first-child {
    margin-top: 30px;
}

.item-left,
.item-right {
    width: 36%;
    padding: 20px;
    border: 1px solid #da7d19;
    border-radius: 10px;
    position: relative;
    font-family: 华文楷体;
}

.item-left {
    float: left;
    transform: translateX(16%);
}

.item-right {
    float: right;
    transform: translateX(-16%);
}

.item-left::after,
.item-right::before {
    content: '';
    display: inline-block;
    width: 12px;
    height: 12px;
    border-bottom: 1px solid #da7d19;
    border-right: 1px solid #da7d19;
    transform: translateY(-50%) rotate(-45deg);
    position: absolute;
    right: -7px;
    top: 30%;
    background-color: #fff;
}

.item-right::before {
    left: -7px;
    transform: translateY(-50%) rotate(135deg);
    top: 70%;
}

h4 {
    font-size: 18px;
    font-weight: 400;
    color: #67492e;
}

.item-left h4 {
    text-align: right;
}

.item-left h4::after,
.item-right h4::before {
    content: '';
    border: 1px solid #ffc107;
    width: 0;
    margin-left: 5px;
}

.item-right h4::before {
    margin-left: 0 !important;
    margin-right: 5px;
}

p {
    font-size: 18px;
    line-height: 28px;
    color: #67492e;
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
