<template>
  <el-card>
    <el-form :model="changePassword" :rules="changePasswordRules" ref="changePasswordRef">
      <!-- 本想设计在修改密码时验证一遍原密码，原密码经过加密，需要后端先新建密码加密接口，传入新密码加密后返回与原密码比对 -->
        <el-form-item label="原密码">
          <el-input v-model="changePassword.originpass" type="password"></el-input>
        </el-form-item>
        <el-form-item has-feedback label="新密码" prop="password">
          <el-input v-model="changePassword.password" type="password"></el-input>
        </el-form-item>
        <el-form-item has-feedback label="确认密码" prop="checkpass">
          <el-input v-model="changePassword.checkpass" type="password"></el-input>
        </el-form-item>
      </el-form>
      <div class="user-edit-buttons">
        <el-button type="primary" @click="changePasswordOk">确 定</el-button>
        <el-button type="danger" @click="changePasswordCancel">取 消</el-button>
      </div>
  </el-card>
  
</template>

<script>
const token = sessionStorage.getItem('token')
export default {
  data() {
    return {
      originPassword: '',
      changePassword: {
        id: 0,
        originpass: '',
        password: '',
        checkPass: '',
      },
      changePasswordRules: {
        originpass: [
          {
            validator: (rule, value, callback) => {
              if (this.changePassword.originpass === '') {
                callback(new Error('请输入原密码'))
              } else {
                callback()
              }
            },
            trigger: 'blur',
          },
        ],
        password: [
          {
            validator: (rule, value, callback) => {
              if (this.changePassword.password == '') {
                callback(new Error('请输入密码'))
              }
              if ([...this.changePassword.password].length < 6 || [...this.changePassword.password].length > 20) {
                callback(new Error('密码应当在6到20位之间'))
              } else {
                callback()
              }
            },
            trigger: 'blur',
          },
        ],
        checkpass: [
          {
            validator: (rule, value, callback) => {
              if (this.changePassword.checkpass == '') {
                callback(new Error('请输入密码'))
              }
              if (this.changePassword.password !== this.changePassword.checkpass) {
                callback(new Error('密码不一致，请重新输入'))
              } else {
                callback()
              }
            },
            trigger: 'blur',
          },
        ],
      },
    }
  },
  created() {
    this.getUserInfo()
  },
  methods: {
    async getUserInfo() {
        const token = sessionStorage.getItem('token')
        const {data:res} = await this.$http.get('user/info', {
          headers: {
                'Authorization': `Bearer ${token}` 
          }
        })
        this.originPassword = res.data.password
        // console.log("原密码：", this.originPassword)
    },
    async changePasswordOk() {
      const res = await this.$http.post('scrypt/pass',{
        originPass: this.changePassword.originpass
      })
      const scryptOrigin = res.data.scryPass
      // console.log("加密后的原密码：", scryptOrigin)
      this.$refs.changePasswordRef.validate(async (valid) => {
        const token = sessionStorage.getItem('token')
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
        if (this.originPassword !== scryptOrigin ) {
          this.$message.error('原密码输入错误')
        } else {
          const { data: res } = await this.$http.put('user/changepw', {
            password: this.changePassword.password,
          }, {
            headers: {
                'Authorization': `Bearer ${token}` 
          }
        })
        // console.log("修改密码后的res:", res)
        if (res.status != 200) return this.$message.error(res.message)
        this.$message.success('修改密码成功')
        }
        
      })
    },
    changePasswordCancel() {
      this.$refs.changePasswordRef.resetFields()
      this.$message.info('已取消')
    },
  }
}

</script>

<style>
.user-edit-buttons {
  text-align: center;
  margin-top: 20px;
}
</style>
