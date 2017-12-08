var editor; // use a global for the submit and return data rendering in the examples
var table; //data_table
var not_allow_edit_field; // table > field : name , not allow to edit
var editor_field; // table > field : name , not allow to edit
var page_now = window.location.pathname;
var form_field_list = [] //page form field

var url_post_data_student = "/post_data_student"
var url_post_data_student_edit = "/post_data_student_edit"
var url_post_data_student_remove = "/post_data_student_remove"
var url_member_student = "/members/students"

var url_post_data_assistant = "/post_data_assistant"
var url_post_data_assistant_edit = "/post_data_assistant_edit"
var url_post_data_assistant_remove = "/post_data_assistant_remove"
var url_member_assistant = "/members/assistants"

var url_post_member_teacher = "/post_data_teacher"
var url_edit_member_teacher = "/post_data_teacher_edit"
var url_delete_member_teacher = "/post_data_teacher_remove"
var url_member_teacher = "/members/teachers"

var url_account_teacher = "/data_account_teacher"
var url_post_account_teacher = "/post_account_teacher"
var url_post_account_teacher_edit = "/post_account_teacher_edit"
var url_post_account_teacher_remove = "/post_account_teacher_remove"

var url_account_student = "/data_account_student"
var url_post_account_student = "/post_account_student"
var url_post_account_student_edit = "/post_account_student_edit"
var url_post_account_student_remove = "/post_account_student_remove"

var url_account_assistant = "/data_account_assistant"
var url_post_account_assistant = "/post_account_assistant"
var url_post_account_assistant_edit = "/post_account_assistant_edit"
var url_post_account_assistant_remove = "/post_account_assistant_remove"

var url_notice = "/data_notice"
var url_post_notice = "/post_notice"
var url_post_notice_edit = "/post_notice_edit"
var url_post_notice_remove_remove = "/post_notice_remove"

var url_login = "/login"
var url_login_check ="/login_check"
var url_post_registry = url_post_data_assistant


var List_teacher;
var List_student;
var List_assistant;
var List_class;


//var selected_datas = ''

var button_click_name = ""

var ajaxBase = {
    dataSrc: "",
    dataType: "json",
    contentType: "application/json; charset=utf-8",
   // processData: "false",
   // headers: { "X-CSRF-TOKEN": token },
    data:  function ( d ) {
        console.log(d)
        console.log(Object.keys(d.data))
        if (d.action == "edit")  {
            let index = Object.keys(d.data)[0]
            return JSON.stringify(d.data[index])
        }else if (d.action == "remove"){
            let rm = []
            let index = Object.keys(d.data)
            console.log(index)
            let j = 0
            for (i in index){
                console.log(index[i])
                rm[j] = d.data[index[i]]
                console.log(d.data[index[i]])
                j += 1
            }
            console.log(JSON.stringify(rm))
            return JSON.stringify(rm)
        }else{
          return JSON.stringify(d.data[0])
        }
    },
    success: function(result){
        setTimeout( function () {
            table.ajax.reload();
        }, 1000 );
    },
    error: function(result){
        console.log(result)
        let respon = result.responseText.toString()
        console.log(respon.search("OK"))
        if(respon.search("OK") >=  0 ){
            $(".DTED_Lightbox_Close").click()
            setTimeout( function () {
                table.ajax.reload();
            }, 1000 );
        }else{

        } 
        
    }

};

var login_success_direct = function(){
    window.location.replace("/")
}

//===================check format=================
var check_null = function(str_val){
    if(str_val  == undefined || str_val == "" || str_val == null){
        return true
    }else{
        return false
    }
}
var check_digits = function(str_val){
    if(str_val.match(/^[0-9]+$/) != null){
        return true
    }else{
        return false
    }
}
var check_Email = function(email) {
    var re = /\S+@\S+\.\S+/;
    return re.test(email);
}

//==============format warning=================
var fill_warning = function(col_name, type){
    if(type == null)
    alert("Please fill out the " + col_name + " to this form!")
    else if(type == "no_digits")
    alert(  "Please fill number to " +col_name +" .")
    else if(type == "no_email")
    alert(  "Please fill the right e-mail.")
    else if(type == "no_same_pass")
    alert(  "Password and Confirm password are not the same.")
}



