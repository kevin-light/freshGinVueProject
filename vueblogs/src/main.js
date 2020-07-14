import Vue from 'vue';
// eslint-disable-next-line import/no-extraneous-dependencies
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue';
import Vuelidate from 'vuelidate';
import axios from 'axios';
import VueAxios from 'vue-axios';
import App from './App.vue';
import router from './router';
import store from './store';
import './assets/scss/index.scss';

Vue.config.productionTip = false;
Vue.use(Vuelidate);

// Install BootstrapVue; Vuelidate; axios; vue-axios
Vue.use(BootstrapVue);
// Optionally install the BootstrapVue icon components plugin
Vue.use(IconsPlugin);
// axios; vue-axios
Vue.use(VueAxios, axios);

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount('#app');
