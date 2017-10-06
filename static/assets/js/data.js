var editor; // use a global for the submit and return data rendering in the examples

$(document).ready(function() {
   editor = new $.fn.dataTable.Editor( {
       ajax: "/json/staff.php",
       table: "#example",
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
                label: "Parents Contect:",
                name: "Student_Parents_Contect"
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

   $('#example').DataTable( {
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
} );