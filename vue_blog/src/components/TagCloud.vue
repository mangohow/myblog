<!--标签云组件-->
<template>
    <div class="wapper">
        <div class="title">
            <i class="iconfont icon-biaoqian"></i>
            <h4>标签</h4>
        </div>
        <div class="content">
            <ul class="clear-fix">
                <li :key="item.id" v-for="item in tags" @click="itemClicked(item.id)">
                    <span :style="{color:$randomColor()}">{{ item.name }}</span>
                </li>
            </ul>
        </div>
    </div>
</template>

<script>
export default {
    name: "TagCloud",
    data() {
      return {
          tags: [
              {id:1, name: "Java"},
              {id:2, name: "C++"},
              {id:3, name: "Golang"},
              {id:4, name: "Nginx"},
              {id:5, name: "Docker"},
              {id:6, name: "Kubernetes"},
              {id:7, name: "Mysql"},
              {id:8, name: "Redis"},
              {id:9, name: "MongoDb"},
              {id:10, name: "Gin"}
          ],
      }
    },
    methods: {
        itemClicked(id) {
            this.$router.push({
                path: "/tags",
                query: {
                    id: id
                }
            })
        },
        // 获取所有博客标签
        async getTagList() {
            const {data: res} = await this.$axios.get("/myblog/tagList");
            if(res.status === 1) {
                this.tags = res.data.length > 0 ? res.data[0] : this.tags;
            } else {
                this.$message.warning("获取标签列表失败！")
            }
        },
    },
    created() {
        this.getTagList()
    }
}
</script>

<style scoped>

ul,
li {
    margin: 0;
    padding: 0;
}

.wapper {
    width: 270px;
    border-radius: 12px;
    background-color: #fff;
    margin-top: 30px;
    user-select: none;
}

.title {
    border-bottom: 1px solid rgba(14, 14, 14, 0.1);
}

.title i {
    color: #fb6c28;
    font-size: 17px;
    margin-left: 10px;
    margin-right: 5px;
}

.title h4 {
    height: 45px;
    line-height: 50px;
    margin: 0;
    font-size: 16px;
    display: inline-block;
}

ul {
    padding: 20px 10px;
}

/*清除浮动*/
.clear-fix::after {
    content: '';
    display: block;
    height: 0;
    clear: both;
    visibility: hidden;
}

ul li {
    list-style: none;
    float: left;
    margin: 5px 8px;
    cursor: pointer;
    font-weight: 500;
    transition: all 0.3s;
}

ul li:hover {
    transform: scale(1.2);
}

</style>

