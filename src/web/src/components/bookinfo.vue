<template>
	<div class="bookinfo">
		<router-link to="">
			<div class="bookinformation" @touchstart="changecolor($event)" @touchend="reinstatecolor($event)">
				<div class="bookimg">
					<img :src="bookmsg.imageurl">
					<!-- <span class="imgnum">3张</span> -->
				</div>
				<ul class="nickname">
					<li>
						{{bookmsg.bookname}}
					</li>
					<li>
						{{bookmsg.author}}
					</li>
					<li>
						{{bookmsg.depreciation}}
					</li>
					<li class="bookstate">
						<span>{{bookmsg.is_borrow == '1' ? '可借阅' : '不可借阅'}}</span>
					</li>
				</ul>		
				<!-- <div class="moreimg">
					<img src="../assets/image/mine/more.png" >
				</div> -->
			</div>
		</router-link>
		<div class="atlas">
			<div class="imgtitle">
				<div class="left">
					<span class="line"></span>
				</div>
				<div class="center">
					<img src="../assets/image/bookinfo/img.png"><span>图片</span>
				</div>
				<div class="right">
					<span class="line"></span>
				</div>
			</div>
			<ul class="imglist" v-for="item in bookmsg.imagelist">
				<li>
					<img :src="item">
				</li>
				<!-- <li>
					<img src="../assets/image/bookstore/text.jpg">
				</li>
				<li>
					<img src="../assets/image/bookstore/text.jpg">
				</li> -->
			</ul>
		</div>
		<div class="profile">
			<div class="imgtitle">
				<div class="left">
					<span class="line"></span>
				</div>
				<div class="center">
					<img src="../assets/image/bookinfo/booklogo.png"><span>图书简介</span>
				</div>
				<div class="right">
					<span class="line"></span>
				</div>
				<!-- <div class="editlogo">
					<img src="../assets/image/bookinfo/edit.png">
				</div> -->
			</div>
			<div class="describe">
				{{bookmsg.describe}}
			</div>
		</div>
		<div class="footerbtn">
			<div class="collect">
				<img src="../assets/image/bookinfo/collect.png">
			</div>
			<button type="button" :disabled="disabled" :class="{'isborrow':bookmsg.is_borrow==2}" @click="exchangebook">{{bookmsg.is_borrow == '1' ? '点击交换' : '无法交换'}}</button>
		</div>
	</div>
</template>
<script>
	export default{
		name:'bookinfo',
		data () {
			return{
				bookmsg: {},
				disabled:false,
			};
		},
		mounted(){
		    this.getbookmsg();
		    this.is_borrow();
		},
		methods:{
			is_borrow () {
				if(this.bookmsg.is_borrow==2){
					this.disabled = true;
				}else{
					this.disabled = false;
				}
			},
			changelist (e) {
				$(".tab").find('div').css({
					background: '#fff',
					color: '#000'
				});
				$(e.currentTarget).css({
					background: '#13b8f5',
					color: '#fff'
				});
			},
			getbookmsg () {
				this.$axios({
		      		method: 'post',
			    	url:'/book/getuserbookinfo',	
			    	data:{
			    		// concern_type: '1',
			    		bookqid: this.$route.params.bookqid,
			    	}
				})
				.then(function(res){
					console.log(res.data.data)
					this.bookmsg = res.data.data;
					//this.bookmsg.concernid = res.data.data.concernid;
				}.bind(this))
				.catch(function (error) {
					console.log(error);
				});
			},
			exchangebook () {
				console.log(this.$route.params.userid)
				this.$router.push({ name: 'getcontact', params:{ bookqid: this.bookmsg.bookqid }})
			}
		}
	}
