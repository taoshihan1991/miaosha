$.leftTime("2021-01-01 00:00:00",function(d){
    if(d.status){
        var $dateShow1=$(".count-time");
        $dateShow1.find(".d").html(d.d);
        $dateShow1.find(".h").html(d.h);
        $dateShow1.find(".m").html(d.m);
        $dateShow1.find(".s").html(d.s);
        //d.status 状态
        //d.d 天
        //d.h 时
        //d.m 分
        //d.s 秒
    }
});