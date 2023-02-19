package customer

import (
	"github.com/go-openapi/runtime/middleware"
	"fmt"

	"github.com/webcustomerapi1/work/db"
	api "github.com/webcustomerapi1/restapi/operations/connection"

)



//
//
//
var DoGetConnectionAll= func(p api.GetconectionParams) middleware.Responder{

	out:=db.TransformConnectionall()
	fmt.Println(out)
	//t1:=string(p.ID)
	//t1, err := strconv.Atoi(p.ID) strconv.Itoa(97)
	

	return api.NewGetconectionOK().WithPayload(out)
}

