<template>
    <div class="time-line">
        <div class="year">{{ year(list[0].createTime) }}</div>
        <ul>
            <li :key="item.id" v-for="item in list">
                <span class="date">{{ date(item.createTime) }}</span>
                <span class="title" @click="gotoBlog(item.id)">{{ item.title }}</span>
            </li>
        </ul>
    </div>
</template>

<script>

import dayjs from 'dayjs'

export default {
    name: "TimeAxis",
    props: ["list"],
    methods: {
        year(t) {
            return dayjs(t).format('YYYY')
        },
        date(t) {
            return dayjs(t).format('MM-DD')
        },
        gotoBlog(id) {
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

<style scoped>

ul, li {
   padding: 0;
    margin: 0;
}

.year {
    font-size: 20px;
    font-weight: 700;
    margin: 20px 0 15px;
    line-height: 1.5;
    font-family: "PingFang SC", "Microsoft YaHei", Lato, sans-serif;
}


ul li {
    position: relative;
    list-style: none;
    margin-bottom: 10px;
    line-height: 28px;
    padding-left: 35px;
}

ul li::before {
    content: '';
    width: 6px;
    height: 6px;
    background-color: #97dffd;
    border-radius: 50%;
    display: inline-block;
    position: absolute;
    left: 10px;
    top: 13px;
}

ul li .date {
    margin-right: 20px;
    color: #00a7e0;
    font-size: 12px;
}

ul li .title {
    position: relative;
    font-size: 14px;
    font-weight: 400;
    cursor: pointer;
    padding-bottom: 6px;
}

ul li .title::before {
    content: '';
    position: absolute;
    width: 0;
    height: 0;
    bottom: 0;
    transition: 0.4s width linear;
}

ul li .title:hover {
    color: #00a7e0;
}

ul li .title:hover::before {
    width: 100%;
    border-bottom: 1px solid #00a7e0;
}

</style>
