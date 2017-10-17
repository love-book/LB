<template>
	<div class="editbookinfo">
		<div class="uploadimg">
			<div class="title">
				图书图片
			</div>
			<div class="bookimg">
				<img src="../assets/image/bookstore/text.jpg" >
			</div>
			<div class="more">
				<img src="../assets/image/mine/more.png" >
			</div>
		</div>
		<ul>
			<li>
				<mt-field label="书名" placeholder="请输入书名" v-model="username"></mt-field>
			</li>
			<li>
				<mt-field label="作者" placeholder="请输入作者" v-model="username"></mt-field>
			</li>
			<li>
				<mt-field label="新旧程度" placeholder="请输入例如：9成新 正版" v-model="username"></mt-field>
			</li>
			<li>
				<mt-field label="图书简介" placeholder="请输入图书简介" v-model="username"></mt-field>
			</li>
		</ul>
		<div class="btn">
			<mt-button type="default">删除本书</mt-button>
			<mt-button type="primary">保存</mt-button>
		</div>
	</div>
</template>
<script>
	export default{
		name:'editbookinfo',
		data () {
			return{
				
			}
		}
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