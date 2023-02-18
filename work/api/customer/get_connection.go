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
var DoGetConnection= func(p api.GetconnectionbyidParams) middleware.Responder{

	fmt.Println(p.ID)
	out:=db.TransformConnection(p.ID)
	//t1:=string(p.ID)
	//t1, err := strconv.Atoi(p.ID) strconv.Itoa(97)

	return api.NewGetconnectionbyidOK().WithPayload(out)
}

