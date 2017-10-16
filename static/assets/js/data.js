var editor; // use a global for the submit and return data rendering in the examples
var table; //data_table
var field_name; // table > field : name , not allow to edit
var editor_field; // table > field : name , not allow to edit

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


var table_m_s = function(){
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
    } );
     table = $('#data_table').DataTable( {
        dom: "Bfrtip",
     //    ajax: "/json/data.txt",
        ajax: "/data_student",
        idSrc:  'Uid',
        columns: [
         //    { data: null, render: function ( data, type, row ) {
         //        // Combine the first and last names into a single table field
         //        return data.first_name+' '+data.last_name;
         //    } },
           
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
        ],
        select: true,
        buttons: [
            { extend: "create", editor: editor },
            { extend: "edit",   editor: editor },
            { extend: "remove", editor: editor }
        ]
    } );
}


var table_m_t = function(){
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
    } );
     table = $('#data_table').DataTable( {
        dom: "Bfrtip",
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
        select: true,
        buttons: [
            { extend: "create", editor: editor },
            { extend: "edit",   editor: editor },
            { extend: "remove", editor: editor }
        ]
    } );
}


var table_m_a = function(){
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
    } );
     table = $('#data_table').DataTable( {
        dom: "Bfrtip",
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
        select: true,
        buttons: [
            { extend: "create", editor: editor },
            { extend: "edit",   editor: editor },
            { extend: "remove", editor: editor }
        ]
    } );
}


var table_c = function(){
    //classes
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
        fields: [ {
                label: "Class Name:",
                name: "Class"
            }, {
                label: "Class Day (ex: Mon.):",
                name: "Day"
            }, {
                label: "Class Time(ex: 13:00) :",
                name: "Time_HR"
            }, {
                label: "Duration(HR ex: 2.5):",
                name: "Duration_HR"
            }, 
            {
                label: "Join Date:",
                name: "Open_Date",
                type: "datetime"
            },
            {
                label: "Teacher:",
                name: "Teacher"

            },{
                 label: "Class Type:",
                 name: "Course_Type"
             },{
                label: "Students:",
                name: "Students"
            },{
                label: "Cost (each student):",
                name: "Cost_Each_Student"
            }
        ]
    } );
     table = $('#data_table').DataTable( {
        dom: "Bfrtip",
     //    ajax: "/json/data.txt",
        ajax: "/data_class",
        idSrc:  'Cid',
        columns: [
            { data: "Class" },
            { data: "Day" },
            { data: "Time_HR" },
            { data: "Duration_HR" },
            { data: "Open_Date" ,  render: $.fn.dataTable.render.number( ',', '.', 0, '$' )},
            { data: "Teacher" },
            { data: "Course_Type" },
            { data: "Students" },
            { data: "Cost_Each_Student" }
     
            //{ data: "UpdateTime" }
         //    { data: "salary", render: $.fn.dataTable.render.number( ',', '.', 0, '$' ) }
        ],
        select: true,
        buttons: [
            { extend: "create", editor: editor },
            { extend: "edit",   editor: editor },
            { extend: "remove", editor: editor }
        ]
    } );
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
        
    default:
    break;
}

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
})

$(".buttons-create").click(function(){
    editor_field = editor.field(field_name);
    editor_field.enable();
})


  

  
} );