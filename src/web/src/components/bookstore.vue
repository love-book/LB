<template>
	<div class="bookstore">
		<mt-loadmore :bottom-method="loadBottom" @bottom-status-change="handleTopChange" :bottom-all-loaded="allLoaded" :auto-fill="false" ref="loadmore">
			<ul class="booklist">
				<li v-for='(item,index) in booklist'>
					<router-link class="booklink" :to="'/book/'+item.bookid">
						<dt>
							<img :src="item.imageurl" height="100%" width="100%">
						</dt>
						<dd>
							<div class="bookmsg">
								<span class="bookname">
									{{item.bookname}}
								</span>
								<span class="author">
									{{item.author}}
								</span>
							</div>
							<p class="describe">
								{{item.describe}}
							</p>
						</dd>
						<!-- <div class="tips">
							<div class="sex">
								<img  v-if="item.gender == 2" src="../assets/image/bookstore/woman.png">
								<img  v-else src="../assets/image/bookstore/man.png">
								<span class="age">{{item.age}}</span>
							</div>
							<div class="distance">
								{{item.radius}}
							</div>
						</div> -->
					</router-link>			
				</li>
			</ul>
		</mt-loadmore>
		<div class="gotop" @click="gotop()">
			<img src="../assets/image/bookcase/gotop.png">
		</div>
	</div>
