<template>
	<div class="editprofile">
		<div class="headportrait" @click="">
			<div class="projectname">
				头像
			</div>
			<div class="headimgwarp">
				<div class="headimg">
					<input id="file" type="file" style="position:absolute;top:0;left:0;width:100%;height:100%;opacity: 0;" accept="image/*" @change="uploadimg($event)"/>
					<img :src="usermsg.imgurl" >
				</div>
			</div>
			<!-- <div class="more">
				<img src="../assets/image/mine/more.png" >
			</div> -->
		</div>
		<ul>
			<li>
				<mt-field label="用户名" placeholder="请输入用户名" v-model="usermsg.nickname"></mt-field>
			</li>
			<li @click="selectsex()">
				<mt-cell title="性别">
				  <span style="color: #707070">{{usermsg.gender==1?'男':'女'}}</span>
				</mt-cell>
			</li>
			<li>
				<mt-cell title="个性签名" is-link to="/signature">
				  <span style="color: #707070">介绍一下自己吧</span>
				</mt-cell>
			</li>
			<li>
				<mt-cell title="手机账号" is-link to="/bindphonenum">
				  <span style="color: #707070">绑定手机账号</span>
				</mt-cell>
			</li>
		</ul>
		<m-picker :slots='slots' :isPicker='isPicker' :indexText='indexText' :datakey='datakey' :valueKey='valueKey' @confirm='pickerConfirm' @cancel='pickerCancel'>
	    </m-picker>
		<div class="btn">
			<mt-button type="primary" @click="save()">保存</mt-button>
		</div>
	</div>
</template>
<script>
	import mPicker from './picker/index';
	import { Toast } from 'mint-ui';
	export default{
		name:'editprofile',
		data () {
			return {
				sex:'男',
			    datakey: '', 
			    slots: [{values: [
		                      {k: '1', v: '男'},
		                      {k: '2', v: '女'},
		                  ]}],
			    isPicker: false,
      			indexText: '请选择',
      			valueKey: 'v',
      			info: {
      				sex:{
      					type:'',
			        	desc:''
      				}
			    },
			    usermsg:{},
			}
		},
		components: {
		    mPicker,
		},
		mounted(){
		    this.getusermsg();
		},
		methods: {
			uploadimg (el) {
				var that = this;
				lrz(el.target.files[0])
		        .then(function (rst) {
		            //console.log(rst)
		            rst.formData.append('base64img', rst.base64); 
		            //console.log(rst.formData);
		            that.$axios({
			      		method: 'post',
				    	url:'/users/uploadfile',	
				    	data: rst.formData,
					})
					.then(function(res){
						console.log(res.data.data)
						that.usermsg.imgurl = res.data.data;
					})
					.catch(function (error) {
						console.log(error);
					});
		        })
		        .catch(function (err) {
		            console.log(err)
		        })
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
			},
		   	selectsex() {
		      this.datakey = 'sex';
		      this.isPicker = true;
		    },
		    pickerConfirm(value, key) {
		      this.isPicker = false;
		      this.usermsg.gender = value[0].k;
		      // this.info[key].type = value[0].k;
		      // this.info[key].desc = value[0].v;
		    },
		    pickerCancel() {
		      this.isPicker = false;
		    },
		    save () {
		    	console.log(this.usermsg.imgurl)
		    	console.log(this.usermsg.nickname)
		    	console.log(this.info.sex.type)
		    	if(this.usermsg.imgurl == ''){
		    		Toast({
					  message: '没有添加头像',
					  position: 'middle',
					  duration: 2000
					});
		    	}else if(this.usermsg.nickname == ''){
		    		Toast({
					  message: '没有添加用户名',
					  position: 'middle',
					  duration: 2000
					});
		    	}else if(this.usermsg.gender == ''){
		    		Toast({
					  message: '没有添加性别',
					  position: 'middle',
					  duration: 2000
					});
		    	}else{
		    		this.$axios({
			      		method: 'post',
				    	url:'/users/updateuser',	
				    	data:{
				    		imgurl: this.usermsg.imgurl,
				    		nickname: this.usermsg.nickname,
				    		gender: this.usermsg.gender,
				    	}
					})
					.then(function(res){
						Toast({
						  message: '修改成功',
						  position: 'bottom',
						  duration: 2000
						});
						setTimeout(() => {
						  this.$router.go(-1);
						}, 2000);
					}.bind(this))
					.catch(function (error) {
						console.log(error);
					});
		    	}
		    }
  		},
	}
</script>
<style scoped>
	.editprofile{
		height:100%;
		width:100%;
		background: #ececec;
		position: relative;
	}
	.btn{
		position: fixed;
		bottom: 0;
		left:0;
		width:100%;
		height:0.9rem;
	}
	.btn button{
		display: block;
		height:100%;
		width:90%;
		margin:0 auto;
		background: #13b7f6;
		color:#fff;
	}
	.headportrait{
		background: #fff;
		display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
	    -webkit-flex-wrap: wrap;
   		flex-wrap: wrap;
   		padding:0.3rem;
   		margin-bottom: 0.3rem;
	}
	.headportrait .projectname{
		box-flex:2;
		-webkit-box-flex:2;
		-moz-box-flex:2;
		flex:2;
		-webkit-flex:2;
		-webkit-align-self: center;
    	align-self: center;
	}
	.headportrait .headimgwarp{
		position: relative;
		box-flex:14;
		-webkit-box-flex:14;
		-moz-box-flex:14;
		flex:14;
		-webkit-flex:14;
		-webkit-align-self: center;
    	align-self: center;
    	display:-webkit-box;
	    display: -moz-box;
	    display: -ms-flexbox;
	    display: -webkit-flex;
	    display: flex;
    	-webkit-flex-direction: row-reverse;
    	flex-direction: row-reverse;
    	padding:0 0.3rem;
	}
	.headportrait .headimg{
		width:1.17rem;
		height:1.17rem;
	}
	.headportrait .headimg{
		width:1.17rem;
		height:1.17rem;
		border-radius:100%;
		overflow:hidden;
	}
	.headportrait .headimg img{
		width:100%;
		height:100%;
	}
	.headportrait .more{
		box-flex:1;
		-webkit-box-flex:1;
		-moz-box-flex:1;
		flex:1;
		-webkit-flex:1;
		-webkit-align-self: center;
    	align-self: center;
	}
	.picker{
		z-index: 9999;
	}
</style>