import Vue from 'vue';
import {
  Button, Layout, Menu, Row, Col, Form, Input, Select, Checkbox, Icon, Tooltip, Dropdown, Avatar,
  List,
} from 'ant-design-vue';
import axios from 'axios';
import VueAxios from 'vue-axios';
import Vuex from 'vuex';
import App from './App.vue';
import router from './router';
import store from './store';

// Vue.use(Vuelidate);
Vue.use(VueAxios, axios);
Vue.use(Vuex);

Vue.component(Button.name, Button);
Vue.component(Form.name, Form);
Vue.component(Row.name, Row);
Vue.component(Layout.name, Layout);
Vue.component(Icon.name, Icon);
Vue.component(Menu.name, Menu);
Vue.component(Menu.Item.name, Menu.Item);
Vue.use(Button);
Vue.use(Menu);
Vue.use(Layout);
Vue.use(Row);
Vue.use(Col);
Vue.use(Form);
Vue.use(Input);
Vue.use(Checkbox);
Vue.use(Select);
Vue.use(Icon);
Vue.use(Tooltip);
Vue.use(Dropdown);
Vue.use(Avatar);
Vue.use(List);

Vue.config.productionTip = false;

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount('#app');
