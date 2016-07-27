package alert

// Auto-generated. DO NOT EDIT.
import (
	"github.com/golang/glog"
	"github.com/xeipuuv/gojsonschema"
)

var acknowledgeRequestSchema *gojsonschema.Schema
var notifyRequestSchema *gojsonschema.Schema

func init() {
	var err error
	acknowledgeRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "comment": {
      "type": "string"
    },
    "phid": {
      "title": "Notification PHID",
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	notifyRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "alert_phid": {
      "type": "string"
    },
    "host_name": {
      "type": "string"
    },
    "service_output": {
      "type": "string"
    },
    "state": {
      "type": "string"
    },
    "state_type": {
      "type": "string"
    },
    "time": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
}

func (m *AcknowledgeRequest) IsValid() (*gojsonschema.Result, error) {
	return acknowledgeRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *AcknowledgeRequest) IsRequest() {}

func (m *NotifyRequest) IsValid() (*gojsonschema.Result, error) {
	return notifyRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *NotifyRequest) IsRequest() {}