</script>
<style>
	.bookinfo{
		height:100%;
		width:100%;
		background: #ececec;
	}
	.isborrow{
		background:#e1e1e1!important;
	}
	.bookinformation{
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    -webkit-align-items: center;
    	align-items: center;
		padding:0.2rem 0.3rem;
		background: #ffffff;
	}
	.bookinformation .bookimg{
		width:1.3rem;
		height:100%;
		position: relative;
	}
	.bookinformation .bookimg img{
		width:100%;
	}
	.bookinformation .bookimg span{
		position: absolute;
		bottom:0.1rem;
		right:0.1rem;
	}
	.bookinformation .nickname{
		box-flex:1;
		-webkit-box-flex:1;
		-moz-box-flex:1;
		flex:1;
		-webkit-flex:1;
		padding:0.2rem 0.3rem;

	}
	.bookinformation .nickname li{
		margin-bottom: 0.1rem;
		font-size: 0.24rem;
	}
	.bookinformation .nickname li:last-child{
		margin-bottom:0;
	}
	.bookinformation .nickname .bookstate span{
		font-size: 0.24rem;
		color:#fe9f5a;
	}
	.bookinformation .moreimg{
		position: relative;
		width:0.2rem;
	}
	.bookinformation .moreimg img{
		width:100%;
		position: absolute;
        top: 50%;
        left: 50%;
        margin-top: -10px; 
        margin-left: -5px;
	}
	.atlas{
		padding:0 0.3rem;
		background: #ffffff;
		margin:0.18rem 0;
	}
	.atlas .imgtitle{
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    height:0.7rem;
	    border-bottom: 1px solid #ebebec;
	}
	.atlas .imgtitle .left{
		box-flex:1;
		-webkit-box-flex:1;
		-moz-box-flex:1;
		flex:1;
		-webkit-flex:1;
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    -webkit-align-items: center;
    	align-items: center;
    	-webkit-flex-direction: row-reverse;
    	flex-direction: row-reverse;
	}
	.atlas .imgtitle .left .line{
		display: inline-block;
		width:0.5rem;
		height: 2px;
		background: #d5d5d7;
	}
	.atlas .imgtitle .right{
		box-flex:1;
		-webkit-box-flex:1;
		-moz-box-flex:1;
		flex:1;
		-webkit-flex:1;
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    -webkit-align-items: center;
    	align-items: center;
    	-webkit-flex-direction: row;
    	flex-direction: row;
	}
	.atlas .imgtitle .right .line{
		display: inline-block;
		width:0.5rem;
		height: 2px;
		background: #d5d5d7;
	}
	.atlas .imgtitle .center{
		width:1.5rem;
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    -webkit-align-items: center;
    	align-items: center;
    	-webkit-justify-content: center;
    	justify-content: center;
	}
	.atlas .imgtitle .center span{
		font-size: 0.25rem;
		color:#000;
	}
	.atlas .imgtitle .center img{
		width:0.3rem;
		height:0.3rem;
		margin-right: 0.15rem;
	}
	.atlas .imglist{
		padding:0.3rem;
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    -webkit-align-items: center;
    	align-items: center;
    	-webkit-justify-content:  space-around;
    	justify-content:  space-around;
	}
	.atlas .imglist li{
		width: 1.2rem;
		height:1.7rem;
	}
	.atlas .imglist li img{
		width:100%;
		height:100%;
	}
	.profile{
		padding:0 0.3rem 0.84rem;
		background: #ffffff;
		margin:0.18rem 0 0;
	}
	.profile .imgtitle{
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    height:0.7rem;
	    border-bottom: 1px solid #ebebec;
	    position: relative;
	}
	.profile .imgtitle .editlogo{
		width:0.32rem;
		height:0.32rem;
		position: absolute;
		right:0;
		bottom:0.2rem;
	}
	.profile .imgtitle .editlogo img{
		width:100%;
		height:100%;
	}
	.profile .imgtitle .left{
		box-flex:1;
		-webkit-box-flex:1;
		-moz-box-flex:1;
		flex:1;
		-webkit-flex:1;
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    -webkit-align-items: center;
    	align-items: center;
    	-webkit-flex-direction: row-reverse;
    	flex-direction: row-reverse;
	}
	.profile .imgtitle .left .line{
		display: inline-block;
		width:0.5rem;
		height: 2px;
		background: #d5d5d7;
	}
	.profile .imgtitle .right{
		box-flex:1;
		-webkit-box-flex:1;
		-moz-box-flex:1;
		flex:1;
		-webkit-flex:1;
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    -webkit-align-items: center;
    	align-items: center;
    	-webkit-flex-direction: row;
    	flex-direction: row;
	}
	.profile .imgtitle .right .line{
		display: inline-block;
		width:0.5rem;
		height: 2px;
		background: #d5d5d7;
	}
	.profile .imgtitle .center{
		width:2rem;
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    -webkit-align-items: center;
    	align-items: center;
    	-webkit-justify-content: center;
    	justify-content: center;
	}
	.profile .imgtitle .center span{
		font-size: 0.25rem;
		color:#000;
	}
	.profile .imgtitle .center img{
		width:0.3rem;
		height:0.4rem;
		margin-right: 0.15rem;
	}
	.profile .describe{
		padding:0.3rem;
		font-size: 0.24rem;
		color:#000;
	}
	.footerbtn{
		position: fixed;
		bottom: 0;
		left:0;
		width:100%;
		height:0.84rem;
		background: #fff;
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    -webkit-align-items: center;
    	align-items: center;
    	-webkit-justify-content: center;
    	justify-content: center;
	}
	.footerbtn .collect{
		box-flex:1;
		-webkit-box-flex:1;
		-moz-box-flex:1;
		flex:1;
		-webkit-flex:1;
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    -webkit-align-items: center;
    	align-items: center;
    	-webkit-justify-content: center;
    	justify-content: center;
		border-right: 1px solid #ccc;
	}
	.footerbtn .collect img{
		width:0.4rem;
		height:0.37rem;
	}
	.footerbtn button{
		box-flex:5;
		-webkit-box-flex:5;
		-moz-box-flex:5;
		flex:5;
		-webkit-flex:5;
		height:100%;
		background: #fff;
	}
	.bookuserstab{
		padding:0.3rem;
		background: #fff;
		margin-bottom: 1rem;
	}
	.bookuserstab .tab{
		margin:0 auto 0.3rem;
		width:3.6rem;
		height:0.6rem;
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    -webkit-align-items: center;
    	align-items: center;
    	-webkit-justify-content: center;
    	justify-content: center;
	}
	.bookuserstab .tab .mybook{
		height:100%;
		box-flex:1;
		-webkit-box-flex:1;
		-moz-box-flex:1;
		flex:1;
		-webkit-flex:1;
		line-height: 0.6rem;
		text-align: center;
		font-size:0.26rem;
		color:#000;
		border: 1px solid #e5e5e6;
		border-top-left-radius:10px;
		border-bottom-left-radius: 10px;
	}
	.bookuserstab .tab .bookuser{
		height:100%;
		box-flex:1;
		-webkit-box-flex:1;
		-moz-box-flex:1;
		flex:1;
		-webkit-flex:1;
		text-align: center;
		font-size:0.26rem;
		color:#000;
		line-height: 0.6rem;
		border: 1px solid #e5e5e6;
		border-top-right-radius:10px;
		border-bottom-right-radius: 10px;
	}
	.bookuserstab .userlist li{
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    -webkit-align-items: center;
    	align-items: center;
    	border:1px solid #e7e7e8;
    	margin-bottom: 0.2rem;
    	padding:0.3rem;
	}
	.bookuserstab .userlist li .headerimg{
		width:0.85rem;
		height:0.85rem;
		border-radius: 100%;
	}
	.bookuserstab .userlist li .headerimg img{
		width:100%;
		height:100%;
	}
	.bookuserstab .userlist li .nickname{
		box-flex:1;
		-webkit-box-flex:1;
		-moz-box-flex:1;
		flex:1;
		-webkit-flex:1;
		padding-left: 0.15rem;
	}
	.bookuserstab .userlist li .nickname h3{
		font-size: 0.28rem;
		margin-bottom: 0.1rem;
	}
	.bookuserstab .userlist li .nickname p{
		font-size: 0.24rem;
	}
	.bookuserstab .userlist li .changemsg{
		width:1.8rem;
		font-size: 0.24rem;
	}
	.bookuserstab .userlist li .change{
		display: block;
		font-size: 0.24rem;
		margin:0.15rem auto 0;
		width:1.56rem;
		height:0.34rem;
		background: #fd9d5c;
		border-radius: 5px;
		color:#fff;
	}
	.bookuserstab .userlist li .msg{
		font-size: 0.24rem;
	}
	.bookuserstab .userlist li .msg span{
		font-size: 0.24rem;
	}
</style>