</template>
<script>
	import { Indicator } from 'mint-ui';
	import { Loadmore } from 'mint-ui';
	export default{
		components:{
	        'mt-loadmore':Loadmore,
	    },
		name: 'bookstore',
	  	data() {  
	    	return {
	    		bottomStatus:'',
	    		pageTotal:0,
			    draw:1,
	    		length:10,
			    allLoaded: false, //是否可以上拉属性，false可以上拉，true为禁止上拉，就是不让往上划加载数据了
			    scrollMode:"auto", //移动端弹性滚动效果，touch为弹性滚动，auto是非弹性滚动
			    loading:false,
				booklist:[
					// {bookid: 1,imageurl: '../assets/image/bookstore/text.jpg', bookname: '世界很大，幸好有你世界很大，幸好有你',gender:'1',age: '25',radius: '100M'},
					// {bookid: 1,imageurl: '../assets/image/bookstore/text.jpg', bookname: '世界很大，幸好有你世界很大，幸好有你',gender:'2',age: '25',radius: '100M'},
					// {bookid: 1,imageurl: '../assets/image/bookstore/text.jpg', bookname: '世界很大，幸好有你世界很大，幸好有你',gender:'1',age: '25',radius: '100M'},
					// {bookid: 1,imageurl: '../assets/image/bookstore/text.jpg', bookname: '世界很大，幸好有你世界很大，幸好有你',gender:'1',age: '25',radius: '100M'},
					// {bookid: 1,imageurl: '../assets/image/bookstore/text.jpg', bookname: '世界很大，幸好有你世界很大，幸好有你',gender:'2',age: '25',radius: '100M'},
					// {bookid: 1,imageurl: '../assets/image/bookstore/text.jpg', bookname: '世界很大，幸好有你世界很大，幸好有你',gender:'2',age: '25',radius: '100M'},
					// {bookid: 1,imageurl: '../assets/image/bookstore/text.jpg', bookname: '世界很大，幸好有你世界很大，幸好有你',gender:'1',age: '25',radius: '100M'},
					// {bookid: 1,imageurl: '../assets/image/bookstore/text.jpg', bookname: '世界很大，幸好有你世界很大，幸好有你',gender:'2',age: '25',radius: '100M'},
					// {bookid: 1,imageurl: '../assets/image/bookstore/text.jpg', bookname: '世界很大，幸好有你世界很大，幸好有你',gender:'1',age: '25',radius: '100M'},
					// {bookid: 1,imageurl: '../assets/image/bookstore/text.jpg', bookname: '世界很大，幸好有你世界很大，幸好有你',gender:'2',age: '25',radius: '100M'},
				]
	    	};  
	  	}, 
	  	mounted(){
		    this.loadPageList();  //初次访问查询列表
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
			loadPageList() {
		      	this.$axios({
		      		method: 'post',
			    	url:'/book/booklist',	
			    	data:{
			    		length: this.length,
			            draw: this.draw,
			    	}
				})
				.then(function(res){
					console.log(res);
					this.pageTotal = res.data.data.pageTotal;
					this.booklist = this.booklist.concat(res.data.data.data);
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
			more:function (){
		      // 分页查询
		      if(this.pageTotal == 1){
		        this.draw = 1;
		        this.allLoaded = true;
		      }else{
		        this.draw = parseInt(this.draw) + 1;
		        this.allLoaded = false;
		        this.loadPageList();
		      }
		    },
		    isHaveMore:function(){
		      // 是否还有下一页，如果没有就禁止上拉刷新
		      //this.allLoaded = false; //true是禁止上拉加载
		      if(this.draw == this.pageTotal){
		        this.allLoaded = true;
		      }
		    },
	  		gotop () {
				document.getElementById('page-wrap').scrollTop = 0;
			},
			scroll () {
		    	scrollTo(0,0);
		    }
	  	},
	  	created: function() {
            
        }
	}
</script>
<style scoped>
	.mint-loadmore{
		padding-bottom: 1.2rem;
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
	.active{
		cursor: default;
	    background-color: #fff;
	    border: 1px solid #ddd;
	    border-bottom-color: transparent;
	}
	.filtratelist{
		position: fixed;
		top:0;
		left:0;
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    width:86.8%;
	    height: 0.8rem;
	    border-bottom: 1px solid #e6e6e6;
	    padding:0.08rem 0.42rem 0;
	    background: rgb(250, 250, 250);
	    z-index: 9999;
	}
	.filtratelist .filtrateitem{
		box-flex:1;
		-webkit-box-flex:1;
		-moz-box-flex:1;
		flex:1;
		-webkit-flex:1;
	    margin-right: 2px;
	    border-radius: 4px 4px 0 0;
		margin-bottom: -1px;
		z-index: 9999999;
	}
	.screentit{
		position: relative;
		height: 100%;
		line-height: 100%;
	}
	.screentit p{
		height: 100%;
		line-height: 0.8rem;
		text-align:center;
	}
	.screentit p img{
		margin-left: 0.2rem;
		display: inline-block;
		vertical-align:middle;
	}
	.conditions{
		z-index: 999999;
		position: absolute;
		top:0.9rem;
		left:0;
		display: none;
		width:6.4rem;
		height: auto;
	}
	.filteroptionslist{
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    -webkit-flex-wrap: wrap;
    	flex-wrap: wrap;
	    height: auto;
	    padding:0 0.28rem;
	    background: rgb(250, 250, 250);
	}
	.filteroptionslist li{
		width: 50%;
	}
	.filteroptionitem{
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    padding:0.2rem 0;
	}
	.filteroptionitem div{
		box-flex:7;
		-webkit-box-flex:7;
		-moz-box-flex:7;
		flex:7;
		-webkit-flex:7;
		color:#000000;
	}
	.filteroptionitem span{
		box-flex:1;
		-webkit-box-flex:1;
		-moz-box-flex:1;
		flex:1;
		-webkit-flex:1;
	}
	.filteroptionitem span img{
		display: none;
		margin-top:0.1rem;
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
		background: #fff;
	}
	.booklist li{
		width:100%;
		height:1.8rem;
		padding:0.1rem 0.3rem;
		border-bottom: 1px solid #e4e5e9;
	}
	.booklist li .booklink{
		height:100%;
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	}
	.booklist li .booklink dt{
		box-flex:1;
		-webkit-box-flex:1;
		-moz-box-flex:1;
		flex:1;
		-webkit-flex:1;
		height:100%;
	}
	.booklist li dd{
		width:4.2rem;
		padding:0.2rem 0 0.1rem 0.2rem;
	}
	.booklist li dd .bookmsg{
		width:100%;
		height:0.5rem;
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    -webkit-align-items: flex-end;
    	align-items: flex-end;
	    margin-bottom: 0.2rem;
	}
	.booklist li dd .bookmsg span{
		box-flex:1;
		-webkit-box-flex:1;
		-moz-box-flex:1;
		flex:1;
		-webkit-flex:1;
	}
	.booklist li dd .bookmsg .bookname{
		font-size:0.28rem;
		text-overflow:ellipsis;
		overflow:hidden;
		white-space:nowrap;
	}
	.booklist li dd .bookmsg .author{
		font-size:0.22rem;
		text-overflow:ellipsis;
		overflow:hidden;
		white-space:nowrap;
	}
	.booklist li dd .describe{
		overflow : hidden;
		text-overflow: ellipsis;
		display: -webkit-box;
		-webkit-line-clamp: 2;
		-webkit-box-orient: vertical;
		width:100%;
		height:0.7rem;
		font-size:0.18rem;
		color:#b2b2b2;
		line-height:0.34rem;
	}
</style>