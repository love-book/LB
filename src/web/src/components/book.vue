<template>
	<div class="book" id="book">
		<div class="bookinformation">
			<div class="bookimg">
				<img :src="bookmsg.imageurl">
				<span class="imgnum">3张</span>
			</div>
			<ul class="nickname">
				<li>
					{{bookmsg.bookname}}
				</li>
				<li>
					{{bookmsg.auhtor}}
				</li>
				<li>
					{{bookmsg.depreciation}}
				</li>
			</ul>		
			<div class="moreimg">
				<img v-if="!collect" @click="isCollect()" src="../assets/image/bookinfo/nocollect.png" >
				<img v-else @click="isCollect()" src="../assets/image/bookinfo/collect.png" >
			</div>
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
			</div>
			<div class="describe">
				<p>
					{{bookmsg.describe}}
				</p>
				<!-- <a v-if="showreadmore" class="readmore" @click="readmore()">展开</a> -->
			</div>
		</div>
		<div class="bookuserstab">
			<div class="tab">
				拥有此书的人
			</div>
			<mt-loadmore :bottom-method="loadBottom" @bottom-status-change="handleTopChange" :bottom-all-loaded="allLoaded" :auto-fill="false" ref="loadmore">
				<ul class="userlist">
					<li v-for='(item,index) in personlist'>
						<router-link :to="{ name: 'otherbookcase', params:{ userid: item.userid }}">
							<div class="headerimg">
								<img :src="item.imgurl" >
							</div>
						</router-link>
						<div class="nickname">
							<h3>
								{{item.nickname}}
							</h3>
							<p>
								{{item.signature}}
							</p>
						</div>		
						<div class="changemsg">
							<p class="msg">
								{{item.radius}} <span>{{item.logintime}}</span>
							</p>
							<button type="button" class="change" @click="gobookinfo(item.bookqid)">点击交换</button>
						</div>
					</li>
				</ul>
			</mt-loadmore>
		</div>
	</div>
</template>
<script>
	import { Indicator } from 'mint-ui';
	import { Loadmore } from 'mint-ui';
	import { Toast } from 'mint-ui';
	export default{
		components:{
	        'mt-loadmore':Loadmore,
	    },
		name:'book',
		data () {
			return{
				bottomStatus:'',
	    		pageTotal:0,
			    allLoaded: false, //是否可以上拉属性，false可以上拉，true为禁止上拉，就是不让往上划加载数据了
			    scrollMode:"auto", //移动端弹性滚动效果，touch为弹性滚动，auto是非弹性滚动
			    loading:false,
				draw:1,
	    		length:2,
				collect:false,
				bookmsg:{},
				personlist:[],
				deletebook:[],
				showreadmore:false,
			};
		},
		mounted(){
		    this.getbookmsg();
		    this.getuserlist();
		    this.scroll();
		},
		methods:{
			// changelist (e) {
			// 	$(".tab").find('div').css({
			// 		background: '#fff',
			// 		color: '#000'
			// 	});
			// 	$(e.currentTarget).css({
			// 		background: '#13b8f5',
			// 		color: '#fff'
			// 	});
			// },
			isCollect () {
				if(this.collect){
					this.collect = false;
					this.deletebook.push(this.bookmsg.concernid)
					console.log(this.deletebook)
					this.$axios({
			      		method: 'post',
				    	url:'/book/delbookconcern',	
				    	data:{
				    		concernid: this.deletebook,
				    	}
					})
					.then(function(res){
						console.log(res)
						Toast({
						  message: '取消收藏',
						  position: 'bottom',
						  duration: 2000
						});
					}.bind(this))
					.catch(function (error) {
						console.log(error);
					});
				}else{
					this.collect = true;
					this.$axios({
			      		method: 'post',
				    	url:'/book/addconcern',	
				    	data:{
				    		concern_type: '1',
				    		userid_from: this.$route.params.bookid,
				    	}
					})
					.then(function(res){
						console.log(res.data.data)
						this.bookmsg.concernid = res.data.data.concernid;
						console.log(this.bookmsg.concernid)
						Toast({
						  message: '收藏成功',
						  position: 'bottom',
						  duration: 2000
						});
					}.bind(this))
					.catch(function (error) {
						console.log(error);
					});
				}
			},
			getbookmsg () {
				this.$axios({
		      		method: 'post',
			    	url:'/bookrack/getbookinfo',	
			    	data:{
			    		bookid: this.$route.params.bookid,
			    	}
				})
				.then(function(res){
					this.bookmsg = res.data.data;
					console.log(this.bookmsg)
					// if(bookmsg.describe.leength>150){
					// 	this.showreadmore = trun;
					// }
					if(this.bookmsg.concernid != ''){
						this.collect = true;
					}
				}.bind(this))
				.catch(function (error) {
					console.log(error);
				});
			},
			handleTopChange(status) {
		        this.bottomStatus = status;
		    },
		    loadBottom() {
		    	setTimeout(function () {
                    this.more();
                    this.$refs.loadmore.onBottomLoaded();
                }.bind(this), 2000);	
			},
			getuserlist() {
				
		      	this.$axios({
		      		method: 'post',
			    	url:'/bookrack/getbookusers',	
			    	data:{
			    		draw:this.draw,
	    				length:this.length,
			    		bookid: this.$route.params.bookid,
			    	}
				})
				.then(function(res){
					console.log(res);
					this.pageTotal = res.data.data.pageTotal;
					this.personlist = this.personlist.concat(res.data.data.data);
			        if(this.pageTotal == 1){
			          this.allLoaded = true;
			        }
					this.$nextTick(function () {
			          // 是否还有下一页，加个方法判断，没有下一页要禁止上拉
			          this.scrollMode = "touch";
			          this.isHaveMore();
			        });
				}.bind(this))
				.catch(function (error) {
					console.log(error);
				});
			},
			more:function (){
		      // 分页查询
		      if(this.pageTotal == 1){
		        this.draw = 1;
		        this.allLoaded = true;
		      }else{
		        this.draw = parseInt(this.draw) + 1;
		        this.allLoaded = false;
		        this.getuserlist();
		      }
		    },
		    isHaveMore:function(){
		      // 是否还有下一页，如果没有就禁止上拉刷新
		      //this.allLoaded = false; //true是禁止上拉加载
		      if(this.draw == this.pageTotal){
		        this.allLoaded = true;
		      }
		    },
		    scroll () {
		    	scrollTo(0,0);
		    },
		    gobookinfo (bookqid) {
		    	this.$router.push({ name: 'bookinfo', params:{ bookqid: bookqid }})
		    },
		    // isshowreadmore(){

		    // },
		    // readmore () {

		    // },
		},
	}
</script>
<style>
	.book{
		width:100%;
		background: #ececec;
	}
	.bookinformation{
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
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
		width:0.5rem;
	}
	.bookinformation .moreimg img{
		width:100%;
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
		padding:0 0.3rem;
		background: #ffffff;
		margin:0.18rem 0;
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
	.profile .describe p{
		font-size: 0.24rem;
		line-height: 0.35rem;
	}
	.profile .describe .readmore{
		display: block;
		width:1rem;
		height:0.5rem;
		line-height:0.5rem;
		text-align: center;
		margin:0.3rem auto 0;
		border:1px solid #ca0c16;
		border-radius: 4px;
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
		border-right: 1px solid #ededee;
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