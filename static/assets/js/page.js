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


// var save_todo = function(string_json_data, url, http_type,  callback){
//     $.ajax({
//         type: http_type,
//         url: url,
//         dataType: "Json",
//         data: string_json_data,
//         success: function(result){
//             console.log("=  save 1===")
//             console.log(result)
//             let str_res = result.toString()
//             console.log("=  save 2===")
//             console.log(str_res)
//             if(str_res.split("warn").length < 2 && str_res.split("error").length < 2)
//                 return callback()
//         },error: function(err){
//             console.log("=  error 1===")
//             console.log(err)
//         }
        
//     })
// }


$(".date_now").ready(function(){
    $(".date_now").html(now_date());
})
