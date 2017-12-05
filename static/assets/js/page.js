var now_date = function(){
    let n = new Date();
    let y = n.getFullYear();
    let m = n.getMonth()+1;;
    let d = n.getDate();
    let dd = n.getDay();
    if(d<10) {
        d = '0'+d
    } 
    
    if(m<10) {
        m = '0'+m
    } 
    return y + "-" + m + "-" + d + "(" + number_date_formate(dd) + ")"
}

var number_date_formate = function(day){
    switch (day){
        case 1:
            return "星期一"
            break;  
        case 2:
            return "星期二"
            break;
        case 3:
            return "星期三"
            break;
        case 4:
            return "星期四"
            break;
        case 5:
            return "星期五"
            break;
        case 6:
            return "星期六"
            break;
        case 0:
            return "星期日"
            break;
        default:
            return "ERROR"

    }
}

$(".date_now").ready(function(){
    $(".date_now").html(now_date());
})
