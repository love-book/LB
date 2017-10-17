<template>
	<div class="navbar" id="navbar">
		<div class="nav">
			<router-link :to="{path:'/tabber/navbar/near',activeClass: 'active',exact: true}">
		        <div class="navlabel">
		          附近
		        </div>
		    </router-link>
		    <router-link :to="{path:'/tabber/navbar/bookstore',activeClass: 'active',exact: true}">
		        <div class="navlabel">
		          书城
		        </div>
		    </router-link>
		</div>
		<div class="navcontainer" id="navcontainer">    
			<transition :name="transitionName">
	        	<router-view class="child-view"></router-view> 
	        </transition>
	    </div>
	</div>
</template>
<script>
	import common from './common.vue'
	export default{
		name:'navbar',
		data() {
			return{
				transitionName: 'slide-left'
			}
		},
		created(){
		    this.getselect();
		},
		methods:{
			getselect() {
				common.navselect = this.selected;
			}
		}
	}
</script>
<style>
	.navbar{
		height:100%;
		background: #ebeeee;
	}
	.active .navlabel{
	    border-bottom: 3px solid #26a2ff;
	    color: #000;
	    margin-bottom: -3px;
	}
	.nav{
		background-color: #fff;
	    display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    text-align: center;
	    margin-bottom: 0.15rem;
	}
	.nav a{
		box-flex:1;
	    -webkit-box-flex:1;
	    -moz-box-flex:1;
	    flex:1;
	    -webkit-flex:1;
		padding: 0 30px;
		font-size: 15px;
		border-right: 1px solid #e4e4e4;
	}
	.nav a .navlabel{
		padding: 17px 0;
    	font-size: 0.3rem;
	}
	.navcontainer{
		background: #ebeeee;
	}
	.mint-navbar{
		margin-bottom: 0.15rem;
	}
	.mint-navbar .navbaritem{
		padding:0 30px;
	}
    .navbaritem.is-selected {
		border-bottom:none!important;
	    color: #000;
	    margin-bottom: -3px;
	}
	.navbaritem.is-selected .mint-tab-item-label{
		border-bottom: 3px solid #26a2ff;
	    color: #000;
	    margin-bottom: -3px;
	}
	.navbaritem .mint-tab-item-label{
		padding:17px 0;
		font-size:0.3rem;	
	}
	.near{
		border-right:1px solid #e4e4e4;
	}
	.child-view {   
	    transition: all .8s ease;
	}  
	.slide-left-enter, .slide-right-leave-active {  
	    opacity: 0;
        -webkit-transform: translate(50px, 0);
        transform: translate(50px, 0);
	}  
	.slide-left-leave-active, .slide-right-enter {  
	    opacity: 0;
        -webkit-transform: translate(-50px, 0);
        transform: translate(-50px, 0); 
	}
</style>