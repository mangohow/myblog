<template>
    <div>
        <el-card>
            <el-row>
                <el-button class="add" type="primary" plain @click="handleAdd">添加资源</el-button>
                <el-button class="show-categories" type="primary" plain @click="showCategories">查看资源分类</el-button>
            </el-row>
            <!-- 列表区域 -->
            <el-table :data="links" border stripe>
                <el-table-column type="index"></el-table-column>
                <el-table-column label="标题" prop="name" width="300px"></el-table-column>
                <el-table-column label="描述" prop="desc"></el-table-column>
                <el-table-column label="类别" prop="category" width="180px"></el-table-column>
                <el-table-column label="地址" prop="url"></el-table-column>
                <el-table-column label="操作"  width="150">
                    <template slot-scope="scope">
                        <el-button size="mini" @click="handleEdit(scope.$index)">编辑</el-button>
                        <el-button size="mini" type="danger" @click="handleDelete(scope.row.id)">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>
            <!-- 分页区域 -->
            <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange"
                           :current-page="queryInfo.pageNum" :page-sizes="[10, 12, 15, 20]" :page-size="queryInfo.pageSize"
                           layout="total, sizes, prev, pager, next, jumper" :total="total">
            </el-pagination>
            <el-dialog title="资源信息" :visible.sync="dialogFormVisible">
                <el-form label-width="80px">
                    <el-form-item label="标题：">
                        <el-input v-model="postInfo.name" autocomplete="off" clearable></el-input>
                    </el-form-item>
                    <el-form-item label="链接：">
                        <el-input v-model="postInfo.url" autocomplete="off" clearable></el-input>
                    </el-form-item>
                    <el-form-item label="类别：">
                        <el-select :value="selectedCategory" @visible-change="getAllCategories" @change="changeCategory" clearable placeholder="资源类别">
                            <el-option v-for="item in categories" :key="item.id" :value="item.name">
                            </el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item label="描述：">
                        <el-input type="textarea" :rows="3" v-model="postInfo.desc" autocomplete="off" clearable></el-input>
                    </el-form-item>
                    <el-form-item label="图标：">
                        <el-row :gutter="22">
                            <el-col :span="20">
                                <el-input v-model="postInfo.icon" autocomplete="off" clearable></el-input>
                            </el-col>
                            <el-col :span="2">
                                <el-upload :on-success="uploadSuccess" :show-file-list="false" class="upload-demo" :action="uploadIcon">
                                    <el-button  type="primary" plain>点击上传</el-button>
                                </el-upload>
                            </el-col>
                        </el-row>
                    </el-form-item>
                </el-form>
                <div slot="footer" class="dialog-footer">
                    <el-button @click="cancel">取 消</el-button>
                    <el-button type="primary" @click="commitLink">确 定</el-button>
                </div>
            </el-dialog>
            <el-dialog title="资源分类信息" :visible.sync="dialogCategoryVisible">
                <el-row>
                    <el-button class="add" type="primary" plain @click="innerVisible=true;categoryPost.id=0">添加资源分类</el-button>
                </el-row>
                <el-table :data="categories" border stripe>
                    <el-table-column type="index"></el-table-column>
                    <el-table-column label="名称" prop="name"></el-table-column>
                    <el-table-column label="操作"  width="150">
                        <template slot-scope="scope">
                            <el-button size="mini" @click="editCategory(scope.$index)">编辑</el-button>
                            <el-button size="mini" type="danger" @click="deleteCategory(scope.row.id)">删除</el-button>
                        </template>
                    </el-table-column>
                </el-table>
                <el-dialog
                    width="40%"
                    title="添加资源分类"
                    :visible.sync="innerVisible"
                    append-to-body>
                    <el-input placeholder="分类名称" v-model="categoryPost.name"></el-input>
                    <div slot="footer" class="dialog-footer">
                        <el-button @click="addCategoryCancel">取 消</el-button>
                        <el-button type="primary" @click="commitCategory">确 定</el-button>
                    </div>
                </el-dialog>
            </el-dialog>
            <div slot="footer" class="dialog-footer">
                <el-button @click="dialogCategoryVisible=false">取 消</el-button>
            </div>
        </el-card>
    </div>
</template>

<script>
import axios from "axios";

