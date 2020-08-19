<template>

  <div id="components-layout-demo-basic">

    <a-layout>
      <a-layout-header>
        <div>
          <a-row>
            <a-col :span="6">
              <h1>BlogsLogo</h1>
            </a-col>
            <a-col :span="6">
              <a-input-search  placeholder="input search text" @search="onSearch" />
            </a-col>
            <a-col :span="6">
              <a-avatar :size="37">
                <a-icon type="undo" />
              </a-avatar>
            </a-col>
            <a-col :span="4">
              <h1>十八岁的我们</h1>
            </a-col>

<!--            <a-dropdown  v-if="userInfo">-->
<!--              <a class="ant-dropdown-link" @click="e => e.preventDefault()">-->

<!--              <a-col :span="2">-->
<!--                <a-avatar :size="37">-->
<!--                  <a-icon  slot="icon" type="user" />-->
<!--                </a-avatar>-->

<!--                <a-menu slot="overlay">-->
<!--                  <a-menu-item>-->
<!--                    <a @click="gologin">退出</a>-->
<!--                  </a-menu-item>-->
<!--                  <a-menu-item>-->
<!--                    <a  @click="goregister">切换账号</a>-->
<!--                  </a-menu-item>-->
<!--                  <a-menu-item>-->
<!--                    <a href="javascript:;">个人空间</a>-->
<!--                  </a-menu-item>-->
<!--                </a-menu>-->
<!--              </a-col>-->
<!--              </a>-->
<!--            </a-dropdown>-->
            <a-dropdown  v-if="userInfo">
              <a class="ant-dropdown-link" @click="e => e.preventDefault()">
                <a-col :span="1" >
                  <a-avatar size="37">
                  <a-icon  slot="icon" type="user" />
                </a-avatar>
                </a-col>
                <a-col :span="1"><em class="userInfoSz">{{ userInfo.name }}</em></a-col>
              </a>
              <a-menu slot="overlay">
                    <a-menu-item>
                      <a @click="logouts">退出</a>
                    </a-menu-item>
                    <a-menu-item>
                      <a  @click="goregister">切换账号</a>
                    </a-menu-item>
                    <a-menu-item>
                      <a href="javascript:;">个人空间</a>
                    </a-menu-item>
              </a-menu>
            </a-dropdown>

            <a-col :span="2" v-if="!userInfo" >
              <a-button type="primary" @click="gologin" v-if="$route.name != 'login'">
                IN

              </a-button>
              <a-button type="primary" @click="goregister" v-if="$route.name != 'register'">
                OUT</a-button>
            </a-col>

          </a-row>
        </div>
      </a-layout-header>
    </a-layout>
  </div>

</template>

<script>
// import { mapState, mapActions } from 'vuex';
import { mapActions } from 'vuex';
// import storageService from '../../service/storageService';

// computed: {
//   userInfo() {
//     return this.$store.state.userModules.userInfo; // 基础用法前端用户信息=实时更新
//   },
// },
// computed: mapState({
//   userInfo: state => state.userModules.userInfo,
// }), // vuex mapState 用法前端用户信息=实时更新
export default {
  name: 'navbar',
  data() {
    return {
      // userInfo: this.userInfo,
    };
  },
  computed: {
    userInfo() {
      return JSON.parse(this.$store.state.userModules.userInfo); // 基础用法前端用户信息=实时更新
    },
  },
  // computed: mapState({ userInfo: 'userInfo' }), // vuex mapState 用法前端用户信息=实时更新
  methods: {
    ...mapActions(['logouts']), // 用户退出
    gologin() {
      this.$router.replace({ name: 'login' });
    },
    goregister() {
      this.$router.push({ name: 'register' });
    },
  },
};

</script>

<style scoped>

  #components-layout-demo-basic {
    text-align: center;
  }
  #components-layout-demo-basic .ant-layout-header,
  #components-layout-demo-basic .ant-layout-footer {
    background: #7dbcea;
    color: #fff;
  }
  .userInfoSz {
    font-size: 27px;
  }
  /*#components-layout-demo-basic .ant-layout-sider {*/
  /*  background: #3ba0e9;*/
  /*  color: #fff;*/
  /*  line-height: 120px;*/
  /*}*/
  /*#components-layout-demo-basic .ant-layout-content {*/
  /*  background: rgba(16, 142, 233, 1);*/
  /*  color: #fff;*/
  /*  min-height: 120px;*/
  /*  line-height: 120px;*/
  /*}*/
  #components-layout-demo-basic > .ant-layout {
    margin-bottom: 48px;
  }
  #components-layout-demo-basic > .ant-layout:last-child {
    margin: 0;
  }
  #components-layout-demo-basic > .a-input-search{
    width: 30px;
  }
</style>
