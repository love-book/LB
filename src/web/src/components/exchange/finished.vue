<template>
	<div class="finished">
		<ul class="toexchangelist">
			<li>
				<div class="title">
					<div class="nickname">
						清歌 & 二狗
					</div>
					<div class="state">
						已结束
					</div>	
				</div>
				<div class="book">
					<div class="bookdetails">
						<dt>
							<img src="../../assets/image/bookstore/text.jpg" height="100%" width="100%">
						</dt>
						<dd>
							世界很大，幸好有你
						</dd>
					</div>
					<div class="changeimg">
						<img src="../../assets/image/exchange/change.png">
					</div>
					<div class="bookdetails">
						<dt>
							<img src="../../assets/image/bookstore/text.jpg" height="100%" width="100%">
						</dt>
						<dd>
							世界很大，幸好有你
						</dd>
					</div>
				</div>
				<div class="operation">
					<mt-button type="default" size="small">删除订单</mt-button>
				</div>
			</li>
		</ul>
	</div>
</template>
<script>
	export default{
		name:'finished',
		data() {  
		    return {  
		      	selected: '1'  
		    };  
		},
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