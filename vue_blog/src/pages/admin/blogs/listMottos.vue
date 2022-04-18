<template>
    <div width="100%">
        <el-card class="box-card">

            <ul class="clear-fix">
                <li :key="item.id" v-for="item in mottos">
                    <div class="box">
                        <div>
                            <h5>{{ dateFormat(item.createTime) }}</h5>
                            <h6>{{ item.id }}</h6>
                        </div>
                        <p class="ch">{{ item.ch }}</p>
                        <p class="en">{{ item.en }}</p>
                    </div>
                </li>
            </ul>
            <div class="clear-fix">
                <a type="button" class="add" @click.prevent="addMotto">添加</a>
                <a type="button" class="del" @click.prevent="deleteMotto">删除</a>
                <a type="button" class="mod" @click.prevent="editMotto">修改</a>
                <input v-model.number="selected" placeholder="ID" type="number" class="num"></input>
            </div>

            <el-dialog title="格言信息" :visible.sync="dialogFormVisible">
                <el-form>
                    <el-form-item label="ch：">
                        <el-input v-model="postInfo.ch" autocomplete="off" clearable></el-input>
                    </el-form-item>
                    <el-form-item label="en：">
                        <el-input v-model="postInfo.en" autocomplete="off" clearable></el-input>
                    </el-form-item>
                </el-form>
                <div slot="footer" class="dialog-footer">
                    <el-button @click="cancel">取 消</el-button>
                    <el-button type="primary" @click="commitMotto">确 定</el-button>
                </div>
            </el-dialog>
        </el-card>
    </div>
</template>

<script>

import dayjs from "dayjs";

export default {
    name: "listMottos",
    data() {
        return {
            postInfo: {
                ch: "",
                en: "",
                id: 0
            },
            mottos: [],
            selected: null,
            dialogFormVisible: false,
        }
    },
    methods: {
        async getMottoList() {
            const {data:res} = await this.$axios.get("/admin/mottoList")
            if(res.status !== 1) {
                this.$message.error("获取列表失败，请重试！")
                return
            }
            const data = res.data.length > 0 ? res.data[0] : []
            this.mottos.splice(0, this.mottos.length)
            this.mottos.push(...data)
        },
        dateFormat(d) {
            return dayjs(d).format("YYYY-MM-DD HH:mm:ss")
        },
        editMotto() {
            const id = this.selected
            if(!this.selected) {
                this.$message.error("请输入要修改的格言编号!")
                return
            }

            this.postInfo.id = id
            const index = this.mottos.findIndex((val) => {
                return val.id === id
            })
            if(index === -1) {
                this.selected = null
                this.$message.error("请输入正确的ID值！")
                return
            }
            this.postInfo.en = this.mottos[index].en
            this.postInfo.ch = this.mottos[index].ch
            this.dialogFormVisible = true
        },
        deleteMotto() {
            if(!this.selected) {
                this.$message.error("请输入要删除的格言编号!")
                return
            }
            const id = this.selected
            const index = this.mottos.findIndex((val) => {
                return val.id === id
            })
            if(index === -1) {
                this.selected = null
                this.$message.error("请输入正确的ID值！")
                return
            }

            this.$messageBox.confirm('是否确认要删除?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            }).then(async () => {
                if(!this.selected) {
                    this.$message.warning("请输入要删除的格言编号！")
                    return
                }
                const id = this.selected
                //删除博客
                const{data: res} = await this.$axios.delete("/admin/deleteMotto", {params: {id:id}});
                if(res.status === 401) {
                    const index = this.mottos.findIndex((val) => {
                        return val.id === id
                    })
                    if(index !== -1) {
                        this.mottos.splice(index, 1)
                        this.selected = null
                        this.$message.success("删除成功");
                    } else {
                        this.$message.success("删除成功");
                    }
                } else {
                    this.$message.error("删除失败");
                }
            }, () => {
                this.$message({type: 'info', message: '已取消删除'});
            });
        },
        addMotto() {
            this.dialogFormVisible = true
        },
        async commitMotto() {
            if(!this.postInfo.en || !this.postInfo.ch) {
                this.$message.error("内容不能为空！")
                return
            }
            let res
            if(this.postInfo.id === 0) {
                res = await this.$axios.post("/admin/addMotto", this.postInfo)
            } else {
                res = await this.$axios.put("/admin/updateMotto", this.postInfo)
            }
            if (res !== 101) {
                this.$message.success("操作成功！")
                await this.getMottoList()
                this.cancel()
            } else {
                this.$message.success("操作失败，请重试！")
            }
        },
        cancel() {
            this.postInfo.en = ""
            this.postInfo.ch = ""
            this.postInfo.id = 0
            this.dialogFormVisible = false
            this.selected = null
        }
    },
    created() {
        this.getMottoList()
    }
}
</script>

<style scoped>

ul {
    padding: 0;
    margin: 0;
    list-style: none;
}

.box {
    position: relative;
    width: 300px;
    height: 240px;
    float: left;
    margin-left: 20px;
    margin-bottom: 20px;
    border: 1px solid #8f8bec;
    border-radius: 10px;
    box-shadow: 0 2px 12px 0 rgb(0 0 0 / 10%);
}

.box:hover {
    transform: translate(-3px, -3px);
    transition: all 0.1s ease;
    box-shadow: 1px 1px 3px 3px #e3e2e2;
}

.box h5 {
    float: left;
    font-size: 14px;
    margin-left: 10px;
    margin-top: 10px;
    color: #67492e;
}

.box h6 {
    float: right;
    font-size: 14px;
    line-height: 35px;
    margin-right: 20px;
}

.box h5::before {
    content: '\e61f';
    font-family: "iconfont";
    margin-right: 3px;
}

.box p {
    margin: 10px 15px;
    font-size: 15px;
}

.box .ch {
    color: #F47E60;
}

.box .en {
    color: #3498DB;
}

.box>div {
    overflow: hidden;
    height: 35px;
}


a {
    border-radius: 5px;
    border: none;
    padding: 5px 16px;
    color: #fff;
    font-size: 14px;
    margin: 3px 5px;
    text-decoration: none;
    display: block !important;
    float: right !important;
}

.num {
    width: 60px;
    height: 30px;
    border: 1px solid skyblue;
    border-radius: 6px;
    box-shadow:none;
    outline: none;
    margin-top: 4px;
    padding-left: 10px;
    display: block !important;
    float: right !important;
    margin-right: 8px;
}

.del {
    background-color: #f56c6c;
}

.del:hover {
    background-color: #fa3a3a;
}

.add,
.mod {
    background-color: #409eff;
}

.add:hover,
.mod:hover {
    background-color: #2385e8;
}

/*清除浮动*/
.clear-fix::after {
    content: '';
    display: block;
    height: 0;
    clear: both;
    visibility: hidden;
}

</style>