<template>
    <div class="wapper">
        <div class="toprow">
            <div class="avatar" :style="imgObj"></div>
            <div class="info">
                <h4 class="nickname">{{ msg.nickname }}</h4>
                <h5 class="date"> {{ dateFormat() }} </h5>
            </div>
            <a @click.prevent="replay">回复</a>
        </div>
        <p><span v-if="msg.parentName" style="color:#4CBF30;cursor:pointer;font-weight: 500;">@{{ msg.parentName }}&nbsp;</span>
            {{ msg.content }}</p>
    </div>
</template>

<script>

import dayjs from "dayjs";

export default {
    name: "MessageItem",
    props: ["msg"],
    computed: {
        imgObj() {
            const x = Math.floor((this.msg.avatar - 1) % 7) + 1
            const y = Math.floor((this.msg.avatar - 1) / 7)
            return {
                backgroundPositionX: x * 48 + 'px' ,
                backgroundPositionY: y * 48 + 'px',
            }
        },
    },
    methods: {
        dateFormat() {
            return dayjs(this.msg.createTime).format('YYYY-MM-DD HH:mm:ss')
        },
        replay() {
            this.$bus.$emit("replay", {
                id: this.msg.id,
                topParentId: this.msg.topParentId === 0 ? this.msg.id : this.msg.topParentId
            })
        }
    },

}
</script>

<style scoped>

.wapper {
    overflow: hidden;
}

.toprow {
    height: 48px;
    overflow: hidden;
    line-height: 48px;
}


.avatar {
    width: 48px;
    height: 48px;
    float: left;
    border-radius: 50%;
    background-image: url("../assets/images/avatars.png");
    transition: all 0.6s linear;
}

.avatar:hover {
    transform: rotate(360deg) ;

}

.info {
    float: left;
    margin-left: 12px;
}


h4 {
    font-size: 13px;
    font-weight: 700;
    margin-top: 8px;
    margin-bottom: 5px;
    cursor: pointer;
}

h5 {
    font-size: 12px;
    font-weight: 400;
    margin: 0;
}

a {
    float: right;
    margin-right: 10px;
    outline: none;
    text-decoration: none;
    color: rgb(88 80 236) !important;
    cursor: pointer;
}

a:hover {
    color: rgb(166, 80, 236) !important;;
}

p {
    display: block;
    margin: 5px 54px 10px;
    line-height: 1.8;
}

</style>