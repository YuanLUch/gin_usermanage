<template>
  <div class="container">
    <div class="registerBox">
      <h2 style="margin-left: 250px;">用户注册</h2>
      <el-form :label-position="labelPosition" ref="registerFormRef" :rules="rules" :model="formdata" class="registerForm">
        <el-form-item prop="username" label="用户名">
          <el-input v-model="formdata.username" placeholder="请输入用户名" type="text">
          </el-input>
        </el-form-item>

        <el-form-item prop="password" label="密码">
          <el-input v-model="formdata.password" placeholder="请输入密码" type="password"></el-input>
        </el-form-item>

        <el-form-item prop="repassword" label="确认密码">
          <el-input v-model="formdata.repassword" placeholder="请输入密码" type="password"></el-input>
        </el-form-item>

        <el-form-item prop="phoneNumber" label="电话号码">
          <el-input v-model="formdata.phoneNumber" placeholder="请输入电话号码" type="text"></el-input>
        </el-form-item>

        <el-form-item prop="email" label="邮箱">
          <el-input v-model="formdata.email" placeholder="请输入邮箱地址" type="email"></el-input>
        </el-form-item>

        <!-- <el-form-item>
            <el-input></el-input>
          </el-form-item> -->

        <div class="buttonContainer">
          <a href="http://localhost:8080/#/" class="registerLink">已有账号?点击登录</a>
          <el-form-item class="registerBtn">
          <el-button type="primary" @click="register">注册</el-button>
          <el-button type="info" @click="resetForm">取消</el-button>
          <!-- <el-button type="info" @click="backLogin">返回登录页</el-button> -->
        </el-form-item>
      </div>
      </el-form>
    </div>
  </div>
</template>
  
<script>
import { number } from 'echarts'

export default {
  data() {
    return {
      labelPosition: 'right',
      formdata: {
        username: '',
        password: '',
        repassword: '',
        phoneNumber: '',
        email: ''
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

        repassword: [
          { required: true, message: '请确认密码', trigger: 'blur' },
          {
            validator: (rule, value, callback) => {
                if (this.formdata.repassword == '') {
                  callback(new Error('请确认密码'))
                }
                if (this.formdata.password !== this.formdata.repassword) {
                  callback(new Error('密码不一致，请重新输入'))
                } else {
                  callback()
                }
              },
            trigger: 'blur',
          },
        ],

        phoneNumber: [
          { required: true, message: '请输入手机号', trigger: 'blur' },
          { pattern: /^1[3|5|7|8|9]\d{9}$/, message: '请输入正确的号码格式', trigger: 'blur' }
        ],

        email: [{
          required: true,//是否必填
          message: '请输入邮箱地址',      //错误提示信息
          trigger: 'blur'              //检验方式（blur为鼠标点击其他地方，）
        },
        {
          type: 'email',                       //要检验的类型（number，email，date等）
          message: '请输入正确的邮箱地址',
          // trigger: ['blur', 'change']（change为检验的字符变化的时候）
          trigger: 'blur'
        }
        ],
      }
    }
  },
  methods: {
    resetForm() {
      this.$refs.registerFormRef.resetFields()
    },
    register() {
      this.$refs.registerFormRef.validate(async valid => {
        if (!valid) return this.$message.error("输入非法数据，请重新输入")
        // const res = await this.$http.post('user/add', this.formdata)
        const res = await this.$http.post('user/add', {
          username: this.formdata.username,
          password: this.formdata.password,
          phone_number: this.formdata.phoneNumber,
          email: this.formdata.email
        })
        // console.log(res)
        // console.log('========')
        // console.log(res.data)
        // console.log(res.data.status)
        if (res.data.status === 200) {
          alert("恭喜您注册成功！")
        } else if(res.data.status === 1001) {
          this.$message.error("用户名已被使用！")
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

.registerBox {
  width: 600px;
  height: 650px;
  background-color: #fff;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  border-radius: 9px;
}

.registerForm {
  width: 100%;
  position: absolute;
  bottom: 2%;
  padding: 0 20px;
  /* box-sizing: border-box;  */
}

.registerBtn {
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