package routes

import (
	"github.com/ArvRao/ecommerce-app/controller"
	"github.com/ArvRao/ecommerce-app/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

var u = controller.NewUserFunc()

func UserRoute(app *fiber.App) {

	app.Get("/", u.First)

	app.Get("/metrics", monitor.New(monitor.Config{Title: "GoCart Metrics Page"}))
	app.Post("/user/registration", u.UserSignup) //json
	app.Post("/user/login", u.UserLogin)         //json
	app.Post("/user/refresh", u.Refresh)

	user := app.Group("/user", middleware.RequreUserAuth)

	user.Get("/home", u.Home)

	user.Post("/verification", u.Verification) //json
	user.Put("/user_account", u.EditUserInfo)  //json
	user.Get("/view_products", p.ViewProducts)

	user.Get("/addtocart/:id", u.AddToCart)
	user.Get("/removefromcart/:id", u.RemoveFromCart)

	user.Get("/get_by_category", p.GetbyCategory)
	app.Get("/search/:key", p.SearchProduct)

	user.Get("/order_from_cart", u.OrderFromCart)
	user.Get("/logout", u.UserLogout)

	user.Get("/checkout", u.Checkout)

	user.Get("/generate_invoice/:order_id", u.GenerateInvoice)

}