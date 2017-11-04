var editor; // use a global for the submit and return data rendering in the examples
var table; //data_table
var field_name; // table > field : name , not allow to edit
var editor_field; // table > field : name , not allow to edit
var url_post_member_teacher = "/post_data_teacher"
var url_edit_member_teacher = "/post_data_teacher_edit"
var url_delete_member_teacher = "/post_data_teacher_remove"
var url_member_teacher = "/members/teachers"

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
}




var get_click_rows_data = function(){
    return table.rows('.selected').data()
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
    }
    console.log(JSON.stringify(rm))
    save_todo( JSON.stringify(rm), url_delete_member_teacher, "PUT")
    
}

//close windows
var close_modal_dialog = function(){
    $('span[aria-hidden="true"]').click()
}

//close windows
var clear_form = function(){
    let pathName = window.location.pathname 
    if(pathName == url_member_teacher){
        
                $("#input_teacher").val("")
                $("#input_phone").val("")
                $("#input_cellphone").val("")
                $("#input_email").val("")
                $("#classes_output").val("")
                $("#join_date").val("")
                $("#input_salary").val("")
            
    }
    //table refesh
    setTimeout( function () {
        table.ajax.reload();
        close_modal_dialog()
    }, 1000 );
}

//click and submit form. step 2
var save_todo = function(string_json_data, url, http_type,  callback){
    $.ajax({
        type: http_type,
        url: url,
        dataType: "Json",
        data: string_json_data,
        success: function(result){
            console.log(result)
            let str_res = result.toString()
            console.log(str_res)
            if(str_res.split("warn").length < 2 && str_res.split("error").length < 2)
                return callback()
        },error: function(err){
            console.log(err)
        }
        
    })
}
//click and submit form. step 1
var save_pre = function(){
    let pathName = window.location.pathname 
    if(pathName == url_member_teacher){
      
        let input_teacher = $("#input_teacher").val()
        let input_phone = $("#input_phone").val()
        let input_cellphone = $("#input_cellphone").val()
        let input_email = $("#input_email").val()
        let classes_output = $("#classes_output").val()
        let join_date = $("#join_date").val()
        let input_salary = $("#input_salary").val()
        let textarea_note = $("#textarea_note").val()
        
        if(check_null(input_teacher)) return fill_warning("teacher name", null)
        if(check_null(input_phone)) return fill_warning("phone" , null)
        if(!check_digits(input_phone)) return fill_warning("phone", "no_digits")
        if(check_null(input_cellphone)) return fill_warning("cellphone", null)
        if(!check_digits(input_cellphone)) return fill_warning("cellphone", "no_digits")
        if(check_null(input_email)) return fill_warning("mail", null)
        if(!check_Email(input_email)) return fill_warning("mail", "no_email")
        if(check_null(classes_output)) return fill_warning("classes", null)
        if(check_null(join_date)) return fill_warning("join date", null)
        if(check_null(input_salary)) return fill_warning("salary", null)
        data = JSON.stringify({
            name: input_teacher,
            phone_home: input_phone,
            phone_cell: input_cellphone,
            email:  input_email,
            identity: "Teacher",
            open_date:join_date,
            classes:classes_output,
            salary_hr:input_salary,
            note:textarea_note
        })
        if(button_click_name == "update")
            return  save_todo(data, url_edit_member_teacher, "PUT", clear_form)
        else if(button_click_name == "add_new")
            return  save_todo(data, url_post_member_teacher, "POST", clear_form)
    }
}

var table_m_s = function(){
    /*
    editor = new $.fn.dataTable.Editor( {
        ajax: {
            create: $.extend( true, {}, ajaxBase, {
                type:"POST",
                url: "/post_data_student"
            }),
            edit: $.extend( true, {}, ajaxBase,{
                type: "PUT",
                url:  "/post_data_student_edit"
            }),
            remove: $.extend( true, {}, ajaxBase,{
                type: "PUT",
                url:  "/post_data_student_remove"
            })
       },
        //ajax: "/post_data_student",
        table: "#data_table",
        idSrc:  'Uid',
        
        fields: [ {
                label: "Name:",
                name: "Name"
            }, {
                label: "Phone (Home):",
                name: "Phone_Home"
            }, {
                label: "Phone (CellPHONE):",
                name: "Phone_Cell"
            }, {
                label: "EMAIL:",
                name: "EMAIL"
            }, {
                label: "Classes:",
                name: "Classes"
            }, {
                label: "Join Date:",
                name: "Open_Date",
                type: "datetime"
            },
             {
                 label: "Parents Contect:",
                 name: "Student_Parents_Contect"
             },{
                 label: "Note:",
                 name: "Note"
             }
            
        ]
    } );*/
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
     
            //{ data: "UpdateTime" }
         //    { data: "salary", render: $.fn.dataTable.render.number( ',', '.', 0, '$' ) }
        ]
        //select: true
        /*buttons: [
            { extend: "create", editor: editor },
            { extend: "edit",   editor: editor },
            { extend: "remove", editor: editor }
        ]*/
    } );
}