//get selected records
var get_click_rows_data = function(){
    return table.rows('.selected').data()
}
//get selected record names
var get_selected_names = function(){
    selected_datas = get_click_rows_data();
    console.log(selected_datas)
    len_datas = selected_datas.length
    str_names = ""
    for (i=0; i < len_datas; i++){
        str_names += selected_datas[i]['Name'] + ", "
    }
    return str_names.substring(0, str_names.length -2) 
}

//pre update
var assign_update_form = function(){
    let data = get_click_rows_data()
    
    let pathName = window.location.pathname 
    if(pathName == url_member_teacher){
        $("#input_teacher").val(data[0]['Name'])
        $("#input_phone").val(data[0]['Phone_Home'])
        $("#input_cellphone").val(data[0]['Phone_Cell'])
        $("#input_email").val(data[0]['EMAIL'])
        $("#classes_output").val(data[0]['Classes'])
        $("#join_date").val(data[0]['Open_Date'])
        $("#input_salary").val(data[0]['Salary_HR'])
        $("#textarea_note").val(data[0]['Note'])
        
    }
}

var delete_records = function(){
    let data = get_click_rows_data()
    let rm =  []
    let pathName = window.location.pathname 
    if(pathName == url_member_teacher){

        for (let i=0 ; i < data.length; i++){
            rm[i] = {
                name: data[i].Name,
                identity: data[i].Identity,
            }
        }
        console.log(JSON.stringify(rm))
        save_todo( JSON.stringify(rm), url_delete_member_teacher, "PUT")
    }
    
    
}

//close windows
var close_modal_dialog = function(){
    $('span[aria-hidden="true"]').click()
}

//close windows
var clear_form = function(){
    // let pathName = window.location.pathname 
    // if(pathName == url_member_teacher){
        
    //             $("#input_teacher").val("")
    //             $("#input_phone").val("")
    //             $("#input_cellphone").val("")
    //             $("#input_email").val("")
    //             $("#classes_output").val("")
    //             $("#join_date").val("")
    //             $("#input_salary").val("")
            
    // }
    for(let i in form_field_list){
        $("#" + form_field_list[i]).val("")
    }
    //table refesh
   
    setTimeout( function () {
        if(page_now != "/login")
            table.ajax.reload();
        close_modal_dialog()
    }, 1000 );
  
    
}


//===========add update option======
var select_options = function(this_id){
    let output_id = $("#"+this_id).parents("p").next().children()[1].id
    console.log(output_id)
    let val_classes = $("#"+output_id).val()
    let new_class = $("#"+this_id).val()
    console.log(val_classes)
    exist = val_classes.split(new_class)
    if (exist.length > 1){
        //remove existed option
       
        let exist_con = val_classes.split(new_class+",")
        if (exist_con.length > 1){
            $("#"+output_id).val(exist_con[0] +  exist_con[1] )
        }else{
            let exist_con = val_classes.split("," + new_class)
            if (exist_con.length > 1){
                $("#"+output_id).val(exist_con[0])
            }else{
                $("#"+output_id).val("")
            }
        }
    }else{
        //add new option
        if (val_classes == "") val_classes = ""
        else val_classes = val_classes + ","
        $("#"+output_id).val(val_classes +  new_class)
    }

}

setTimeout(function(){
    //
    $("#select_classes").click(function(){
        console.log($(this).attr('id'))
        return select_options($(this).attr('id'))
    })
     //
    

},2000)



//click and submit form. step 2
var save_todo = function(string_json_data, url, http_type,  callback){
    $.ajax({
        type: http_type,
        url: url,
        dataType: "Json",
        data: string_json_data,
        success: function(result){
            console.log("=  save 1===")
            console.log(result)
            let str_res = result.toString()
            console.log("=  save 2===")
            console.log(str_res)
            console.log(result.error)
           
            if(result.error != undefined && result.error != ""){
                alert(result.error)
            }
            if(str_res.split("warn").length < 2 && str_res.split("error").length < 2)
                return callback()
        },error: function(err){
            console.log("=  error 1===")
            console.log(err)
        }
        
    })
}

