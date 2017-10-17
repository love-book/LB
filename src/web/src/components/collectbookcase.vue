<template>
	<div class="collectbookcase">
		<mt-loadmore :bottom-method="loadBottom" @bottom-status-change="handleTopChange" :bottom-all-loaded="allLoaded" :auto-fill="false" ref="loadmore">
			<ul class="bookcaselist">
				<li v-for="(item,index) in bookcaselist">
					<router-link :to="{ name: 'otherbookcase', params:{ userid: item.userid_from}}">
						<div class="title">
							<div class="headerimg">
								<img :src="item.books.imgurl" >
							</div>
							<div class="nickname">
								<h3>
									{{item.books.nickname}}
								</h3>
								<p>
									{{item.books.signature == '' ? '该用户还未添加个性签名' : item.books.signature}}
								</p>
							</div>		
							<div class="moreimg">
								<img src="../assets/image/mine/more.png" >
							</div>
						</div>
					</router-link>
					<div class="book">
						<div class="bookdetails" v-for="book in item.booksList">
							<dt>
								<img :src="book.imageurl" height="100%" width="100%">
							</dt>
							<dd>
								{{book.bookname}}
							</dd>
						</div>
						<!-- <div class="bookdetails">
							<dt>
								<img src="../assets/image/bookstore/text.jpg" height="100%" width="100%">
							</dt>
							<dd>
								世界很大，幸好有你
							</dd>
						</div>
						<div class="bookdetails">
							<dt>
								<img src="../assets/image/bookstore/text.jpg" height="100%" width="100%">
							</dt>
							<dd>
								世界很大，幸好有你
							</dd>
						</div> -->
					</div>
					<div class="operation">
						<mt-button type="default" size="small" @click="cancel(item.concernid)">取消收藏</mt-button>
					</div>
				</li>
			</ul>
		</mt-loadmore>
	</div>
</template>
<script>
	import { Toast } from 'mint-ui';
	export default{
		name:'collectbookcase',
		data() {
			return{
				bottomStatus:'',
	    		pageTotal:0,
			    draw:1,
	    		length:10,
			    allLoaded: false, //是否可以上拉属性，false可以上拉，true为禁止上拉，就是不让往上划加载数据了
			    scrollMode:"auto", //移动端弹性滚动效果，touch为弹性滚动，auto是非弹性滚动
			    loading:false,		
				isShow:false,
				bookcaselist: [],
				concernid:[],
			}
		},
		mounted(){
		    this.getbookcaselist();
		    this.scroll();
		},
		methods:{
			getbookcaselist () {
				this.$axios({
		      		method: 'post',
			    	url:'/book/concernuserlist',	
			    	data:{
			    		draw: this.draw,
			    		length:this.length,
			    	}
				})
				.then(function(res){
					console.log(res)
					this.pageTotal = res.data.data.pageTotal;
					this.bookcaselist = this.bookcaselist.concat(res.data.data.data);
					// if(this.bookcaselist.length>12){
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
		        this.getbookcaselist();
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
			cancel (concernid) {
				this.concernid.push(concernid);
				this.$axios({
		      		method: 'post',
			    	url:'/book/delbookconcern',	
			    	data:{
			    		concernid:this.concernid,
			    	}
				})
				.then(function(res){
					console.log(res)
					this.bookcaselist = [];
					this.getbookcaselist();
					Toast({
					  message: '删除成功',
					  position: 'middle',
					  duration: 2000
					});
				}.bind(this))
				.catch(function (error) {
					console.log(error);
				});
			},
		},
	}
</script>
<style scoped>
	.collectbookcase{
		height:100%;
		width:100%;
		background: #ececec;
	}
	.bookcaselist{
		background: #ececec;
	}
	.bookcaselist li{
		background: #fff;
		padding:0 0.3rem;
	}
	.bookcaselist li .title{
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    padding:0.3rem 0;
	}
	.bookcaselist li .title .headerimg{
		width:0.85rem;
		height:0.85rem;
	}
	.bookcaselist li .title .headerimg img{
		width:100%;
		height:100%;
	}
	.bookcaselist li .title .moreimg{
		position: relative;
		width:0.2rem;
	}
	.bookcaselist li .title .nickname{
		box-flex:1;
		-webkit-box-flex:1;
		-moz-box-flex:1;
		flex:1;
		-webkit-flex:1;
		padding:0 0.3rem;
	}
	.bookcaselist li .title .nickname h3{
		font-size: 0.24rem;
		margin-bottom: 0.1rem;
	}
	.bookcaselist li .title .nickname p{
		font-size: 0.24rem;
	}
	.bookcaselist li .title .moreimg img{
		width:100%;
		position: absolute;
        top: 50%;
        left: 50%;
        margin-top: -10px; 
        margin-left: -5px;
	}
	.bookcaselist li .book{
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    -webkit-justify-content: space-around;
    	justify-content: space-around;
	    padding:0.3rem;
	    border:1px solid #f0f0f1;
	}
	.bookcaselist li .book .bookdetails{
		width:1.5rem;
	}
	.bookcaselist li .book .bookdetails dt{
		height:2rem;
	}
	.bookcaselist li .book .bookdetails dd{
		font-size:0.2rem;
	}
	.bookcaselist li .operation{
		padding:0.3rem 0;
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
</style>