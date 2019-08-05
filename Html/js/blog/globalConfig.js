const ServerIp = 'https://www.printf520.com:8080'
const ImgServerIp = 'https://www.printf520.com'
//
// const ServerIp = 'http://127.0.0.1:8080'
// const ImgServerIp = 'http://www.printf520.com'


function timetrans(date){
    var date = new Date(date*1000);//如果date为13位不需要乘1000
    var Y = date.getFullYear() + '-';
    var M = (date.getMonth()+1 < 10 ? '0'+(date.getMonth()+1) : date.getMonth()+1) + '-';
    var D = (date.getDate() < 10 ? '0' + (date.getDate()) : date.getDate()) ;
    // var h = (date.getHours() < 10 ? '0' + date.getHours() : date.getHours()) + ':';
    // var m = (date.getMinutes() <10 ? '0' + date.getMinutes() : date.getMinutes()) + ':';
    // var s = (date.getSeconds() <10 ? '0' + date.getSeconds() : date.getSeconds());
    return Y+M+D
}

String.prototype.replaceAll = function(search, replacement) {
    var target = this;
    return target.replace(new RegExp(search, 'g'), replacement);
};

function isMobile() {
    var userAgentInfo = navigator.userAgent;

    var mobileAgents = [ "Android", "iPhone", "SymbianOS", "Windows Phone", "iPad","iPod"];

    var mobile_flag = false;

    //根据userAgent判断是否是手机
    for (var v = 0; v < mobileAgents.length; v++) {
        if (userAgentInfo.indexOf(mobileAgents[v]) > 0) {
            mobile_flag = true;
            break;
        }
    }

    var screen_width = window.screen.width;
    var screen_height = window.screen.height;

    //根据屏幕分辨率判断是否是手机
    if(screen_width < 500 && screen_height < 800){
        mobile_flag = true;
    }

    return mobile_flag;
}

function doShake(o){
    var $panel = $("#"+o);
    box_left = 0;
    //box_left = $panel.css('left');
    //box_left = ($(window).width() -  $panel.width()) / 2;
    $panel.css({'left': box_left,'position':'relative'});
    for(var i=1; 4>=i; i++){
        $panel.animate({left:box_left-(20-5*i)},30);
        $panel.animate({left:box_left+(20-5*i)},30);
    }

}