var schema_mapping = function(field_name){
    var schema_field = {
        "input_teacher" : "name",
        "input_student" : "name",
        "input_assistant" : "name",
        "username" : "name",
        "input_phone" : "phone_home" ,          
        "input_cellphone" : "phone_cell" ,             
        "input_email" : "email" ,  
        "password" : "password",               
        "input_password" : "password",              										  
        "join_date" : "open_date" ,                      
        "input_salary" : "salary_hr" ,             
        "input_parents_contect" : "student_parents_contect", 
        "textarea_note" : "note"
    }
    if (schema_field.hasOwnProperty(field_name))
        return schema_field[field_name]
    else{
        alert("FATAL: schema_mapping")
    }
}

var input_check_and_return_json = function(field_list){
    let len_field_list = field_list.length
    let count = 0
    let res_json = {}
    console.log(field_list)
    for (let i in field_list){
        let field_name = field_list[i]
        let field_val =  $("#"+ field_name).val()
        if (field_name != "input_password_re")
            res_json[schema_mapping(field_name)] = field_val
            alert(field_name)
        switch (field_name){
            
            case "input_assistant":
            case "input_student":
            case "input_teacher":
            case "username":
                if(check_null(field_val)) fill_warning("name", null)
                else count += 1
                break;
            case "password":
            case "input_password_re":
                if(check_null(field_val)) fill_warning("password", null)
                else count += 1
                break;
            case "input_password":
                if(check_null(field_val)) fill_warning("password", null)
                else if(field_val !=  $("#input_password_re").val()) fill_warning("password", "no_same_pass")
                else count += 1
                break;
            case "input_phone":
                if(check_null(field_val)) return fill_warning("phone" , null)
                else if(!check_digits(field_val)) return fill_warning("phone", "no_digits")
                else count += 1
                break;
            case "input_cellphone":
                if(check_null(field_val)) return fill_warning("cellphone", null)
                else if(!check_digits(field_val)) return fill_warning("cellphone", "no_digits")
                else count += 1
                break;
            case "input_email":
                if(check_null(field_val)) return fill_warning("mail", null)
                else if(!check_Email(field_val)) return fill_warning("mail", "no_email")
                else count += 1
                break;
            case "classes_output":
                if(check_null(field_val)) return fill_warning("classes", null)
                else count += 1
                break;
            case "join_date":
                if(check_null(field_val)) return fill_warning("join date", null)
                else count += 1
                break;
            case "input_salary":
                if(check_null(field_val)) return fill_warning("salary", null)
                else count += 1
                break;
            default:
                count += 1
                break;
        }

    }
    if (count == len_field_list){
        
        return  JSON.stringify(res_json)
    }else{
        return "fail"
    }
}


//click and submit form. step 1
var save_pre = function(){
    let pathName = window.location.pathname 
    let url_post = ""
    var field_json_str = ""
    console.log(pathName)
    console.log(url_member_teacher  )
    if(pathName == url_member_teacher){
        form_field_list = ["input_teacher", "input_phone", "input_cellphone", "input_email", "join_date", "textarea_note"]
        field_json_str = input_check_and_return_json(form_field_list)
        console.log(field_json_str)
        if(button_click_name == "update")
            url_post = url_edit_member_teacher
        else if(button_click_name == "add_new")
            url_post = url_post_member_teacher
            console.log(url_post)
        // let input_teacher = $("#input_teacher").val()
        // let input_phone = $("#input_phone").val()
        // let input_cellphone = $("#input_cellphone").val()
        // let input_email = $("#input_email").val()
        // let classes_output = $("#classes_output").val()
        // let join_date = $("#join_date").val()
        // let input_salary = $("#input_salary").val()
        // let textarea_note = $("#textarea_note").val()
        
        // if(check_null(input_teacher)) return fill_warning("teacher name", null)
        // if(check_null(input_phone)) return fill_warning("phone" , null)
        // if(!check_digits(input_phone)) return fill_warning("phone", "no_digits")
        // if(check_null(input_cellphone)) return fill_warning("cellphone", null)
        // if(!check_digits(input_cellphone)) return fill_warning("cellphone", "no_digits")
        // if(check_null(input_email)) return fill_warning("mail", null)
        // if(!check_Email(input_email)) return fill_warning("mail", "no_email")
        // if(check_null(classes_output)) return fill_warning("classes", null)
        // if(check_null(join_date)) return fill_warning("join date", null)
        // if(check_null(input_salary)) return fill_warning("salary", null)
        // data = JSON.stringify({
        //     name: input_teacher,
        //     phone_home: input_phone,
        //     phone_cell: input_cellphone,
        //     email:  input_email,
        //     identity: "Teacher",
        //     open_date:join_date,
        //     classes:classes_output,
        //     salary_hr:input_salary,
        //     note:textarea_note
        // })
       
    }else if(pathName == url_login || pathName == url_member_assistant){
            if(button_click_name == "login"){
                form_field_list = ["username", "password"]
                field_json_str = input_check_and_return_json(form_field_list) 
                // let str_json = JSON.stringify({
                //     "name" : $("#username").val(),
                //     "password" : $("#password").val()
                // })
                console.log(form_field_list)
                console.log(field_json_str)
                return save_todo(field_json_str, url_login_check, "POST", login_success_direct)
            }else{
                form_field_list = ["input_assistant", "input_email", "input_password", "input_password_re", "input_phone", "input_cellphone", "join_date", "input_salary","textarea_note"]
                field_json_str = input_check_and_return_json(form_field_list) 
                if(button_click_name == "update")
                    url_post = url_post_data_assistant_edit
                else if(button_click_name == "add_new")
                    url_post = url_post_data_assistant  

            }
               
    }
    console.log(field_json_str)
    console.log(url_post)
    if(field_json_str != "fail" && field_json_str != "" && url_post != ""){
        console.log(field_json_str)
        if(button_click_name == "update")
            return  save_todo(field_json_str, url_post, "PUT", clear_form)
        else if(button_click_name == "add_new")
            return  save_todo(field_json_str, url_post, "POST", clear_form)
    }else{
        //return alert("Something wrong! Please refill your form!")
    }
}

