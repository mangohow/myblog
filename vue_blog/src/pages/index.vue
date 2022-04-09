<template>

    <div>
        <transition appear
            name="animate__animated animate__bounce animate__slow"
            enter-active-class="animate__slideInDown"
            leave-active-class="animate__slideOutUp"
        >
            <NavigationBar class="navibar" v-show="showNavi"></NavigationBar>
        </transition>
        <BackToTop v-show="showCat"></BackToTop>
        <keep-alive include="Home">
            <router-view ref="home"></router-view>
        </keep-alive>
    </div>

</template>



<script>

import NavigationBar from "../components/NavigationBar";
import BackToTop from "../components/BackToTop";
import "jquery"

import "animate.css"
// import "../assets/js/sakura"
// import "../assets/js/love"

export default {
    components: {NavigationBar, BackToTop},
    data() {
        return {
            keyWord: "",
            showCat: false,
            showNavi: true,
            preTop: 0,
        }
    },
    methods: {
        handleScroll() {
            this.showNavi = this.preTop >= window.scrollY;
            this.preTop = window.scrollY
            const w = window.screen.height / 2
            this.showCat = window.scrollY > w;
        }
    },
    mounted() {
        document.addEventListener("scroll", this.handleScroll, true)
    }
}
</script>



<style scoped>

.animate__animated{
    animation-duration: 1.5s !important;
}

.navibar {
    position: fixed;
    top: 0;
}

</style>
