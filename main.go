package main

import (
	_middleware "final/app/middleware"
	"final/app/routes"
	_adminUseCase "final/business/admins"
	_productUseCase "final/business/products"
	_transactionUseCase "final/business/transactions"
	_userUseCase "final/business/users"
	_adminController "final/controllers/admins"
	_productController "final/controllers/products"
	_transactionController "final/controllers/transactions"
	_userController "final/controllers/users"
	_adminDB "final/drivers/database/admins"
	_productDB "final/drivers/database/products"
	_transactionDB "final/drivers/database/transactions"
	_userDB "final/drivers/database/users"
	_mysqlDriver "final/drivers/mysql"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(&_userDB.User{},
		&_transactionDB.Shipment{}, &_productDB.Product_type{}, &_transactionDB.Payment_Method{},
		&_transactionDB.Transaction{}, &_transactionDB.Transaction_Detail{}, &_adminDB.Admin{})
}

func main() {
	configDB := _mysqlDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	configJWTAdmin := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.admin`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}
	Conn := configDB.InitialDb()

	dbMigrate(Conn)
	e := echo.New()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	userRepository := _userDB.NewMysqlRepository(Conn)
	userUseCase := _userUseCase.NewUserUseCase(userRepository, timeoutContext, configJWT)
	userController := _userController.NewUserController(userUseCase)

	adminRepository := _adminDB.NewMysqlRepository(Conn)
	adminUseCase := _adminUseCase.NewAdminUseCase(adminRepository, timeoutContext, configJWTAdmin)
	adminController := _adminController.NewAdminController(adminUseCase)

	transactionRepository := _transactionDB.NewMysqlRepository(Conn)
	transactionUseCase := _transactionUseCase.NewTransactionUseCase(transactionRepository, timeoutContext, configJWT)
	transactionController := _transactionController.NewTransactionController(transactionUseCase)

	productRepository := _productDB.NewMysqlRepository(Conn)
	productUseCase := _productUseCase.NewProductUseCase(productRepository, timeoutContext)
	productController := _productController.NewProductController(productUseCase)

	routesInit := routes.ControllerList{
		UserController:        *userController,
		TransactionController: *transactionController,
		ProductController:     *productController,
		JWTMiddleware:         configJWT.Init(),
		AdminController:       *adminController,
		JWTAdmin:              configJWTAdmin.Init(),
	}
	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}

//  $ ssh -i "C:\Users\remus\Downloads\kampus_merdeka.pem" ec2-user@ec2-3-21-206-55.us-east-2.compute.amazonaws.com

// [ec2-user@ip-172-31-31-116 ~]$ docker run -p 8000:8000 --name kampus_merdeka 19
// 4517/kampus_merdeka:1.0.0
// "host": "kampusmerdeka.cscvpk8eja5o.us-east-2.rds.amazonaws.com",
// "port": "3306",
// "user": "admin",
