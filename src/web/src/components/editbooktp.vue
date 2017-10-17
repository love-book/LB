<template>
	<div class="editbookinfo">
		<div class="uploadimg">
			<div class="title">
				图书图片
			</div>
			<div class="bookimg">
				<input id="file" type="file" style="position:absolute;top:0;left:0;width:100%;height:100%;opacity: 0;" accept="image/*" @change="uploadimg($event)"/>
				<img id="url" src="../assets/image/bookstore/add.png" height="120" width="120" >
			</div>
		</div>
		<ul>
			<li>
				<mt-field label="书名" placeholder="请输入书名" v-model="bookname"></mt-field>
			</li>
			<li>
				<mt-field label="作者" placeholder="请输入作者" v-model="author"></mt-field>
			</li>
			<li>
				<mt-field label="图书简介" placeholder="请输入图书简介" type="textarea" rows="4" v-model="describe"></mt-field>
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
				bookname:'',
				author:'',
				describe:'',
				imageurl:'',
			}
		},
		mounted(){
		    
		},
		methods:{
			uploadimg (el) {
				console.log(el)
				var that = this;
				lrz(el.target.files[0])
		        .then(function (rst) {
		            //console.log(rst)
		            rst.formData.append('base64img', rst.base64); 
		            //console.log(rst.formData);
		            that.$axios({
			      		method: 'post',
				    	url:'/users/uploadfile',	
				    	data: rst.formData,
					})
					.then(function(res){
						console.log(res.data.data)
						that.imageurl = res.data.data;
	                	$("#url").attr({
	                		src: res.data.data,
	                	});
					})
					.catch(function (error) {
						console.log(error);
					});
		        })
		        .catch(function (err) {
		            console.log(err)
		        })
			},
			confirm () {
				console.log(this.bookname)
				console.log(this.author)
				console.log(this.describe)
				console.log(this.imageurl)
				if(this.bookname == ''){
					Toast({
					  message: '没有添加书名',
					  position: 'middle',
					  duration: 2000
					});
				}else if(this.author == ''){
					Toast({
					  message: '没有添加作者',
					  position: 'middle',
					  duration: 2000
					});
				}else if(this.describe == ''){
					Toast({
					  message: '没有添加图书描述',
					  position: 'middle',
					  duration: 2000
					});
				}else if(this.imageurl == ''){
					Toast({
					  message: '没有添加图书图片',
					  position: 'middle',
					  duration: 2000
					});
				}else{
					this.$axios({
			      		method: 'post',
				    	url:'/book/bookadd',	
				    	data:{
				    		imageurl: this.imageurl,
				    		bookname: this.bookname,
				    		author: this.author,
				    		describe: this.describe,
				    	}
					})
					.then(function(res){
						console.log(res)
						Toast({
						  message: '添加成功',
						  position: 'bottom',
						  duration: 2000
						});
						setTimeout(() => {
						  this.$router.go(-1);
						}, 2000);
					}.bind(this))
					.catch(function (error) {
						console.log(error);
					});
				}
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
		position: relative;
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