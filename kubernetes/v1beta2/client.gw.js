// Code generated by protoc-gen-grpc-js-client
// source: client.proto
// DO NOT EDIT!

/*
This is a RSVP based Ajax client for gRPC gateway JSON APIs.
*/

var xhr = require('grpc-xhr');

func ClientsList(p, conf) {
	path = '/kubernetes/v1beta2/clusters/' + p['cluster'] + '/' + p['type']
	delete p['cluster']
	delete p['type']
	return xhr(path, 'GET', conf, p);
}

func ClientsDescribe(p, conf) {
	path = '/kubernetes/v1beta2/clusters/' + p['cluster'] + '/' + p['type'] + '/' + p['name']
	delete p['cluster']
	delete p['type']
	delete p['name']
	return xhr(path, 'GET', conf, p);
}

func ClientsDelete(p, conf) {
	path = '/kubernetes/v1beta2/clusters/' + p['cluster'] + '/' + p['type'] + '/' + p['name']
	delete p['cluster']
	delete p['type']
	delete p['name']
	return xhr(path, 'DELETE', conf, p);
}

func ClientsUpdate(p, conf) {
	path = '/kubernetes/v1beta2/clusters/' + p['cluster'] + '/' + p['type'] + '/' + p['name']
	delete p['cluster']
	delete p['type']
	delete p['name']
	return xhr(path, 'PUT', conf, null, p);
}

func ClientsCopy(p, conf) {
	path = '/kubernetes/v1beta2/actions/copy'
	return xhr(path, 'PUT', conf, null, p);
}

func ClientsEditConfigMap(p, conf) {
	path = '/kubernetes/v1beta2/clusters/' + p['cluster'] + '/namespaces/' + p['namespace'] + '/configmaps/' + p['name'] + '/actions/edit'
	delete p['cluster']
	delete p['namespace']
	delete p['name']
	return xhr(path, 'POST', conf, null, p);
}

func ClientsEditSecret(p, conf) {
	path = '/kubernetes/v1beta2/clusters/' + p['cluster'] + '/namespaces/' + p['namespace'] + '/secrets/' + p['name'] + '/actions/edit'
	delete p['cluster']
	delete p['namespace']
	delete p['name']
	return xhr(path, 'POST', conf, null, p);
}

module.exports = {
  Clients: {
      List: ClientsList,
      Describe: ClientsDescribe,
      Delete: ClientsDelete,
      Update: ClientsUpdate,
      Copy: ClientsCopy,
      EditConfigMap: ClientsEditConfigMap,
      EditSecret: ClientsEditSecret
  }
}
