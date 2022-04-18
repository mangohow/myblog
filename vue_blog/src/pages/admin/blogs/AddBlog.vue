
<template>

    <div width="100%">
        <el-card class="box-card">
            <el-row :gutter="10">
                <el-col :span="2">
                    <el-select :value="postInfo.flag" @change="changeWriteType" clearable>
                        <el-option v-for="item in writeTypes" :key="item.id" :value="item.type">
                        </el-option>
                    </el-select>
                </el-col>
                <el-col :span="4">
                    <el-select :value="selectedType" @visible-change="getAllTypes" @change="changeType" clearable placeholder="分类">
                        <el-option v-for="item in types" :key="item.id" :value="item.name">
                        </el-option>
                    </el-select>
                </el-col>
                <el-col :span="16">
                    <el-input v-model="postInfo.title" placeholder="标题" clearable></el-input>
                </el-col>
            </el-row>

            <el-row class="rowDown" :gutter="10">
                <el-col :span="6">
                    <el-select class="tag_select" v-model="selectedTag" @visible-change="getAllTags" clearable multiple filterable default-first-option placeholder="标签">
                        <el-option v-for="item in tags" :label="item.name" :key="item.id" :value="item.name">
                        </el-option>
                    </el-select>
                </el-col>
                <el-col :span="14">
                    <el-input v-model="postInfo.firstPicture" placeholder="首图引用地址" clearable></el-input>
                </el-col>
                <el-col :span="2">
                    <el-upload :on-success="uploadSuccess" :show-file-list="false" class="upload-demo" :action="uploadFirstPic">
                        <el-button  type="primary" plain>点击上传</el-button>
                        <div slot="tip" class="el-upload__tip">只能上传jpg/png文件</div>
                    </el-upload>
                </el-col>
            </el-row>

            <!-- markdown编辑器 -->
            <mavon-editor style="min-height:600px;font-size:18px" fontSize="16px" class="markdown"
            :ishljs="true" :scrollStyle="true" v-model="postInfo.content" ref=md @imgAdd="$imgAdd"/>

            <!-- 博客描述 -->
            <el-input class="textarea" type="textarea" :rows="5" placeholder="博客描述..." v-model="postInfo.description">
            </el-input>

            <el-checkbox v-model="postInfo.appreciation" label="赞赏"></el-checkbox>
            <el-checkbox v-model="postInfo.commentabled" label="评论"></el-checkbox>
            <el-checkbox v-model="postInfo.recommend" label="推荐"></el-checkbox>
            <el-checkbox v-model="postInfo.shareStatement" label="转载"></el-checkbox>

            <el-row class="endbtns" type="flex" justify="end">
                <el-button type="primary" plain @click="saveBlog(false)">保存</el-button>
                <el-button type="success" plain @click="publishBlog">发布</el-button>
            </el-row>
        </el-card>
    </div>

</template>



<script>
import axios from "axios";

