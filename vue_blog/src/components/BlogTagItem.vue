<!--博客标签组件-->
<template>
    <a :style="{backgroundColor: $randomColor()}" :class="{active: isActive}">
        {{ taginfo.name }}
    </a>
</template>

<script>



export default {
    name: "TagsDisplay",
    props: ["taginfo", "activeId"],
    methods: {
        itemClicked(id) {
            this.$emit("trigger", id)
        },
    },
    computed: {
        isActive() {
            return this.activeId == this.taginfo.id
        }
    }
}
</script>

<style scoped>

a {
    float: left;
    font-size: 14px;
    font-weight: 600;
    line-height: 32px;
    color: #fff;
    border-radius: 8px;
    margin: 6px 8px;
    padding: 0 15px;
    cursor: pointer;
    list-style: none;
    transition: all 0.3s linear;
    outline: none;
    text-decoration: none;
}

.active {
    transform: scale(1.2);
}

a:hover {
    transform: scale(1.1);
}


</style>
