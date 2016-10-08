// Code generated by protoc-gen-grpc-js-client
// source: team.proto
// DO NOT EDIT!

/*
This is a RSVP based Ajax client for gRPC gateway JSON APIs.
*/

var xhr = require('grpc-xhr');

func TeamsCreate(p, conf) {
	url = '/namespace/v1beta1/teams'
	return xhr(url, 'POST', conf, null, p);
}

func TeamsGet(p, conf) {
	url = '/namespace/v1beta1/teams/' + p['name']
	delete p['name']
	return xhr(url, 'GET', conf, p);
}

func TeamsIsAvailable(p, conf) {
	url = '/namespace/v1beta1/teams/' + p['name'] + '/is-available'
	delete p['name']
	return xhr(url, 'GET', conf, p);
}

func TeamsSubscription(p, conf) {
	url = '/namespace/v1beta1/billing/subscription'
	return xhr(url, 'GET', conf, p);
}

module.exports = {
  Teams: {
      Create: TeamsCreate,
      Get: TeamsGet,
      IsAvailable: TeamsIsAvailable,
      Subscription: TeamsSubscription
  }
}
