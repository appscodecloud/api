package billing

// Auto-generated. DO NOT EDIT.
import (
	"github.com/golang/glog"
	"github.com/xeipuuv/gojsonschema"
)

var subscriptionSubscribeRequestSchema *gojsonschema.Schema

func init() {
	var err error
	subscriptionSubscribeRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "product_phid": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
}

func (m *SubscriptionSubscribeRequest) IsValid() (*gojsonschema.Result, error) {
	return subscriptionSubscribeRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *SubscriptionSubscribeRequest) IsRequest() {}

