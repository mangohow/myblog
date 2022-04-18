
<template>

    <div>
        <!-- 卡片视图区域 -->
        <el-card class="box-card">
            <!-- 搜索栏 -->
            <el-row :gutter="15">
                <el-col :span="10">
                    <el-input v-model="inputType" @blur="checkTypeExist(false)" @focus="checkTypeExist(true)" placeholder="分类名称" clearable></el-input>
                </el-col>
                <el-col :span="4">
                    <el-button type="primary" @click="addType" :disabled="addBtnDisabled" icon="el-icon-edit">{{btnName}}</el-button>
                </el-col>
            </el-row>

            <!-- 分类列表区域 -->
            <el-table :data="types" border stripe>
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
            inputType: "",
            inputTypeId: 0,
            types: [],
            total: 0,
            addBtnDisabled: false,
            queryInfo: {
                pageNum: 1,
                pageSize: 5,
            }
        }
    },
    created() {
        this.getTypeList();   //获取类型列表
    },
    methods: {
        getTypeList: async function() {
            const{data: res} = await this.$axios.get("/admin/getPageTypes", {params: this.queryInfo})
            if(res.status == 1 && res.data.length > 1) {
                this.types = res.data[0];
                this.total = res.data[1]
            } else {
                this.$message.error("查询列表失败")
            }
        },
        checkTypeExist: async function(focused) {
            if(focused) {     //获得焦点
                this.addBtnDisabled = false;
            }
            // 失去焦点
            if(this.inputType === "") {
                return;
            }

            const{data: res} = await this.$axios.get("/admin/checkTypeExist", {params: {typeName:this.inputType} });
            if(res.status == 1) {
                if(res.data[0] == true) {
                    this.addBtnDisabled = true;
                } else {
                    this.addBtnDisabled = false;
                }
            }
        },
        handleEdit: function(id, name) {
            this.inputType = name;
            this.inputTypeId = id;
            this.btnName = "提交";
        },
        handleDelete: async function(typeId) {
            if(typeId <= 0) return;
            this.$messageBox.confirm('此操作将永久删除该文件, 是否继续?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            }).then(async () => {
                const {data:res} = await this.$axios.delete("/admin/deleteType", {params: {id: typeId}});
                if(res.status === 401) {
                    this.$message.success(res.message);
                } else {
                    this.$message.error("删除失败");
                }
                if (this.queryInfo.pageNum === Math.ceil(this.total / this.queryInfo.pageSize) && this.types.length === 1) {
                    this.queryInfo.pageNum -= 1
                    if(this.queryInfo.pageNum <= 0) {
                        this.queryInfo.pageNum = 1
                        return
                    }
                }

                await this.getTypeList();
            }, () => {
                this.$message({type: 'info', message: '已取消删除'});
            });
        },
        handleSizeChange: function(pageSize) {
            this.queryInfo.pageSize = pageSize;
            this.getTypeList();
        },
        handleCurrentChange: function(newPage) {
            this.queryInfo.pageNum = newPage;
            this.getTypeList();
        },
        addType: async function() {
            if (this.inputType === "") {
                this.$message.error("请输入类型名！")
                return
            }
            if(this.btnName === "提交") {
                // 修改Type name
                const {data:res} = await this.$axios.put("/admin/updateType", {id: this.inputTypeId, name: this.inputType});
                if(res.status === 101) {
                    this.$message.success(res.message);
                } else {
                    this.$message.error("修改失败")
                }
                this.inputType = "";
                this.inputTypeId = 0;
                this.btnName = "新增";
            } else {
                // 新增type
                const {data:res} = await this.$axios.post("/admin/addType", {name: this.inputType});
                if(res.status === 101) {
                    this.$message.success(res.message);
                } else {
                    this.$message.error("修改失败")
                }
                this.inputType = "";
            }
            this.getTypeList();
        }
    }

}
</script>



<style scoped>



</style>
