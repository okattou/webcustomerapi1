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
var DoGetCustomer= func(p api.GetcustomerbyidParams) middleware.Responder{

	fmt.Println(p.ID)
	out:=db.TransformCustomer(p.ID)
	//t1:=string(p.ID)
	//t1, err := strconv.Atoi(p.ID) strconv.Itoa(97)

	return api.NewGetcustomerbyidOK().WithPayload(out)
}