var table_m_s = function(){
   
     table = $('#data_table').DataTable( {
　　　　  //  dom: "Bfrtip",
         //   ajax: "/json/data.txt",
        ajax: "/data_student",
        idSrc:  'Uid',
        columns: [
       
          
            { data: "Name" },
            { data: "Phone_Home" },
            { data: "Phone_Cell" },
            { data: "EMAIL" },
            { data: "Classes" },
            { data: "Open_Date" },
            { data: "Student_Parents_Contect" },
            { data: "Note" }

        ]
    } );
}


var table_m_t = function(){
     table = $('#data_table').DataTable( {
        //dom: "Bfrtip",
     //    ajax: "/json/data.txt",
        ajax: "/data_teacher",
        idSrc:  'Uid',
        columns: [
            { data: "Name" },
            { data: "Phone_Home" },
            { data: "Phone_Cell" },
            { data: "EMAIL" },
            { data: "Classes" },
            { data: "Open_Date"},
            { data: "Salary_HR" ,  render: $.fn.dataTable.render.number( ',', '.', 0, '$' )},
            { data: "Note" }
        ],
    } );
}


var table_m_a = function(){
    
     table = $('#data_table').DataTable( {
      //  dom: "Bfrtip",
     //    ajax: "/json/data.txt",
        ajax: "/data_assistant",
        idSrc:  'Uid',
        columns: [
            { data: "Name" },
            { data: "Phone_Home" },
            { data: "Phone_Cell" },
            { data: "EMAIL" },
            { data: "Open_Date"},
            { data: "Salary_HR" ,  render: $.fn.dataTable.render.number( ',', '.', 0, '$' )},
            { data: "Note" }
     
            //{ data: "UpdateTime" }
         //    { data: "salary", render: $.fn.dataTable.render.number( ',', '.', 0, '$' ) }
        ],
    } );
}


var table_c = function(){
     table = $('#data_table').DataTable( {
        ajax: "/data_class",

        columns: [
            { data: "Class_Name" },
            { data: "Class" },
            { data: "Day" },
            { data: "Time_HR" },
            { data: "Duration_HR" },
            { data: "Open_Date"  },
            { data: "Teacher" },
            { data: "Course_Type" },
            { data: "Students" },
            { data: "Cost_Each_Student", render: $.fn.dataTable.render.number( ',', '.', 0, '$' ) },
            { data: "Charge_Times", render: $.fn.dataTable.render.number( ',', '.', 0) },
            //{ data: "UpdateTime" }
         //    { data: "salary", render: $.fn.dataTable.render.number( ',', '.', 0, '$' ) }
        ],
    } );
}

