<template>
    <div class="bg">
        <span class="title">{{ info.title }}</span>
        <span class="desc">{{ info.desc }}</span>
    </div>
</template>

<script>
export default {
    name: "TitleArea",
    props: ["info"]
}
</script>

<style scoped>

.bg {
    height: 400px;
    width: 100%;
    position: relative;
    text-align: center;
    z-index: 0;
    user-select: none;
    font-family: '楷体';
}

.bg::before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    z-index: -1;
    background-image: url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAQAAAAECAYAAACp8Z5+AAAABmJLR0QA/wD/AP+gvaeTAAAACXBIWXMAAAsTAAALEwEAmpwYAAAAKUlEQVQImU3IMREAIAgAwJfNkQCEsH8cijjpMf6vnXlQaIiJFx+omEBfmqIEZLe2jzcAAAAASUVORK5CYII=);
}

.title {
    font-size: 36px;
    font-weight: 500;
    color: #fff;
    display: block;
    padding-top: 160px;
}

.desc {
    display: inline-block;
    margin-top: 8px;
    color: #fff;
    font-size: 18px;
}

</style>
