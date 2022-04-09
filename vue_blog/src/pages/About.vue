
<template>

    <div>
        <div class="top-about" align="center">关于我</div>


        <div class="main-box">
            <div style="float:left;font-size:18px;font-weight:bold;">关于作者</div>
            <div style="float:right;">2021/04/08</div>
            <hr style="margin-top:40px;">
            <div class="title">作者简介</div>
            <div class="content">本人，男，90后。</div>
            <div class="title">作者经历</div>
            <div class="content">
                目前是杭电的一名研究生，菜鸡一枚。
            </div>
            <div class="title">本站介绍</div>
            <div class="content">本站使用的技术：后端:golang(gin + sqlx) + mysql ,  前端: Vue + Element Ui</div>
        </div>
    </div>

</template>



<script>
export default {
    data() {
        return {

        }
    },
    created() {
        window.scrollTo(0,0);
    },
    methods: {

    }
}
</script>



<style lang="less" scoped>

.top-about {
    font-size: 450%;
    color: #ffffff;
    margin-bottom: 30px;
    bottom: 0 !important;
    right: 0 !important;
    font-family:'STXingkai';
    opacity: 0.8;
}

.main-box {
    width: 76%;
    margin-left: 12%;
    background-color: #FFF;
    border-radius: .28571429rem;
    border: 1px solid rgba(34,36,38,.15);
    opacity: 0.8;
    border-top: 1px solid rgba(34,36,38,.15);
    min-height: 700px;
    padding: 24px 24px;


}

.title {
    font-size: 20px;
    font-weight: bold;
    margin: 10px 14px;
    margin-top: 50px;
    padding: 0 18px;
    border-left: 5px solid #3690cf;
}

.content {
    font-size: 17px;
    margin: 20px 80px;
    font-weight: 520;
}
</style>
