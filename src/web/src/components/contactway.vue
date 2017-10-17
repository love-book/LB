<template>
	<div class="contactway">
		<div class="headportrait">
			<div class="headerimg">
				<img :src="usermsg.imgurl" >
			</div>
			<div class="nickname">
				<h3>
					{{usermsg.nickname}}
				</h3>
				<p>
					{{usermsg.signature === '' ? '该用户还未添加个性签名' : usermsg.signature}}
				</p>
			</div>	
		</div>
		<ul class="contactlist">
			<li>
				<div class="name">
					手机号
				</div>
				<div class="num">
					{{usermsg.telphone}}
				</div>
			</li>
			<li>
				<div class="name">
					微信号
				</div>
				<div class="num">
					{{usermsg.wechat}}
				</div>
			</li>
			<li>
				<div class="name">
					QQ号
				</div>
				<div class="num">
					{{usermsg.qq}}
				</div>
			</li>
		</ul>
	</div>
</template>
<script>
	export default{
		name:'contactway',
		data() {
			return{
				usermsg:{}
			}
		},
		mounted(){
		  	this.getcontact();
		},
		methods:{
			getcontact () {
				this.$axios({
		      		method: 'post',
			    	url:'/users/userinfo',	
			    	data:{
			    		userid : this.$route.params.userid,
			    	}
				})
				.then(function(res){
					console.log(res)
					this.usermsg = res.data.data;
					console.log(this.usermsg)
				}.bind(this))
				.catch(function (error) {
					console.log(error);
				});
			}
		}
	}
</script>
<style scoped>
	.contactway{
		height:100%;
		width:100%;
		background: #ececec;
	}
	.headportrait{
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    padding:0.3rem;
	    background: #fff;
	    margin-bottom: 0.2rem;
	}
	.headportrait .headerimg{
		width:1.15rem;
		height:1.15rem;
	}
	.headportrait .headerimg img{
		width:100%;
		height:100%;
	}
	.headportrait .nickname{
		box-flex:1;
		-webkit-box-flex:1;
		-moz-box-flex:1;
		flex:1;
		-webkit-flex:1;
		padding:0.2rem 0.3rem;
	}
	.headportrait .nickname h3{
		font-size: 0.24rem;
		margin-bottom: 0.1rem;
	}
	.headportrait .nickname p{
		font-size: 0.24rem;
	}
	.contactlist{
		background: #fff;
		padding:0 0.3rem;
	}
	.contactlist li{
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    -webkit-align-items: center;
    	align-items: center;
	    padding:0.2rem 0;
	    border-bottom: 1px solid #e6e6e8;
	}
	.contactlist li:last-child{
		border-bottom: none;
	}
	.contactlist li .name{
		box-flex:1;
		-webkit-box-flex:1;
		-moz-box-flex:1;
		flex:1;
		-webkit-flex:1;
		text-align: center;
		font-size: 0.26rem;
	}
	.contactlist li .num{
		box-flex:4;
		-webkit-box-flex:4;
		-moz-box-flex:4;
		flex:4;
		-webkit-flex:4;
		font-size: 0.22rem;
		color:#ababab;
	}
</style>