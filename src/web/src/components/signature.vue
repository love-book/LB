<template>
	<div class="signature">
		<div class="textarea">
			<mt-field placeholder="自我介绍" type="textarea" rows="4" v-model="signature"></mt-field>
			<div class="words">
				<span>{{b}}/10</span>
			</div>
		</div>
		<div class="btn">
			<mt-button type="primary" @click="confirm()">完成</mt-button>
		</div>
	</div>
</template>
<script>
	import { Toast } from 'mint-ui';
	export default{
		name:'signature',
		data() {
			return{
				signature:'',
			}
		},
		computed: {
		    b: function () {
		    	console.log(this.signature.length)
		    	if(this.signature.length >= 10){
		    		return 10;
		    		$('textarea').returnValue = false; 
		    	}else{
		    		return this.signature.length;
		    	}
		    }
		},
		mounted(){
		    this.getusermsg();
		},
		methods: {
			getusermsg () {
				this.$axios({
		      		method: 'post',
			    	url:'/users/userinfo',	
			    	data:{
			    		
			    	}
				})
				.then(function(res){
					console.log(res);
					this.signature = res.data.data.signature;
				}.bind(this))
				.catch(function (error) {
					console.log(error);
				});
			},
			confirm () {
				console.log(this.signature)
				this.$axios({
		      		method: 'post',
			    	url:'/users/userinfo',	
			    	data:{
			    		signature : this.signature
			    	}
				})
				.then(function(res){
					console.log(res);
					Toast({
					  message: '修改成功',
					  position: 'bottom',
					  duration: 2000
					});
				}.bind(this))
				.catch(function (error) {
					console.log(error);
				});
			}
  		},
	}
</script>
<style scoped>
	.signature{
		height:100%;
		width:100%;
		background: #ececec;
		overflow: hidden;
		position: relative;
	}
	.btn{
		position: fixed;
		bottom: 0;
		left:0;
		width:100%;
		height:0.9rem;
	}
	.btn button{
		display: block;
		height:100%;
		width:90%;
		margin:0 auto;
	}
	.textarea{
		margin-top:0.1rem;
		background: #fff;
	}
	.textarea .words{
		padding:0.1rem 0.2rem;
		text-align: right;
	}
</style>