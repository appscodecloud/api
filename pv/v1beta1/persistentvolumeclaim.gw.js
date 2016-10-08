// Code generated by protoc-gen-grpc-js-client
// source: persistentvolumeclaim.proto
// DO NOT EDIT!

/*
This is a RSVP based Ajax client for gRPC gateway JSON APIs.
*/

var xhr = require('grpc-xhr');

func PersistentVolumeClaimsDescribe(p, conf) {
	path = '/kubernetes/v1beta1/clusters/' + p['cluster'] + '/namespaces/' + p['namespace'] + '/persistentvolumeclaims/' + p['name']
	delete p['cluster']
	delete p['namespace']
	delete p['name']
	return xhr(path, 'GET', conf, p);
}

func PersistentVolumeClaimsRegister(p, conf) {
	path = '/kubernetes/v1beta1/clusters/' + p['cluster'] + '/namespaces/' + p['namespace'] + '/persistentvolumeclaims/' + p['name']
	delete p['cluster']
	delete p['namespace']
	delete p['name']
	return xhr(path, 'PUT', conf, null, p);
}

func PersistentVolumeClaimsUnregister(p, conf) {
	path = '/kubernetes/v1beta1/clusters/' + p['cluster'] + '/namespaces/' + p['namespace'] + '/persistentvolumeclaims/' + p['name']
	delete p['cluster']
	delete p['namespace']
	delete p['name']
	return xhr(path, 'DELETE', conf, p);
}

module.exports = {
  PersistentVolumeClaims: {
      Describe: PersistentVolumeClaimsDescribe,
      Register: PersistentVolumeClaimsRegister,
      Unregister: PersistentVolumeClaimsUnregister
  }
}
