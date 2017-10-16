package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/fatih/structs"
	"github.com/kataras/iris"
)

func writePathHandler(ctx iris.Context) {
	/*
		get/post data router
	*/
	//PrintFields()
	page_now := ctx.Path()
	//fmt.Println(page_now)
	switch page_now {
	case "/data_student":
		res := Data_find_Person_list("Student")
		ctx.Writef("{\"data\": ")
		ctx.JSON(res)
		ctx.Writef("}")
	case "/data_teacher":
		res := Data_find_Person_list("Teacher")
		ctx.Writef("{\"data\": ")
		ctx.JSON(res)
		ctx.Writef("}")
	case "/data_assistant":
		res := Data_find_Person_list("Assistant")
		ctx.Writef("{\"data\": ")
		ctx.JSON(res)
		ctx.Writef("}")
	case "/data_class":
		fmt.Println("data_class")
		res := Data_find_Class_list()
		ctx.Writef("{\"data\": ")
		ctx.JSON(res)
		fmt.Println(res)
		ctx.Writef("}")
	default:
		ctx.Writef("")
	}

	//ctx.Writef("Hello from %s.", ctx.Path())
	//ctx.Header().Set("Content-Type", "application/json; charset=utf-8")
	//j, _ := json.Marshal(res)
	//w.Write(j)

	//fmt.Println(res)

}

func postHandler(ctx iris.Context) {

	page_now := ctx.Path()
	c := &Person{}
	c_multi := []Person{}
	identity := "Student"
	//fmt.Println(page_now)
	switch page_now {
	case "/post_data_student", "/post_data_student_edit":
		c = &Person{}
		identity = "Student"
	case "/post_data_student_remove":
		c_multi = []Person{}
		identity = "Student"
	case "/post_data_teacher", "/post_data_teacher_edit":
		c = &Person{}
		identity = "Teacher"
	case "/post_data_teacher_remove":

		c_multi = []Person{}
		identity = "Teacher"
		fmt.Println(identity)
		fmt.Println(identity)
	case "/post_data_assistant", "/post_data_assistant_edit":
		c = &Person{}
		identity = "Assistant"
	case "/post_data_assistant_remove":
		c_multi = []Person{}
		identity = "Assistant"
	default:
		ctx.Writef("fail")
	}

	//c := &Person{}
	//post jason data and read it
	fmt.Println(&c_multi)
	if strings.Contains(page_now, "remove") {
		if err := ctx.ReadJSON(&c_multi); err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.WriteString(err.Error())
			ctx.Writef("fail")
		}
	} else {
		if err := ctx.ReadJSON(c); err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.WriteString(err.Error())
			ctx.Writef("fail")
		}
	}
	// fmt.Println(c_multi)
	fmt.Println("identity")
	fmt.Println(identity)
	res := ""
	if strings.Contains(page_now, "edit") {
		res = Data_update_Person(*c, identity)
	} else if strings.Contains(page_now, "remove") {
		res = Data_remove_Person(c_multi, identity)
	} else if strings.Contains(page_now, "data") {
		res = Data_insert_Person(*c, identity)
	} else {
		ctx.Writef("fail")
	}
	// switch page_now {
	// case "/post_data_student":
	// 	res = Data_insert_Person(*c, identity)
	// case "/post_data_student_edit":
	// 	res = Data_update_Person(*c, identity)
	// case "/post_data_student_remove":
	// 	res = Data_remove_Person(c_multi, identity)

	// default:
	// 	ctx.Writef("fail")
	// }

	//res := Data_insert_Person(*c, identity)
	fmt.Println(page_now)
	if res != "success" {
		ctx.Writef(string("{ \"error\": \"" + res + "\" }"))
	} else {
		if strings.Contains(page_now, "remove") {
			fmt.Println("=======remove")
			ctx.WriteString("OK")
			//ctx.Writef("")
		} else {
			m := structs.Map(*c)
			b, err := json.Marshal(m)
			fmt.Println("=======m")
			fmt.Println(m)
			fmt.Println("=======b")
			fmt.Println(b)
			if err != nil {
				fmt.Println(err.Error())
				fmt.Println("=======error")
				ctx.Writef("fail")
			}
			ctx.Writef(string(b))
		}

	}

}
