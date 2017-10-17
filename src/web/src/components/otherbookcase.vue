<template>
	<div class="otherbookcase">
		<div class="personalinformation">
			<div class="collect">
				<img v-if="!collect" @click="isCollect()" src="../assets/image/bookinfo/nocollect.png" >
				<img v-else @click="isCollect()" src="../assets/image/bookinfo/collect.png" >
			</div>
			<div class="headimg">
				<img :src="personmsg.imgurl" height="88" width="88">
			</div>
			<h3>
				{{personmsg.nickname}}
			</h3>
			<ul class="informationlist">
				<li class="sex">
					<img v-if="personmsg.gender === 1" src="../assets/image/bookcase/b_man.png" height="30" width="30">
					<img v-else src="../assets/image/bookcase/b_woman.png" height="30" width="30">
				</li>
				<li class="age">
					{{personmsg.age}}
				</li>
				<li class="address">
					{{personmsg.address}}
				</li>
			</ul>
			<p>
				{{personmsg.signature == '' ? '该用户还未添加个性签名' : personmsg.signature}}
			</p>
			<div class="cf"></div>
		</div>
		<div class="main_warp">
			<div class="main_title">
				<div class="titel_name">
					<span class="line"></span>
					<p>
						个人书架
					</p>
				</div>
			</div>
			<mt-loadmore :bottom-method="loadBottom" @bottom-status-change="handleTopChange" :bottom-all-loaded="allLoaded" :auto-fill="false" ref="loadmore">
				<ul class="booklist">
					<li v-for='(item,index) in booklist' :class="{'isMargin': isMargin(index)}">
						<div>
							<router-link :to="{ name: 'bookinfo', params:{ bookqid: item.bookqid }}">
								<dt>
									<img :src="item.imageurl" height="100%" width="100%">
								</dt>
								<dd>
									{{item.bookname}}
								</dd>
							</router-link>
						</div>
					</li>
				</ul>
			</mt-loadmore>
		</div>
		<div class="gotop" @click="gotop()">
			<img src="../assets/image/bookcase/gotop.png">
		</div>
	</div>
