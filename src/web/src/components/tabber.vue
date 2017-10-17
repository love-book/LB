<template>  
  <div class="page-tabbar">  
    <div class="page-wrap" id="page-wrap">
      <transition :name="transitionName">      
        <router-view class="child-view"></router-view>
      </transition> 
    </div>  
    <!-- <mt-tabbar fixed>  
      <mt-tab-item id="1" href="#/tabber/navbar" @click.native="getselect()">  
          <img slot="icon" src="../assets/image/tabber/bookshop.png">  
          书城  
      </mt-tab-item>  
      <mt-tab-item class="bookcase" id="2" href="#/tabber/container2" @click.native="getselect()">  
        <img slot="icon" src="../assets/image/tabber/bookcase.png">  
      </mt-tab-item>  
      <mt-tab-item id="3" href="#/tabber/container3" @click.native="getselect()">  
        <img slot="icon" src="../assets/image/tabber/mine.png">  
        我的  
      </mt-tab-item>  
    </mt-tabbar> --> 
    <div class="tabbar"> 
      <router-link :to="{path:'/tabber/navbar',activeClass: 'active',exact: true}">
        <div class="icon">
          <img src="../assets/image/tabber/bookshop.png">
        </div>
        <div class="label">
          书城
        </div>
      </router-link>
      <router-link :to="{path:'/tabber/container2',activeClass: 'active',exact: true}">
        <div class="icon bookcaseicon">
          <img src="../assets/image/tabber/bookcase.png">
        </div>
      </router-link>
      <router-link :to="{path:'/tabber/container3',activeClass: 'active',exact: true}">
        <div class="icon">
          <img src="../assets/image/tabber/mine.png">
        </div>
        <div class="label">
          我的
        </div>
      </router-link>
    </div>
  </div>  
</template>  
  
<script>
import bookstore from '../components/bookstore'
import mine from '../components/mine' 
import bookcase from '../components/bookcase.vue'   
import common from './common.vue'
export default {  
  name: 'page-tabbar',  
  data() {  
    return {  
      transitionName: 'slide-left'  
    };  
  },
  created(){
      this.getselect();
  },
  methods:{
    getselect() {
      common.tabselect = this.selected;
    }
  },
  components:{
    bookstore,mine,bookcase
  }  
};  
</script>  
  
<style scoped>  
  .active{
    color:blue;
  }
  .page-tabbar {  
    overflow: hidden;
    height: 100%;  
  }  
  
  .page-wrap {  
    overflow: auto;  
    height: 100%;  
  }
  .tabbar{
    display:-webkit-box;
    display: -moz-box;
    display: -ms-flexbox;
    display: -webkit-flex;
    display: flex;
    height: 1rem;
    right: 0;
    bottom: 0;
    left: 0;
    position: fixed;
    z-index: 1;
    background-color: #fafafa;
    text-align: center;
  }
  .tabbar a{
    display: block;
    padding: 5px 0;
    box-flex:1;
    -webkit-box-flex:1;
    -moz-box-flex:1;
    flex:1;
    -webkit-flex:1;
    text-decoration: none;
  }
  .tabbar a .icon{
    width: 24px;
    height: 30px;
    margin: 0 auto 0.05rem;
  }
  .tabbar a .icon img{
    height:100%;
    width:100%;
  }
  .tabbar a .bookcaseicon{
    width: 60px;
    height: 30px;
    margin: 0.1rem auto 0.05rem;
  }
  .child-view {   
     transition: all .8s ease;
  }  
  .slide-left-enter, .slide-right-leave-active {  
    opacity: 0;
    -webkit-transform: translate(100%, 0);
    transform: translate(100%, 0);
  }  
  .slide-left-leave-active, .slide-right-enter {  
    opacity: 0;
    -webkit-transform: translate(-100%, 0);
    transform: translate(-100% 0); 
  }  




</style>  