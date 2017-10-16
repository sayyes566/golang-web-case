package main

import (
	"github.com/kataras/iris"
)

type page struct {
	PTitle     string
	Javascript string
}

func main() {
	app := iris.New()
	tmpl := iris.HTML("./templates", ".html")
	// tmpl.Layout("layouts/layout.html")
	tmpl.Layout("layouts/main_layout.html")
	tmpl.AddFunc("greet", func(s string) string {
		return "Greetings " + s + "!"
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

	app.Get("/data_class", writePathHandler)
	app.Post("/post_data_class", postHandler)
	app.Put("/post_data_class_edit", postHandler)
	app.Put("/post_data_class_remove", postHandler)

	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("", page{PTitle: "HOME-Victory"})
		if err := ctx.View("main_content.html"); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Writef(err.Error())
		}
	})
	app.Get("/calendar", func(ctx iris.Context) {
		ctx.ViewData("", page{PTitle: "Calendar-Victory"})
		if err := ctx.View("page_calendar.html"); err != nil {
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
				//Javascript: " <script src=\"/js/data.js\"></script>",
			})
			if err := ctx.View("page_members_students.html"); err != nil {
				ctx.StatusCode(iris.StatusInternalServerError)
				ctx.Writef(err.Error())
			}
		})
		p_members.Get("/teachers", func(ctx iris.Context) {
			//ctx.View("page1.html")
			ctx.ViewData("", page{PTitle: "Teachers-Members-Victory"})
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