</template>
<script>
	import { Indicator } from 'mint-ui';
	import { Loadmore } from 'mint-ui';
	import { Toast } from 'mint-ui';
	export default{
		name: "otherbookcase",
		data() {
			return {
				bottomStatus:'',
	    		pageTotal:0,
			    allLoaded: false, //是否可以上拉属性，false可以上拉，true为禁止上拉，就是不让往上划加载数据了
			    scrollMode:"auto", //移动端弹性滚动效果，touch为弹性滚动，auto是非弹性滚动
			    loading:false,
				draw:1,
	    		length:10,
				collect:false,
				personmsg:{},
				booklist:[],
				deletebookcase:[],
			}
		},
		methods: {
			isMargin(index){
				if((index+1)%3==0){
					return true;
				}else{
					return false;
				}
			},
			gotop () {
				scrollTo(0,0);
			},
			isCollect () {
				if(this.collect){
					this.collect = false;
					this.deletebookcase.push(this.personmsg.concernid)
					//console.log(this.deletebookcase)
					this.$axios({
			      		method: 'post',
				    	url:'/book/delbookconcern',	
				    	data:{
				    		concernid: this.deletebookcase,
				    	}
					})
					.then(function(res){
						//console.log(res)
						Toast({
						  message: '取消此收藏书架',
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
				    		concern_type: '2',
				    		userid_from: this.$route.params.userid,
				    	}
					})
					.then(function(res){
						//console.log(res.data.data)
						this.personmsg.concernid = res.data.data.concernid;
						//console.log(this.personmsg.concernid)
						Toast({
						  message: '收藏此书架成功',
						  position: 'bottom',
						  duration: 2000
						});
					}.bind(this))
					.catch(function (error) {
						console.log(error);
					});
				}
			},
			getpersonmsg () {
				this.$axios({
		      		method: 'post',
			    	url:'/users/userinfo',	
			    	data:{
			    		userid: this.$route.params.userid,
			    	}
				})
				.then(function(res){
					//console.log(res)
					this.personmsg = res.data.data;
					//console.log(this.personmsg)
					if(this.personmsg.concernid != ''){
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
			getbooklist() {
		      	this.$axios({
		      		method: 'post',
			    	url:'/bookrack/mybookrack',	
			    	data:{
			    		draw:this.draw,
	    				length:this.length,
			    		userid: this.$route.params.userid,
			    	}
				})
				.then(function(res){
					console.log(res);
					this.pageTotal = res.data.data.pageTotal;
					this.booklist = this.booklist.concat(res.data.data.data);
					if(this.booklist.length>9){
						this.show = true;
						$(".mint-loadmore-bottom").css('display','block');
						this.allLoaded = false;
					}else{
						$(".mint-loadmore-bottom").css('display','none');
						this.allLoaded = true;
					}
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
		        this.getbooklist();
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
		    }
		},
		mounted(){
		    this.getpersonmsg();
		    this.getbooklist();
		    this.scroll();
		},
	}
	
</script>
<style scoped>
	.isMargin{
		margin-right: 0!important;
	}
	.textcenter{
		text-align: center!important;
	}
	.gotop{
		width:0.6rem;
		height: 0.6rem;
		position: fixed;
		bottom: 1.2rem;
		right:0.5rem;
	}
	.gotop img{
		width:100%;
	}
	.personalinformation{
		position: relative;
		height:auto;
		background: url(../assets/image/bookstore/text.jpg) no-repeat;
		background-size: 100% 100%;
		color:#fff;
		overflow: hidden;
	}
	.personalinformation .collect{
		position: absolute;
		left:0.3rem;
		top:0.3rem;
		width:0.5rem;
		height:0.5rem;
	}
	.personalinformation .collect img{
		width:100%;
	}
	.personalinformation .headimg{
		width:1.45rem;
		height:1.45rem;
		margin:0.3rem auto;
		border-radius:100%;
		overflow: hidden;
	}
	.personalinformation .headimg img{
		width:100%;
		height: 100%;
	}
	.personalinformation h3,.personalinformation p{
		text-align: center;
		margin-bottom: 0.12rem;
		font-size: 0.25rem;
	}
	.personalinformation .informationlist{
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    width:2.2rem;
	    margin:0 auto 0.12rem;
	}
	.personalinformation .informationlist li{
		margin-right: 0.15rem;
		font-size: 0.25rem;
	}
	.personalinformation .informationlist .sex{
		text-align: center;
	}
	.personalinformation .informationlist .sex img{
		width:0.24rem;
		height:0.24rem;
		display: inline-block;
		vertical-align: middle;
	}
	.main_warp{
		padding:0 0.4rem;
	}
	.main_title{
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    -webkit-justify-content: space-between;
    	justify-content: space-between;
	    height: 0.3rem;
	    padding:0.18rem 0;
	}
	.main_title .titel_name{
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	}
	.main_title .titel_name .line{
		height: 100%;
		width:0.04rem;
		background: #35b7f2;
	}
	.main_title p{
		height:100%;
		line-height:0.3rem;
		font-size:0.25rem;
		margin-left: 0.08rem
	}
	.booklist{
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    -webkit-flex-wrap: wrap;
    	flex-wrap: wrap;
	    -webkit-justify-content: flex-start;
    	justify-content: flex-start;
	}
	.booklist li{
		margin-top: 0.17rem;
		width:1.5rem;
		margin-right: 0.51rem;
	}
	.booklist li dt{
		width:100%;
		height:2rem;
	}
	.booklist li .add{
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    -webkit-justify-content: center;
    	justify-content: center;
    	-webkit-align-items: center;
    	align-items: center;
		border:1px solid #CCC;
		-moz-box-shadow:0 0 30px rgba(204,204,204,0.5);
	    -webkit-box-shadow:0 0 30px rgba(204,204,204,0.5);
	    box-shadow:0 0 30px rgba(204,204,204,0.5);
	}
	.booklist li dd{
		color:#000000;
		font-size: 0.22rem;
		text-align: left;
		padding:0.17rem 0;
	}
	.booklist li .tips{
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	}
	.booklist li .tips .sex{
		width: 35%;
		height:0.3rem;
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    -webkit-align-items: center;
    	align-items: center;
		border:1px solid #c0dcf9;
		border-radius: 4px;
		padding:0.05rem;
		margin-right: 0.1rem;
	}
	.booklist li .tips .sex img{
		width:0.2rem;
		height:0.2rem;
		margin-right: 0.06rem;
		vertical-align:middle
	}
	.booklist li .tips .sex .age{
		line-height: 0.3rem;
		height:100%;
		font-size: 0.15rem;
	}
	.booklist li .tips .distance{
		width: 55%;
		height:0.3rem;
		line-height: 0.3rem;
		border:1px solid #b4b1e4;
		border-radius: 4px;
		font-size: 0.15rem;
		padding:0.05rem;
	}
</style>