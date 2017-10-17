<template>
	<div class="toexchange">
		<mt-loadmore :bottom-method="loadBottom" @bottom-status-change="handleTopChange" :bottom-all-loaded="allLoaded" :auto-fill="false" ref="loadmore">
			<ul class="toexchangelist">
				<li v-for="(item,index) in awaitlist">
					<div class="title">
						<div class="nickname">
							{{item.user_from.nickname}} & {{item.user_to.nickname}}
						</div>
						<div class="state">
							待交换
						</div>	
					</div>
					<div class="book">
						<div class="bookdetails">
							<dt>
								<img :src="item.book_from.imageurl" height="100%" width="100%">
							</dt>
							<dd>
								{{item.book_from.bookname}}
							</dd>
						</div>
						<div class="changeimg">
							<img src="../../assets/image/exchange/change.png">
						</div>
						<div class="bookdetails">
							<dt>
								<img :src="item.book_to.imageurl" height="100%" width="100%">
							</dt>
							<dd>
								{{item.book_to.bookname}}
							</dd>
						</div>
					</div>
					<div class="operation">
						<mt-button type="default" size="small" @click="tocontactway(item.user_to.userid)">联系方式</mt-button>
						<mt-button type="default" size="small" @click="cancel(item.orderid)">取消交换</mt-button>
						<mt-button type="default" size="small" @click="confirm(item.orderid)">确认交换</mt-button>
					</div>
				</li>
			</ul>
		</mt-loadmore>
	</div>
</template>
<script>
	export default{
		name:'toexchange',
		data() {  
		    return {  
		      	selected: '1',
		      	bottomStatus:'',
	    		pageTotal:0,
			    draw:1,
	    		length:10,
			    allLoaded: false, //是否可以上拉属性，false可以上拉，true为禁止上拉，就是不让往上划加载数据了
			    scrollMode:"auto", //移动端弹性滚动效果，touch为弹性滚动，auto是非弹性滚动
			    loading:false,
		      	awaitlist:[],  
		    };  
		},
		mounted(){
		    this.getawaitlist();
		    this.scroll();
		},
		methods:{
			getawaitlist () {
				this.$axios({
		      		method: 'post',
			    	url:'/bookorder/orderList',	
			    	data:{
			    		draw: this.draw,
			    		length:this.length,
			    		order_state:'0',
			    	}
				})
				.then(function(res){
					console.log(res)
					this.pageTotal = res.data.data.pageTotal;
					this.awaitlist = this.awaitlist.concat(res.data.data.data);
					// if(this.booklist.length>12){
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
	  		handleTopChange(status) {
		        this.bottomStatus = status;
		    },
		    loadBottom() {
		    	setTimeout(function () {
                    this.more();
                    this.$refs.loadmore.onBottomLoaded();
                }.bind(this), 2000);	
			},
			more:function (){
		      // 分页查询
		      if(this.pageTotal == 1){
		        this.draw = 1;
		        this.allLoaded = true;
		      }else{
		        this.draw = parseInt(this.draw) + 1;
		        this.allLoaded = false;
		        this.getawaitlist();
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
			tocontactway (userid) {
				this.$router.push({ name: 'contactway',params:{userid:userid}})
			},
			confirm (orderid) {
				this.$axios({
		      		method: 'post',
			    	url:'/bookorder/orderupdate',	
			    	data:{
			    		orderid:orderid,
			    		order_state:'1',
			    	}
				})
				.then(function(res){
					console.log(res.data.data.data)
					this.awaitlist = [];
					this.getawaitlist();
				}.bind(this))
				.catch(function (error) {
					console.log(error);
				});
			},
			cancel (orderid) {
				this.$axios({
		      		method: 'post',
			    	url:'/bookorder/orderupdate',	
			    	data:{
			    		orderid:orderid,
			    		order_state:'2',
			    	}
				})
				.then(function(res){
					console.log(res.data.data.data)
					this.awaitlist = [];
					this.getawaitlist();
				}.bind(this))
				.catch(function (error) {
					console.log(error);
				});
			},
		}
	}
</script>
<style scoped>
	.toexchangelist{
		overflow: auto;
	}
	.toexchangelist li{
		padding:0 0.25rem;
		background: #fff;
		margin-bottom:0.05rem;
	}
	.toexchangelist li .title{
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    -webkit-justify-content: space-between;
    	justify-content: space-between;
    	height:0.75rem;
    	line-height: 0.75rem;
	}
	.toexchangelist li .title .nickname{
		font-size:0.26rem;
		color:#424242;
	}
		.toexchangelist li .title .state{
		font-size:0.18rem;
		color:#888;
	}
	.book{
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    -webkit-align-items: center;
    	align-items: center;
	    border:1px solid #e3e6eb;
	    padding:0.25rem 0.72rem;
	}
	.book .bookdetails{
		width:1.5rem;
	}
	.book .bookdetails dt{
		width:100%;
		height:2rem;
	}
	.book .bookdetails dt img{
		width:100%;
		height:100%;
	}
	.book .bookdetails dd{
		color:#000000;
		font-size: 0.22rem;
		text-align: center;
		margin:0.17rem 0;
		height:0.6rem;
		display: -webkit-box;
		-webkit-box-orient: vertical;
		-webkit-line-clamp: 2;
		overflow: hidden;
	}
	.book .changeimg{
		width:0.5rem;
		height:0.5rem;
		margin:0 0.4rem;
	}
	.book .changeimg img{
		width:100%;
		height:100%;
	}
	.operation{
		padding:0.25rem 0;
		text-align: right;
	}
</style>