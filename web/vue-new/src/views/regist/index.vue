<template>
    <div class="regist-container">
      <el-form ref="registForm" :model="registForm" :rules="registRules" class="regist-form" auto-complete="on" label-position="left">
  
        <div class="title-container">
          <h3 class="title">用户注册</h3>
        </div>
  
        <el-form-item prop="username">
          <span class="svg-container">
            <svg-icon icon-class="user" />
          </span>
          <el-input
            ref="username"
            v-model="registForm.username"
            placeholder="用户名"
            name="username"
            type="text"
            tabindex="1"
            auto-complete="on"
          />
        </el-form-item>

        <el-form-item prop="email">
          <span class="svg-container">
            <svg-icon icon-class="email" />
          </span>
          <el-input
            ref="email"
            v-model="registForm.email"
            placeholder="邮箱"
            name="email"
            type="text"
            tabindex="1"
            auto-complete="on"
          />
        </el-form-item>
  
        <el-form-item prop="password">
          <span class="svg-container">
            <svg-icon icon-class="password" />
          </span>
          <el-input
            :key="passwordType"
            ref="password"
            v-model="registForm.password"
            :type="passwordType"
            placeholder="密码"
            name="password"
            tabindex="2"
            auto-complete="on"
            @keyup.enter.native="handleregist"
          />
          <span class="show-pwd" @click="showPwd">
            <svg-icon :icon-class="passwordType === 'password' ? 'eye' : 'eye-open'" />
          </span>
        </el-form-item>
  
        <el-button :loading="loading" type="primary" style="width:100%;margin-bottom:30px;" @click.native.prevent="handleregist">登录</el-button>
      </el-form>
    </div>
  </template>
  
  <script>
  import { validUsername } from '@/utils/validate'
  import { tSTypeLiteral } from '@babel/types'
  
  export default {
    name: 'regist',
    data() {
      const validateUsername = (rule, value, callback) => {
        if (!validUsername(value)) {
          callback(new Error('请输入正确的用户名'))
        } else {
          callback()
        }
      }
      const validatePassword = (rule, value, callback) => {
        if (value.length < 6) {
          callback(new Error('密码至少为6位'))
        } else {
          callback()
        }
      }
      return {
        registForm: {
          username: 'admin',
          password: '123456',
          email:"111111111@qq.com"
        },
        registRules: {
          // username: [{ required: true, trigger: 'blur', validator: validateUsername }],
          password: [{ required: true, trigger: 'blur', validator: validatePassword }]
        },
        loading: false,
        passwordType: 'password',
        redirect: undefined
      }
    },
    watch: {
      $route: {
        handler: function(route) {
          this.redirect = route.query && route.query.redirect
        },
        immediate: true
      }
    },
    methods: {
      showPwd() {
        if (this.passwordType === 'password') {
          this.passwordType = ''
        } else {
          this.passwordType = 'password'
        }
        this.$nextTick(() => {
          this.$refs.password.focus()
        })
      },
      handleregist() {
        // this.$refs.registForm.validate(valid => {
        //   if (valid) {
        //     this.loading = true
        //     this.$store.dispatch('user/regist', this.registForm).then(() => {
        //       console.log('dashboard:',this.redirect)
        //       this.$router.push({ path: '/dashboard' || '/' })
        //       this.loading = false
        //     }).catch(() => {
        //       this.loading = false
        //     })
        //   } else {
        //     console.log('error submit!!')
        //     return false
        //   }
        // })
        // console.log('dashboard:',this.redirect)
        this.$router.push({ path: '/dashboard' || '/' })
      }
    }
  }
  </script>
  
  <style lang="scss">
  /* 修复input 背景不协调 和光标变色 */
  /* Detail see https://github.com/PanJiaChen/vue-element-admin/pull/927 */
  
  $bg:#283443;
  $light_gray:#fff;
  $cursor: #fff;
  
  @supports (-webkit-mask: none) and (not (cater-color: $cursor)) {
    .regist-container .el-input input {
      color: $cursor;
    }
  }
  
  /* reset element-ui css */
  .regist-container {
    .el-input {
      display: inline-block;
      height: 47px;
      width: 85%;
  
      input {
        background: transparent;
        border: 0px;
        -webkit-appearance: none;
        border-radius: 0px;
        padding: 12px 5px 12px 15px;
        color: $light_gray;
        height: 47px;
        caret-color: $cursor;
  
        &:-webkit-autofill {
          box-shadow: 0 0 0px 1000px $bg inset !important;
          -webkit-text-fill-color: $cursor !important;
        }
      }
    }
  
    .el-form-item {
      border: 1px solid rgba(255, 255, 255, 0.1);
      background: rgba(0, 0, 0, 0.1);
      border-radius: 5px;
      color: #454545;
    }
  }
  </style>
  
  <style lang="scss" scoped>
  $bg:#242b33;
  $dark_gray:#889aa4;
  $light_gray:#eee;
  
  .regist-container {
    min-height: 100%;
    width: 100%;
    background-color: $bg;
    overflow: hidden;
  
    // background-image: url('../../assets/logo.png');
    // background-size: 80%;
    // background-position: 50%;
  
    display: flex;
    align-items: center;
    .regist-form {
      position: relative;
      width: 520px;
      max-width: 100%;
      padding: 20px 50px 5px;
      margin: 0 auto;
      overflow: hidden;
      background-color: #042333;
      border-radius: 15px;
      opacity: 0.8;
    }
  
    .tips {
      font-size: 14px;
      color: #fff;
      margin-bottom: 10px;
  
      span {
        &:first-of-type {
          margin-right: 16px;
        }
      }
    }
  
    .svg-container {
      padding: 6px 5px 6px 15px;
      color: $dark_gray;
      vertical-align: middle;
      width: 30px;
      display: inline-block;
    }
  
    .title-container {
      position: relative;
  
      .title {
        font-size: 26px;
        color: $light_gray;
        margin: 0px auto 40px auto;
        text-align: center;
        font-weight: bold;
      }
    }
  
    .show-pwd {
      position: absolute;
      right: 10px;
      top: 7px;
      font-size: 16px;
      color: $dark_gray;
      cursor: pointer;
      user-select: none;
    }
  }
  </style>
  