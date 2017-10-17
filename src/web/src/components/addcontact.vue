<template>
	<div class="addcontact">
		<div class="headportrait">
			<div class="headerimg">
				<img src="../assets/image/mine/header.png" >
			</div>
			<div class="nickname">
				<h3>
					清歌
				</h3>
				<p>
					快去添加个性签名吧！
				</p>
			</div>	
		</div>
		<ul>
			<li>
				<mt-field label="QQ号" placeholder="请输入QQ号" v-model="username"></mt-field>
			</li>
			<li>
				<mt-field label="手机号" placeholder="请输入手机号" v-model="username"></mt-field>
			</li>
			<li>
				<mt-field label="微信号" placeholder="请输入微信号" v-model="username"></mt-field>
			</li>
		</ul>
		<div class="hints">
			<img src="../assets/image/bookinfo/collect.png">三种联系方式只填一种即可确认提交后双方将获得彼此联系方式
		</div>
		<div class="btn">
			<mt-button type="primary">确认提交</mt-button>
		</div>
	</div>
</template>
<script>
	export default{
		name:'addcontact',
		data () {
			return{

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