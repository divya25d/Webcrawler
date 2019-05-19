package httplib

import (

	"../services"

	"fmt"
	"github.com/astaxie/beego"

)


//MainController of the HTTP server
type MainController struct {
	beego.Controller
}

//To do health check
func (this *MainController) GetStatus() {
	this.Ctx.ResponseWriter.WriteHeader(201)
	this.Ctx.WriteString("Service is up & running")
}

//To get domain results
func (this *MainController) GetDomainDetails() {
	url := this.Ctx.Input.Header("url")

	getDomainDetails, err := services.GetDomainDetails(url)

	if err != nil {
		this.Ctx.WriteString(fmt.Sprintf("%v", err))
		return
	}

	this.Ctx.ResponseWriter.WriteHeader(201)
	this.Ctx.WriteString(string(getDomainDetails))
}

//Run main function that starts the HTTP server
func Run(config string) {
	fmt.Println("Starting the HTTP server at port ", config)
	beego.Run(":" + config)
}
