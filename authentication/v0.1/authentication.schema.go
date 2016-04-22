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
  "properties":{
    "namespace":{
      "type":"string"
    },
    "secret":{
      "type":"string"
    },
    "username":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
}

func (m *ValidateRequest) InValid() (*gojsonschema.Result, error) {
	return validateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
