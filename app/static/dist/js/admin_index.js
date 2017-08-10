var contentHtml = '<div class="col-sm-9 col-sm-offset-3 col-md-8 col-md-offset-2 main">';
 contentHtml += '<h2 class="sub-header">list</h2>'
    +'<div class="table-responsive">'
    +'<table id="table-admin-index" class="table table-striped">'
    +'</table>'
    +'</div>'
    +'<div class="table-responsive">'
    +'<table id="table-account" class="table table-striped">'
    +'</table>'
    +'</div>'
    +'</div>';


$(".content-wrapper").append(contentHtml);


var table  = $("#table-admin-index").DataTable({
    "processing": true,
    "serverSide": true,
    "searching": false,
    "ordering": false,
    "bLengthChange": false, //去掉每页显示多少条数据方法
    "iDisplayLength": 10,
    "stateSave": false,
    "stripeClasses": [ 'strip1', 'strip2', 'strip3' ],//斑马线
    //不管表格中有没有数据都固定300的高度，具体是啥单位，大家自己试试啊 ：）
    //"scrollY": 400,
    "scrollCollapse": true,
    //当处理大数据时，延迟渲染数据，有效提高Datatables处理能力
    "destory": true,
    "ajax": {
        "url": '/admin/user/index.html',
        "type": "post",  // method  , by default get
        "error": function () {  // error handling

        },
        "data": function (d) {

        }
    },
    "language": {//对表格国际化
        "sLengthMenu": "每页显示 _MENU_条",
        "sZeroRecords": "没有找到符合条件的数据",
        "sInfo": "当前第 _START_ - _END_ 条　共计 _TOTAL_ 条",
        "sInfoEmpty": "木有记录",
        "sInfoFiltered": "(从 _MAX_ 条记录中过滤)",
        "sSearch": "搜索：",
        "sEmptyTable": "表中数据为空",
        "sLoadingRecords": "载入中...",
        "sInfoThousands": ",",
        "oPaginate": {
            "sFirst": "首页",
            "sPrevious": "上页",
            "sNext": "下页",
            "sLast": "末页"
        },
        "oAria": {
            "sSortAscending": ": 以升序排列此列",
            "sSortDescending": ": 以降序排列此列"
        }
    },
    "columns": [
        {"data":null,title:"排序","bSortable": false},
        {"data": "userid", title:"用户编号"},
        {"data": "openid", title:"OpenID"},
        {"data": "nickname", title:"用户昵称"},
        {"data": "telphone", title:"手机号"},
       /* {"data": "status",title:"是否开奖",render: function(data, type, row, meta) {
            if(data=="2"){
                return "已开奖";
            }else{
                return "未开奖";
            }
        }},*/
        {"data": "created_at",title:"注册时间",render: function(data, type, row, meta) {
            return getLocalTime(data);
        }},
    ],
    "columnDefs": [],
    "order": [[1, 'asc']],//排序按第二列为准，第一列是排序
    "headerCallback": function( thead, data, start, end, display ) {
        //可以分别打印 thead, data, start, end, display 看看究竟是什么
        //$(thead).find('th').eq(0).html( '显示 '+(end-start)+' 条记录' );
        //$(thead).find('th').eq(0).html( '排序' );
        //console.log(data);
        //console.log(display);
    },
});

//添加序号
//参考：http://datatables.club/blog/2016/07/10/add-index-for-table.html
//不管是排序，还是分页，还是搜索最后都会重画，这里监听draw事件即可
table.on('draw',function() {
    table.column(0, {
        search: 'applied',
        order: 'applied'
    }).nodes().each(function(cell, i) {
        //i 从0开始，所以这里先加1
        i = i+1;
        //服务器模式下获取分页信息，使用 DT 提供的 API 直接获取分页信息
        var page = table.page.info();
        //当前第几页，从0开始
        var pageno = page.page;
        //每页数据
        var length = page.length;
        //行号等于 页数*每页数据长度+行号
        var columnIndex = (i+pageno*length);
        cell.innerHTML = columnIndex;
    });
});

//timestap to date 2017/6/29 下午2:45
function getLocalTime(nS) {
    //var d = new Date( data * 1000 );
    //return d.getFullYear() +'-'+ (d.getMonth()+1) +'-'+ d.getDate()+' '+ d.getHours()+':'+ d.getMinutes()+':'+ d.getSeconds();
    return new Date(parseInt(nS) * 1000).toLocaleString().replace(/:\d{1,2}$/,' ');
}
