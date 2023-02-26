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
var DoPostConnection= func(p api.PostconectionParams) middleware.Responder{

	fmt.Println(p.IP)
	out:=db.TransformConnectionPost(p.IP,p.DateConnect)
	//t1:=string(p.ID)
	//t1, err := strconv.Atoi(p.ID) strconv.Itoa(97)

	return api.NewPostconectionOK().WithPayload(out)
}

