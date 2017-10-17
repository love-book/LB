<template>
	<div class="shelvesrecord">
		<mt-loadmore :bottom-method="loadBottom" @bottom-status-change="handleTopChange" :bottom-all-loaded="allLoaded" :auto-fill="false" ref="loadmore">
			<ul class="shelveslist">
				<li class="shelveslist_li" v-for="(item,index) in shelveslist">
					<div class="bookimg">
						<img :src="item.imageurl">
					</div>
					<ul class="bookinfo">
						<li>
							{{item.bookname}}							
						</li>
						<li>
							{{item.auhtor}}
						</li>
						<li>
							{{item.depreciation}}
						</li>
						<li>
							{{changetime(parseInt(item.create_time*1000))}}
						</li>
					</ul>	
				</li>
			</ul>
		</mt-loadmore>
	</div>
</template>
<script>
	export default{
		name:'shelvesrecord',
		data () {
			return{
				bottomStatus:'',
	    		pageTotal:0,
			    allLoaded: false, //是否可以上拉属性，false可以上拉，true为禁止上拉，就是不让往上划加载数据了
			    scrollMode:"auto", //移动端弹性滚动效果，touch为弹性滚动，auto是非弹性滚动
			    loading:false,
				draw:1,
	    		length:3,
				shelveslist:[],
			}
		},
		methods:{
			changetime (date) {
				var oDate = new Date(date);
				function p(s) {
			        return s < 10 ? '0' + s: s;
			    }
				return oDate.getFullYear()+'-'+(oDate.getMonth()+1)+'-'+oDate.getDate()+' '+oDate.getHours()+':'+p(oDate.getMinutes())+':'+p(oDate.getSeconds());
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
			getshelveslist() {
		      	this.$axios({
		      		method: 'post',
			    	url:'/bookrack/getmybooklist',	
			    	data:{
			    		draw:this.draw,
	    				length:this.length,
			    	}
				})
				.then(function(res){
					console.log(res);
					this.pageTotal = res.data.data.pageTotal;
					this.shelveslist = this.shelveslist.concat(res.data.data.data);
					// if(this.shelveslist.length>9){
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
		        this.getshelveslist();
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
		mounted () {
		    this.getshelveslist();
		    this.scroll();
		},
	}
</script>
<style scoped>
	.shelvesrecord{
		height:100%;
		width:100%;
		background: #ececec;
	}
	.shelveslist{
		background: #ececec;
		padding:0.3rem 0.2rem;
	}
	.shelveslist .shelveslist_li{
		padding:0.3rem;
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    background: #fff;
	    margin-bottom: 0.1rem;
	}
	.shelveslist .shelveslist_li .bookimg{
		width:1.8rem;
		height:2.4rem;
	}
	.shelveslist .shelveslist_li .bookimg img{
		width:100%;
		height:100%;
	}
	.shelveslist .shelveslist_li .bookinfo{
		padding:0 0.2rem;
		box-flex:1;
		-webkit-box-flex:1;
		-moz-box-flex:1;
		flex:1;
		-webkit-flex:1;
	}
	.shelveslist .shelveslist_li .bookinfo li{
		margin-bottom: 0.1rem;
		text-overflow:ellipsis;
		overflow:hidden;
		white-space:nowrap;
		width:3.2rem;
	}
</style>