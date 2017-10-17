<template>
	<div class="mine">
		<router-link to="/editprofile">
			<div class="personalCenter" @touchstart="changecolor($event)" @touchend="reinstatecolor($event)">
				<div class="headerimg">
					<img :src="usermsg.imgurl" >
				</div>
				<div class="nickname">
					<h3>
						{{usermsg.nickname}}
					</h3>
					<p>
						{{usermsg.signature === '' ? '该用户还未添加个性签名' : usermsg.signature}}
					</p>
				</div>		
				<div class="moreimg">
					<img src="../assets/image/mine/more.png" >
				</div>
			</div>
		</router-link>
		<div class="collect">
			<router-link to="/collectbooks" class="collectbooks">
				<div @touchstart="changecolor($event)" @touchend="reinstatecolor($event)">
					<img src="../assets/image/mine/book.png" >
					收藏的图书
				</div>
			</router-link>
			<router-link to="/collectbookcase" class="collectbookcase">
				<div @touchstart="changecolor($event)" @touchend="reinstatecolor($event)">
					<img src="../assets/image/mine/bookcase.png" >
					收藏的书架
				</div>
			</router-link>
		</div>
		<ul class="minelist">
			<li>
				<mt-cell title="消息" is-link to="/message">
				  <img slot="icon" src="../assets/image/mine/news.png" width="24" height="24" style="display:inline-block;">
				  <mt-badge type="error">10</mt-badge>
				</mt-cell>
			</li>
			<li>
				<mt-cell title="交换记录" is-link to="/exchange">
				  <img slot="icon" src="../assets/image/mine/exchange.png" width="24" height="24" style="display:inline-block;">
				</mt-cell>
			</li>
			<li>
				<mt-cell title="上架记录" is-link to="/shelvesrecord">
				  <img slot="icon" src="../assets/image/mine/shelves.png" width="24" height="24" style="display:inline-block;">
				</mt-cell>
			</li>
			<li>
				<mt-cell title="意见箱" is-link to="/ideasbox">
				  <img slot="icon" src="../assets/image/mine/Suggestion.png" width="24" height="24" style="display:inline-block;">
				</mt-cell>
			</li>
			<li>
				<mt-cell title="关于我们" is-link to="/aboutus">
				  <img slot="icon" src="../assets/image/mine/aboutus.png" width="24" height="24" style="display:inline-block;">
				</mt-cell>
			</li>
		</ul>
	</div>
</template>
<script>
	export default{
		name: 'bookstore',
		data() {  
		    return {  	 
		      	usermsg:{},  
		    };  
		},
		mounted(){
		    this.getusermsg();
		},
		methods: {
			changecolor (e) {
				var el = e.currentTarget;
				$(el).css("background","#eee");
			},
			reinstatecolor (e){
				var el = e.currentTarget;
				$(el).css("background","#ffffff");
			},
			getusermsg () {
				this.$axios({
		      		method: 'post',
			    	url:'/users/userinfo',	
			    	data:{
			    		
			    	}
				})
				.then(function(res){
					console.log(res);
					this.usermsg = res.data.data;
				}.bind(this))
				.catch(function (error) {
					console.log(error);
				});
			}
		}
	}
</script>
<style scoped>
	.mine{
		height:100%;
		width:100%;
		background: #ececec;
	}
	.personalCenter{
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
		width:90%;
		padding:0.3rem;
		background: #ffffff;
	}
	.personalCenter .headerimg{
		width:1.5rem;
		height:1.5rem;
		border-radius: 100%;
		overflow: hidden;
	}
	.personalCenter .headerimg img{
		width:100%;
		height:100%;
	}
	.personalCenter .nickname{
		box-flex:1;
		-webkit-box-flex:1;
		-moz-box-flex:1;
		flex:1;
		-webkit-flex:1;
		padding:0.2rem 0.3rem;
	}
	.personalCenter .nickname h3{
		margin-bottom: 0.2rem;
		color:#424242;
	}
	.personalCenter .moreimg{
		position: relative;
		width:0.2rem;
	}
	.personalCenter .moreimg img{
		width:100%;
		position: absolute;
        top: 50%;
        left: 50%;
        margin-top: -10px; 
        margin-left: -5px;
	}
	.collect{
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    height:0.75rem;
	    width:100%;
	    background:#fff;
	    margin:0.18rem 0;
	}
	.collect .collectbooks{
		box-flex:1;
		-webkit-box-flex:1;
		-moz-box-flex:1;
		flex:1;
		-webkit-flex:1;
		text-align: center;
	}
	.collect .collectbooks div{
		width:100%;
		height:100%;
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    -webkit-align-items: center;
    	align-items: center;
    	-webkit-justify-content: center;
    	justify-content: center;
	}
	.collect .collectbooks img{
		display: inline-block;
		vertical-align: middle;
		width:0.37rem;
		height:0.45rem;
		margin-right: 0.22rem;
	}
	.collect .collectbookcase{
		box-flex:1;
		-webkit-box-flex:1;
		-moz-box-flex:1;
		flex:1;
		-webkit-flex:1;
		border-left: 1px solid #ebebeb;
		text-align: center;
	}
	.collect .collectbookcase div{
		width:100%;
		height:100%;
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    -webkit-align-items: center;
    	align-items: center;
    	-webkit-justify-content: center;
    	justify-content: center;
	}
	.collect .collectbookcase img{
		display: inline-block;
		vertical-align: middle;
		width:0.37rem;
		height:0.45rem;
		margin-right: 0.22rem;
	}
	.minelist li{
		margin-bottom: 2px;
	}
</style>