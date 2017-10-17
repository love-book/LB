<template>
	<div class="delbook">
		<mt-loadmore :bottom-method="loadBottom" @bottom-status-change="handleTopChange" :bottom-all-loaded="allLoaded" :auto-fill="false" ref="loadmore">
			<ul class="booklist">
				<li v-for='(item,index) in booklist'>
					<dt @click="isChack($event,item.bookqid)">
						<img :src="item.imageurl" height="100%" width="100%">
						<div class="selected" ref="chack"  datashow="false" >
							<img src="../assets/image/bookinfo/check.png" height="100%" width="100%">
						</div>
					</dt>
					<dd>
						{{item.bookname}}
					</dd>
				</li>
			</ul>
		</mt-loadmore>
		<div class="delbtn">
			<mt-button type="primary" @click="deletebook()" :disabled="isdisable()">删除</mt-button>
		</div>
	</div>
</template>
<script>
	import { Toast } from 'mint-ui';
	export default{
		name:'delbook',
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
				booklist: [],
				bookqid:[],
			}
		},
		mounted(){
		    this.getbooklist();
		    this.scroll();
		},
		methods:{
			isdisable() {
				console.log(this.bookqid)
				if(this.bookqid.length>0){
					return false;
				}else{
					return true;
				}
			},
			getbooklist () {
				this.$axios({
		      		method: 'post',
			    	url:'/bookrack/mybookrack',	
			    	data:{
			    		draw: this.draw,
			    		length:this.length,
			    	}
				})
				.then(function(res){
					console.log(res.data.data.data)
					this.pageTotal = res.data.data.pageTotal;
					this.booklist = this.booklist.concat(res.data.data.data);
					console.log(this.booklist)
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
			deletebook () {
				if(this.bookqid.length<=0){
					Toast({
					  message: '没有选择图书',
					  position: 'middle',
					  duration: 2000
					});
				}else{
					this.$axios({
			      		method: 'post',
				    	url:'/bookrack/bookrackdelete',	
				    	data:{
				    		bookqid: this.bookqid,
				    	}
					})
					.then(function(res){
						this.bookqid = [];
						Toast({
						  message: '删除成功',
						  position: 'middle',
						  duration: 2000
						});
						this.$router.back(-1);
					}.bind(this))
					.catch(function (error) {
						console.log(error);
					});
				}
			},
			isChack (e,bookqid) {
				var el = e.currentTarget;
				if($(el).find('.selected').attr('datashow') == "false"){
					$(el).find('.selected').attr('datashow',true);
					$(el).find('.selected').find('img').css({
						display: 'block'
					});
					this.bookqid.push(bookqid);
				}else{
					$(el).find('.selected').attr('datashow',false);					
					$(el).find('.selected').find('img').css({
						display: 'none'
					});
					var that = this;
					$.each(this.bookqid,function(k,v){
						if(v == bookqid){
							that.bookqid.splice(k, 1);
						}
					})
				}
			},
		},
	}
</script>
<style scoped>
	.isdisable{

	}
	.delbook{
		position: relative;
	}
	.delbtn{
		position: fixed;
		bottom: 0;
		left:0;
		width:100%;
		height:0.9rem;
		background: #fff;
	}
	.delbtn button{
		display: block;
		height:100%;
		width:90%;
		margin:0 auto;
		background: #13b7f6;
		color:#fff;
		border-radius: 5px;
	}
	.booklist{
		padding:0.3rem;
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    -webkit-flex-wrap: wrap;
   		flex-wrap: wrap;
	}
	.booklist li{
		width:1.6rem;
		margin-right: 0.33rem;
		margin-bottom: 0.2rem;
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
	.booklist li dt .selected img{
		display: none;
	}
	.booklist li dd{
		padding:0.1rem 0;
		font-size: 0.24rem;
		color:#000;
	}
</style>