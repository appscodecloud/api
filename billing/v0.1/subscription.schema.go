package billing

import (
	"github.com/xeipuuv/gojsonschema"
	"log"
)

// Auto-generated. DO NOT EDIT.

var subscriptionCloseRequestSchema *gojsonschema.Schema
var subscriptionQutaRequestSchema *gojsonschema.Schema
var subscriptionCreateRequestSchema *gojsonschema.Schema
var subscriptionDescribeRequestSchema *gojsonschema.Schema
var subscriptionOpenRequestSchema *gojsonschema.Schema

func init() {
	var err error
	subscriptionCloseRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "properties":{
    "object_phid":{
      "type":"string"
    },
    "product_type":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	subscriptionQutaRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "definitions":{
    "billingResource":{
      "default":"USER",
      "enum":[
        "USER",
        "CLUSTER",
        "NODE",
        "DB",
        "CI"
      ],
      "type":"string"
    }
  },
  "properties":{
    "count":{
      "type":"integer"
    },
    "object_phid":{
      "type":"string"
    },
    "resource":{
      "$ref":"#/definitions/billingResource"
    },
    "subresource":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	subscriptionCreateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "properties":{
    "author_phid":{
      "type":"string"
    },
    "product_phid":{
      "type":"string"
    },
    "type":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	subscriptionDescribeRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "properties":{
    "time":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
	subscriptionOpenRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "properties":{
    "author_phid":{
      "type":"string"
    },
    "metadata":{
      "type":"string"
    },
    "object_phid":{
      "type":"string"
    },
    "product_type":{
      "type":"string"
    },
    "subscription_phid":{
      "type":"string"
    }
  },
  "type":"object"
}`))
	if err != nil {
		log.Fatal(err)
	}
}

func (m *SubscriptionCloseRequest) InValid() (*gojsonschema.Result, error) {
	return subscriptionCloseRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *SubscriptionQutaRequest) InValid() (*gojsonschema.Result, error) {
	return subscriptionQutaRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *SubscriptionCreateRequest) InValid() (*gojsonschema.Result, error) {
	return subscriptionCreateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *SubscriptionDescribeRequest) InValid() (*gojsonschema.Result, error) {
	return subscriptionDescribeRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *SubscriptionOpenRequest) InValid() (*gojsonschema.Result, error) {
	return subscriptionOpenRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
