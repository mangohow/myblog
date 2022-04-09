<!--博客类型标签组件-->
<template>
    <a class="label" :class="{active: isActive}">
        <span class="item-name">{{ typeinfo.name }}</span>
        <span :style="{backgroundColor: itemColor}" class="item-count">{{ typeinfo.count }}</span>
    </a>
</template>

<script>


export default {
    name: "BlogTypeLabel",
    props: ["typeinfo", "activeId"],
    data() {
        return {
            itemColor: "#f37d7c",
        }
    },
    computed:{
        isActive() {
            return (this.activeId == this.typeinfo.id)
        }
    },
    created() {
        this.itemColor = this.$randomColor();
    }
}
</script>

<style scoped>

.label {
    display: flex;
    margin: 4px 8px 10px;
    padding: 8px 14px;
    align-items: center;
    text-align: center;
    color: #666;
    background-color: #fff;
    width: auto;
    border-radius: 0.25rem;
    font-weight: 500;
    border: None;
    height: 45px;
    font-size: 13px;
    box-shadow: 0 1px 8px 0 rgba(0, 0, 0, 0.1);
    cursor: pointer;
    text-decoration:none;
}

.label:hover, .active {
    background-color: #F37D7C;
    color: #fff;
}

.item-name {
    display: inline-block;
    text-align: center;
    box-sizing: border-box;
}

.item-count{
    display: inline-block;
    text-align: center;
    margin-left: 4px;
    background-color: #f37d7c;
    border-radius: 0.25rem;
    color: white;
    width: 1.2rem;
    height: 1.2rem;
    font-size: 0.7rem;

}

</style>
