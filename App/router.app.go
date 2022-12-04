package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	controller "github.com/vandenbill/brand-commerce/Controller"
	middleware "github.com/vandenbill/brand-commerce/Middleware"
	repo "github.com/vandenbill/brand-commerce/Repo"
	service "github.com/vandenbill/brand-commerce/Service"
	"gorm.io/gorm"
	"os"
)

func NewRouter(r *gin.Engine, db *gorm.DB) {
	var (
		secretKey      string                    = os.Getenv("SECRET_KEY")
		jwtService     service.JwtService        = service.NewJwtServiceImpl(secretKey)
		userRepo       repo.UserRepo             = repo.NewUserRepoImpl(db)
		authService    service.AuthService       = service.NewAuthServiceImpl(userRepo)
		userService    service.UserService       = service.NewUserServicImpl(userRepo)
		authController controller.AuthController = controller.NewAuthControllerImpl(authService, userService, jwtService)
		userController controller.UserController = controller.NewUserControllerImpl(userService)
	)

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", authController.Register)
			auth.POST("/login", authController.Login)
		}

		user := api.Group("/user")
		{
			user.GET("/", middleware.AuthorizeJWT(jwtService), userController.FindByID)
			user.PATCH("/", middleware.AuthorizeJWT(jwtService), userController.Update)
			user.DELETE("/", middleware.AuthorizeJWT(jwtService), userController.Delete)
		}

		product := api.Group("/product")
		{
			product.GET("/", middleware.AuthorizeJWT(jwtService), func(ctx *gin.Context) {
				fmt.Println(ctx.Get("sub"))
			})
		}
	}
}
