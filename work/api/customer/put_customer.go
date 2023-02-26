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
var DoPutCustomer= func(p api.PutcustomerbyidParams) middleware.Responder{

	fmt.Println(p)
	out:=db.TransformCustomerPut()

	return api.NewPutcustomerbyidOK().WithPayload(out)
}

