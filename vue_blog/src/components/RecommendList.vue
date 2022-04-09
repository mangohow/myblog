<!--最新推荐组件-->
<template>
    <div class="wapper">
        <div class="title">
            <i class="iconfont" :class="blogList.icon"></i>
            <h4>{{ blogList.title }}</h4>
        </div>
        <ul class="list">
            <li class="title-ellipsis" :key="item.id" v-for="item in blogList.list" @click="itemClicked(item.id)">
                <span>{{ item.title }}</span>
            </li>

        </ul>
    </div>
</template>

<script>

import "../assets/icon/iconfont.css"

export default {
    name: "NewRecommend",
    props: ["blogList"],
    methods: {
        itemClicked(id) {
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
    margin: 0;
    padding: 0;
}

.wapper {
    width: 270px;
    border-radius: 12px;
    background-color: #fff;
    margin-top: 30px;
    user-select: none;
}

.title {
    border-bottom: 1px solid rgba(14, 14, 14, 0.1);
}

.title i {
    color: #fb6c28;
    font-size: 17px;
    margin-left: 10px;
    margin-right: 5px;
}

.title h4 {
    height: 45px;
    line-height: 50px;
    margin: 0;
    font-size: 16px;
    display: inline-block;
}

ul {
    padding-top: 10px;
    padding-bottom: 10px;
}

ul li {
    list-style: none;
    height: 36px;
    line-height: 36px;
    padding-left: 15px;
    cursor: pointer;
    font-size: 14px;
}

ul li span {
    padding-bottom: 3px;
    position: relative;
    color: #606266;

}

.title-ellipsis {
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    width: 230px;
}

ul li span::before {
    content: '';
    position: absolute;
    width: 0;
    height: 0;
    bottom: 0;
}

ul li:hover span {
    color: #fb6c28;
}

ul li:hover span::before {
    width: 100%;
    transition: 0.4s width ease-in;
    border-bottom: 1px solid #fb6c28;
}


</style>
