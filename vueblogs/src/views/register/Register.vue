<template>
  <div class="ant-form-rg" >
  <a-form :form="form" @submit="handleSubmit" >
    <a-form-item v-bind="tailFormItemLayout">
       <h2 class="abb-fh">注册</h2>
    </a-form-item>
    <a-form-item v-bind="formItemLayout">
      <span slot="label">
        姓  名&nbsp;
        <a-tooltip title="What do you want to call you?">
          <a-icon type="question-circle-o" />
        </a-tooltip>
      </span>
      <a-input
        v-decorator="[
          'nickname',
          {
            rules: [{ required: false, message: 'Please input your nickname!', whitespace: true }],
          },
        ]"
      />
    </a-form-item>
    <a-form-item v-bind="formItemLayout" label="手机号 ">
      <a-input
        v-decorator="[
          'phone',
          {
            rules: [{ required: true, len:11, message: 'Please input your phone number len 11 !'} ],
          },
        ]"
        style="width: 100%"
      >
        <a-select
          slot="addonBefore"
          v-decorator="['prefix', { initialValue: '86' }]"
          style="width: 70px"
        >
          <a-select-option value="86">
            +86
          </a-select-option>
          <a-select-option value="87">
            +87
          </a-select-option>
        </a-select>
      </a-input>
    </a-form-item>
    <a-form-item v-bind="formItemLayout" label="密  码  " has-feedback>
      <a-input
        v-decorator="[
          'password',
          {
            rules: [
              {
                required: true,
                message: 'Please input your password!',
              },
              {
                validator: validateToNextPassword,
              },
            ],
          },
        ]"
        type="password"
      />
    </a-form-item>
    <a-form-item v-bind="formItemLayout" label="确认密码" has-feedback>
      <a-input
        v-decorator="[
          'confirm',
          {
            rules: [
              {
                required: true,
                message: 'Please confirm your password!',
              },
              {
                validator: compareToFirstPassword,
              },
            ],
          },
        ]"
        type="password"
        @blur="handleConfirmBlur"
      />
    </a-form-item>
    <a-form-item v-bind="tailFormItemLayout">
      <a-checkbox v-decorator="['agreement', { valuePropName: 'checked' }]">
        I have read the
        <a href="">
          agreement
        </a>
      </a-checkbox>
    </a-form-item>
    <a-form-item v-bind="tailFormItemLayout">
      <a class="abb-fh" >
      <a-button type="primary" html-type="submit" htmlType="submit" >
        Register
      </a-button>
      </a>
    </a-form-item>
  </a-form>
  </div>
</template>

<script>
// import { required, minLength, maxLength } from 'vuelidate/lib/validators';  //引入validate验证数据格式

import { mapActions } from 'vuex';

export default {
  data() {
    return {
      user: {
        Name: '',
        Telephone: '',
        Password: '',
      },

      confirmDirty: false,
      formItemLayout: {
        labelCol: {
          xs: { span: 24 },
          sm: { span: 8 },
        },
        wrapperCol: {
          xs: { span: 24 },
          sm: { span: 5 },
        },
      },
      tailFormItemLayout: {
        wrapperCol: {
          xs: {
            span: 24,
            offset: 24,
          },
          sm: {
            span: 16,
            offset: 9,
          },
        },
      },
    };
  },
  // validations: {   validate验证数据格式
  //   user: {
  //     telephone: {
  //       required,
  //       minLength: minLength(11),
  //       maxLength: maxLength(11),
  //     },
  //   },
  // },
  beforeCreate() {
    this.form = this.$form.createForm(this, { name: 'register' });
  },
  methods: {
    // validateState(name) {
    //   // 这里是es6 解构赋值
    //   const { $dirty, $error } = this.$v.user[name];
    //   return $dirty ? !$error : null;
    // },
    // // 请求服务端 demo1
    // userService.register({ ...formVal }).then((res) => {
    //   // const api = 'http://localhost:8090/api/auth/register';
    //   // this.axios.post(api, { ...formVal }).then((res) => {
    //   // 保存token
    //   // console.log(res.data, '---0');
    //   // localStorage.setItem('token', res.data.data.token);
    //   storageService.set(storageService.USER_TOKEN, res.data.data.token);
    //   userService.info().then((response) => {
    //     // 保存用户信息
    //     storageService.set(storageService.USER_INFO, JSON.stringify(response.data.data.user));
    //     // 跳转到主页
    //     this.$router.replace({ name: 'Home' });
    //   });
    // // 请求服务端 demo2
    // userService.register(this.user).then((res) => {
    //   this.$store.commit('userModules/SET_TOKEN', res.data.data.token);
    //   return userService.info();
    // }).then((response) => {
    //   this.$store.commit('userModules/SET_USERINFO', response.data.data.token);
    //   // 跳转到主页
    //   this.$router.replace({ name: 'Home' });
    // }).catch((error) => {
    //   console.log('err', error.response.data.msg, '---e1');
    // });
    // vuex 请求
    // userService.register(this.user).then((res) => {
    //   this.USER_TOKEN(res.data.data.token);
    //   return userService.info();
    // }).then((response) => {
    //   this.SET_USERINFO(response.data.data.user);
    //   // 跳转到主页
    //   this.$router.replace({ name: 'Home' });
    // }).catch((error) => {
    //   console.log('err', error.response.data.msg, '---e1');
    // });

    ...mapActions(['registers']),
    handleSubmit(e) {
      e.preventDefault();
      this.form.validateFieldsAndScroll((err, values) => {
        if (!err) {
          this.user.Name = values.nickname;
          this.user.Telephone = values.phone;
          this.user.Password = values.password;
          console.log(this.user, '---1', values.nickname);
          // 请求     //...mapActions函数
          this.registers(this.user).then(() => {
            // 跳转到主页
            this.$router.replace({ name: 'Home' });
          }).catch((error) => {
            console.log('err', error.response.data.msg, '---e1');
          });
        }
      });
    },
    // gohome() {
    //   this.$router.replace({ name: 'Home' });
    // },
    handleConfirmBlur(e) {
      const { value } = e.target;
      this.confirmDirty = this.confirmDirty || !!value;
    },
    compareToFirstPassword(rule, value, callback) {
      const { form } = this;
      if (value && value !== form.getFieldValue('password')) {
        callback('Two passwords that you enter is inconsistent!');
      } else {
        callback();
      }
    },
    validateToNextPassword(rule, value, callback) {
      const { form } = this;
      if (value && this.confirmDirty) {
        form.validateFields(['confirm'], { force: true });
      }
      callback();
    },

    //   register() {
    //     // 验证数据
    //     console.log(this.form.validateFieldsAndScroll, '---p1');
    //
    //     // 请求服务端
    //
    //     const api = 'http://localhost:8090/api/auth/register';
    //     this.axios.post(api, { ...this.user }).then((res) => {
    //       // 保存token
    //       console.log(res.data, '---0');
    //     }).catch((err) => {
    //       console.log('err', err.response.data.msg, '---e1');
    //     });
    //   },
    // },
    //   phoneCheck(rule, value, callbackFn) {
    //     const reg = /^[1][3,4,5,6,7,8,9][0-9]{9}$/;
    //     if (!reg.test(value)) {
    //       callbackFn('请输入正确的手机号码');
    //       return;
    //     }
    //     callbackFn();
    //   },
  },
};
</script>

<style lang="scss">
  .ant-form-rg{
    padding-top: 7%;
    padding-left: 70px;
  }
  .abb-fh{
    padding-left: 50px;
  }
</style>
