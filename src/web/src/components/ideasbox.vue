<template>
	<div class="ideasbox">
		<div class="warmtips">
			<p>
				一个建议或许是成就一个出色产品的开端！
			</p>
			<p>
				若有好的建议，请告知我们！感恩！感谢！
			</p>
		</div>
		<h3>问题反馈</h3>
		<mt-field placeholder="请输入您在使用中出现的问题，和您对我们的产品的建议" type="textarea" rows="5"></mt-field>
		<div class="screenshot">
			<p>
				上传照片(问题截图)
			</p>
			<ul class="imglist">
				<li>
					<img src="../assets/image/ideasbox/upload.png" height="100%" width="100%" alt="">
				</li>
				<li>
					<img src="../assets/image/ideasbox/upload.png" height="100%" width="100%" alt="">
				</li>
				<li>
					<img src="../assets/image/ideasbox/upload.png" height="100%" width="100%" alt="">
				</li>
			</ul>
		</div>
		<div class="phonenum">
			<mt-field placeholder="请输入您的任意联系方式"></mt-field>
		</div>
		<div class="hints">
			其他意见或咨询请联系yao_han613@163.com
		</div>
		<div class="btn">
			<mt-button type="primary">完成</mt-button>
		</div>
	</div>
</template>
<script>
	export default{
		name:'ideasbox',
		data() {
			return{

			}
		}
	}
</script>
<style scoped>
	.ideasbox{
		width:100%;
		height:100%;
		background: url(../assets/image/ideasbox/lamp.png) no-repeat;
		background-size: 100% 100%;
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
	.warmtips{
		width:3.6rem;
		height:1.3rem;
		margin-top:2rem;
		margin-left: 2.1rem;
		padding:0.2rem;
	}
	.warmtips p{
		font-size: 0.22rem;
		text-indent: 0.3rem;
	}
	h3{
		padding:0.2rem 0.3rem;
		margin-top:0.4rem;
	}
	.screenshot{
		background: #fff;
		border-top:1px solid #e6e6e6;
		padding:0.2rem 0.3rem;
	}
	.screenshot p{
		padding-bottom: 0.2rem;
	}
	.imglist{
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    -webkit-flex-wrap: wrap;
   		flex-wrap: wrap;
	}
	.imglist li{
		width:1.2rem;
		height:1.2rem;
		margin-right: 0.2rem;
		margin-bottom: 0.2rem;
	}
	.phonenum{
		margin-top:0.2rem;
	}
	.hints{
		position: fixed;
		bottom: 12%;
		left:0;
    	width:100%;
    	text-align: center;
	    font-size:0.2rem;
	    color:#7e7f81;
	    margin:0 auto;
	}
</style>