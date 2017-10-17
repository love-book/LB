require('es6-promise').polyfill()
import axios from 'axios'

function getQueryString(name) {  
    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)", "i");  
    var r = window.location.search.substr(1).match(reg);  
    if (r != null) return unescape(r[2]); return null;   
}  
var code = getQueryString("code");
var token = ''; 
// var Axios = axios.create({
	  
// });
console.log(code);
// axios({
//   method: 'post',
//   url: 'http://app.kasoly.com/v1/app/getusertokenbycode',
//   data: {
//     code: code,
//   }
// })
// .then(function(res){
// 	console.log(res)
// 	token = res.data;
// 	console.log(token)
// 	Axios = axios.create({
// 	  baseURL: 'http://app.kasoly.com/v1', 
// 	  headers: {'token': token},
// 	})
// })
// .catch(function (error) {
// 	console.log(error);
// });
// $.ajax({
// 	url: 'http://app.kasoly.com/v1/app/getusertokenbycode',
// 	type: 'post',
// 	data: {code: code},
// 	async: false,
// })
// .done(function(res) {
// 	console.log(res)
// 	token = res.data;
// 	console.log(token)
// 	Axios = axios.create({
// 	  baseURL: 'http://app.kasoly.com/v1', 
// 	  headers: {'token': token},
// 	})
// })
// .fail(function() {
// 	console.log("error");
// })
// .always(function() {
// 	console.log("complete");
// });

var Axios = axios.create({
	baseURL: 'http://app.kasoly.com/v1', 
	headers: {'token': 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBpZCI6IjQzMzk5OTc3MjI5MDk2MDc5MzY7b1g4dkt3dWVUSE9DM3dyVWttMmVKQm5tLW02QTvljJfkuqzluII7aHR0cDovL3d4LnFsb2dvLmNuL21tb3Blbi9rcnN0VzVyZlQxZzRDMHBzbFIzVFJOZFF1a3VDT25hUTFxdHF5ZE80UFJGRGliSlBqMDRmcmVpYVR0aWFlNTRsVWFMb0dYVUJ6VFZLSm9BSWliVTV4dkFoRHAzMnhYQ2ozaGJILzA7bG92ZWJvb2siLCJleHAiOjE1MDc5NjU4MjksImlzcyI6IjQzMzk5OTc3MjI5MDk2MDc5MzY7b1g4dkt3dWVUSE9DM3dyVWttMmVKQm5tLW02QTvljJfkuqzluII7aHR0cDovL3d4LnFsb2dvLmNuL21tb3Blbi9rcnN0VzVyZlQxZzRDMHBzbFIzVFJOZFF1a3VDT25hUTFxdHF5ZE80UFJGRGliSlBqMDRmcmVpYVR0aWFlNTRsVWFMb0dYVUJ6VFZLSm9BSWliVTV4dkFoRHAzMnhYQ2ozaGJILzA7bG92ZWJvb2siLCJuYmYiOjE1MDc4Nzk0Mjl9.lBC_6qAb5TF4Vn0OGRPbru9hHxuAKvNNSI0nU5f6-QY'},
})



export default{
	install(Vue) {
	    Object.defineProperty(Vue.prototype, '$axios', { value: Axios })
	}
}





























