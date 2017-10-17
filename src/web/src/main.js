// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import './assets/js/jquery-1.9.1.min.js'
import AxiosPlugin from './assets/js/axios.js'
import MintUI from 'mint-ui'
import 'mint-ui/lib/style.css'
import './assets/css/common.css'
import App from './App'

// import axios from 'axios'
import router from './router'
import './assets/js/rem.js'
import './assets/js/touch.js'
import lrz from 'lrz'


Vue.config.productionTip = false
Vue.use(MintUI)
Vue.use(AxiosPlugin)
// axios.defaults.headers['token'] = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBpZCI6IjQzMzk5OTc3MjI5MDk2MDc5MzYtb1g4dkt3dWVUSE9DM3dyVWttMmVKQm5tLW02QSIsImV4cCI6MTUwNTcyNTkwMiwiaXNzIjoiNDMzOTk5NzcyMjkwOTYwNzkzNi1vWDh2S3d1ZVRIT0Mzd3JVa20yZUpCbm0tbTZBIiwibmJmIjoxNTA1NzIyMzAyfQ.m4a_Q7nco1xyQVJ_XMZ7r7z08xdbTYQ33Kf8bL0b3nw';
// axios.defaults.baseURL = 'http://api.kasoly.com/v1';
// Vue.prototype.$axios = axios;
/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  data:{
  	select1:1,
  },
  template: '<App/>',
  components: { App },
})
