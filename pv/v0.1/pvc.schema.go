package pv

// Auto-generated. DO NOT EDIT.
import (
    "github.com/appscode/api/dtypes"
    "github.com/xeipuuv/gojsonschema"
    "log"
)
var pVCUnregisterRequestSchema *gojsonschema.Schema
var pVCRegisterRequestSchema *gojsonschema.Schema
var pVCDescribeRequestSchema *gojsonschema.Schema

func init() {
	var err error
	pVCUnregisterRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cluster": {
      "type": "string"
    },
    "name": {
      "type": "string"
    },
    "namespace": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	pVCRegisterRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cluster": {
      "type": "string"
    },
    "name": {
      "type": "string"
    },
    "namespace": {
      "type": "string"
    },
    "size": {
      "type": "integer"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	pVCDescribeRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cluster": {
      "type": "string"
    },
    "name": {
      "type": "string"
    },
    "namespace": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
}

func (m *PVCUnregisterRequest) IsValid() (*gojsonschema.Result, error) {
	return pVCUnregisterRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *PVCUnregisterRequest) IsRequest() {}

func (m *PVCRegisterRequest) IsValid() (*gojsonschema.Result, error) {
	return pVCRegisterRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *PVCRegisterRequest) IsRequest() {}

func (m *PVCDescribeRequest) IsValid() (*gojsonschema.Result, error) {
	return pVCDescribeRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *PVCDescribeRequest) IsRequest() {}

func (m *PVCDescribeResponse) SetStatus(s *dtypes.Status) {
   m.Status = s
}

