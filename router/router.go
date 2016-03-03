package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pasarsakomandiri/controllers"
	"github.com/pasarsakomandiri/router/middleware"
)

func Initialize(r *gin.Engine)  {
	//GLOBAL PAGES
	r.GET("/redirected", controllers.Redirected)

	//LOGIN PAGE
	r.GET("/", middleware.DisallowAuthenticated(), controllers.LoginPage)

	//USER PAGES
	r.GET("/user/user_auth", middleware.DisAllowAnon(), controllers.UserSessionRedirect)
	r.GET("/user/register",middleware.DisAllowAnon(), controllers.UserRegisterPages)

	//ADMIN PAGES
	r.GET("/admin", middleware.AllowOnlyAdministrator(), controllers.AdminPage)

	//DEVICE PAGES
	r.GET("/device/device_group", middleware.AllowOnlyAdministrator(), controllers.DeviceGroupPage)
	r.GET("/formval", middleware.DisAllowAnon(), controllers.GetDeviceType)
	r.GET("/device", middleware.DisAllowAnon(), controllers.DeviceListPages)

	//user register
	r.GET("/user_list", middleware.DisAllowAnon(), controllers.UserListPages)
	r.GET("/update_user", middleware.DisAllowAnon(), controllers.UserEditPages)
	r.GET("/get_id", middleware.DisAllowAnon(), controllers.UserGetInfoAPI)

	//CASHIER PAGE
	r.GET("/cashier", middleware.DisAllowAnon(), controllers.CashierPage)

	//USER API
	r.POST("/api/user/login", middleware.DisallowAuthenticated(), controllers.LoginAPI)
	r.GET("/api/user/logout", middleware.DisAllowAnon(), controllers.LogoutAPI)
	r.POST("/api/user/register", middleware.DisAllowAnon(), controllers.RegisterUser)
	r.GET("/api/user/role_list",middleware.DisAllowAnon(), controllers.UserGetAllRoleAPI)
	r.GET("/api/user/all_user", middleware.DisAllowAnon(), controllers.UserGetAllAPI)
	r.POST("/api/user/update", middleware.DisAllowAnon(), controllers.UserUpdateAPI)

	//DEVICE API
	r.POST("/api/device/create_device", middleware.DisAllowAnon(), controllers.DeviceRegisterAPI)
	r.GET("/api/device/device_list", middleware.DisAllowAnon(), controllers.GetDeviceListAPI)
	r.GET("api/device/devices_by_type", middleware.DisAllowAnon(), controllers.DeviceGetFromTypeAPI)
	r.GET("/api/device/device_info", middleware.DisAllowAnon(), controllers.DeviceGetDeviceInfoAPI)
	r.POST("/api/device/create_device_group", middleware.DisAllowAnon(), controllers.DeviceGroupRegisterAPI)
	r.GET("/api/device/device_group_list", middleware.DisAllowAnon(), controllers.DeviceGroupGetAllAPI)
	r.GET("/api/device/check_out_device", controllers.DeviceContactCheckOut)
	r.POST("/api/device/delete_device_group",middleware.DisAllowAnon(), controllers.DeviceGroupDeleteAPI)
	r.POST("/api/device/delete", middleware.DisAllowAnon(), controllers.DeleteDeviceList)

	//PARKING API
	r.GET("/api/parking/checkIn", controllers.ParkingCheckIn) //authenticated in controllers
	r.POST("/api/parking/checkOut", controllers.ParkingCheckOut) //authenticated in controllers
	r.GET("/api/parking/getTicketInfo", middleware.DisAllowAnon(), controllers.ParkingGetTicketInfo)
	r.GET("/api/parking/vehicleAll", middleware.DisAllowAnon(), controllers.VehicleGetAll)
	r.GET("/api/parking/test/checkIn", controllers.ParkingCheckIn)
	//PARKING PRICE
	r.GET("/parking_price", middleware.DisAllowAnon(), controllers.PriceConfigPage)
	r.POST("/api/parking_price", middleware.DisAllowAnon(), controllers.PriceRegister)
	r.GET("/api/parking/price_list", middleware.DisAllowAnon(), controllers.PriceGetAll)
	//PARKING TRANSACTIONS API
	r.GET("/parking/transactions", middleware.DisAllowAnon(), controllers.ParkingTransactionsPage)
	r.GET("/transactions_tabel", middleware.DisAllowAnon(), controllers.ParkingTransactionsGetAll)
	r.GET("/tglparking", middleware.DisAllowAnon(), controllers.ParkingTransTgl)

	//CAMERA API
	//r.GET("/camera/takepicture", controllers.IpCamTakePicture)

	//API
	r.GET("/api/create_super_user", controllers.SecretCreateSuperUser)
	r.GET("/api/user_since", controllers.UserSinceAPI)
}