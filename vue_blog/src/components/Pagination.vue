<!--分页导航栏-->
<template>
    <div class="wapper">
        <ul>
            <li v-show="hasPrevPage" @click="changePage(-1)" class="finger">👈</li>
            <li @click="toPage(p)" class="page-num" :key="p" v-for="p in showPages" :class="{pselected: pageInfo.pageNum === p}">{{ p }}</li>
            <li v-show="show3dot" class="page-num">...</li>
            <li v-show="show3dot" @click="toPage(pageInfo.pages)" class="page-num">{{ pageInfo.pages }}</li>
            <li v-show="hasNextPage" @click="changePage(1)" class="finger">👉</li>
        </ul>
    </div>
</template>

<script>
export default {
    name: "Paginations",
    props: ["pageInfo"],
    data() {
        return {
            maxPages: 8
        }
    },
    methods: {
        changePage(step) {
            const nextPage = this.pageInfo.pageNum + step
            this.toPage(nextPage)
        },
        toPage(pageNum) {
            if( pageNum <= 0 || pageNum > this.pages || pageNum === this.pageInfo.pageNum ) {
                return
            }
            this.$emit("jumpPage", pageNum)
        }
    },
    computed: {
        hasPrevPage() {
            return this.pageInfo.pageNum > 1
        },
        hasNextPage() {
            return this.pageInfo.pages > this.pageInfo.pageNum
        },
        show3dot() {
            return this.pageInfo.pages > this.maxPages
        },
        showPages() {
            const bach = Math.floor(this.pageInfo.pageNum / (this.maxPages + 1))   // 0
            const index = this.pageInfo.pages % this.maxPages   // 1
            const start = bach * this.maxPages + 1   // 1
            const end = ((bach + 1) * this.maxPages) <= this.pageInfo.pages ? this.maxPages : index //8
            const arr = [];
            for (let i = 0; i < end; i++) {
                arr.push(start + i)
            }
            return arr
        },
    }
}
</script>

<style lang="less" scoped>

ul {
    text-align: center;
    margin: 30px 0;
    padding: 0;

    li {
        display: inline-block;
        user-select: none;
    }
}

.finger {
    width: 48px;
    height: 48px;
    font-size: 20px;
    border-radius: 50%;
    cursor: pointer;
    transition: all 0.5s ease;
    background-color: rgba(255, 255, 255, 0.2);
    margin: 0 10px;
    line-height: 40px;
}


.pselected,
.finger:hover,
.page-num:hover {
    box-shadow: 0 14px 28px rgb(0 0 0 / 25%), 0 10px 10px rgb(0 0 0 / 22%) !important;
    background: #F37D7C !important;
    color: #fff;
}

.page-num:hover {
    transform: scale(1.1);
}

.page-num {
    width: 30px;
    height: 30px;
    border-radius: 50%;
    background-color: rgba(255, 255, 255, 0.2);
    cursor: pointer;
    margin-left: 10px;
    font-size: 14px;
    line-height: 30px;
    font-weight: bold;
    transition: all 0.5s;
}


ul li:nth-child(2) {
    margin-left: 0;
}


</style>
