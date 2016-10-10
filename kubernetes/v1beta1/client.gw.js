// Code generated by protoc-gen-grpc-js-client
// source: client.proto
// DO NOT EDIT!

/*
This is a RSVP based Ajax client for gRPC gateway JSON APIs.
*/

var xhr = require('grpc-xhr');

function ClientsCopy(p, conf) {
    path = '/kubernetes/v1beta1/actions/copy'
    return xhr(path, 'PUT', conf, null, p);
}

function ClientsNodes(p, conf) {
    path = '/kubernetes/v1beta1/client/' + p['cluster'] + '/nodes'
    delete p['cluster']
    return xhr(path, 'GET', conf, p);
}

function ClientsApps(p, conf) {
    path = '/kubernetes/v1beta1/client/' + p['cluster'] + '/apps'
    delete p['cluster']
    return xhr(path, 'GET', conf, p);
}

function ClientsApp(p, conf) {
    path = '/kubernetes/v1beta1/client/' + p['cluster'] + '/apps/' + p['namespace'] + '/' + p['name']
    delete p['cluster']
    delete p['namespace']
    delete p['name']
    return xhr(path, 'GET', conf, p);
}

function ClientsAppPods(p, conf) {
    path = '/kubernetes/v1beta1/client/' + p['cluster'] + '/apps/' + p['namespace'] + '/' + p['name'] + '/pods'
    delete p['cluster']
    delete p['namespace']
    delete p['name']
    return xhr(path, 'GET', conf, p);
}

function ClientsNamespaces(p, conf) {
    path = '/kubernetes/v1beta1/client/' + p['cluster'] + '/namespaces'
    delete p['cluster']
    return xhr(path, 'GET', conf, p);
}

function ClientsSecrets(p, conf) {
    path = '/kubernetes/v1beta1/client/' + p['cluster'] + '/secrets'
    delete p['cluster']
    return xhr(path, 'GET', conf, p);
}

function ClientsSecret(p, conf) {
    path = '/kubernetes/v1beta1/client/' + p['cluster'] + '/secrets/' + p['namespace'] + '/' + p['name']
    delete p['cluster']
    delete p['namespace']
    delete p['name']
    return xhr(path, 'GET', conf, p);
}

function ClientsJobs(p, conf) {
    path = '/kubernetes/v1beta1/client/' + p['cluster'] + '/jobs'
    delete p['cluster']
    return xhr(path, 'GET', conf, p);
}

function ClientsPods(p, conf) {
    path = '/kubernetes/v1beta1/client/' + p['cluster'] + '/pods'
    delete p['cluster']
    return xhr(path, 'GET', conf, p);
}

function ClientsServices(p, conf) {
    path = '/kubernetes/v1beta1/client/' + p['cluster'] + '/services'
    delete p['cluster']
    return xhr(path, 'GET', conf, p);
}

function ClientsReplicationControllers(p, conf) {
    path = '/kubernetes/v1beta1/client/' + p['cluster'] + '/rcs'
    delete p['cluster']
    return xhr(path, 'GET', conf, p);
}

function ClientsConfigMaps(p, conf) {
    path = '/kubernetes/v1beta1/client/' + p['cluster'] + '/configmaps'
    delete p['cluster']
    return xhr(path, 'GET', conf, p);
}

function ClientsConfigMap(p, conf) {
    path = '/kubernetes/v1beta1/client/' + p['cluster'] + '/configmaps/' + p['namespace'] + '/' + p['name']
    delete p['cluster']
    delete p['namespace']
    delete p['name']
    return xhr(path, 'GET', conf, p);
}

module.exports = {
    Clients: {
        Copy: ClientsCopy,
        Nodes: ClientsNodes,
        Apps: ClientsApps,
        App: ClientsApp,
        AppPods: ClientsAppPods,
        Namespaces: ClientsNamespaces,
        Secrets: ClientsSecrets,
        Secret: ClientsSecret,
        Jobs: ClientsJobs,
        Pods: ClientsPods,
        Services: ClientsServices,
        ReplicationControllers: ClientsReplicationControllers,
        ConfigMaps: ClientsConfigMaps,
        ConfigMap: ClientsConfigMap
    }
};