package billing

import (
	"github.com/xeipuuv/gojsonschema"
	"log"
)

// Auto-generated. DO NOT EDIT.

var chargeCalculateRequestSchema *gojsonschema.Schema

func init() {
	var err error
	chargeCalculateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "properties":{
    "charge_type":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
}

func (m *ChargeCalculateRequest) InValid() (*gojsonschema.Result, error) {
	return chargeCalculateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
