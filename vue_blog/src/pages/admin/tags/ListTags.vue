
<template>

    <div>
        <!-- 卡片视图区域 -->
        <el-card class="box-card">
            <!-- 搜索栏 -->
            <el-row :gutter="15">
                <el-col :span="10">
                    <el-input v-model="inputTag" @blur="checkTagExist(false)" @focus="checkTagExist(true)" placeholder="标签名称" clearable></el-input>
                </el-col>
                <el-col :span="4">
                    <el-button type="primary" @click="addTag" :disabled="addBtnDisabled" icon="el-icon-edit">{{btnName}}</el-button>
                </el-col>
            </el-row>

            <!-- 分类列表区域 -->
            <el-table :data="tags" border stripe>
                <el-table-column type="index"></el-table-column>
                <el-table-column label="名称" prop="name"></el-table-column>
               <el-table-column align="center" label="操作"  width="150">
                    <template slot-scope="scope">
                        <el-button size="mini" @click="handleEdit(scope.row.id, scope.row.name)">编辑</el-button>
                        <el-button size="mini" type="danger" @click="handleDelete(scope.row.id)">删除</el-button>
                    </template>
               </el-table-column>
            </el-table>

            <!-- 分页区域 -->
             <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange"
            :current-page="queryInfo.pageNum" :page-sizes="[5, 8, 10, 12, 15]" :page-size="queryInfo.pageSize"
            layout="total, sizes, prev, pager, next, jumper" :total="total">
            </el-pagination>

        </el-card>


    </div>

</template>



<script>
export default {
    data() {
        return {
            btnName: "新增",
            inputTag: "",
            inputTagId: 0,
            tags: [],
            total: 0,
            addBtnDisabled: false,
            queryInfo: {
                pageNum: 1,
                pageSize: 6,
            }
        }
    },
    created() {
        this.getTagList();   //获取类型列表
    },
    methods: {
        getTagList: async function() {
            const{data: res} = await this.$axios.get("/admin/getPageTags", {params: this.queryInfo})
            if(res.status == 1 && res.data.length > 1) {
                this.tags = res.data[0];
                this.total = res.data[1]
            } else {
                this.$message.error("查询列表失败")
            }
        },
        checkTagExist: async function(focused) {
            if(focused) {     //获得焦点
                this.addBtnDisabled = false;
            }
            // 失去焦点
            if(this.inputTag === "") {
                return;
            }

            const{data: res} = await this.$axios.get("/admin/checkTagExist", {params: {tagName:this.inputTag} });
            if(res.status == 1) {
                if(res.data[0] == true) {
                    this.addBtnDisabled = true;
                } else {
                    this.addBtnDisabled = false;
                }
            }
        },
        handleEdit: function(id, name) {
            this.inputTag = name;
            this.inputTagId = id;
            this.btnName = "提交";
        },
        handleDelete: async function(tagId) {
            if(tagId <= 0) return;
            this.$messageBox.confirm('此操作将永久删除该文件, 是否继续?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            }).then(async () => {
                const {data:res} = await this.$axios.delete("/admin/deleteTag", {params: {id: tagId}});
                if(res.status === 401) {
                    this.$message.success(res.message);
                } else {
                    this.$message.error("删除失败");
                }
                if (this.queryInfo.pageNum === Math.ceil(this.total / this.queryInfo.pageSize) && this.tags.length === 1) {
                    this.queryInfo.pageNum -= 1
                    if(this.queryInfo.pageNum <= 0) {
                        this.queryInfo.pageNum = 1
                        return
                    }
                }

                await this.getTagList();
            }, () => {
                this.$message({type: 'info', message: '已取消删除'});
            });
        },
        handleSizeChange: function(pageSize) {
            this.queryInfo.pageSize = pageSize;
            this.getTagList();
        },
        handleCurrentChange: function(newPage) {
            this.queryInfo.pageNum = newPage;
            this.getTagList();
        },
        addTag: async function() {
            if (this.inputTag === "") {
                this.$message.error("请输入标签名！")
                return
            }
            if(this.btnName === "提交") {
                // 修改Type name
                const {data:res} = await this.$axios.put("/admin/updateTag", {id: this.inputTagId, name: this.inputTag});
                if(res.status === 101) {
                    this.$message.success(res.message);
                } else {
                    this.$message.error("修改失败")
                }
                this.inputTag = "";
                this.inputTagId = 0;
                this.btnName = "新增";
            } else {
                // 新增type
                const {data:res} = await this.$axios.post("/admin/addTag", {name: this.inputTag});
                if(res.status === 101) {
                    this.$message.success(res.message);
                } else {
                    this.$message.error("修改失败")
                }
                this.inputTag = "";
            }
            this.getTagList();
        }
    }

}
</script>



<style scoped>



</style>
