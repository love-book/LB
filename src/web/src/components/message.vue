<template>
	<div class="message">
		<ul class="messagelist">
			<li v-for="(item,index) in newslist" v-show="item.order_type ==2 && item.order_state == 0?false:true">
				<div v-if="item.order_type ==2 && item.order_state == 1?true:false">
					<div class="title">
						<div class="accept">
							{{item.order_desc}}
						</div>	
						<div class="date">
							{{changetime(parseInt(item.update_time*1000))}}
						</div>
					</div>
					<div class="usermsg">
						<div class="userdetails">
							<dt>
								<img :src="item.user_from.imgurl" height="100%" width="100%">
							</dt>
							<dd>
								{{item.user_from.nickname}}
							</dd>
							<p>
								{{item.user_from.signature}}
							</p>
						</div>
						<div class="changeimg">
							<img src="../assets/image/exchange/change.png">
						</div>
						<div class="userdetails">
							<dt>
								<img :src="item.user_to.imgurl" height="100%" width="100%">
							</dt>
							<dd>
								{{item.user_to.nickname}}
							</dd>
							<p>
								{{item.user_to.signature}}
							</p>
						</div>
					</div>
					<div class="operation">	
						<mt-button type="default" size="small" @click="tocontactway(item.user_to.userid)">查看联系方式</mt-button>
					</div>
				</div>
				<div v-else-if="item.order_type ==2 && item.order_state == 2?true:false">
					<div class="title">
						<div class="headerimg">
							<img :src="item.user_from.imgurl" >
						</div>
						<div class="nickname">
							<h3>
								{{item.user_from.nickname}}
							</h3>
							<p>
								{{item.user_from.signature}}
							</p>
						</div>		
						<div class="date">
							{{changetime(parseInt(item.update_time*1000))}}
						</div>
					</div>
					<div class="book">
						<p class="prompting">
							{{item.order_desc}}
						</p>
					</div>
					<div class="operation">	
						<mt-button type="default" size="small" @click="tobookstore()">看看其他</mt-button>
					</div>
				</div>
				<div v-else="item.order_type===1">
					<div class="title">
						<div class="headerimg">
							<img :src="item.user_to.imgurl" >
						</div>
						<div class="nickname">
							<h3>
								{{item.user_to.nickname}}
							</h3>
							<p>
								{{item.user_to.signature}}
							</p>
						</div>		
						<div class="date">
							{{changetime(parseInt(item.update_time*1000))}}
						</div>
					</div>
					<div class="book">
						<p class="prompting">
							{{item.order_desc}}
						</p>
						<div class="bookdetails">
							<dt>
								<img :src="item.book_from.imageurl" height="100%" width="100%">
							</dt>
							<dd>
								{{item.book_from.bookname}}
							</dd>
						</div>
					</div>
					<div class="operation">
						<mt-button v-if="item.order_state==0" type="default" size="small"  @click="reject(item.newid)">拒绝</mt-button>
						<mt-button v-if="item.order_state==0" type="default" size="small" @click="choosebook(item.user_to.userid)">去看看</mt-button>
						<mt-button v-if="item.order_state==1" type="default" disabled size="small">已同意</mt-button>
						<mt-button v-if="item.order_state==2" type="default" disabled size="small">已拒绝</mt-button>
					</div>
				</div>
			</li>
		</ul>
	</div>
