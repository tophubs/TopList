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