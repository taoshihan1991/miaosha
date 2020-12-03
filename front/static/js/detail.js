function fix(num, length) {
    console.log(num);
    // 数字转化为字符串  进行拼接
    return num.toString().length<length?'0'+num:num;
}

$(function () {
    $.ajax({
        type: "GET",
        url: api+"/product?id=1",
        success: function(res) {
            var data=res.data;
            $("#title").html(data.title);
            $("#price").html(data.price);
            $.leftTime(countTime,function(d){
                if(d.status){
                    var $dateShow1=$(".count-time");
                    $dateShow1.find(".d").html(d.d);
                    $dateShow1.find(".h").html(d.h);
                    $dateShow1.find(".m").html(d.m);
                    $dateShow1.find(".s").html(d.s);
                }
            });
        }
    });

})
