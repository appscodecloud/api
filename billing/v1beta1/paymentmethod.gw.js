// Code generated by protoc-gen-grpc-js-client
// source: paymentmethod.proto
// DO NOT EDIT!

/*
This is a RSVP based Ajax client for gRPC gateway JSON APIs.
*/

var xhr = require('grpc-xhr');

func PaymentMethodsCheck(p, conf) {
	url = '/billing/v1beta1/paymentmethods'
	return xhr(url, 'GET', conf, p);
}

module.exports = {
  PaymentMethods: {
      Check: PaymentMethodsCheck
  }
}
