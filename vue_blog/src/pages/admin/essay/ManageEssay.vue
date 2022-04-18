<template>
    <div>
        <el-row>
            <el-button class="add" type="primary" plain @click="handleAdd">添加随笔</el-button>
        </el-row>
        <el-table :data="essays" border stripe>
            <el-table-column type="index"></el-table-column>
            <el-table-column label="内容" prop="content"></el-table-column>
            <el-table-column label="创建时间" width="300px">
                <template slot-scope="scope">
                    <i class="el-icon-time"></i>
                    <span style="margin-left: 10px">{{ dateFormat(scope.row.createTime) }}</span>
                </template>
            </el-table-column>
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
        <el-dialog title="添加资源分类"  :visible.sync="dialogFormVisible">
            <el-input type="textarea" :rows="10" clearable autocomplete="off" placeholder="随心所欲..." v-model="postInfo.content"></el-input>
            <div slot="footer" class="dialog-footer">
                <el-button @click="cancel">取 消</el-button>
                <el-button type="primary" @click="commit">确 定</el-button>
            </div>
        </el-dialog>
    </div>
</template>

<script>

import dayjs from "dayjs";

export default {
    name: "ManageEssay",
    data() {
        return {
            essays: [],
            total: 0,
            dialogFormVisible: false,
            queryInfo: {
                pageNum: 1,
                pageSize: 10
            },
            postInfo: {
                id: 0,
                content: ""
            }
        }
    },
    methods: {
        async getEssayList() {
            const {data:res} = await this.$axios.get("/admin/essayList", {params: this.queryInfo});
            if(res.status !== 1) {
                this.$message.error("获取列表失败，请重试！")
                return
            }
            const data = res.data.length > 0 ? res.data[0] : []
            this.essays.splice(0, this.essays.length)
            this.essays.push(...data)
            this.total = res.data.length > 1 ? res.data[1] : 0
        },
        handleAdd() {
            this.dialogFormVisible = true
            this.postInfo.id = 0
        },
        handleEdit(index) {
            this.postInfo.content = this.essays[index].content
            this.postInfo.id = this.essays[index].id
            this.dialogFormVisible = true
        },
        async handleDelete(id) {
            this.$messageBox.confirm('确认删除该链接?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            }).then(async () => {
                //删除
                const {data:res} = await this.$axios.delete("/admin/deleteEssay", {params: {id: id}});
                console.log(res)
                if (res.status !== 401) {
                    this.$message.error("删除失败，请重试！")
                } else {
                    this.$message.success("删除成功！")
                }
                if (this.queryInfo.pageNum === Math.ceil(this.total / this.queryInfo.pageSize) && this.essays.length === 1) {
                    this.queryInfo.pageNum -= 1
                    if(this.queryInfo.pageNum <= 0) {
                        this.queryInfo.pageNum = 1
                        return
                    }
                }

                //刷新列表
                await this.getEssayList();
            }, () => {
                this.$message({type: 'info', message: '已取消删除'});
            });
        },
        handleSizeChange: function(pagesize) {     // 监听pagesize 改变的事件
            this.queryInfo.pageSize = pagesize;
            this.getEssayList();
        },
        handleCurrentChange: function(newPage) {  // 页码值发送变化
            this.queryInfo.pageNum = newPage;
            this.getEssayList();
        },
        dateFormat(d) {
            return dayjs(d).format("YYYY-MM-DD HH:mm:ss")
        },
        cancel() {
            this.dialogFormVisible = false
        },
        async commit() {
            let res
            if(this.postInfo.id === 0) {
                res = await this.$axios.post("/admin/addEssay", this.postInfo)
            } else {
                res = await this.$axios.put("/admin/updateEssay", this.postInfo)
            }
            if (res.data.status !== 101) {
                this.$message.error("操作失败，请重试！")
            } else {
                this.cancel()
                this.postInfo.content = ""
                await this.getEssayList()
                this.$message.success("操作成功！")
            }
        }
    },
    created() {
        this.getEssayList()
    }
}
</script>

<style scoped>

.add {
    float: right !important;
    margin-right: 50px;
}

</style>