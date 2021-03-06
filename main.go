package main

import (
	"fmt"
	"time"

	"github.com/kataras/iris"
	session "github.com/kataras/iris/sessions"
)

type page struct {
	PTitle      string
	Permission  string
	User        string
	To_user_msg string
	Javascript  string
	Something   string
	SomeNumber  string
}

type page_classes struct {
	Full_Name string
}

var sess = session.New(session.Config{
	Cookie:  ".cookiesession.id",
	Expires: time.Minute,
})

//Set and Get
func h(ctx iris.Context) {
	session := sess.Start(ctx)
	session.Set("key", "value")

	value := session.GetString("key")
	if value == "" {
		ctx.WriteString("NOT_OK")
		return
	}

	ctx.WriteString(value)
}

func main() {
	// vars := map[string]interface{}{
	// 	"something":  "some value",
	// 	"someNumber": 23,
	// }

	app := iris.New()
	tmpl := iris.HTML("./templates", ".html")
	// tmpl.Layout("layouts/layout.html")
	tmpl.Layout("layouts/main_layout.html")
	// t, _ := tmpl.Templates.Parse("page_members_students.html")
	// t.Execute(io.Writer, vars)
	tmpl.AddFunc("greet", func(s string) string {
		return "Greetings " + s + "!"
	})
	tmpl.AddFunc("greetee", func(s string) string {
		return "Greetings " + s + "!"
	})
	tmpl.AddFunc("confirm_del_title", func(s string) string {
		return "Delete confirm"
	})
	tmpl.AddFunc("confirm_del_body", func(s string) string {
		return ""
	})
	tmpl.AddFunc("confirm_del_submit", func(s string) string {
		return "Delete"
	})

	tmpl.AddFunc("classes", func(word string) string {
		switch word {
		case "Full Name":
			return "課程名稱"
		case "Day":
			return "上課星期"
		case "Time":
			return "上課時間"
		case "Duration":
			return "上課時數(小時)"
		case "Join Date":
			return "課程成立日"
		case "Teacher":
			return "老師"
		case "Type":
			return "課程類型"
		case "Students":
			return "學生"
		case "Cost":
			return "課程價格(一期ㄧ人)"
		case "Periodical Payment":
			return "下次收費時間"
		case "Auto Calculate":
			return "自動更新"

		}
		return "Delete"
	})

	app.RegisterView(tmpl)

	// register static assets request path and system directory
	app.StaticWeb("/js", "./static/assets/js")
	app.StaticWeb("/css", "./static/assets/css")
	app.StaticWeb("/pic", "./static/assets/pic")
	app.StaticWeb("/json", "./static/assets/json")

	app.Get("/data_student", writePathHandler)
	app.Post("/post_data_student", postHandler)
	app.Put("/post_data_student_edit", postHandler)
	app.Put("/post_data_student_remove", postHandler)

	app.Get("/data_teacher", writePathHandler)
	app.Post("/post_data_teacher", postHandler)
	app.Put("/post_data_teacher_edit", postHandler)
	app.Put("/post_data_teacher_remove", postHandler)

	app.Get("/data_assistant", writePathHandler)
	app.Post("/post_data_assistant", postHandler)
	app.Put("/post_data_assistant_edit", postHandler)
	app.Put("/post_data_assistant_remove", postHandler)

	app.Get("/data_student_name_list", writePathHandler)
	app.Get("/data_teacher_name_list", writePathHandler)
	app.Get("/data_assistant_name_list", writePathHandler)

	// app.Get("/get_teachers", writePathHandler)
	// app.Get("/get_students", writePathHandler)

	app.Get("/data_class", writePathHandler)
	app.Post("/post_data_class", post_class_handler)
	app.Put("/post_data_class_edit", post_class_handler)
	app.Put("/post_data_class_remove", post_class_handler)

	app.Get("/data_account_student", writePathHandler)            //view
	app.Post("/post_account_student", post_account_handler)       //add
	app.Put("/post_account_student_edit", post_account_handler)   //edit
	app.Put("/post_account_student_remove", post_account_handler) //remove

	app.Get("/data_account_teacher", writePathHandler)            //view
	app.Post("/post_account_teacher", post_account_handler)       //add
	app.Put("/post_account_teacher_edit", post_account_handler)   //edit
	app.Put("/post_account_teacher_remove", post_account_handler) //remove

	app.Get("/data_account_assistant", writePathHandler)            //view
	app.Post("/post_account_assistant", post_account_handler)       //add
	app.Put("/post_account_assistant_edit", post_account_handler)   //edit
	app.Put("/post_account_assistant_remove", post_account_handler) //remove

	app.Get("/send_mail", writePathHandler)
	app.Get("/data_notice", writePathHandler)            //view
	app.Post("/post_notice", post_account_handler)       //add
	app.Put("/post_notice_edit", post_account_handler)   //edit
	app.Put("/post_notice_remove", post_account_handler) //remove

	app.Post("/login_check", post_login_Handler)
	app.Get("/login", func(ctx iris.Context) {

		ctx.ViewLayout(iris.NoLayout)
		ctx.ViewData("", page{PTitle: "Login-Victory"})
		if err := ctx.View("index.html"); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Writef(err.Error())
		}

	})

	// app.Get("/logout", func(ctx iris.Context) {
	// 	// removes all entries
	// 	sess.Start(ctx).Clear()
	// 	ctx.Redirect("/login")
	// })

	app.Get("/", func(ctx iris.Context) {
		session := sess.Start(ctx)
		name := session.GetString("name")
		permission := session.GetString("permission")
		fmt.Println("name:" + name)
		ctx.ViewData("", page{PTitle: "HOME-Victory", Permission: permission, User: name, To_user_msg: to_user_msg(permission)})
		if err := ctx.View("page_index.html"); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Writef(err.Error())
		}
	})
	app.Get("/calendar", func(ctx iris.Context) {
		session := sess.Start(ctx)
		value := session.GetString("key")
		// if value == "" {
		// 	ctx.WriteString("NOT_OK")
		// }

		ctx.ViewData("", page{PTitle: "Calendar-Victory" + value})
		if err := ctx.View("page_calendar.html"); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Writef(err.Error())
		}
	})
	app.Get("/mailer", func(ctx iris.Context) {
		ctx.ViewData("", page{PTitle: "Mailer-Victory"})
		if err := ctx.View("page_gmail.html"); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Writef(err.Error())
		}
	})
	// set a layout for a party, .Layout should be BEFORE any Get or other Handle party's method
	p_members := app.Party("/members").Layout("layouts/main_layout.html")
	{ // both of these will use the layouts/mylayout.html as their layout.
		p_members.Get("/students", func(ctx iris.Context) {
			//ctx.View("page1.html")
			ctx.ViewData("", page{
				PTitle: "Students-Members-Victory",
				// Something:  "some value",
				// SomeNumber: "23",
			})
			if err := ctx.View("page_members_students.html"); err != nil {
				ctx.StatusCode(iris.StatusInternalServerError)
				ctx.Writef(err.Error())
			}
		})
		p_members.Get("/teachers", func(ctx iris.Context) {
			//ctx.View("page1.html")

			ctx.ViewData("", page{
				PTitle: "Teachers-Members-Victory",
				// modal_title:  "Delete Confirm",
				// modal_body:   "Do you want to delete? ",
				// modal_submit: "Delete",
			})
			// ctx.ViewDate("", page{modal_title: "Delete Confirm"})
			// ctx.ViewDate("", page{modal_body: "Do you want to delete? "})
			// ctx.ViewDate("", page{modal_submit: "Delete"})
			if err := ctx.View("page_members_teachers.html"); err != nil {
				ctx.StatusCode(iris.StatusInternalServerError)
				ctx.Writef(err.Error())
			}
		})
		p_members.Get("/assistants", func(ctx iris.Context) {
			//ctx.View("page1.html")
			ctx.ViewData("", page{PTitle: "Assistants-Members-Victory"})
			if err := ctx.View("page_members_assistants.html"); err != nil {
				ctx.StatusCode(iris.StatusInternalServerError)
				ctx.Writef(err.Error())
			}
		})
	}
	app.Get("/classes", func(ctx iris.Context) {
		ctx.ViewData("", page{PTitle: "Classes-Victory"})
		if err := ctx.View("page_classes.html"); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Writef(err.Error())
		}
	})
	app.Get("/notices", func(ctx iris.Context) {
		ctx.ViewData("", page{PTitle: "Notices-Victory"})
		if err := ctx.View("page_notices.html"); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Writef(err.Error())
		}
	})
	p_accounts := app.Party("/accounts").Layout("layouts/main_layout.html")
	{
		p_accounts.Get("/students", func(ctx iris.Context) {
			//ctx.View("page1.html")
			ctx.ViewData("", page{PTitle: "Students-Account-Victory"})
			if err := ctx.View("page_accounts_students.html"); err != nil {
				ctx.StatusCode(iris.StatusInternalServerError)
				ctx.Writef(err.Error())
			}
		})

		p_accounts.Get("/teachers", func(ctx iris.Context) {
			ctx.ViewData("", page{PTitle: "Teachers-Account-Victory"})
			if err := ctx.View("page_accounts_teachers.html"); err != nil {
				ctx.StatusCode(iris.StatusInternalServerError)
				ctx.Writef(err.Error())
			}
		})
		p_accounts.Get("/assistants", func(ctx iris.Context) {
			ctx.ViewData("", page{PTitle: "Assistants-Account-Victory"})
			if err := ctx.View("page_accounts_assistants.html"); err != nil {
				ctx.StatusCode(iris.StatusInternalServerError)
				ctx.Writef(err.Error())
			}
		})
	}

	// remove the layout for a specific route
	app.Get("/nolayout", func(ctx iris.Context) {
		ctx.ViewLayout(iris.NoLayout)
		if err := ctx.View("page1.html"); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Writef(err.Error())
		}
	})

	// set a layout for a party, .Layout should be BEFORE any Get or other Handle party's method
	my := app.Party("/my").Layout("layouts/mylayout.html")
	{ // both of these will use the layouts/mylayout.html as their layout.
		my.Get("/", func(ctx iris.Context) {
			ctx.View("page1.html")
		})
		my.Get("/other", func(ctx iris.Context) {
			ctx.View("page1.html")
		})
	}

	// http://localhost:8080
	// http://localhost:8080/nolayout
	// http://localhost:8080/my
	// http://localhost:8080/my/other
	app.Run(iris.Addr(":8080"))
}
