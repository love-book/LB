<template>
	<div class="bookcase">
		<div class="personalinformation">
			<div class="settings">
				<router-link to="/delbook">
					<img src="../assets/image/bookinfo/collect.png">
				</router-link>
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
				{{personmsg.signature === '' ? '该用户还未添加个性签名' : personmsg.signature}}
			</p>
			<div class="cf"></div>
		</div>
		<div class="main_warp" id="main_warp">
			<div class="main_title">
				<div class="titel_name">
					<span class="line"></span>
					<p>
						个人书架
					</p>
				</div>
				<!-- <router-link to="">
		  			<div class="checkmore">
		  				查看更多<img src="../assets/image/bookcase/more.png">
		  			</div>
		  		</router-link> -->
			</div>
			<mt-loadmore :bottom-method="loadBottom" @bottom-status-change="handleTopChange" :bottom-all-loaded="allLoaded" :auto-fill="false" ref="loadmore">
				<ul class="booklist">
					<li v-for='(item,index) in booklist' :class="{'isMargin': isMargin(index)}">
						<div v-if="!item.type">
							<router-link :to="'/book/'+item.bookid">
								<dt>
									<img :src="item.imageurl" height="100%" width="100%">
								</dt>
								<dd>
									{{item.bookname}}
								</dd>
							</router-link>
						</div>
						<div v-else class="addwarp" @click="addbook()">
							<dt class="add">
								<img src="../assets/image/bookstore/add.png" height="50" width="50">
							</dt>
							<dd class="textcenter">
								{{item.bookname}}
							</dd>
							<div class="layer" v-if="isShow">
								<div class="angle">
									<img src="" alt="">
								</div>
								<ul class="uploadList">
									<li class="scan" @click="richscan()">
										<img src="../assets/image/bookinfo/edit.png" height="27" width="27" alt="">
										<p>
											扫一扫上传
										</p>
									</li>
									<li @click="takepictures()">
										<img src="../assets/image/bookinfo/edit.png" height="27" width="27" alt="">
										<p>
											拍照上传
										</p>
									</li>
								</ul>
							</div>
						</div>	
					</li>
				</ul>
			</mt-loadmore>
			
		</div>
		<router-link to="" class="declaration">
		  		服务声明
		</router-link>
		<div class="gotop" @click="gotop()">
			<img src="../assets/image/bookcase/gotop.png">
		</div>
	</div>