var table_a_s = function(){
    table = $('#data_table').DataTable( {
       ajax: url_account_student,
       select: true,
       columns: [
           { data: "Student" },
           { data: "Class" },
           { data: "Date_Start_Study" },
           { data: "Date_Next_Pay" }
       ],
   } );
}
var table_a_t = function(){
    table = $('#data_table').DataTable( {
       ajax: url_account_teacher,
       select: true,
       columns: [
        { data: "Teacher" },
        { data: "Class" },
        { data: "Date_Next_Earn" },
        { data: "HR_Salary" },
        { data: "Earn_Next" }
       ],
   } );
}
var table_a_a = function(){
    table = $('#data_table').DataTable( {
       ajax: url_account_assistant,
       idSrc:  'Cid',
       select: true,
       columns: [
           { data: "Assistant" },
           { data: "HR_Work_Total" },
           { data: "HR_Salary" },
           { data: "HR_Work_No_Gain" } //hr is not this column
       ],
   } );
}

var table_n = function(){
    table = $('#data_table').DataTable( {
       ajax: url_notice,
       select: true,
       columns: [
        { data: "Target" },
        { data: "Title" },
        { data: "Content" },
        { data: "UpdateTime" } 
       ],
   } );
}

//var Name_List = ""
var get_name_list = function(identity){
    let url = ""
    switch (identity){
        case "student":
            url = "/data_student_name_list"
            break;
        case "teacher":
            url = "/data_teacher_name_list"
            break;
        case "assistant":
            url = "/data_assistant_name_list"
            break;
    }
    let jqxhr = $.getJSON( url, function() {
            console.log( "success" );
      })
        .done(function() {
            switch (identity){
                case "student":
                    List_student  = JSON.parse(jqxhr.responseText).data
                    break;
                case "teacher":
                    List_teacher = JSON.parse(jqxhr.responseText).data
                    break;
                case "assistant":
                    List_assistant = JSON.parse(jqxhr.responseText).data
                    break;
                   
            }
            console.log( "second success" );
        })
        .fail(function() {
            console.log( "error" );
        })
        .always(function() {
            console.log( "complete" );
        });
      
        
}