var table_m_t = function(){
    /*
    editor = new $.fn.dataTable.Editor( {
       // ajax: "/json/staff.php",
        ajax: {
                create: $.extend( true, {}, ajaxBase, {
                    type:"POST",
                    url: "/post_data_teacher"
                }),
                edit: $.extend( true, {}, ajaxBase,{
                    type: "PUT",
                    url:  "/post_data_teacher_edit"
                }),
                remove: $.extend( true, {}, ajaxBase,{
                    type: "PUT",
                    url:  "/post_data_teacher_remove"
                })
        },
        table: "#data_table",
        idSrc:  'Uid',
        fields: [ {
                label: "Name:",
                name: "Name"
            }, {
                label: "Phone (Home):",
                name: "Phone_Home"
            }, {
                label: "Phone (CellPHONE):",
                name: "Phone_Cell"
            }, {
                label: "EMAIL:",
                name: "EMAIL"
            }, {
                label: "Classes:",
                name: "Classes"
            },
         //    }, {
         //        label: "Start date:",
         //        name: "start_date",
         //        type: "datetime"
         //    }, 
            {
                label: "Join Date:",
                name: "Open_Date",
                type: "datetime"
            },
             {
                 label: "Salary (HR):",
                 name: "Salary_HR"
 
             },{
                 label: "Note:",
                 name: "Note"
             }
             // ,{
             //     label: "UpdateTime:",
             //     name: "UpdateTime"
             // }
        ]
    } );*/
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
      //  select: true,
      /*  buttons: [
            { extend: "create", editor: editor },
            { extend: "edit",   editor: editor },
            { extend: "remove", editor: editor }
        ]*/
    } );
}


var table_m_a = function(){
    /*
    editor = new $.fn.dataTable.Editor( {
        //ajax: "/json/staff.php",
        ajax: {
            create: $.extend( true, {}, ajaxBase, {
                type:"POST",
                url: "/post_data_assistant"
            }),
            edit: $.extend( true, {}, ajaxBase,{
                type: "PUT",
                url:  "/post_data_assistant_edit"
            }),
            remove: $.extend( true, {}, ajaxBase,{
                type: "PUT",
                url:  "/post_data_assistant_remove"
            })
       },
        table: "#data_table",
        idSrc:  'Uid',
        fields: [ {
                label: "Name:",
                name: "Name"
            }, {
                label: "Phone (Home):",
                name: "Phone_Home"
            }, {
                label: "Phone (CellPHONE):",
                name: "Phone_Cell"
            }, {
                label: "EMAIL:",
                name: "EMAIL"
            }, 
            {
                label: "Join Date:",
                name: "Open_Date",
                type: "datetime"
            },
            {
                label: "Salary (HR):",
                name: "Salary_HR"

            },{
                 label: "Note:",
                 name: "Note"
             }
        ]
    } );*/
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
      //  select: true,
       /* buttons: [
            { extend: "create", editor: editor },
            { extend: "edit",   editor: editor },
            { extend: "remove", editor: editor }
        ]*/
    } );
}


var table_c = function(){
    //classes
    /*
    editor = new $.fn.dataTable.Editor( {
        //ajax: "/json/staff.php",
        ajax: {
            create: $.extend( true, {}, ajaxBase, {
                type:"POST",
                url: "/post_data_class"
            }),
            edit: $.extend( true, {}, ajaxBase,{
                type: "PUT",
                url:  "/post_data_class_edit"
            }),
            remove: $.extend( true, {}, ajaxBase,{
                type: "PUT",
                url:  "/post_data_class_remove"
            })
       },
        table: "#data_table",
        idSrc:  'Cid',
        fields: [{
                label: "Class Name (full name) ",
                name: "Class_Name"
            }, {
                label: "Class Name（unique code name)",
                name: "Class",
                def: "must input this field, ex: ENG12"
            }, {
                label: "Class Day",
                name: "Day",
                type: "select",
                options: [
                    { label: "Monday", value: "Monday" },
                    { label: "Tuesday", value: "Tuesday" },
                    { label: "Wednesday", value: "Wednesday" },
                    { label: "Thursday", value: "Thursday" },
                    { label: "Friday", value: "Friday" },
                    { label: "Saturday", value: "Saturday" },
                    { label: "Sunday", value: "Sunday" }
                ]
            }, {
                label: "Class Time",
                name: "Time_HR",
                className: "timepicker" 
            }, {
                label: "Duration(HR ex: 2.5)",
                name: "Duration_HR",
                type: "select",
                options: [
                    { label: "1 Hour", value: "1" },
                    { label: "1.5 Hours", value: "1.5" },
                    { label: "2 Hours", value: "2" },
                    { label: "2.5 Hours", value: "2.5" },
                    { label: "3 Hours", value: "3" },
                    { label: "3.5 Hours", value: "3.5" },
                    { label: "4 Hours", value: "4" },
                    { label: "4.5 Hours", value: "4.5" },
                    { label: "5 Hours", value: "5" },
                    { label: "5.5 Hours", value: "5.5" },
                    { label: "6 Hours", value: "6" },
                    { label: "6.5 Hours", value: "6.5" },
                    { label: "7 Hours", value: "7" },
                    { label: "7.5 Hours", value: "7" },
                    { label: "8 Hours", value: "8" }
                   
                ]
            }, 
            {
                label: "Join Date",
                name: "Open_Date",
                type: "datetime"
            },
            {
                label: "Teacher",
                name: "Teacher"

            },{
                 label: "Class Type",
                 name: "Course_Type",
                 def: "English, Mathematics or Japanese"
                
             },{
                label: "Students",
                name: "Students"
            },{
                label: "Cost (each student)",
                name: "Cost_Each_Student_Str"
            },{
                label: "Charge Times: (ex: 5 )",
                name: "Charge_Times",
                def: 8
                
            }
        ]
    } );*/
     table = $('#data_table').DataTable( {
      //  dom: "Bfrtip",
     //    ajax: "/json/data.txt",
        ajax: "/data_class",
        idSrc:  'Cid',
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
        /*
        select: true,
        buttons: [
            { extend: "create", editor: editor },
            { extend: "edit",   editor: editor },
            { extend: "remove", editor: editor }
        ]*/
    } );
}

