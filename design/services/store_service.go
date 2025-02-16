package services

import (
	"aion/design/errors"

	. "goa.design/goa/v3/dsl"
)

var FileStoreService = Service("FileStore", func() {
	Description("The FileStore service handles file management")

	HTTP(func() {
		Path("/file-store")
	})

	Error("unauthorized", errors.Unauthorized, "Autenticazione fallita")
	Error("internalServerError", errors.InternalServerError, "Errore nel server")
	Error("notFound", errors.NotFound, "Dato non trovato all'interno del sistema")
	Error("badRequest", errors.BadRequest, "Parametri non validi")
})
