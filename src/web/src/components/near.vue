<template>
	<div class="near">
		<div class="header">
			<div class="blueline"></div>
			<img class="place" src="../assets/image/icon/place.png" height="37" width="24">
			<div>
				北京
			</div>
			<div class="screening" @click="open()">
				筛选
				<img src="../assets/image/icon/more.png" alt="">
			</div>
			<mt-popup class="search" v-model="popupVisible" position="right" :modal=false>
				<ul class="sex">
					<li v-for="(item,index) in sexlist" @click="selected(index)" :class="{'selected':item.show}">
						<img v-if="item.type==1 && !item.show" src="../assets/image/icon/man.png" alt="">
						<img v-else-if="item.type==1 && item.show" src="../assets/image/icon/baiman.png" alt="">
						<img v-else-if="item.type==2 && !item.show" src="../assets/image/icon/woman.png" alt="">
						<img v-else-if="item.type==2 && item.show" src="../assets/image/icon/baiwoman.png" alt="">
						{{item.desc}}						
					</li>
					<!-- <li @click="selected" sex-data="2">
						<img src="../assets/image/icon/woman.png" alt="">女
					</li>
					<li @click="selected" sex-data="0">
						不限
					</li> -->
				</ul>
				<ul class="searchlist">
					<li @click="selectage">
						<div class="title">
							年龄
						</div>
						<div class="elect">
							<p>
								{{info.age.desc}}
							</p>
							<img src="../assets/image/icon/more.png" alt="">
						</div>
					</li>
					<li @click="selectdistance">
						<div class="title">
							距离
						</div>
						<div class="elect">
							<p>
								{{info.distance.desc}}
							</p>
							<img src="../assets/image/icon/more.png" alt="">
						</div>
					</li>
					<li @click="selectonlinetime">
						<div class="title">
							上线时间
						</div>
						<div class="elect">
							<p>
								{{info.onlinetime.desc}}
							</p>
							<img src="../assets/image/icon/more.png" alt="">
						</div>
					</li>
				</ul>
				<m-picker :slots='ageslot' :isPicker='agePicker' :indexText='indexText' :datakey='datakey' :valueKey='valueKey' @confirm='pickerConfirm' @cancel='pickerCancel'>
			    </m-picker>
			    <m-picker :slots='distanceslot' :isPicker='distancePicker' :indexText='indexText' :datakey='datakey' :valueKey='valueKey' @confirm='pickerConfirm' @cancel='pickerCancel'>
			    </m-picker>
			    <m-picker :slots='linetimeslot' :isPicker='linetimePicker' :indexText='indexText' :datakey='datakey' :valueKey='valueKey' @confirm='pickerConfirm' @cancel='pickerCancel'>
			    </m-picker>
			  	<button id="searchbtn" type="button" @click="close()">完成</button>
			</mt-popup>
		</div>
		<div class="cf"></div>
		
		<mt-loadmore :bottom-method="loadBottom" @bottom-status-change="handleTopChange" :bottom-all-loaded="allLoaded" :auto-fill="false" ref="loadmore">
			<ul class="userlist">
				<li v-for='(item,index) in userlist' :class="{'isMargin': isMargin(index)}">
					<router-link :to="{ name: 'otherbookcase', params:{ userid: item.userid }}">
						<img class="headimg" :src="item.imgurl" alt="">
						<div class="usermsg">
							<p class="nickname">
								{{item.nickname}}
							</p>
							<img v-if="item.gender == 2" src="../assets/image/icon/woman.png" alt="">
							<img  v-else src="../assets/image/icon/man.png" alt="">
						</div>
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
	import mPicker from './picker/index';
	
	export default{
		name:'near',
		data() {
			return{
				popupVisible:false,
				bottomStatus:'',
	    		pageTotal:0,
			    draw:1,
	    		length:1,
	    		// gender:'',
	    		// age:'',
	    		// radius:'',
			    allLoaded: false, //是否可以上拉属性，false可以上拉，true为禁止上拉，就是不让往上划加载数据了
			    scrollMode:"auto", //移动端弹性滚动效果，touch为弹性滚动，auto是非弹性滚动
			    loading:false,
			    showPicker: false,
			    datakey: '', 
			    ageslot: [{values: [
		      				  {k: '', v: '不限'},
		                      {k: '1', v: '0岁-17岁'},
		                      {k: '2', v: '18岁-23岁'},
		                      {K: '3', v: '18岁-23岁'},
		                      {k: '4', v: '24岁-30岁'},
		                      {k: '5', v: '31岁-40岁'},
		                      {K: '6', v: '40岁以上'},
		                  ]}],
			    distanceslot:[{values: [
		      				  {k: '', v: '不限'},
		                      {k: '1', v: '1km以内'},
		                      {k: '2', v: '1km-3km'},
		                      {K: '3', v: '3km-10km'},
		                      {k: '4', v: '10km-50km'},
		                      {k: '5', v: '50km以上'},
		                  ]}],
			    linetimeslot:[{values: [
		      				  {k: '', v: '不限'},
		                      {k: '1', v: '<1天'},
		                      {k: '2', v: '<3天'},
		                      {K: '3', v: '<7天'},
		                      {k: '4', v: '<30天'},
		                  ]}],
			    agePicker: false,
			    distancePicker: false,
			    linetimePicker: false,
      			indexText: '请选择',
      			valueKey: 'v',
      			info: {
      				sex:{
      					type:'',
			        	desc:''
      				},
			        age:{
			        	type:'',
			        	desc:''
			        },
			        distance:{
			        	type:'',
			        	desc:''
			        },
			        onlinetime:{
			        	type:'',
			        	desc:''
			        }
			    },
			    sexlist:[
			    	{
			    		type:'1',
			    		desc:'男',
			    		show:false
			    	},
			    	{
			    		type:'2',
			    		desc:'女',
			    		show:false
			    	},
			    	{
			    		type:'',
			    		desc:'不限',
			    		show:false
			    	},
			    ],
			    userlist:[],
			}
		},
		components: {
		    mPicker,
		},
		mounted(){
		    this.loadPageList();  //初次访问查询列表
		    this.scroll();
		},
		methods:{
			open () {
				this.popupVisible = true;
			},
			close () {
				console.log(this.info);
				this.popupVisible = false;
				this.userlist = [];
				this.draw = 1;
				this.loadPageList();
			},
			selected(index) {
				for(var i=0;i<this.sexlist.length;i++){
	  				this.sexlist[i].show = false;
	  			}
	  			this.sexlist[index].show = true;	
	  			this.info.sex.type = this.sexlist[index].type;
	  			this.info.sex.desc = this.sexlist[index].desc;
			},
			
			selectage() {
		      this.datakey = 'age';
		      this.agePicker = true;
		    },
		    selectdistance() {
		      this.datakey = 'distance';
		      this.distancePicker = true;
		    },
		    selectonlinetime() {
		      this.datakey = 'onlinetime';
		      this.linetimePicker = true;
		    },
			pickerConfirm(value, key) {
		      this.agePicker = false;
		      this.distancePicker = false;
		      this.linetimePicker = false;
		      console.log(value);

		      this.info[key].type = value.k;
		      this.info[key].desc = value.v;
		      console.log(this.info.age)
		    },
		    pickerCancel() {
		      this.agePicker = false;
		      this.distancePicker = false;
		      this.linetimePicker = false;
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
			loadPageList() {
		      	this.$axios({
		      		method: 'post',
			    	url:'/users/getusersbylocaltion',	
			    	data:{
			    		length: this.length,
			            draw: this.draw,
			            gender: this.info.sex.type,
			            age: this.info.age.type,
			            radius: this.info.distance.type,
			    	}
				})
				.then(function(res){
					console.log(res);
					console.log(res.data.data);
					this.pageTotal = res.data.data.pageTotal;
					this.userlist = this.userlist.concat(res.data.data.data);
					// if(this.userlist.length>9){
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
		    },
		    isMargin(index){
				if((index+1)%3==0){
					return true;
				}else{
					return false;
				}
			},
		}
	}
</script>
<style scoped>
	.isMargin{
		margin-right: 0!important;
	}
	.selected{
		background: #01a4df;
		color:#fff;
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
	.mint-loadmore{
		padding-bottom: 1.2rem;
	}
	.header{
		height:0.6rem;
		background: #fff;
		margin-bottom:0.15rem;
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    -webkit-align-items: center;
    	align-items: center;
	}
	.blueline{
		width:0.3rem;
		height:100%;
		background: #12b5f6;
	}
	.place{
		width:0.24rem;
		height:0.37rem;
		margin:0 0.2rem;
	}
	.screening{
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
	    -webkit-justify-content: flex-end;
    	justify-content: flex-end;
    	-webkit-align-items: center;
    	align-items: center;
	}
	.screening img{	
		width:0.15rem;
		height:0.25rem;
		margin:0 0.2rem;
	}
	.search{
		height:100%;
		width:100%;
		background: #ebeeee;
	}
	.search .sex{
		height:0.9rem;
		margin:0.2rem 0;
		background: #fff;
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
    	-webkit-align-items: center;
    	align-items: center;
	}
	.search .sex li{
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
	    -webkit-justify-content: center;
    	justify-content: center;
    	-webkit-align-items: center;
    	align-items: center;
    	height:100%;
	}
	.search .sex li img{
		width:0.3rem;
		height:0.3rem;
		margin-right: 0.2rem;
	}
	.search .searchlist{
		background: #fff;
	}
	.search .searchlist li{
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    -webkit-justify-content: space-between;
    	justify-content: space-between;
    	-webkit-align-items: center;
    	align-items: center;
    	padding:0.2rem 0.3rem;
	}
	.search .searchlist li img{
		width:0.16rem;
		height:0.25rem;
	}
	.search .searchlist li .elect{
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    -webkit-align-items: center;
    	align-items: center;
	}
	.search .searchlist li .elect p{
		color:#aaa;
		margin-right: 0.2rem;
	}
	#searchbtn{
		width:90%;
		height:0.9rem;
		display: block;
		background: #12b5f6;
		color:#fff;
		border-radius:5px;
		font-size:0.3rem;
		margin:0.5rem auto;
	}
	.userlist{
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
	.userlist li{
		position: relative;
		width:2rem;
		height:2.2rem;
		margin-right: 0.19rem;
		margin-bottom: 0.15rem;
	}
	.userlist li .headimg{
		width:100%;
		height:100%;
	}
	.usermsg{
		position: absolute;
		left:0;
		bottom:0;
		width:100%;
		height:0.35rem;
		background:rgba(0,0,0,.3);
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    -webkit-align-items: center;
    	align-items: center;
	}
	.usermsg p{
		box-flex:5;
		-webkit-box-flex:5;
		-moz-box-flex:5;
		flex:5;
		-webkit-flex:5;
		font-size:0.18rem;
		color:#fff;
		text-align: center;
		text-overflow:ellipsis;
		overflow:hidden;
		white-space:nowrap;
	}
	.usermsg img{
		box-flex:1;
		-webkit-box-flex:1;
		-moz-box-flex:1;
		flex:1;
		-webkit-flex:1;
		width:0.32rem;
		height:0.32rem;
	}
</style>