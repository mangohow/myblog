
<template>

    <div class="login_container ">
        <div class="login_box">
            <div class="login_font">管理后台登录</div>
            <!-- 登陆区域 -->
            <el-form ref="loginFormRef" :model="loginForm" :rules="loginFormRules" label-width="0px" class="login_form">
                <!-- 用户名 -->
                <el-form-item prop="username">
                    <el-input v-model="loginForm.username" prefix-icon="el-icon-user-solid" placeholder="请输入用户名"></el-input>
                </el-form-item>
                <!-- 密码 -->
                <el-form-item prop="password">
                    <el-input v-model="loginForm.password" @keyup.enter.native="login" prefix-icon="el-icon-lock" type="password" placeholder="请输入密码"></el-input>
                </el-form-item>
                <!-- 按钮区域 -->
                <el-form-item class="btns">
                    <el-button type="primary" @click="login">登录</el-button>
                    <el-button type="primary" @click="resetLoginForm">重置</el-button>
                </el-form-item>
            </el-form>
        </div>
    </div>

</template>



<script>

import crypot from 'crypto'

export default {
    data() {
        return {
            // 登录表单的数据绑定对象
            loginForm: {
                username: "",
                password: ""
            },
            //表单的验证规则对象
            loginFormRules: {
                // 验证用户名
                username: [
                    { required: true, message: "请输入用户名", trigger: "blur"},
                    { min: 3, max: 11, message: "长度在 3 到 10 个字符之间", trigger: "blur"}
                ],
                // 验证密码
                password: [
                    { required: true, message: "请输入密码", trigger: "blur"},
                    { min: 6, max: 15, message: "长度在 6 到 15 个字符之间", trigger: "blur"}
                ]
            }
        }
    },
    methods: {
        resetLoginForm() {
            this.$refs.loginFormRef.resetFields();
        },
        login() {   //表单预校验
            this.$refs.loginFormRef.validate(async (valid) =>{
                if(!valid) return;     //预验证没有通过
                let md5 = crypot.createHash('md5');
                md5.update(this.loginForm.password);
                const passwd = md5.digest('hex')
                const forms = {
                    username: this.loginForm.username,
                    password: passwd
                }

                const {data: res} = await this.$axios.post("/admin/login", forms);
                if(res.status == 200) {   //登录失败
                    return this.$message.error(res.message);
                }

                if (res.data.length <= 0) {
                    return this.$message.error(res.message);
                }
                const token = res.data[0]
                const id = res.data[1]
                this.$message.success(res.message);
                // 登录成功之后:
                    //1 将登陆成功的Token保存到客户端的sessionStorage中，token只应在当前网站
                    //    打开期间生效，所以将Token保存到客户端的sessionStorage中
                    // sessionStorage 是会话期间的存储 localStorage是持久化的存储
                    //2 通过编程式导航跳转到后台主页.路由地址为home
                window.sessionStorage.setItem("token", token);
                window.sessionStorage.setItem("userId", id);
                await this.$router.push("/manageHome")
            });
        }
    }
}
</script>



<style lang="less" scoped>

.login_container {
    background: url("~@/assets/images/back7.jpg");
    width:100%;
    height: 100%;
    background-size:100% 100%;
}

.login_box {
    width: 450px;
    height: 300px;
    background-color: rgba(255, 255, 255, .3);
    border-radius: 3px;
    position: absolute;
    left: 50%;
    top: 50%;
    transform: translate(-50%, -60%);
}

.login_form {
    position: absolute;
    bottom: 0;
    width: 100%;
    padding: 0 20px;
    box-sizing: border-box;
}

.el-input {
    opacity: 0.5;
}

.btns {
    display: flex;
    justify-content: flex-end;
}

.login_font {
    font-size: 30px;
    font-weight: bold;
    color: #00B5AD;
    width: 100%;
    padding-left: 130px;
    padding-top: 30px;
}
</style>
