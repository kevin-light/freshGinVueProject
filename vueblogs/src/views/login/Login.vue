<template>
  <div class="ant-form-rg" >
    <a-form :form="form" @submit="handleSubmit" >
      <a-form-item v-bind="tailFormItemLayout">
        <h2 class="abb-fh">登录</h2>
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
            ],
          },
        ]"
          type="password"
        />
      </a-form-item>
      <a-form-item v-bind="tailFormItemLayout">
        <a class="abb-fh" >
          <a-button type="primary" html-type="submit" htmlType="submit" >
            Login
          </a-button>
        </a>
      </a-form-item>
    </a-form>
  </div>
</template>

<script>

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
  beforeCreate() {
    this.form = this.$form.createForm(this, { name: 'register' });
  },
  methods: {
    ...mapActions(['logins']),
    handleSubmit(e) {
      e.preventDefault();
      this.form.validateFieldsAndScroll((err, values) => {
        if (!err) {
          this.user.Telephone = values.phone;
          this.user.Password = values.password;
          console.log(this.user, '---l1', values.nickname);
          // 请求     //...mapActions函数
          this.logins(this.user).then(() => {
            // 跳转到主页
            this.$router.replace({ name: 'Home' });
          }).catch((error) => {
            console.log('err', error.response.data.msg, '---e1');
          });
        }
      });
    },
    handleConfirmBlur(e) {
      const { value } = e.target;
      this.confirmDirty = this.confirmDirty || !!value;
    },
    // compareToFirstPassword(rule, value, callback) {
    //   const { form } = this;
    //   if (value && value !== form.getFieldValue('password')) {
    //     callback('Two passwords that you enter is inconsistent!');
    //   } else {
    //     callback();
    //   }
    // },
    // validateToNextPassword(rule, value, callback) {
    //   const { form } = this;
    //   if (value && this.confirmDirty) {
    //     form.validateFields(['confirm'], { force: true });
    //   }
    //   callback();
    // },
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
