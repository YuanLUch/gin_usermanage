<template>
<div class="edit-user-container">
  <!-- <div>
      用户角色：
        <span>{{ userInfo.role == '1' ? '用户' : '管理员' }}</span>
    </div> -->
  <el-card>
    <el-form :model="userInfo" :rules="userRules" ref="editUserRef" label-width="100px">
      <el-form-item label="用户名" prop="username">
        <el-input v-model="userInfo.username"></el-input>
      </el-form-item>

    <el-form-item label="电话号码" prop="phone_number">
      <el-input v-model="userInfo.phone_number"></el-input>
    </el-form-item>

    <el-form-item label="邮箱" prop="email">
      <el-input v-model="userInfo.email"></el-input>
    </el-form-item>

    </el-form>
   
    <div class="user-edit-buttons">
        <el-button type="primary" @click="editUserOk">确 定</el-button>
        <el-button type="danger" @click="editUserCancel">取 消</el-button>
    </div>
  </el-card>
</div>
</template>

<script>
  export default {
    data() {
      return {
        userInfo: {
          username: '',
          phone_number: '',
          email: '',
          role: ''
        },
        userRules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          {
            validator: (rule, value, callback) => {
              if (!value) {
                callback(new Error('请输入用户名'));
              }
              if (value.length < 4 || value.length > 12) {
                callback(new Error('用户名应当在4到12个字符之间'));
              } else {
                callback();
              }
            },
            trigger: 'blur',
          },
        ],
        phone_number: [
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
              'Authorization': `Bearer ${token}` // 在请求头中携带 token
          }
        })
        this.userInfo = res.data
      },
      editUserOk() {
       this.$refs.editUserRef.validate(async (valid) => {
        const token = sessionStorage.getItem('token')
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
        const { data: res } = await this.$http.put('user/edit', {
          username: this.userInfo.username,
          password: this.userInfo.password,
          phone_number: this.userInfo.phone_number,
          email: this.userInfo.email,
          role: this.userInfo.role,
        }, {
          headers: {
                'Authorization': `Bearer ${token}` 
          }
        })
        if (res.status != 200) return this.$message.error(res.message)
        this.$message.success('更新用户信息成功')
        this.getUserInfo()
      })
    },
    editUserCancel() {
      this.$refs.editUserRef.resetFields()
      this.$message.info('编辑已取消')
    },
    }
  }
</script>

<style>
.edit-user-container{
  margin-top: 50px;
}
.user-edit-buttons {
  text-align: center;
  margin-top: 20px;
}

</style>