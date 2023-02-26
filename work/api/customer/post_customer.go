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
var DoPostCustomer= func(p api.PostcustomerParams) middleware.Responder{

	fmt.Println(p)
	out:=db.TransformCustomerPost()

	return api.NewPostcustomerOK().WithPayload(out)
}

