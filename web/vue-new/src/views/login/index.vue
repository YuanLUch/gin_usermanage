<template>
  <div class="container">
    <div class="loginBox">
      <h2 style="margin-left: 170px;">用户登录</h2>
      <el-form ref="loginFormRef" :rules="rules" :model="formdata" class="loginForm">
        <el-form-item prop="username" label="用户名">
          <el-input v-model="formdata.username" placeholder="请输入用户名" type="text">
          </el-input>
        </el-form-item>
        <el-form-item prop="password" label="密码">
          <el-input v-model="formdata.password" placeholder="请输入密码" type="password"></el-input>
        </el-form-item>
        <!-- <el-form-item>
          <el-input></el-input>
        </el-form-item> -->
       <div class="buttonContainer">
        <a href="http://localhost:8080/#/register" class="registerLink">没有账号？点击注册</a>
        <el-form-item class="loginBtn">
          <el-button type="primary" @click="login">登录</el-button>
          <el-button type="info" @click="resetForm">取消</el-button>
        </el-form-item>
      </div>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      formdata: {
        username: '',
        password: ''
      },
      rules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          {
            min: 4,
            max: 12,
            message: '用户名必须在4到12个字符之间',
            trigger: 'blur'
          }
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' },
          {
            min: 6,
            max: 20,
            message: '密码必须在6到20个字符之间',
            trigger: 'blur'
          }
        ],
      }
    }
  },
  methods: {
    resetForm() {
      this.$refs.loginFormRef.resetFields()
    },
    login() {
      this.$refs.loginFormRef.validate(async valid => {
        if (!valid) return this.$message.error("输入非法数据，请重新输入")
        const { data: res } = await this.$http.post('login', this.formdata)

        // console.log(res.username)
        // console.log(res.data.message)
        if (res.status !== 200) return this.$message.error(res.message)
        // window.sessionStorage.setItem('username', res.username)
        // window.sessionStorage.setItem('user_id', res.ID)
        window.sessionStorage.setItem('token', res.token)
        // window.sessionStorage.setItem('role', res.role)
        
        if (res.role === 2) {
          this.$message.success('管理员登录成功')
          this.$router.push('/users')
        } else {
          this.$router.push('/dashboard')
          this.$message.success("登录成功")
        }
        
      })
    }
  }
}
</script>

<style scoped>
.container {
  height: 100%;
  background-color: #282c34;
}

.loginBox {
  width: 450px;
  height: 300px;
  background-color: #fff;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  border-radius: 9px;
}

.loginForm {
  width: 100%;
  position: absolute;
  bottom: 0%;
  padding: 0 20px;
  box-sizing: border-box;
}

.loginBtn {
  display: flex;
  justify-content: flex-end;
}

.buttonContainer {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.registerLink {
  font-size: 15px; /* 设置字体大小 */
  text-decoration: underline; /* 添加底部横线 */
  color: rgb(15, 123, 191); /* 设置颜色为蓝色 */
}
</style>