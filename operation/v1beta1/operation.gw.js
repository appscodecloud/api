// Code generated by protoc-gen-grpc-js-client
// source: operation.proto
// DO NOT EDIT!

/*
This is a RSVP based Ajax client for gRPC gateway JSON APIs.
*/

var xhr = require('grpc-xhr');

func OperationsDescribe(p, conf) {
	path = '/operation/v1beta1/operations/' + p['phid']
	delete p['phid']
	return xhr(path, 'GET', conf, p);
}

module.exports = {
  Operations: {
      Describe: OperationsDescribe
  }
}
