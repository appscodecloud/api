package artifactory

import (
	"github.com/xeipuuv/gojsonschema"
	"log"
)

// Auto-generated. DO NOT EDIT.

var listRequestSchema *gojsonschema.Schema
var searchRequestSchema *gojsonschema.Schema
var describeRequestSchema *gojsonschema.Schema

func init() {
	var err error
	listRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "type": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	searchRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "query": {
      "type": "string"
    },
    "type": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	describeRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "id": {
      "type": "string"
    },
    "type": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		log.Fatal(err)
	}
}

func (m *ListRequest) IsValid() (*gojsonschema.Result, error) {
	return listRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ListRequest) IsRequest() {}

func (m *SearchRequest) IsValid() (*gojsonschema.Result, error) {
	return searchRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *SearchRequest) IsRequest() {}

func (m *DescribeRequest) IsValid() (*gojsonschema.Result, error) {
	return describeRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *DescribeRequest) IsRequest() {}
