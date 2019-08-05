$(document).ready(function () {
    allArticle();
});

function myTimeTrans(date){
    var date = new Date(date*1000);//如果date为13位不需要乘1000
    var Y = date.getFullYear() + '-';
    var M = (date.getMonth()+1 < 10 ? '0'+(date.getMonth()+1) : date.getMonth()+1) + '-';
    var D = (date.getDate() < 10 ? '0' + (date.getDate()) : date.getDate()) ;
    var h = (date.getHours() < 10 ? '0' + date.getHours() : date.getHours()) + ':';
    var m = (date.getMinutes() <10 ? '0' + date.getMinutes() : date.getMinutes()) + ':';
    var s = (date.getSeconds() <10 ? '0' + date.getSeconds() : date.getSeconds());
    return Y+M+D+" "+h+m+s
}

//绘制单个div
function setDiv(item){
    var content = item.article_content;
    content = content.replaceAll("!"," ");
    content = content.replaceAll("100%","100%;");
    content = content.replaceAll("30%","30%;");
    content = content.replaceAll("50%","50%;");
    var date = myTimeTrans(item.article_create_at);
    var div = '<div class="item-box">\n' +
        '\t\t\t\t\t\t\t\t\t\t<div class="item">\n' +
        '\t\t\t\t\t\t\t\t\t\t\t<div class="whisper-title">\n' +
        '\t\t\t\t\t\t\t\t\t\t\t\t<i class="layui-icon layui-icon-date"></i><span class="hour">'+date+'</span><span class="date"></span>\n' +
        '\t\t\t\t\t\t\t\t\t\t\t</div>\n' +
        '\t\t\t\t\t\t\t\t\t\t\t<p class="text-cont">\n' +
        '\t\t\t\t\t\t\t\t\t\t\t\t'+content+'' +
        '\t\t\t\t\t\t\t\t\t\t\t</p>\n' +
        '\t\t\t\t\t\t\t\t\t\t\t<div class="img-box">\n' +
        '\t\t\t\t\t\t\t\t\t\t\t\t\n' +
        '\t\t\t\t\t\t\t\t\t\t\t</div>\n' +
        '\t\t\t\t\t\t\t\t\t\t\t<div class="op-list">\n' +
        '\t\t\t\t\t\t\t\t\t\t\t\t<p class="like"><i class="layui-icon layui-icon-praise"></i><span>'+item.up+'</span></p>\n' +
        '\t\t\t\t\t\t\t\t\t\t\t\t<p class="edit"><i class="layui-icon layui-icon-reply-fill"></i><span>1200</span></p>\n' +
        '\t\t\t\t\t\t\t\t\t\t\t\t<p class="off"><span>展开</span><i class="layui-icon layui-icon-down"></i></p>\n' +
        '\t\t\t\t\t\t\t\t\t\t\t</div>\n' +
        '\t\t\t\t\t\t\t\t\t\t</div>\n' +
        '\t\t\t\t\t\t\t\t\t\t<div class="review-version layui-hide">\n' +
        '\t\t\t\t\t\t\t\t\t\t\t<div class="form">\n' +
        '\t\t\t\t\t\t\t\t\t\t\t\t<div class="comment a_comment" id="myComment">\n' +
        '\t\t\t\t\t\t\t\t\t\t\t\t\t<h2>发表评论</h2>\n' +
        '\n' +
        '\t\t\t\t\t\t\t\t\t\t\t\t\t<div>\n' +
        '\n' +
        '\t\t\t\t\t\t\t\t\t\t\t\t\t\t<div class="row">\n' +
        '\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t<div class="col-md-6">\n' +
        '\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t<div class="form-group">\n' +
        '\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t<input type="text" class="form-control input-lg" name="name" id="name" placeholder="您的称呼" required="required" value="">\n' +
        '\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t<input type="hidden" class="form-control input-lg" name="comment_id" id="comment_id" value="0">\n' +
        '\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t</div>\n' +
        '\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t</div>\n' +
        '\t\t\t\t\t\t\t\t\t\t\t\t\t\t</div>\n' +
        '\t\t\t\t\t\t\t\t\t\t\t\t\t\t<input type="hidden" value="" name="id" id="articleId">\n' +
        '\t\t\t\t\t\t\t\t\t\t\t\t\t\t<div class="row">\n' +
        '\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t<div class="col-md-6">\n' +
        '\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t<div class="form-group">\n' +
        '\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t<input type="email" class="form-control input-lg" name="email" id="email" placeholder="接收回复的邮箱" required="required" value="">\n' +
        '\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t</div>\n' +
        '\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t</div>\n' +
        '\t\t\t\t\t\t\t\t\t\t\t\t\t\t</div>\n' +
        '\t\t\t\t\t\t\t\t\t\t\t\t\t\t<div class="row">\n' +
        '\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t<div class="col-md-12">\n' +
        '\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t<div class="form-group">\n' +
        '\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t<textarea name="content" id="message" class="form-control" rows="4" cols="25" required="required" placeholder="评论内容"></textarea>\n' +
        '\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t</div>\n' +
        '\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t<button   onclick="addArticleComment()" class="btn btn-4 btn-block save" style="background-color: #f0efee">\n' +
        '\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t立即评论\n' +
        '\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t</button>\n' +
        '\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t</div>\n' +
        '\t\t\t\t\t\t\t\t\t\t\t\t\t\t</div>\n' +
        '\t\t\t\t\t\t\t\t\t\t\t\t\t</div>\n' +
        '\n' +
        '\t\t\t\t\t\t\t\t\t\t\t\t</div>\n' +
        '\t\t\t\t\t\t\t\t\t\t\t</div>\n' +
        '\t\t\t\t\t\t\t\t\t\t\t<div class="list-cont">\n' +
        '\t\t\t\t\t\t\t\t\t\t\t\t<div class="cont">\n' +
        '\t\t\t\t\t\t\t\t\t\t\t\t\t<div class="img">\n' +
        '\t\t\t\t\t\t\t\t\t\t\t\t\t\t<img src="other/res/img/header.png" alt="">\n' +
        '\t\t\t\t\t\t\t\t\t\t\t\t\t</div>\n' +
        '\t\t\t\t\t\t\t\t\t\t\t\t\t<div class="text">\n' +
        '\t\t\t\t\t\t\t\t\t\t\t\t\t\t<p class="tit"><span class="name">Mr.W</span><span class="data">2018/06/06</span></p>\n' +
        '\t\t\t\t\t\t\t\t\t\t\t\t\t\t<p class="ct">敢问大师，师从何方？上古高人呐逐一地看完你的作品后，我的心久久 不能平静！这世间怎么可能还有如此精辟的作品？我不敢相信自己的眼睛。自从改革开放以后，我就以为再也不会有任何作品能打动我，没想到今天看到这个如此精妙绝伦的作品好厉害！</p>\n' +
        '\t\t\t\t\t\t\t\t\t\t\t\t\t</div>\n' +
        '\t\t\t\t\t\t\t\t\t\t\t\t</div>\n' +
        '\t\t\t\t\t\t\t\t\t\t\t</div>\n' +
        '\t\t\t\t\t\t\t\t\t\t</div>\n' +
        '\t\t\t\t\t\t\t\t\t</div>';
    return div
}
//循环加载到页面
function getnoApplicationData(){
    var html = ''
    for(var i = 0;i<data.length;i++){
        html += setDiv(data[i])
    }
    noApplicationRecord.innerHTML = html
}

