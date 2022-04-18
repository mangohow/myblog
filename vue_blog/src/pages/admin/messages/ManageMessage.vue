<template>
    <div>
        <el-table :data="messages" border stripe>
            <el-table-column type="index"></el-table-column>
            <el-table-column label="昵称" prop="nickname" width="180px"></el-table-column>
            <el-table-column label="邮箱" prop="email" width="200px"></el-table-column>
            <el-table-column label="IP" prop="ip" width="180px"></el-table-column>
            <el-table-column label="创建时间" width="200px">
                <template slot-scope="scope">
                    <i class="el-icon-time"></i>
                    <span style="margin-left: 10px">{{ dateFormat(scope.row.createTime) }}</span>
                </template>
            </el-table-column>
            <el-table-column label="内容" prop="content"></el-table-column>
            <el-table-column label="状态" width="100px">
                <template slot-scope="scope">
                    <span style="margin-left: 10px">{{ statusStr(scope.row.status)}}</span>
                </template>
            </el-table-column>
            <el-table-column label="操作"  width="150">
                <template slot-scope="scope">
                    <el-button size="mini" :disabled="scope.row.status === 1" @click="handleStatus(scope.row.id, 1)">通过</el-button>
                    <el-button size="mini" :disabled="scope.row.status === 2" type="danger" @click="handleStatus(scope.row.id, 2)">禁止</el-button>
                </template>
            </el-table-column>
        </el-table>
        <!-- 分页区域 -->
        <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange"
                       :current-page="queryInfo.pageNum" :page-sizes="[10, 12, 15, 20]" :page-size="queryInfo.pageSize"
                       layout="total, sizes, prev, pager, next, jumper" :total="total">
        </el-pagination>
    </div>
</template>

<script>

import dayjs from "dayjs";

export default {
    name: "ManageMessage",
    data() {
        return {
            messages: [],
            total: 0,
            queryInfo: {
                pageNum: 1,
                pageSize: 10
            }
        }
    },
    methods: {
        async getMessageList() {
            const {data:res} = await this.$axios.get("/admin/msgList", {params: this.queryInfo});
            if(res.data.length > 1) {
                this.messages.splice(0, this.messages.length)
                this.messages.push(...res.data[0])
                this.total = res.data[1]
            }
        },
        async handleStatus(id, status) {
            const {data:res} = await this.$axios.put("/admin/updateMsgStatus", {id: id, status: status});
            if (res.status !== 101) {
                status === 1 ? this.$message.error("评论通过失败，请重试！") : this.$message.error("评论禁止失败，请重试！")
            } else {
                status === 1 ? this.$message.success("评论通过成功！") : this.$message.success("评论禁止成功！")
                await this.getMessageList()
            }
        },
        handleSizeChange: function(pagesize) {     // 监听pagesize 改变的事件
            this.queryInfo.pageSize = pagesize;
            this.getMessageList();
        },
        handleCurrentChange: function(newPage) {  // 页码值发送变化
            this.queryInfo.pageNum = newPage;
            this.getMessageList();
        },
        dateFormat(d) {
            return dayjs(d).format("YYYY-MM-DD HH:mm:ss")
        },
        statusStr(status) {
            if (status === 0) {
                return "未审核"
            } else if (status === 1) {
                return "通过"
            } else {
                return "禁止"
            }
        }
    },
    created() {
        this.getMessageList()
    }
}
</script>

<style scoped>

</style>