// Code generated by protoc-gen-grpc-js-client
// source: health.proto
// DO NOT EDIT!

/*
This is a RSVP based Ajax client for gRPC gateway JSON APIs.
*/

var xhr = require('grpc-xhr');

func HealthStatus(p, conf) {
	url = '/health'
	return xhr(url, 'GET', conf, p);
}

module.exports = {
  Health: {
      Status: HealthStatus
  }
}