var Name_List = ""
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
            Name_List = JSON.parse(jqxhr.responseText).data
            console.log( JSON.parse(jqxhr.responseText).data);
            
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

   

let page_now = window.location.pathname;
$(".nav-item").removeClass("active");
switch (page_now){
    case "/members/students":
        table_m_s()
        $(".members").addClass("active");
        field_name = "Name"
        break;
    case "/members/teachers":
        table_m_t()
        $(".members").addClass("active");
        field_name = "Name"
        break;
    case "/members/assistants":
        table_m_a()
        $(".members").addClass("active");
        field_name = "Name"
        break;
    case "/classes":
        table_c()
        $(".classes").addClass("active");
        field_name = "Class"
        break;
    case "/accounts":
    /*** */
        table_a()
        $(".accounts").addClass("active");
        field_name = "Class"
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
    
    //add new (form function)
     //teacher
    $("#select_classes").click(function(){
        let val_classes = $("#classes_output").val()
        let new_class = $(this).val()
        exist = val_classes.split(new_class)
        if (exist.length > 1){
            //remove existed option
            let exist_con = val_classes.split(new_class+",")
            if (exist_con.length > 1){
                $("#classes_output").val(exist_con[0] +  exist_con[1] )
            }else{
                let exist_con = val_classes.split("," + new_class)
                if (exist_con.length > 1){
                    $("#classes_output").val(exist_con[0])
                }else{
                    $("#classes_output").val("")
                }
            }
        }else{
            //add new option
            if (val_classes == "") val_classes = ""
            else val_classes = val_classes + ","
            $("#classes_output").val(val_classes +  new_class)
        }
        
    })

    $('button[type="button"]').click( function (domObj) {
        //console.log($this.attr("class"))
        console.log($(this).attr("class"))
        let className = $(this).attr("class")
        let pathName = window.location.pathname
       
        selected_row_num = table.rows('.selected').data().length
        if (className.split("update").length > 1   ){
            button_click_name = "update"
            assign_update_form()
        }else if(className.split("delete").length > 1){
            button_click_name = "delete"
            delete_records()
        }else if(className.split("add_new").length > 1){
            button_click_name = "add_new"
           
            
        }else if(className.split("save").length > 1){
            
            save_pre()
            
        }
            console.log(5)
        console.log( table.rows('.selected').data().length +' row(s) selected' );
    } );


    
},1400)


// $(".nav-item a").on("click", function(){
//     alert($(this).parent().html())
//   $(".nav-item").find(".active").removeClass("active");
//   $(this).parent().addClass("active");
// });

$(".buttons-edit").click(function(){
    editor_field = editor.field(field_name);
    editor_field.disable();
    check_select_record_num = editor.field(field_name).s.multiValues
    check_select_record_num = Object.keys(check_select_record_num).length
    if(check_select_record_num > 1){
        $(".DTED_Lightbox_Close").click()
        alert("Edit function just allows one record.")
    }
    switch (page_now){
        case "/classes":
            $('#DTE_Field_Time_HR').timepicker();
            break;
    }
   
})

$(".buttons-create").click(function(){
    editor_field = editor.field(field_name);
    editor_field.enable();
    switch (page_now){
        case "/classes":
            $('#DTE_Field_Time_HR').timepicker();
            get_name_list("student")
            setTimeout( function () {
                $("#DTE_Field_Students").val(Name_List)
            }, 2000 );
           
            break;
    }
})



  
} );