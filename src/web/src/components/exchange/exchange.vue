<template>
	<div class="exchange">
		<mt-navbar v-model="selected">
		  <mt-tab-item id="1">待交换</mt-tab-item>
		  <mt-tab-item id="2">已交换</mt-tab-item>
		  <mt-tab-item id="3">已拒绝</mt-tab-item>
		</mt-navbar>
		<mt-tab-container v-model="selected">
		  <mt-tab-container-item id="1">
		    <toexchange></toexchange>
		  </mt-tab-container-item>
		  <mt-tab-container-item id="2">
		    <exchanged></exchanged>
		  </mt-tab-container-item>
		  <mt-tab-container-item id="3">
		    <reject></reject>
		  </mt-tab-container-item>
		</mt-tab-container>
	</div>
</template>
<script>
	import toexchange from './toexchange'
	import exchanged from './exchanged'
	import reject from './reject'
	export default{
		name:'exchange',
		data() {  
		    return {  
		      	selected: '1'  
		    };  
		},
		components:{
		    toexchange,exchanged,reject
		  } 
	}
</script>
<style scoped>
	.exchange{
		height:100%;
		background: #ebeeee;
	}
	.mint-navbar{
		margin-bottom: 0.2rem;
	}
	.mint-navbar .mint-tab-item.is-selected{
		border-bottom: 3px solid #c4d3df;
	    color: #8a8a8a;
	    margin-bottom: -3px;
	}
	.navbar .mint-tab-item-label{
		font-size:0.26rem;
	}
</style>