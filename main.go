package main

import (
	"github.com/kataras/iris"
)

type page struct {
	PTitle string
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

	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("", page{PTitle: "HOME-W-LIAO"})
		if err := ctx.View("main_content.html"); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Writef(err.Error())
		}
	})
	app.Get("/calendar", func(ctx iris.Context) {
		ctx.ViewData("", page{PTitle: "Calendar-W-LIAO"})
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
			ctx.ViewData("", page{PTitle: "Students-Members-W-LIAO"})
			if err := ctx.View("page_members_students.html"); err != nil {
				ctx.StatusCode(iris.StatusInternalServerError)
				ctx.Writef(err.Error())
			}
		})
		p_members.Get("/teachers", func(ctx iris.Context) {
			//ctx.View("page1.html")
			ctx.ViewData("", page{PTitle: "Teachers-Members-W-LIAO"})
			if err := ctx.View("page_members_teachers.html"); err != nil {
				ctx.StatusCode(iris.StatusInternalServerError)
				ctx.Writef(err.Error())
			}
		})
		p_members.Get("/assistants", func(ctx iris.Context) {
			//ctx.View("page1.html")
			ctx.ViewData("", page{PTitle: "Assistants-Members-W-LIAO"})
			if err := ctx.View("page_members_assistants.html"); err != nil {
				ctx.StatusCode(iris.StatusInternalServerError)
				ctx.Writef(err.Error())
			}
		})
	}

	// app.Get("/calendar", func(ctx iris.Context) {
	// 	ctx.ViewData("", page{PTitle: "Calendar-W-LIAO"})
	// 	if err := ctx.View("page_calendar.html"); err != nil {
	// 		ctx.StatusCode(iris.StatusInternalServerError)
	// 		ctx.Writef(err.Error())
	// 	}
	// })
	// app.Get("/calendar", func(ctx iris.Context) {
	// 	ctx.ViewData("", page{PTitle: "Calendar-W-LIAO"})
	// 	if err := ctx.View("page_calendar.html"); err != nil {
	// 		ctx.StatusCode(iris.StatusInternalServerError)
	// 		ctx.Writef(err.Error())
	// 	}
	// })
	// app.Get("/calendar", func(ctx iris.Context) {
	// 	ctx.ViewData("", page{PTitle: "Calendar-W-LIAO"})
	// 	if err := ctx.View("page_calendar.html"); err != nil {
	// 		ctx.StatusCode(iris.StatusInternalServerError)
	// 		ctx.Writef(err.Error())
	// 	}
	// })
	// app.Get("/calendar", func(ctx iris.Context) {
	// 	ctx.ViewData("", page{PTitle: "Calendar-W-LIAO"})
	// 	if err := ctx.View("page_calendar.html"); err != nil {
	// 		ctx.StatusCode(iris.StatusInternalServerError)
	// 		ctx.Writef(err.Error())
	// 	}
	// })
	// app.Get("/calendar", func(ctx iris.Context) {
	// 	ctx.ViewData("", page{PTitle: "Calendar-W-LIAO"})
	// 	if err := ctx.View("page_calendar.html"); err != nil {
	// 		ctx.StatusCode(iris.StatusInternalServerError)
	// 		ctx.Writef(err.Error())
	// 	}
	// })
	// app.Get("/calendar", func(ctx iris.Context) {
	// 	ctx.ViewData("", page{PTitle: "Calendar-W-LIAO"})
	// 	if err := ctx.View("page_calendar.html"); err != nil {
	// 		ctx.StatusCode(iris.StatusInternalServerError)
	// 		ctx.Writef(err.Error())
	// 	}
	// })
	// app.Get("/calendar", func(ctx iris.Context) {
	// 	ctx.ViewData("", page{PTitle: "Calendar-W-LIAO"})
	// 	if err := ctx.View("page_calendar.html"); err != nil {
	// 		ctx.StatusCode(iris.StatusInternalServerError)
	// 		ctx.Writef(err.Error())
	// 	}
	// })
	// app.Get("/calendar", func(ctx iris.Context) {
	// 	ctx.ViewData("", page{PTitle: "Calendar-W-LIAO"})
	// 	if err := ctx.View("page_calendar.html"); err != nil {
	// 		ctx.StatusCode(iris.StatusInternalServerError)
	// 		ctx.Writef(err.Error())
	// 	}
	// })

	// app.Get("/", func(ctx iris.Context) {
	// 	if err := ctx.View("page1.html"); err != nil {
	// 		ctx.StatusCode(iris.StatusInternalServerError)
	// 		ctx.Writef(err.Error())
	// 	}
	// })

	// // remove the layout for a specific route
	// app.Get("/nolayout", func(ctx iris.Context) {
	// 	ctx.ViewLayout(iris.NoLayout)
	// 	if err := ctx.View("page1.html"); err != nil {
	// 		ctx.StatusCode(iris.StatusInternalServerError)
	// 		ctx.Writef(err.Error())
	// 	}
	// })

	// // set a layout for a party, .Layout should be BEFORE any Get or other Handle party's method
	// my := app.Party("/my").Layout("layouts/mylayout.html")
	// { // both of these will use the layouts/mylayout.html as their layout.
	// 	my.Get("/", func(ctx iris.Context) {
	// 		ctx.View("page1.html")
	// 	})
	// 	my.Get("/other", func(ctx iris.Context) {
	// 		ctx.View("page1.html")
	// 	})
	// }

	// http://localhost:8080
	// http://localhost:8080/nolayout
	// http://localhost:8080/my
	// http://localhost:8080/my/other
	app.Run(iris.Addr(":8080"))
}
