<template>
	<div class="choosesex">
		<h3>
			请选择您的性别！
		</h3>
		<div class="select">
			<div class="sex">
				<img src="../assets/image/common/woman.png">
				<span class="woman">女</span>
			</div>
			<div class="line">
				
			</div>
			<div class="sex">
				<img src="../assets/image/common/man.png">
				<span class="man">男</span>
			</div>
		</div>
	</div>
</template>
<script>
	export default{
		name:'choosesex',
		data() {
			return{

			}
		}
	}
</script>
<style scoped>
	.choosesex{
		height:100%;
		width:100%;
		overflow: hidden;
		background: #ccc;
	}
	h3{
		text-align: center;
		margin-top:3.8rem;
		margin-bottom: 0.8rem;
	}
	.select{
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
    	align-items:center;
	    height:1.75rem;
	}
	.select .sex{
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
	    flex-direction:column;
		-webkit-justify-content: center;
    	justify-content: center;
    	align-items:center;
	}
	.select .sex img{
		width:0.85rem;
		height:0.85rem;
	}
	.select .sex .woman{
		color:#f28ea6;
	}
	.select .sex .man{
		color:#71bdf1;
	}
	.select .line{
		height:60%;
		width:2px;
		background: #dbdae0;
	}
</style>