<!--资源库标签小组件-->
<template>
    <a :href="links.url" :title="links.url" target="_blank">
        <span class="img" :style="bgcObj"></span>
        <div>
            <strong>{{ links.name }}</strong>
            <p :title="links.desc">{{ links.desc }}</p>
        </div>
    </a>
</template>

<script>
export default {
    name: "ResourceLabel",
    props: ["links"],
    computed: {
        bgcObj() {
            if (this.links.icon.trim() === "") {
                return {
                    backgroundColor: this.$randomColor()
                }
            }
            return {
                backgroundImage: 'url(' + this.links.icon + ')'
            }
        }
    }
}
</script>

<style scoped>

a {
    outline: none;
    text-decoration: none;
    display: inline-block;
    background-color: #fff;
    border: 1px solid #d0d0d0;
    border-radius: 3px;
    width: 360px;
    height: 80px;
    padding: 20px;
    transition: all 0.3s linear;
}

a:hover {
    transform: translateY(-5px);
    box-shadow: 0 5px 10px 3px rgba(0, 0, 0, .3);
}

.img {
    width: 40px;
    height: 40px;
    float: left;
    border-radius: 50%;
    background-size: cover;
}

div {
    float: left;
    margin-left: 12px;
    font-size: 13px;
    font-family: Arimo,"Helvetica Neue",Helvetica,Arial,sans-serif;
}

strong {
    font-weight: 700;
    color: #373e4a;
}

p {
    font-weight: 500;
    color: #979898;
    margin: 0;
    white-space: nowrap;
    overflow: hidden;
    max-width: 250px;
    text-overflow: ellipsis;
}

</style>