<template>
	<div class="bindphonenum">
		<div>
			<mt-field placeholder="请输入手机号" v-model="formData.phone"></mt-field>
			<mt-field placeholder="请输入验证码" v-model="formData.code">
				<mt-button size="small" @click="getCode(formData)" :disabled="!show">
					<span v-show="show">获取验证码</span>
		         	<span v-show="!show">剩余{{count}}s</span>
				</mt-button>
			</mt-field>
		</div>
		<div class="btn">
			<mt-button type="primary" @click="subCode(formData)">完成绑定</mt-button>
		</div>
	</div>
</template>
<script>
	import { Toast } from 'mint-ui';
	const TIME_COUNT = 60;
	export default{
		name:'bindphonenum',
		data() {
			return{
				formData: {
		          phone:'',
		          code:'',
		        },
		        show: true,
		        count: '',
		        timer: null,
			}
		},
		methods:{
			getCode(formData){
	            if (!this.timer) {
	            	if(formData.phone==''){
	            		Toast({
						  message: '没有填写手机号',
						  position: 'middle',
						  duration: 2000
						});
	            	}else{
	            		this.count = TIME_COUNT;
		                this.show = false;

		                this.$axios({
				      		method: 'post',
					    	url:'/app/phonecode',	
					    	data:{
					    		telphone: formData.phone,
					    	}
						})
						.then(function(res){
							console.log(res)
						}.bind(this))
						.catch(function (error) {
							console.log(error);
						});

		                this.timer = setInterval(() => {
		                  if (this.count > 0 && this.count <= TIME_COUNT) {
		                    this.count--;
		                  } else {
		                    this.show = true;
		                    clearInterval(this.timer);
		                    this.timer = null;
		                  }
		                }, 1000)
	            	}
	            }
	        },
	        subCode (formData) {
	        	this.$axios({
		      		method: 'post',
			    	url:'/app/checkphonecode',	
			    	data:{
			    		code: formData.code,
			    		telphone: formData.phone,
			    	}
				})
				.then(function(res){
					console.log(res)
					if(res.data.status){
						Toast({
						  message: '绑定成功',
						  position: 'middle',
						  duration: 2000
						});
					}
				}.bind(this))
				.catch(function (error) {
					console.log(error);
				});
	        } 
		}
	}
</script>
<style scoped>
	.bindphonenum{
		height:100%;
		width:100%;
		background: #ececec;
		position: relative;
	}
	.btn{
		width:100%;
		height:0.9rem;
		margin-top: 0.3rem;
	}
	.btn button{
		display: block;
		height:100%;
		width:90%;
		margin:0 auto;
		background: #13b7f6;
		color:#fff;
	}
</style>