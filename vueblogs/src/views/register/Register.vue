<template>
  <div class="register">
    <b-row class="mt-5">
      <b-col md="8" offset-md="2" lg="6" offset-lg="3">
        <b-card title="注册">
          <b-form>
            <b-form-group lable="姓名">
              <b-form-input
                v-model="$v.user.name.$model"
                type="text"
                placeholder="请输入您的名称（选填）"
              ></b-form-input>
            </b-form-group>
            <b-form-group lable="手机号">
              <b-form-input
                v-model="$v.user.telephone.$model"
                type="number"
                placeholder="请输入您的手机号"
                :state="validateStatus('telephone')"
              ></b-form-input>
              <!--              <b-form-text id="password-help-block" text-variant="danger"-->
              <!--                           v-if="showTelephoneValidate">手机号必须是11位-->
              <!--              </b-form-text>   方法一 -->
<!--              <b-form-invalid-feedback :state="validation">-->
              <b-form-invalid-feedback :state="validateStatus('telephone')">
                手机号不符合规则
              </b-form-invalid-feedback>
<!--                <b-form-valid-feedback :state="validation">-->
<!--                  Looks Good.-->
<!--                </b-form-valid-feedback>-->
            </b-form-group>
            <b-form-group lable="姓密码名">
              <b-form-input
                :state="validateStatus('password')"
                v-model="$v.user.password.$model"
                type="password"
                placeholder="请输入您的密码"
              ></b-form-input>
              <b-form-invalid-feedback :state="validateStatus('password')">
                密码必须大于6位
              </b-form-invalid-feedback>
            </b-form-group>
            <b-form-group>
              <b-button @click="register" variant="outline-primary" block>注册</b-button>
            </b-form-group>
          </b-form>
        </b-card>
      </b-col>
    </b-row>
  </div>
</template>
<script>
import { required, minLength } from 'vuelidate/lib/validators';
import customValidator from '@/helper/validator';
import { mapActions } from 'vuex';

export default {
  data() {
    return {
      user: {
        name: '',
        telephone: '',
        password: '',
      },
      showTelephoneValidate: false,
      validation: null,
    };
  },
  validations: {
    user: {
      name: {

      },
      telephone: {
        required,
        telephone: customValidator.telephoneValidator,
        // between: between(11, 11),
        // minLength: minLength(11),
        // maxLength: maxLength(11),
      },
      password: {
        minLength: minLength(6),
      },
    },
  },
  methods: {
    ...mapActions('userModule', { userRegister: 'register' }), // 将 `this.userModule()` 映射为 `this.$store.dispatch('userModule')`
    validateStatus(name) {
      // 这里是ex6 解构赋值
      const { $dirty, $error } = this.$v.user[name];
      return $dirty ? !$error : null;
    },
    register() {
    // axios 验证数据
      this.$v.user.$touch(); // 没有数据触发验证
      if (this.$v.user.$anyError) {
        return;
      }
      // 请求API方式一：嵌套调用，请求API方式2：链式调用，
      // userService.register(this.user).then((res) => {
      //   // 保存token
      //   this.$store.commit('userModule/SET_TOKEN', res.data.data.token);
      //   userService.info().then((response) => {
      //     // 保存用户信息
      //     this.$store.commit('userModule/SET_USERINFO', response.data.data.user);
      //     // 跳转主页
      //     this.$router.replace({ name: 'Home' });
      //   });
      //   return userService.info();
      // 请求API方式2：链式调用，
      // userService.register(this.user).then((res) => {
      //   // 保存token
      //   this.$store.commit('userModule/SET_TOKEN', res.data.data.token);
      //   return userService.info();
      // }).then((response) => {
      //   // 保存用户信息
      //   this.$store.commit('userModule/SET_USERINFO', response.data.data.user);
      //   // 跳转主页
      //   this.$router.replace({ name: 'Home' });
      // })
      // 请求API方式3：链式调用+map函数，
      // userService.register(this.user).then((res) => {
      //   // 保存token
      //   this.SET_TOKEN(res.data.data.token);
      //   return userService.info();
      // }).then((response) => {
      //   // 保存用户信息
      //   this.SET_USERINFO(response.data.data.user);
      //   // 跳转主页
      //   this.$router.replace({ name: 'Home' });
      // })

      // 请求API方式3：链式调用+map函数 => 重构合并到user.js
      this.userRegister(this.user).then(() => {
        console.log(this.user, '---7');
        // 跳转主页
        this.$router.replace({ name: 'Home' });
      }).catch((err) => {
        this.$bvToast.toast(err.response.data.msg, {
          title: '数据验证错误',
          variant: 'danger',
          solid: true,
        });
      });
    },
  },
};
</script>

<style lang="scss" scoped>

</style>