</template>
<script>
	import { Toast } from 'mint-ui';
	export default{
		name:'message',
		data() {
			return{
				bottomStatus:'',
				pageTotal:0,
				newslist:[],
				length:4,
				draw:1,
				allLoaded: false, //是否可以上拉属性，false可以上拉，true为禁止上拉，就是不让往上划加载数据了
			    scrollMode:"auto", //移动端弹性滚动效果，touch为弹性滚动，auto是非弹性滚动
			    loading:false,
			}
		},
		mounted(){
		    this.getnewslist();  //初次访问查询列表
		    this.scroll();
		},
		methods:{
			handleTopChange(status) {
		        this.bottomStatus = status;
		    },
		    loadBottom() {
		    	setTimeout(function () {
                    this.more();
                    this.$refs.loadmore.onBottomLoaded();
                }.bind(this), 2000);	
			},
			getnewslist() {
		      	this.$axios({
		      		method: 'post',
			    	url:'/booknews/newsList',	
			    	data:{
			    		length: this.length,
			            draw: this.draw,
			    	}
				})
				.then(function(res){
					console.log(res);
					console.log(res.data.data.data);
					this.pageTotal = res.data.data.pageTotal;
					this.newslist = this.newslist.concat(res.data.data.data);
					console.log(this.newslist)
					// if(this.newslist.length>9){
					// 	this.show = true;
					// 	$(".mint-loadmore-bottom").css('display','block');
					// 	this.allLoaded = false;
					// }else{
					// 	$(".mint-loadmore-bottom").css('display','none');
					// 	this.allLoaded = true;
					// }
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
		        this.getnewslist();
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
		    changetime (date) {
				var oDate = new Date(date);
				function p(s) {
			        return s < 10 ? '0' + s: s;
			    }
				return oDate.getFullYear()+'-'+(oDate.getMonth()+1)+'-'+oDate.getDate();
			},
			tobookstore () {
				this.$router.push({ path: '/tabber/navbar/bookstore'})
			},
			tocontactway (userid) {
				this.$router.push({ name: 'contactway',params:{userid:userid}})
			},
			choosebook(userid) {
				this.$router.push({ name: 'choosebook',params:{userid:userid}})
			},
			reject (newid) {
				this.$axios({
		      		method: 'post',
			    	url:'/booknews/refuselibraryrequest',	
			    	data:{
			    		newid: newid,
			    	}
				})
				.then(function(res){
					console.log(res);
					this.newslist = [];
					this.getnewslist();
					Toast({
					  message: '已拒绝',
					  position: 'middle',
					  duration: 2000
					});
				}.bind(this))
				.catch(function (error) {
					console.log(error);
				});
			}
		}
	}
</script>
<style scoped>
	.message{
		height:100%;
		width:100%;
		background: #ececec;
	}
	.messagelist{
		background: #ececec;
	}
	.messagelist li{
		background: #fff;
		padding:0.3rem;
		margin-bottom: 0.08rem;
	}
	.messagelist li .title{
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    -webkit-align-items: center;
    	align-items: center;
    	padding:0 0 0.3rem;
	}
	.messagelist li .title .headerimg{
		width:0.7rem;
		height:0.7rem;
		border-radius:100%;
		overflow: hidden;
	}
	.messagelist li .title .headerimg img{
		width:100%;
		height:100%;
	}
	.messagelist li .title .nickname{
		box-flex:1;
		-webkit-box-flex:1;
		-moz-box-flex:1;
		flex:1;
		-webkit-flex:1;
		padding:0 0 0 0.3rem;
	}
	.messagelist li .title .nickname h3{
		font-size:0.24rem;
		margin-bottom: 0.1rem;
	}
	.messagelist li .title .nickname p{
		font-size:0.24rem;
	}
	.messagelist li .title .accept{
		box-flex:1;
		-webkit-box-flex:1;
		-moz-box-flex:1;
		flex:1;
		-webkit-flex:1;
		text-align: center;
		font-size: 0.28rem;
	}
	.messagelist li .title .date{
		width:1.3rem;
		font-size:0.24rem;
		color:#979797;
	}
	.messagelist li .book{
		border:1px solid #f0f0f1;
		padding:0.2rem;
	}
	.messagelist li .book .prompting{
		font-size:0.24rem;
		color:#000;
	}
	.messagelist li .book .bookdetails{
		width:1.6rem;
		margin-top: 0.25rem;
	}
	.messagelist li .book .bookdetails dt img{
		width:100%;
	}
	.messagelist li .book .bookdetails dd{
		font-size: 0.24rem;
		color:#000;
	}
	.messagelist li .operation{
		padding:0.3rem 0 0;
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    -webkit-flex-direction: row-reverse;
    	flex-direction: row-reverse;
    	-webkit-align-items: center;
    	align-items: center;
	}
	.messagelist li .operation .hints{
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    -webkit-align-items: center;
    	align-items: center;
	    font-size:0.24rem;
	}
	.messagelist li .operation .hints img{
		width:0.4rem;
		height:0.4rem;
	}
	.mint-button--small{
		margin-left:0.2rem;
	}
	.messagelist li .usermsg{
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    -webkit-align-items: center;
    	align-items: center;
    	-webkit-justify-content: center;
    	justify-content: center;
    	padding:0.2rem;
    	border: 1px solid #f0f0f1;
	}
	.messagelist li .usermsg .userdetails{
		width:1.7rem;
	}
	.messagelist li .usermsg .userdetails dt{
		width:0.75rem;
		height:0.75rem;
		margin:0 auto;
	}
	.messagelist li .usermsg .userdetails dd{
		text-align: center;
		font-size:0.26rem;
		margin:0.2rem 0;
	}
	.messagelist li .usermsg .userdetails p{
		font-size: 0.24rem;
	}
	.messagelist li .usermsg .changeimg{
		width:0.5rem;
		height:0.5rem;
		margin:0 0.2rem;
	}
	.messagelist li .usermsg .changeimg img{
		width:100%;
		height:100%;
	}
</style>