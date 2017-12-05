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
	page_now := ctx.Path()
	//fmt.Println(page_now)
	switch page_now {
	// case "/send_mail":
	// 	res := send_mail()
	// 	ctx.JSON(res)
	case "/data_student":
		res := person_find_list("Student")
		ctx.Writef("{\"data\": ")
		ctx.JSON(res)
		ctx.Writef("}")
	case "/data_teacher":
		res := person_find_list("Teacher")
		ctx.Writef("{\"data\": ")
		ctx.JSON(res)
		ctx.Writef("}")
	case "/data_assistant":
		res := person_find_list("Assistant")
		ctx.Writef("{\"data\": ")
		ctx.JSON(res)
		ctx.Writef("}")
	case "/data_class":
		fmt.Println("data_class")
		res := class_find_list()
		ctx.Writef("{\"data\": ")
		ctx.JSON(res)
		fmt.Println(res)
		ctx.Writef("}")
	case "/data_student_name_list":
		fmt.Println("data_test")
		res := person_find_Name_List("Student")
		ctx.Writef("{\"data\": ")
		ctx.JSON(res)
		fmt.Println(res)
		ctx.Writef("}")
	case "/data_teacher_name_list":
		fmt.Println("data_test")
		res := person_find_Name_List("Teacher")
		ctx.Writef("{\"data\": ")
		ctx.JSON(res)
		fmt.Println(res)
		ctx.Writef("}")
	case "/data_assistant_name_list":
		fmt.Println("data_test")
		res := person_find_Name_List("Assistant")
		ctx.Writef("{\"data\": ")
		ctx.JSON(res)
		fmt.Println(res)
		ctx.Writef("}")
	case "/data_account_student":
		fmt.Println("data_test")
		res := account_find_student_list()
		ctx.Writef("{\"data\": ")
		ctx.JSON(res)
		fmt.Println(res)
		ctx.Writef("}")
	case "/data_account_teacher":
		fmt.Println("data_test")
		res := account_find_teacher_list()
		ctx.Writef("{\"data\": ")
		ctx.JSON(res)
		fmt.Println(res)
		ctx.Writef("}")
	case "/data_account_assistant":
		fmt.Println("data_test")
		res := account_find_assistant_list()
		ctx.Writef("{\"data\": ")
		ctx.JSON(res)
		fmt.Println(res)
		ctx.Writef("}")
	case "/data_notice":
		fmt.Println("notices")
		res := notice_find_list()
		ctx.Writef("{\"data\": ")
		ctx.JSON(res)
		fmt.Println(res)
		ctx.Writef("}")
	// case "/get_teachers":
	// 	fmt.Println("teachers")
	// 	res := teachers_get()
	// 	ctx.Writef("{\"data\": ")
	// 	ctx.JSON(res)
	// 	fmt.Println(res)
	// 	ctx.Writef("}")
	// case "/get_students":
	// 	fmt.Println("students")
	// 	res := students_get()
	// 	ctx.Writef("{\"data\": ")
	// 	ctx.JSON(res)
	// 	fmt.Println(res)
	// 	ctx.Writef("}")

	default:
		ctx.Writef("")
	}

}

func post_class_handler(ctx iris.Context) {

	page_now := ctx.Path()
	c := &Class{}
	c_multi := []Class{}

	switch page_now {
	case "/post_data_class", "/post_data_class_edit":
		c = &Class{}
	case "/post_data_class_remove":
		c_multi = []Class{}
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

	res := ""

	if strings.Contains(page_now, "edit") {
		res = class_update(*c)
	} else if strings.Contains(page_now, "remove") {
		res = class_remove(c_multi)
	} else if strings.Contains(page_now, "data") {
		res = class_insert(*c)
	} else {
		ctx.Writef("fail")
	}

	//res := person_insert(*c, identity)
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

func post_account_handler(ctx iris.Context) {

	//initial
	page_now := ctx.Path()
	res := ""

	//page = * student
	if strings.Contains(page_now, "student") {
		c := &Student_Payment{}
		c_multi := []Student_Payment{}
		//get post value
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
		//action CRUD
		if strings.Contains(page_now, "edit") {
			res = account_student_update(*c)
		} else if strings.Contains(page_now, "remove") {
			res = account_student_remove(c_multi)
		} else if strings.Contains(page_now, "data") {
			res = account_student_insert(*c)
		} else {
			ctx.Writef("fail")
		}

	} else if strings.Contains(page_now, "assistant") {
		c := &Assistant_Salary{}
		c_multi := []Assistant_Salary{}
		//get post value
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
		//action CRUD
		if strings.Contains(page_now, "edit") {
			res = account_assistant_update(*c)
		} else if strings.Contains(page_now, "remove") {
			res = account_assistant_remove(c_multi)
		} else if strings.Contains(page_now, "data") {
			res = account_assistant_insert(*c)
		} else {
			ctx.Writef("fail")
		}
	} else if strings.Contains(page_now, "teacher") {
		c := &Teacher_Salary{}
		c_multi := []Teacher_Salary{}
		//get post value
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
		//action CRUD
		if strings.Contains(page_now, "edit") {
			res = account_teacher_update(*c)
		} else if strings.Contains(page_now, "remove") {
			res = account_teacher_remove(c_multi)
		} else if strings.Contains(page_now, "data") {
			res = account_teacher_insert(*c)
		} else {
			ctx.Writef("fail")
		}
	} else {
		ctx.Writef("fail")
	}

	//result error handle
	if res != "success" {
		ctx.Writef(string("{ \"error\": \"" + res + "\" }"))
	} else {
		if strings.Contains(page_now, "remove") {
			fmt.Println("=======remove")
			ctx.WriteString("OK")
			//ctx.Writef("")
		} else {
			// m := structs.Map(*c)
			// b, err := json.Marshal(m)
			// fmt.Println("=======m")
			// fmt.Println(m)
			// fmt.Println("=======b")
			// fmt.Println(b)
			// if err != nil {
			// 	fmt.Println(err.Error())
			// 	fmt.Println("=======error")
			// 	ctx.Writef("fail")
			// }
			// ctx.Writef(string(b))
		}

	}
}
func postBattery(ctx iris.Context) {
	var p Power
	fmt.Println(p.Power)
	ctx.Writef("{\"battery\": ")
	ctx.JSON(p.Power)
	ctx.Writef("}")

}
func powerPost(ctx iris.Context) {
	// power := &Power_{}

	var p Power
	fmt.Println(p.Power)
	if err := ctx.ReadJSON(&p); err != nil {
		fmt.Println(11)
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		fmt.Println(err.Error())
	} else {
		fmt.Println(2)

		fmt.Println(p.Power)
		//fmt.Println(ctx.ReadJSON(&power))

	}
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
	case "/post_data_assistant", "/post_data_assistant_edit":
		c = &Person{}
		identity = "Assistant"
	case "/post_data_assistant_remove":
		c_multi = []Person{}
		identity = "Assistant"

	// case "/post_data_class", "/post_data_class_edit":
	// 	c = &Class{}
	// case "/post_data_class_remove":
	// 	c_multi = []Class{}
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

	// member pages
	if strings.Contains(page_now, "edit") {
		res = person_update(*c, identity)
	} else if strings.Contains(page_now, "remove") {
		res = person_remove(c_multi, identity)
	} else if strings.Contains(page_now, "data") {
		res = person_insert(*c, identity)
	} else {
		ctx.Writef("fail")
	}

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
