<template>
    <div>
        <TitleArea class="title-area" :info="info"></TitleArea>
        <div class="main">

        </div>
    </div>

</template>

<script>

import TitleArea from "../components/TitleArea";

export default {
    name: "PhotoWall",
    components: { TitleArea },
    data() {
        return {
            info: {
                title: "照片墙",
                desc: "好看的照片"
            }
        }
    }
}
</script>

<style scoped>

.title-area {
    background: url("../assets/images/back1_3840_800.png") 0 0 / cover  no-repeat;
}

.main {
    background: url("../assets/images/back9.png") 0 0 / cover  no-repeat;
    background-attachment: fixed;
    width: 100%;
    height: 2000px;
}


</style>
