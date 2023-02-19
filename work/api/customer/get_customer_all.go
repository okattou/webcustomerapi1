package customer

import (
	"github.com/go-openapi/runtime/middleware"
	"fmt"

	"github.com/webcustomerapi1/work/db"
	api "github.com/webcustomerapi1/restapi/operations/customer"
)


//
//
//
var DoGetCustomerAll= func(p api.GetcustomerParams) middleware.Responder{

	out:=db.TransformCustomerall()
	fmt.Println(out)
	return api.NewGetcustomerOK().WithPayload(out)
}
