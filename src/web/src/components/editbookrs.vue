<template>
	<div class="editbookinfo">
		<div class="uploadimg">
			<div class="title">
				图书图片
			</div>
			<div class="bookimg">
				<img :src="bookmsg.imageurl" >
			</div>
			<!-- <div class="more">
				<img src="../assets/image/mine/more.png" >
			</div> -->
		</div>
		<ul>
			<li>
				<mt-field label="书名" placeholder="请输入书名" disabled="true" v-model="bookmsg.bookname"></mt-field>
			</li>
			<li>
				<mt-field label="作者" placeholder="请输入作者" disabled="true" v-model="bookmsg.author"></mt-field>
			</li>
			<li>
				<mt-field label="图书简介" placeholder="请输入图书简介" disabled="true" type="textarea" rows="4" v-model="bookmsg.describe"></mt-field>
			</li>
		</ul>
		<div class="btn">
			<mt-button type="default" @click="cancel()">删除本书</mt-button>
			<mt-button type="primary" @click="confirm()">保存</mt-button>
		</div>
	</div>
</template>
<script>
	import { Toast } from 'mint-ui';
	export default{
		name:'editbookinfo',
		data () {
			return{
				bookmsg:{},
				isbn:'',
			}
		},
		mounted(){
		    this.getbookmsg();
		},
		methods:{
			getbookmsg () {
				this.$axios({
		      		method: 'post',
			    	url:'/bookrack/getbookbysn',	
			    	data:{
			    		isbn:this.$route.params.isbn,
			    	}
				})
				.then(function(res){
					console.log(res)
					this.bookmsg = res.data.data;
					if(res.data.status){
						this.bookmsg = res.data.data;
					}else{
						Toast({
						  message: '没有此书数据,请拍照添加',
						  position: 'bottom',
						  duration: 2000
						});
						that.$router.go(-1);
					}
				}.bind(this))
				.catch(function (error) {
					console.log(error);
				});
			},
			confirm () {
				this.$axios({
		      		method: 'post',
			    	url:'/bookrack/bookrackadd',	
			    	data:{
			    		bookid:this.bookmsg.bookid,
			    	}
				})
				.then(function(res){
					console.log(res)
					Toast({
					  message: '添加成功',
					  position: 'bottom',
					  duration: 2000
					});
					this.$router.go(-1);
				}.bind(this))
				.catch(function (error) {
					console.log(error);
				});
			},
			cancel () {
				this.$router.go(-1);
			},
		},
	}
</script>
<style scoped>
	.editbookinfo{
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
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    -webkit-flex-wrap: wrap;
   		flex-wrap: wrap;
	}
	.btn button{
		box-flex:1;
		-webkit-box-flex:1;
		-moz-box-flex:1;
		flex:1;
		-webkit-flex:1;
		display: block;
		height:100%;
		width:90%;
		margin:0 auto;
	}
	.uploadimg{
		background: #fff;
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    -webkit-flex-wrap: wrap;
   		flex-wrap: wrap;
   		padding:0.3rem;
   		margin-bottom: 0.3rem;
	}
	.uploadimg .title{
		box-flex:4;
		-webkit-box-flex:4;
		-moz-box-flex:4;
		flex:4;
		-webkit-flex:4;
		-webkit-align-self: center;
    	align-self: center;
	}
	.uploadimg .bookimg{
		box-flex:14;
		-webkit-box-flex:14;
		-moz-box-flex:14;
		flex:14;
		-webkit-flex:14;
		-webkit-align-self: center;
    	align-self: center;
    	display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
    	-webkit-flex-direction: row-reverse;
    	flex-direction: row-reverse;
    	padding:0 0.3rem;
	}
	.uploadimg .bookimg img{
		width:1.2rem;
		height:1.5rem;
	}
	.uploadimg .more{
		box-flex:1;
		-webkit-box-flex:1;
		-moz-box-flex:1;
		flex:1;
		-webkit-flex:1;
		-webkit-align-self: center;
    	align-self: center;
	}
</style>