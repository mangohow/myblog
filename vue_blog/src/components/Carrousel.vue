<template>
    <div class="father">
        <div class="section">
            <div><span>博客开发</span></div>
            <div><span>SpringBoot使用教程</span></div>
            <div><span>SpringBoot使用教程</span></div>
            <div><span>SpringBoot使用教程</span></div>
            <div><span>SpringBoot使用教程</span></div>
            <div><span>SpringBoot使用教程</span></div>
        </div>
    </div>
</template>

<script>
export default {
    name: "Carrousel"
}
</script>

<style scoped>



@keyframes rotate {
    0% {
        transform: rotateY(0);
    }

    100% {
        transform: rotateY(360deg);
    }
}

.father {
    perspective: 1000px
}

.section {
    position: relative;
    width: 300px;
    height: 200px;
    margin: 300px auto;
    transform-style: preserve-3d;
    animation: rotate 15s linear infinite;
}

.section:hover {
    animation-play-state: paused;
}

.section div {
    position: absolute;
    top: 0;
    left: 0;
    width: 300px;
    height: 185px;
    cursor: pointer;
    border-radius: 2px;
}

.section div:nth-child(1) {
    background: url("~@/assets/images/44.png") 0 0 / cover no-repeat;
    transform: translateZ(400px);
}

.section div:nth-child(2) {
    background: url("~@/assets/images/60.jpg") 0 0 / cover no-repeat;
    transform: rotateY(60deg) translateZ(400px);
}

.section div:nth-child(3) {
    background: url("~@/assets/images/73.png") 0 0 / cover no-repeat;
    transform: rotateY(120deg) translateZ(400px);
}

.section div:nth-child(4) {
    background: url("~@/assets/images/back1.png") 0 0 / cover no-repeat;
    transform: rotateY(180deg) translateZ(400px);
}

.section div:nth-child(5) {
    background: url("~@/assets/images/back2.png") 0 0 / cover no-repeat;
    transform: rotateY(240deg) translateZ(400px);
}

.section div:nth-child(6) {
    background: url("~@/assets/images/back4.png") 0 0 / cover no-repeat;
    transform: rotateY(300deg) translateZ(400px);
}

.section div span {
    display: block;
    position: absolute;
    bottom: 0;
    left: 0;
    height: 30px;
    width: 100%;
    background-color: rgba(0, 0, 0, .2);
    color: #fff;
    font-family: STSong;
    font-size: 10px;
    line-height: 30px;
    padding-left: 20px;
}

</style>