</template>
<script>
	import { Loadmore } from 'mint-ui';
	import { Toast } from 'mint-ui';
	export default{
		name: "bookcase",
		data() {
			return {
				isShow:false,
				bottomStatus:'',
	    		pageTotal:0,
			    allLoaded: false, //是否可以上拉属性，false可以上拉，true为禁止上拉，就是不让往上划加载数据了
			    scrollMode:"auto", //移动端弹性滚动效果，touch为弹性滚动，auto是非弹性滚动
			    loading:false,
				draw:1,
	    		length:1,
				personmsg:{},
				booklist:[
					{type:2, bookname: '点击添加'},
					
				]
			}
		},
		methods: {
			wechat () {
				this.$axios({
		      		method: 'post',
			    	url:'/app/getwxconfig',	
			    	data:{
			    		url: 'http://api.kasoly.com/'
			    	}
				})
				.then(function(res){
					if (res.status == 200) {
	                    wx.config({
	                        debug: false, // 开启调试模式,调用的所有api的返回值会在客户端alert出来，若要查看传入的参数，可以在pc端打开，参数信息会通过log打出，仅在pc端时才会打印。
	                        appId: res.data.data.appId, // 必填，公众号的唯一标识
	                        timestamp: res.data.data.timestamp, // 必填，生成签名的时间戳
	                        nonceStr: res.data.data.nonceStr, // 必填，生成签名的随机串
	                        signature: res.data.data.signature,// 必填，签名，见附录1
	                        jsApiList: ['scanQRCode'] // 必填，需要使用的JS接口列表，所有JS接口列表见附录2
	                    });
	                }
				}.bind(this))
				.catch(function (error) {
					console.log(error);
				});
			},
			isMargin(index){
				if((index+1)%3==0){
					return true;
				}else{
					return false;
				}
			},
			gotop () {
				document.getElementById('main_warp').scrollTop = 0;
			},
			addbook () {
				if(this.isShow){
					this.isShow = false;
				}else{
					this.isShow = true;
				}
			},
			getpersonmsg () {
				this.$axios({
		      		method: 'post',
			    	url:'/users/userinfo',	
			    	data:{
			    		
			    	}
				})
				.then(function(res){
					console.log(res)
					this.personmsg = res.data.data;
					//console.log(this.personmsg)
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
			    	}
				})
				.then(function(res){
					//console.log(res);
					this.pageTotal = res.data.data.pageTotal;
					this.booklist = this.booklist.concat(res.data.data.data);
					// if(this.booklist.length>9){
					// 	this.show = true;
					// 	this.allLoaded = false;
					// }else{
					// 	$(".mint-loadmore-bottom").css('display','none');
					// 	this.allLoaded = true;
					//  	$(".mint-loadmore-bottom").css('display','block');
					//}
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
		    richscan () {
		    	var that = this;
		    	wx.scanQRCode({
				    needResult: 1, // 默认为0，扫描结果由微信处理，1则直接返回扫描结果，
				    scanType: ["qrCode","barCode"], // 可以指定扫二维码还是一维码，默认二者都有
				    success: function (res) {
					    var result = res.resultStr; // 当needResult 为 1 时，扫码返回的结果
					    var arr =[];
					    arr = result.split(',');
					    var isbn = arr[1];
					    that.$router.push({ name: 'editbookrs', params:{ isbn: isbn}});
					}
				});
		    },
		    takepictures () {
		    	this.$router.push({ name: 'editbooktp'});
		    },
		},
		mounted () {
			this.getpersonmsg();
		    this.getbooklist();
		    this.scroll();
		    this.wechat();
		},
	}
</script>
<style scoped>
	.bookcase{
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    flex-direction:column;
	    height:100%;
	}
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
	}
	.personalinformation .settings{
		position: absolute;
		right:0.3rem;
		top:0.3rem;
		width:0.5rem;
		height:0.5rem;
	}
	.personalinformation .settings img{
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
		box-flex:1;
		-webkit-box-flex:1;
		-moz-box-flex:1;
		flex:1;
		-webkit-flex:1;
		overflow: auto;
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    flex-direction:column;
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
	.checkmore{
		font-size: 0.18rem;
		color:#a5a5a5;
	}
	.checkmore img{
		display: inline-block;
		width:0.1rem;
		height:0.18rem;
		margin-left: 0.05rem;	
	}
	.mint-loadmore{
		/*box-flex:1;
		-webkit-box-flex:1;
		-moz-box-flex:1;
		flex:1;
		-webkit-flex:1;*/
		/*margin-bottom: 1rem;*/
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
		/*padding-bottom:2rem;*/
		overflow: auto;
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
	.declaration{
		display: block;
		width:100%;
		text-align:center;
		padding-bottom:1.2rem;
		font-size: 0.2rem;
		color:#a5a5a5;
	}
	.addwarp{
		position: relative;
	}
	.addwarp .layer{
		z-index: 9999;
		position: absolute;
		top:1.4rem;
		left:0.2rem;
	}
	.addwarp .layer .uploadList{
		border-radius: 5px;
		background: #a3a3a3;
		overflow: hidden;
	}
	.addwarp .layer .uploadList li{
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    -webkit-align-items: center;
    	align-items: center;
		background: #a3a3a3;
		margin:0;
		padding:0.1rem 0.1rem;
	}
	.addwarp .layer .uploadList li img{
		width:0.24rem;
		height:0.24rem;
		margin-right: 0.2rem;
	}
	.addwarp .layer .uploadList .scan{
		border-bottom: 1px solid #c1c1c1;
	}
	.addwarp .layer .uploadList li p{
		font-size: 0.2rem;
		color:#fff;
	}
</style>