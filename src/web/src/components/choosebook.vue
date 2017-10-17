<template>
	<div class="choosebook">
		<div class="personalinformation">
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
							<dt @click="isChack(item.bookqid)">
								<img src="../assets/image/bookstore/text.jpg" height="100%" width="100%">
								<div class="selected">
									<img v-if="item.bookqid === bookqid?true:false" src="../assets/image/bookinfo/check.png" height="100%" width="100%">
								</div>
							</dt>
							<dd>
								{{item.bookname}}
							</dd>
						</div>
					</li>
				</ul>
			</mt-loadmore>
		</div>
		<div class="button">
			<mt-button size="normal" type="danger" @click="cancel()">没有想要的</mt-button>
			<mt-button size="normal" type="primary" @click="confirm()">确认换书</mt-button>
		</div>
	</div>
</template>
<script>
	import { Indicator } from 'mint-ui';
	import { Loadmore } from 'mint-ui';
	import { Toast } from 'mint-ui';
	export default{
		name: "choosebook",
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
				isShow:false,
				booklist: [
// 　　　　　　　　　　{bookname: '世界很大，幸好有你世界很大，幸好有你', imgisShow: false},
// 　　　　　　　　　　{bookname: '狼图腾', imgisShow: false},
// 　　　　　　　　　　{bookname: '解忧杂货铺', imgisShow: false},
// 　　　　　　　　　　{bookname: '逻辑思维', imgisShow: false},
// 					{bookname: '随遇而安', imgisShow: false}
　　　　　　　　],
				bookqid:'',
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
		    },
		    isChack (bookqid) {
				this.bookqid = bookqid;
			},
			confirm () {
				console.log(this.bookqid)
				if(this.bookqid == ''){
					Toast({
					  message: '未选中书籍',
					  position: 'bottom',
					  duration: 2000
					});
				}else{
					this.$router.push({ name: 'getcontact', params:{ bookqid: this.bookqid}})
				}
			},
			cancel () {
				this.$router.go(-1);
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
	.choosebook{
		position: relative;
		height:100%;
	}
	.isMargin{
		margin-right: 0!important;
	}
	.textcenter{
		text-align: center!important;
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
		position: relative;
	}
	.booklist li dt .selected{
		position: absolute;
		right:0.1rem;
		bottom:0.1rem;
		width:0.4rem;
		height:0.4rem;
		border-radius: 100%;
		border:1px solid #898a8c;
		z-index: 9999;
	}
	.booklist li dd{
		color:#000000;
		font-size: 0.22rem;
		text-align: left;
		padding:0.17rem 0;
	}
	.button{
		position: fixed;
		left:0;
		bottom:0;
		width:100%;
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	}
	.button .mint-button{
		width:50%;
		border-radius: 0;
	}
</style>