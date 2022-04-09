<template>
    <div class="summarize">
        <div class="summ-title">总览</div>
        <ul class="summ-list">
            <li>
                <i class="iconfont icon-navicon-wzgl"></i>
                <div>
                    <span>80</span>
                    <span>文章总计</span>
                </div>
            </li>
            <li>
                <i class="iconfont icon-biaoqian1"></i>
                <div>
                    <span>30</span>
                    <span>标签总计</span>
                </div>
            </li>
            <li>
                <i class="iconfont icon-pinglun"></i>
                <div>
                    <span>50</span>
                    <span>评论总计</span>
                </div>
            </li>
        </ul>
    </div>
</template>

<script>
export default {
    name: "BlogOverview"
}
</script>

<style scoped>

.summarize {
    padding: 30px 20px;
}

.summ-title,
.chart-title,
.article-title {
    font-size: 30px;
    margin-bottom: 20px;
    padding-left: 15px;
}

.summ-title::before,
.chart-title::before,
.article-title::before {
    content: '\e607';
    font-family: "iconfont";
    padding-right: 10px;
    color: #3498DB;
}

.chart-title::before {
    content: '\e602';
}

.article-title::before {
    content: '\e660';
}


.summ-list>li {
    display: inline-block;
    width: 225px;
    height: 100px;
    border: 2px solid;
    margin: 10px 20px;
    border-radius: 5px;
    text-align: center;
    list-style: none;
}

.summ-list>li:first-child {
    border-color: #3f83f8;
}

.summ-list>li:first-child i {
    color: #3f83f8;
}

.summ-list li:first-child span:first-child {
    color: #3f83f8;
}

.summ-list>li:nth-child(2) {
    border-color: #FE9E0B;
}
.summ-list>li:nth-child(2) i {
    color: #FE9E0B;
}

.summ-list li:nth-child(2) span:first-child {
    color: #FE9E0B;
}

.summ-list>li:nth-child(3) {
    border-color: #0E9F6E;
}
.summ-list>li:nth-child(3) i {
    color: #0E9F6E;
}

.summ-list li:nth-child(3) span:first-child {
    color: #0E9F6E;
}

.summ-list i {
    float: left;
    font-size: 60px !important;
    line-height: 100px;
    height: 100%;
    width: 80px;
    padding-left: 20px;
}

.summ-list li>div {
    float: right;
    padding: 15px 40px 0 15px;
}

.summ-list li>div span {
    display: block;
}

.summ-list li>div span:first-child {
    font-size: 30px;
    font-weight: 700;
}

</style>