export default {
    name: "manageResLink",
    data() {
        return {
            links: [],
            dialogFormVisible: false,
            dialogCategoryVisible: false,
            innerVisible: false,
            total: 0,
            selectedCategory: "",
            postInfo: {
                id: 0,
                name: "",
                desc: "",
                url: "",
                categoryId: 0,
                icon: "",
            },
            categoryPost: {
                id: 0,
                name: ""
            },
            queryInfo: {
                pageNum: 1,
                pageSize: 10
            },
            categories: [],
            uploadIcon: axios.defaults.baseURL + "/admin/uploadIcon"
        }
    },
    methods: {
        async getLinkList() {
            const {data:res} = await this.$axios.get("/admin/pageLinks", {params: this.queryInfo});
            if(res.status !== 1) {
                this.$message.error("获取列表失败，请重试！")
                return
            }

            let links, categories, count
            if (res.data.length > 2 ) {
                links = res.data[0]
                categories = res.data[1]
                count = res.data[2]
            } else {
                this.$message.error("获取列表失败，请重试！")
                return
            }

            this.total = count
            this.categories = categories
            const m = new Map()
            this.links.splice(0, this.links.length)
            categories.forEach((val => {
                m.set(val.id, val.name)
            }))
            links.forEach((val) => {
                this.links.push({...val, category: m.get(val.categoryId)})
            })
        },
        async getAllCategories() {
            const {data:res} = await this.$axios.get("/admin/categories");
            if (res.status !== 1) {
                this.$message.error("获取列表失败，请重试！")
                return
            }
            this.categories = res.data.length > 0 ? res.data[0] : []
        },
        changeCategory(name) {
            this.selectedCategory = name
            const val = this.categories.find(item => {
                return item.name === name
            })
            this.postInfo.categoryId = val.id
        },
        async handleAdd() {
            this.postInfo.id = 0
            this.dialogFormVisible = true
        },
        async handleEdit(index) {
            this.dialogFormVisible = true
            this.postInfo = {...this.links[index]}
            const val = this.categories.find(item => {
                return item.id === this.links[index].categoryId
            })
            this.selectedCategory = val.name
        },
        async handleDelete(id) {
            this.$messageBox.confirm('确认删除该链接?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            }).then(async () => {
                //删除博客
                const {data:res} = await this.$axios.delete("/admin/deleteLink", {params: {id: id}});
                if (res.status !== 401) {
                    this.$message.error("删除失败，请重试！")
                } else {
                    this.$message.success("删除成功！")
                }
                if (this.queryInfo.pageNum === Math.ceil(this.total / this.queryInfo.pageSize) && this.links.length === 1) {
                    this.queryInfo.pageNum -= 1
                    if(this.queryInfo.pageNum <= 0) {
                        this.queryInfo.pageNum = 1
                        return
                    }
                }

                //刷新列表
                await this.getLinkList();
            }, () => {
                this.$message({type: 'info', message: '已取消删除'});
            });
        },
        handleSizeChange: function(pagesize) {     // 监听pagesize 改变的事件
            this.queryInfo.pageSize = pagesize;
            this.getLinkList();
        },
        handleCurrentChange: function(newPage) {  // 页码值发送变化
            this.queryInfo.pageNum = newPage;
            this.getLinkList();
        },
        cancel() {
            this.postInfo = {
                id: 0,
                name: "",
                desc: "",
                url: "",
                categoryId: 0,
                icon: "",
            }
            this.dialogFormVisible = false
            this.selectedCategory = ""
        },
        async commitLink() {
            let res
            if(this.postInfo.id === 0) {
                res = await this.$axios.post("/admin/addLink", this.postInfo)
            } else {
                res = await this.$axios.put("/admin/updateLink", this.postInfo)
            }
            if (res.data.status !== 101) {
                this.$message.error("操作失败，请重试！")
            } else {
                this.cancel()
                await this.getLinkList()
                this.$message.success("操作成功！")
            }
        },
        uploadSuccess(response) {
            this.postInfo.icon = response;
        },
        showCategories() {
            this.dialogCategoryVisible = true
        },
        addCategoryCancel() {
            this.categoryPost.name = ''
            this.categoryPost.id = 0
            this.innerVisible = false
        },
        async commitCategory() {
            let res
            if(this.categoryPost.id === 0) {
                res = await this.$axios.post("/admin/addCategory", this.categoryPost);
            } else {
                res = await this.$axios.put("/admin/updateCategory", this.categoryPost)
            }

            if (res.data.status === 101) {
                this.$message.success("操作成功")
                this.addCategoryCancel()
                await this.getAllCategories()
            } else {
                this.$message.error("添加失败，请重试！")
            }
        },
        editCategory(index) {
            this.categoryPost.name = this.categories[index].name
            this.categoryPost.id = this.categories[index].id
            this.innerVisible = true
        },
        async deleteCategory(id) {
            this.$messageBox.confirm('确认删除该链接?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            }).then(async () => {
                const {data:res} = await this.$axios.delete("/admin/deleteCategory", {params: {id: id}});
                if (res.status === 401) {
                    this.$message.success("删除成功")
                } else {
                    this.$message.error("添加失败，请重试！")
                }
                //刷新列表
                await this.getAllCategories();
            }, () => {
                this.$message({type: 'info', message: '已取消删除'});
            });

        }
    },
    created() {
        this.getLinkList()
    }
}
</script>

<style scoped>

.add {
    float: right !important;
    margin-right: 50px;
}

.show-categories {
    float: right;
    margin-right: 30px;
}

</style>