package payloads

import (
	. "goa.design/goa/v3/dsl"
)

var CreateEventPayload = Type("CreateEventPayload", func() {
	Attribute("name", String, "Name", func() {
		MinLength(1)
		MaxLength(100)
		Example("Football Match")
	})
	Attribute("description", String, "Description", func() {
		Example("A friendly football match")
	})
	Attribute("start_date", String, "Start Date", func() {
		Format(FormatDateTime)
		Example("2022-06-01T15:00:00Z")
	})
	Attribute("end_date", String, "End Date", func() {
		Format(FormatDateTime)
		Example("2022-06-01T17:00:00Z")
	})
	Attribute("tipology", String, "Event Typology", func() {
		Enum("Allenamento", "Trasferta", "Formazione", "Partecipazione Evento Sportivo", "Manifestazione Pubblica", "Riunione")
		Example("Allenamento")
	})
	Attribute("location_id", String, "Location ID", func() {
		Format(FormatUUID)
		Example("f47ac10b-58cc-4372-a567-0e02b2c3d479")
	})
	Attribute("responsible_id", String, "Responsible User ID", func() {
		Format(FormatUUID)
		Example("d47ac10b-58cc-4372-a567-0e02b2c3d480")
	})
	Required("name", "start_date", "end_date", "tipology", "location_id", "responsible_id")
})

var UpdateEventPayload = Type("UpdateEventPayload", func() {
	Attribute("id", String, "Event ID", func() {
		Format(FormatUUID)
		Example("f47ac10b-58cc-4372-a567-0e02b2c3d479")
	})
	Attribute("name", String, "Name", func() {
		MinLength(1)
		MaxLength(100)
		Example("Football Match")
	})
	Attribute("description", String, "Description", func() {
		Example("A friendly football match")
	})
	Attribute("start_date", String, "Start Date", func() {
		Format(FormatDateTime)
		Example("2022-06-01T15:00:00Z")
	})
	Attribute("end_date", String, "End Date", func() {
		Format(FormatDateTime)
		Example("2022-06-01T17:00:00Z")
	})
	Attribute("tipology", String, "Event Typology", func() {
		Enum("Allenamento", "Trasferta", "Formazione", "Partecipazione Evento Sportivo", "Manifestazione Pubblica", "Riunione")
		Example("Allenamento")
	})
	Attribute("location_id", String, "Location ID", func() {
		Format(FormatUUID)
		Example("f47ac10b-58cc-4372-a567-0e02b2c3d479")
	})
	Attribute("responsible_id", String, "Responsible User ID", func() {
		Format(FormatUUID)
		Example("d47ac10b-58cc-4372-a567-0e02b2c3d480")
	})
	Required("id")
})
