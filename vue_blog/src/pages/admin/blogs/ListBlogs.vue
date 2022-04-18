
<template>

    <div>
        <!-- 卡片视图区域 -->
        <el-card class="box-card">
            <!-- 搜索栏 -->
            <el-row :gutter="15">
                <el-col :span="8">
                    <el-input v-model="queryInfo.blogTitle" placeholder="标题" clearable></el-input>
                </el-col>

                <el-col :span="4">
                    <el-select :value="selectedType" @visible-change="getAllTypes" @change="changeType" clearable  placeholder="分类">
                        <el-option v-for="item in types" :key="item.id" :value="item.name">
                        </el-option>
                    </el-select>
                </el-col>

                <el-col :span="2">
                    <!-- 推荐选项 -->
                    <el-checkbox v-model="queryInfo.recommended" label="推荐" border></el-checkbox>
                </el-col>

                <el-col :span="4">
                    <el-button @click="clear" type="info">clear</el-button>
                    <el-button type="primary" @click="getBlogList" icon="el-icon-search">搜索</el-button>
                </el-col>
            </el-row>

            <!-- 博客列表区域 -->
            <el-table :data="blogs" border stripe>
                <el-table-column type="index"></el-table-column>
                <el-table-column label="标题" prop="title"></el-table-column>
                <el-table-column label="类型" prop="typename" width="150"></el-table-column>
                <el-table-column label="推荐" width="100">
                    <template slot-scope="scope">
                        {{scope.row.recommend ? '是' : '否'}}
                    </template>
                </el-table-column>
                <el-table-column label="状态" width="100">
                    <template slot-scope="scope">
                        {{scope.row.published ? '发布' : '草稿'}}
                    </template>
                </el-table-column>
                <el-table-column label="发布时间" width="200">
                    <template slot-scope="scope">
                        <i class="el-icon-time"></i>
                        <span style="margin-left: 10px">{{ dateFormat(scope.row.createTime) }}</span>
                    </template>
                </el-table-column>
               <el-table-column label="操作"  width="150">
                    <template slot-scope="scope">
                        <el-button size="mini" @click="handleEdit(scope.row.id)">编辑</el-button>
                        <el-button size="mini" type="danger" @click="handleDelete(scope.row.id)">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>

            <!-- 分页区域 -->
             <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange"
            :current-page="queryInfo.pageNum" :page-sizes="[8, 10, 12, 15]" :page-size="queryInfo.pageSize"
            layout="total, sizes, prev, pager, next, jumper" :total="total">
            </el-pagination>
        </el-card>
    </div>

</template>



<script>
import dayjs from "dayjs";

export default {
    data() {
        return {
            types: [],
            selectedType: "",
            blogs: [],
            total: 0,
            queryInfo: {          //获取博客列表的数据
                pageNum: 1,
                pageSize: 8,
                blogTitle: "",
                typeId: 0,
                recommended: false,
            }
        }
    },
    created() {
        this.getBlogList();   //获取博客列表
    },
    methods: {
        clear: function() {
            this.queryInfo.blogTitle = "";
            this.queryInfo.typeId = 0;
            this.queryInfo.recommended = false;
            this.selectedType = "";
        },
        getBlogList: async function() {
            if(this.queryInfo.blogTitle === "") {
                this.queryInfo.blogTitle = null;
            }
            if(this.queryInfo.typeId === 0) {
                this.queryInfo.typeId = null;
            }
            if(this.queryInfo.recommended === false) {
                this.queryInfo.recommended = null;
            }

            const{data: res} = await this.$axios.get("/admin/searchBlogs",{params: this.queryInfo});
            if(res.status === 1) {
                if (res.data.length > 0) {
                    this.blogs = [...res.data[0]];
                }
                if (res.data.length > 1) {
                    this.total = res.data[1]
                }
            } else {
                this.$message.error("查询失败,请重新查询");
            }

        },

        getAllTypes: async function(change) {
            if(change === true) {
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
            const item = this.types.find(item => {
                return item.name === name
            });

            if (item) {
                this.queryInfo.typeId = item.id;
            }
        },
        handleEdit: function(id) {
            this.$router.push({
                path: "/addBlog",
                query: {
                    id: id
                }
            });
        },
        handleDelete: async function(id) {
            this.$messageBox.confirm('确认删除该博客?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            }).then(async () => {
                //删除博客
                const{data: res} = await this.$axios.delete("/admin/deleteBlog", {params: {id:id}});
                if(res.status === 401) {
                    this.$message.success("删除成功");
                } else {
                    this.$message.error("删除失败");
                }
                if (this.queryInfo.pageNum === Math.ceil(this.total / this.queryInfo.pageSize) && this.blogs.length === 1) {
                    this.queryInfo.pageNum -= 1
                    if(this.queryInfo.pageNum <= 0) {
                        this.queryInfo.pageNum = 1
                        return
                    }
                }

                //刷新博客列表
                await this.getBlogList();
            }, () => {
                this.$message({type: 'info', message: '已取消删除'});
            });

        },
        handleSizeChange: function(pagesize) {     // 监听pagesize 改变的事件
            this.queryInfo.pageSize = pagesize;
            this.getBlogList();
        },
        handleCurrentChange: function(newPage) {  // 页码值发送变化
            this.queryInfo.pageNum = newPage;
            this.getBlogList();
        },
        dateFormat(d) {
            return dayjs(d).format("YYYY-MM-DD HH:mm:ss")
        }
    }
}
</script>



<style lang="less" scoped>



</style>
