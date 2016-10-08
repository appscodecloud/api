// Code generated by protoc-gen-grpc-js-client
// source: acmeuser.proto
// DO NOT EDIT!

/*
This is a RSVP based Ajax client for gRPC gateway JSON APIs.
*/

var xhr = require('grpc-xhr');

func AcmeUsersRegister(p, conf) {
	url = '/certificate/v1beta1/acme-users'
	return xhr(url, 'PUT', conf, p);
}

module.exports = {
  AcmeUsers: {
      Register: AcmeUsersRegister
  }
}
