


$(function () {
    $.ajax({
        type: "GET",
        url: api+"/product?id=1",
        success: function(res) {
            var product=res.data.product;
            $("#title").html(product.title);
            $("#price").html(product.price);
            $("#storge").html(product.storge);
            if (typeof product.saletime=="undefined" || product.saletime==""){
                buytime=nowTimestamp+3600*24*1000;
            }else{
                buytime=parseInt(product.saletime);
            }
            countDown(buytime);
        }
    });
    function countDown(buytime) {
        var timer = setInterval(function() {
            var NowTime = new Date(nowTimestamp);
            var EndTime = new Date(buytime);
            if(EndTime < NowTime) {
                EndTime = NowTime;
            }
            var t = EndTime - NowTime;

            var d = 0;
            var h = 0;
            var m = 0;
            var s = 0;
            if(t >= 0) {
                d = Math.floor((t / 1000 / 3600) / 24);
                h = Math.floor((t / 1000 / 3600) % 24);
                m = Math.floor((t / 1000 / 60) % 60);
                s = Math.floor(t / 1000 % 60);
                if(d < 10) {
                    d = "0" + d;
                }
                if(h < 10) {
                    h = "0" + h;
                }
                if(m < 10) {
                    m = "0" + m;
                }
                if(s < 10) {
                    s = "0" + s;
                }
                $('#d').html(d);
                $('#h').html(h);
                $('#m').html(m);
                $('#s').html(s);
            }
            if(t < 1000) {
                //window.location.reload();
                clearInterval(timer);
                timer = null;
                $(".buyBtn a").removeClass("btn-disabled");
                $(".buyBtn a").addClass("btn-primary");
                return false;
            }
            nowTimestamp=nowTimestamp+1000;
        }, 1000)
    }
    $(".buyBtn").on("click"," .btn-primary",function () {
        $(".buyBtn a").removeClass("btn-primary");
        $(".buyBtn a").addClass("btn-disabled");
        $.ajax({
            type: "GET",
            url: api+"/buy?id=1",
            success: function(res) {
                if(res.code==200){
                    url=res.data.url;
                    $.ajax({
                        type: "GET",
                        url: api+url,
                        success: function(res) {
                            $(".buyBtn a").removeClass("btn-disabled");
                            $(".buyBtn a").addClass("btn-primary");
                            if(res.code==200){
                                alert(res.msg);
                            }else{
                                alert(res.msg);
                            }
                        }
                    });
                }else{
                    alert(res.msg);
                }
            }
        });
    });
})
