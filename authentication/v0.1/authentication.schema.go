package authentication

import (
	"github.com/xeipuuv/gojsonschema"
	"log"
)

// Auto-generated. DO NOT EDIT.

var validateRequestSchema *gojsonschema.Schema

func init() {
	var err error
	validateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "namespace": {
      "type": "string"
    },
    "secret": {
      "type": "string"
    },
    "username": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
}

func (m *ValidateRequest) IsValid() (*gojsonschema.Result, error) {
	return validateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ValidateRequest) IsRequest() {}

