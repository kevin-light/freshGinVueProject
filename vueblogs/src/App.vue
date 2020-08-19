<template>
  <div id="app">
    <div id="nav">
      <Navbar />    <!--    title组件-->
    </div>
    <div class="routers">
      <router-view/>    <!--    // 路由组件-->
  </div>
  </div>
</template>

<script>
import Navbar from './views/layout/Navbar.vue';
// import Demo from './views/Demotesting.vue';

export default {
  components: { Navbar },
  data() {

  },
};
</script>
<style lang="scss">

</style>
