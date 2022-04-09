<!--导航按钮组件-->
<template>
    <router-link :to="rout" class="nav-item" :class="{active: isActive}">
        <i class="iconfont" :class="icoClass"></i>
        {{ indexName }}
    </router-link>
</template>

<script>

import "../assets/icon/iconfont.css"

export default {
    name: "NaviItem",
    props: ["indexName", "icoClass", "rout", "activeRoute"],
    computed: {
        isActive() {
          return this.rout === this.activeRoute
        }
    }

}
</script>

<style lang="less" scoped>

i {
    color: #fda085 !important;
}

.nav-item {
    line-height: 40px;
    cursor: pointer;
    padding-bottom: 5px;
    text-decoration:none;
    color: black;
    font-weight: 40;
}

.nav-item:hover, .active {
    font-weight: bold;
    color: #fda085;
    border-bottom: 2px solid #fda085;
}

.nav-item:hover .active {
    border-bottom: 4px solid #fda085;
}

</style>