var getParam = function(name){
    var search = document.location.search;
    var pattern = new RegExp("[?&]"+name+"\=([^&]+)", "g");
    var matcher = pattern.exec(search);
    var items = null;
    if(null != matcher){
        try{
            items = decodeURIComponent(decodeURIComponent(matcher[1]));
        }catch(e){
            try{
                items = decodeURIComponent(matcher[1]);
            }catch(e){
                items = matcher[1];
            }
        }
    }
    return items;
};

function setPageNumber(currentPage,totalPage,type) {
    var lastPage = currentPage - 1
    var nextPage = currentPage + 1
    var addCode = ''
    if (lastPage <= 0) {
        if (type !== null) {
            addCode = '<li class="next"><a href="mine.html?type='+type+'&page='+nextPage+'">&rarr; 下一页</a></li>'
        } else {
            addCode = '<li class="next"><a href="mine.html?page='+nextPage+'">&rarr; 下一页</a></li>'
        }
    } else if (nextPage > totalPage) {
        if (type !== null) {
            addCode = '<li class="previous"><a href="mine.html?type='+type+'&page='+lastPage+'">&larr; 上一页</a></li>'
        } else {
            addCode = '<li class="previous"><a href="mine.html?page='+lastPage+'">&larr; 上一页</a></li>'

        }

    } else {
        if (type !== null) {
            addCode = '<li class="previous"><a href="mine.html?type='+type+'&page='+lastPage+'">&larr; 上一页</a></li>\n' +
                '\t\t\t\t\t\t\t\t<li class="next"><a href="mine.html?type='+type+'&page='+nextPage+'">下一页 &rarr;</a></li>'
        } else {
            addCode = '<li class="previous"><a href="mine.html?page='+lastPage+'">&larr; 上一页</a></li>\n' +
                '\t\t\t\t\t\t\t\t<li class="next"><a href="mine.html?page='+nextPage+'">下一页 &rarr;</a></li>'
        }
    }
    var pageHtml = document.getElementById("isSayPageCode")
    pageHtml.innerHTML = addCode
}


function allArticle() {
    var type = getParam("type")
    if (type !== null&&type !== undefined){
        console.log(type)
    } else {
        type = null
    }
    var noApplicationRecord = document.getElementById('myIsSay')
    $.ajax({
        url:ServerIp+"/GetArticles",
        type:"POST",
        async:true,
        data:{page:function () {
                var a = null;
                if (getParam("page") !== null&&getParam("page") !== undefined){
                    a = getParam("page")
                }else {
                    a = 1
                }
                return a;
            },type:22,isSay:1},
        timeout:5000,
        dataType:'json',
        success:function (data) {
            setPageNumber(data['Data']['currentPage'],data['Data']['totalPage'],type)
            if (data['Code']!=0){
                alert("请重试");
            }else {
                var html = ''
                for (var i in data['Data']['rows']){
                    html += setDiv(data['Data']['rows'][i])
                }
                noApplicationRecord.innerHTML = html
            }
        },
        error:function () {
            console.log("失败");
        }
    })
}




