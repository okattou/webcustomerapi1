package customer

import (
	"github.com/go-openapi/runtime/middleware"

	"github.com/webcustomerapi1/work/db"
	api "github.com/webcustomerapi1/restapi/operations/connection"
)


//
//
//
var DoPutConnection= func(p api.PutconnectionbyidParams) middleware.Responder{
	out:=db.TransformConnectionPut(p.IP,p.DateConnect)
	return api.NewPutconnectionbyidOK().WithPayload(out)
}