$(document).ready(function() {

   


$(".nav-item").removeClass("active");
switch (page_now){
    case "/members/students":
        table_m_s()
        $(".members").addClass("active");
        not_allow_edit_field = "Name"
        break;
    case "/members/teachers":
        table_m_t()
        $(".members").addClass("active");
        not_allow_edit_field = "Name"
        break;
    case "/members/assistants":
        table_m_a()
        $(".members").addClass("active");
        not_allow_edit_field = "Name"
        break;
    case "/classes":
        table_c()
        $(".classes").addClass("active");
        not_allow_edit_field = "Class"
        break;
    case "/accounts/students":
        table_a_s()
        $(".accounts").addClass("active");
        not_allow_edit_field = "Student"
        break;
    case "/accounts/teachers":
        table_a_t()
        $(".accounts").addClass("active");
        not_allow_edit_field = "Teacher"
        break;
    case "/accounts/assistants":
        table_a_a()
        $(".accounts").addClass("active");
        not_allow_edit_field = "Assistant"
        break;
    case "/notices":
        table_n()
        $(".notices").addClass("active");
        not_allow_edit_field = "Title"
        break;
    case "/calendar":
        $(".notices").addClass("active");
        //not_allow_edit_field = "Class"
        break;

        
    default:
    break;
}

setTimeout(function(){
    //page load watting 1 sec
    $('#data_table tbody')
    .on( 'click', 'tr', function () {
        $(this).toggleClass('selected');
        let selected_num = table.rows('.selected').data().length
        if(selected_num > 1){
            $(".delete").removeClass("btn-outline-secondary")
            $(".delete").addClass("btn-primary")
            $(".update").removeClass("btn-primary")
            $(".update").addClass("btn-outline-secondary")
            $('.delete').prop("disabled", false); 
            $('.update').prop("disabled", true); 
        }else if(selected_num == 1){
            $(".delete").removeClass("btn-outline-secondary")
            $(".delete").addClass("btn-primary")
            $(".update").removeClass("btn-outline-secondary")
            $(".update").addClass("btn-primary")
            $('.delete').prop("disabled", false); 
            $('.update').prop("disabled", false); 
            
        }else if(selected_num == 0){
            $(".update").removeClass("btn-primary")
            $(".update").addClass("btn-outline-secondary")
            $(".delete").removeClass("btn-primary")
            $(".delete").addClass("btn-outline-secondary")
            $('.update').prop("disabled", true); 
            $('.delete').prop("disabled", true); 
        }
    } );
    
    

    $('button[type="button"]').click( function (domObj) {
        //console.log($this.attr("class"))
        let className = $(this).attr("class")
        let pathName = window.location.pathname
        if(page_now != url_login)
            selected_row_num = table.rows('.selected').data().length
        if (className.split("update").length > 1   ){
            button_click_name = "update"
            console.log(1)
            assign_update_form()
        }else if(className.split("delete").length > 1){
            button_click_name = "delete"
            let msg = "Do you want to delete "
            $(".modal_confirm_body").html(msg + get_selected_names() + " ?")
          
           // delete_records()
        }else if(className.split("confirm_btn").length > 1){
            button_click_name = "confirm_btn"
            console.log("button_click_name")
            console.log(button_click_name)
            delete_records()
            clear_form()

        }
        else if(className.split("add_new").length > 1){
            button_click_name = "add_new"
        }else if(className.split("save").length > 1){
            console.log(page_now)
            save_pre()  
        }else if(className.split("login").length > 1){
            button_click_name = "login"
            // console.log(className)
            // let str_json = JSON.stringify({
            //     "name" : $("#username").val(),
            //     "password" : $("#password").val()
            // })
            save_pre() 
            
        }
        
        if(page_now != url_login){
        console.log( className )
        console.log( table.rows('.selected').data() )
        console.log( table.rows('.selected').data().length +' row(s) selected' );
        }
    } );
  
},1400)


// $(".nav-item a").on("click", function(){
//     alert($(this).parent().html())
//   $(".nav-item").find(".active").removeClass("active");
//   $(this).parent().addClass("active");
// });

// $(".buttons-edit").click(function(){
//     editor_field = editor.field(not_allow_edit_field);
//     editor_field.disable();
//     check_select_record_num = editor.field(not_allow_edit_field).s.multiValues
//     check_select_record_num = Object.keys(check_select_record_num).length
//     if(check_select_record_num > 1){
//         $(".DTED_Lightbox_Close").click()
//         alert("Edit function just allows one record.")
//     }
//     switch (page_now){
//         case "/classes":
//             $('#DTE_Field_Time_HR').timepicker();
//             break;
//     }
   
// })

$(".add_new").click(function(){
    // editor_field = editor.field(not_allow_edit_field);
    // editor_field.enable();
    switch (page_now){
        case "/classes":
           // $('#join_date').timepicker();
            console.log("time")
            var res_teachers = ""
            get_name_list("teacher") 
             //====select teacher
            setTimeout( function () {
                let teachers = List_teacher.split(",")
                let str_option = ""
                for (let i in teachers){
                    str_option += `<option value="`+teachers[i]+`">` + teachers[i] + `</option>`
                }
                $("#select_teachers").html(str_option)
                setTimeout( function () {
                    $("#select_teachers").click(function(){
                        return select_options($(this).attr('id'))
                    })
                }, 1000 );
               
            }, 1000 );
             //====select student
            get_name_list("student") 
            setTimeout( function () {
                let students = List_student.split(",")
                let str_option = ""
                for (let i in students){
                    str_option += `<option value="`+students[i]+`">` + students[i] + `</option>`
                }
                $("#select_students").html(str_option)
                setTimeout( function () {
                    $("#select_students").click(function(){
                        return select_options($(this).attr('id'))
                    })
                }, 1000 );
               
            }, 1000 );
            break;
        case "/members/students":
            get_name_list("class") 
            setTimeout( function () {
                let teachers = List_teacher.split(",")
                let str_option = ""
                for (let i in teachers){
                    str_option += `<option value="`+teachers[i]+`">` + teachers[i] + `</option>`
                }
                $("#select_teachers").html(str_option)
                setTimeout( function () {
                    $("#select_teachers").click(function(){
                        return select_options($(this).attr('id'))
                    })
                }, 1000 );
               
            }, 1000 );
            break;
    }
})



  
} );