export default {
    data() {
        return {
            markdownText: "",
            writeTypes: [
                {id: 1, type: "原创"},
                {id: 2, type: "转载"},
                {id: 3, type: "翻译"}
            ],
            types: [],
            selectedType: "",
            tags: [],
            selectedTag: [],
            postInfo: {
                id: 0,
                userId: 0,
                title: "",
                firstPicture: "",
                flag: "原创",           //原创|转载|翻译
                typeId: 0,
                published: false,
                content: "",
                commentabled: false,       //是否可以评论
                appreciation: false,       //是否开启赞赏
                recommend: false,          //是否推荐
                shareStatement: false,    //是否可以转载
                description: "",
                tagIds: [],
            },
            uploadFirstPic: axios.defaults.baseURL + "/admin/uploadImages"
        }
    },
    created() {
        this.checkEditOrAdd();
        this.getAllTypes(true)
        this.getAllTags(true)
    },
    methods: {
        changeWriteType: function(name) {
            this.postInfo.flag = name;
        },
        getAllTypes: async function(change) {
            if(change === true) {    //下拉框出现时为true 隐藏时为false
                const{data: res} = await this.$axios.get("/admin/getAllTypes");
                if(res.status === 1) {
                    this.types = res.data.length > 0 ? res.data[0] : [];
                } else {
                    this.$message.error(res.message);
                }
            }
        },
        changeType: function(name) {
            this.selectedType = name;
            var item = this.types.find(item => {
                return item.name === name
            })
            if (item) {
                this.postInfo.typeId = item.id;
            }

        },
        getAllTags: async function(change) {
            if(change == true) {    //下拉框出现时为true 隐藏时为false
                const{data: res} = await this.$axios.get("/admin/getAllTags");
                if(res.status == 1) {
                    this.tags = res.data.length > 0 ? res.data[0] : [];
                } else {
                    this.$message.error(res.message);
                }
            }
        },

        checkEditOrAdd: async function() {      //编辑博客或是新增博客
            const blogId = this.$route.query.id
            if(typeof blogId === "undefined") {
                this.postInfo.id = null
                return
            }

            this.postInfo.id = blogId;
            const {data: res} = await this.$axios.get("/admin/getFullBlog", {params: {id: blogId}});
            if(!res.status) {
                this.$message.error("出现了一点小问题，请重试");
                await this.$router.push("/listBlogs");
                return
            }

            let blogs, tags
            if (res.data.length > 1) {
                blogs = res.data[0]
                tags = res.data[1]
            } else {
                this.$message.error("出现了一点小问题，请重试");
                await this.$router.push("/listBlogs");
                return
            }

            this.postInfo = { ...this.postInfo, ...blogs}
            this.postInfo.tagIds = tags.map(x => {return x.id});
            this.markdownText = this.postInfo.content;
            this.selectedType = blogs.typename;
            this.selectedTag = tags.map(x => {return x.name});
        },
        setPostInfo: function() {
            const userId = window.sessionStorage.getItem("userId");
            this.postInfo.userId = parseInt(userId);
            for (let i = 0; i < this.selectedTag.length; i++) {
                this.postInfo.tagIds[i] = this.tags.find(item => {return item.name === this.selectedTag[i]}).id
            }

            //如果没有填写博客描述，就把标题作为博客描述
            if(!this.postInfo.description) {
                this.postInfo.description = this.postInfo.title;
            }
        },
        saveBlog: async function(published) {
            this.setPostInfo();
            this.postInfo.published = published;
            const {data: res} = await this.$axios.put("/admin/updateBlog",this.postInfo);
            if(res.status === 101) {
                this.$message.success(res.message);
                await this.$router.push("/listBlogs");
            } else {
                this.$message.error("保存失败,请重试");
            }
        },
        publishBlog: async function() {
            await this.saveBlog(true);
        },
        //图片上传
        // 绑定@imgAdd event
        $imgAdd(pos, $file){
            // 第一步.将图片上传到服务器.
            const formdata = new FormData();
            formdata.append('file', $file);
           this.$axios({
               url: '/admin/saveImages',
               method: 'post',
               data: formdata,
               headers: { 'Content-Type': 'multipart/form-data' },
           }).then((url) => {
               // 第二步.将返回的url替换到文本原位置![...](./0) -> ![...](url)
               // $vm.$img2Url 详情见本页末尾
               this.$refs.md.$img2Url(pos, url.data);
           })
        },
        uploadSuccess: function(response) {
            this.postInfo.firstPicture = response;
        }

    }
}
</script>



<style scoped>

.markdown {
    margin-top: 15px;
}

.rowDown {
    margin-top: 10px;
}

.textarea {
    margin-top: 15px;
}

.el-checkbox-group {
    margin-top: 15px;
}

.endbtns {
    margin-top: 15px;
    margin-bottom: 15px;
}

.tag_select {
    width: 300px;
}


</style>
