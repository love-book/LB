<template>
	<div class="addcontact">
		<div class="headportrait">
			<div class="headerimg">
				<img :src="usermsginfo.imgurl" >
			</div>
			<div class="nickname">
				<h3>
					{{usermsginfo.nickname}}
				</h3>
				<p>
					{{usermsginfo.signature === '' ? '该用户还未添加个性签名' : usermsginfo.signature}}
				</p>
			</div>	
		</div>
		<ul>
			<li>
				<mt-field label="QQ号" placeholder="请输入QQ号" v-model="qq"></mt-field>
			</li>
			<li>
				<mt-field label="手机号" placeholder="请输入手机号" v-model="phone"></mt-field>
			</li>
			<li>
				<mt-field label="微信号" placeholder="请输入微信号" v-model="wxno"></mt-field>
			</li>
		</ul>
		<div class="hints">
			<img src="../assets/image/bookinfo/collect.png">三种联系方式只填一种即可确认提交后双方将获得彼此联系方式
		</div>
		<div class="btn">
			<mt-button type="primary" @click="borrowbook()">确认提交</mt-button>
		</div>
	</div>
</template>
<script>
	import { Toast } from 'mint-ui';
	export default{
		name:'addcontact',
		data () {
			return{
				qq:'',
				phone:'',
				wxno:'',
				usermsginfo:{},
			}
		},
		mounted(){
		  	this.getusermsg();
		},
		methods:{
			getusermsg() {
				this.$axios({
		      		method: 'post',
			    	url:'/users/userinfo',	
			    	data:{
			    		
			    	}
				})
				.then(function(res){
					console.log(res)
					this.usermsginfo = res.data.data;
					this.qq = this.usermsginfo.qq;
					this.phone = this.usermsginfo.telphone;
					this.wxno = this.usermsginfo.wechat;
				}.bind(this))
				.catch(function (error) {
					console.log(error);
				});
			},
			borrowbook () {
				if(this.qq != '' || this.phone!= '' || this.wxno != ''){
					console.log(this.wxno)
					console.log(this.qq)
					this.$axios({
			      		method: 'post',
				    	url:'/booknews/libraryrequest',	
				    	data:{
				    		bookqid: this.$route.params.bookqid,
			            	qq: this.qq,
			            	telphone: this.phone,
			            	wechat: this.wxno,
				    	}
					})
					.then(function(res){
						console.log(res)
						if(res.status == 200){
							Toast({
							  message: res.data.msg,
							  position: 'middle',
							  duration: 2000
							});
						}else{
							Toast({
							  message: '操作失败',
							  position: 'middle',
							  duration: 2000
							});
						}
						
					}.bind(this))
					.catch(function (error) {
						console.log(error);
					});
				}else{
					Toast({
					  message: '请填写联系方式',
					  position: 'bottom',
					  duration: 2000
					});
				} 	
			}
		}
	}
</script>
<style scoped>
	.addcontact{
		height:100%;
		width:100%;
		background: #ececec;
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
	.headportrait{
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    padding:0.3rem;
	    background: #fff;
	    margin-bottom: 0.3rem;
	}
	.headportrait .headerimg{
		width:1.15rem;
		height:1.15rem;
		border-radius:100%;
		overflow: hidden;
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
	.hints{
		position: fixed;
		bottom: 12%;
		left:20%;
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    -webkit-align-items: center;
    	align-items: center;
    	width:60%;
	    font-size:0.2rem;
	    color:#7e7f81;
	}
	.hints img{
		width:0.3rem;
		height:0.3rem;
		margin-right: 0.2rem;
	}
</style>