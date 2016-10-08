// Code generated by protoc-gen-grpc-js-client
// source: snapshot.proto
// DO NOT EDIT!

/*
This is a RSVP based Ajax client for gRPC gateway JSON APIs.
*/

var xhr = require('grpc-xhr');

func SnapshotsList(p, conf) {
	path = '/kubernetes/v1beta1/clusters/' + p['cluster'] + '/databases/' + p['uid'] + '/snapshots'
	delete p['cluster']
	delete p['uid']
	return xhr(path, 'GET', conf, p);
}

func SnapshotsBackupSchedule(p, conf) {
	path = '/kubernetes/v1beta1/clusters/' + p['cluster'] + '/databases/' + p['uid'] + '/actions/schedule-backup'
	delete p['cluster']
	delete p['uid']
	return xhr(path, 'POST', conf, null, p);
}

func SnapshotsBackupUnschedule(p, conf) {
	path = '/kubernetes/v1beta1/clusters/' + p['cluster'] + '/databases/' + p['uid'] + '/actions/unschedule-backup'
	delete p['cluster']
	delete p['uid']
	return xhr(path, 'PUT', conf, p);
}

func SnapshotsRestore(p, conf) {
	path = '/kubernetes/v1beta1/clusters/' + p['cluster'] + '/databases/' + p['uid'] + '/actions/restore'
	delete p['cluster']
	delete p['uid']
	return xhr(path, 'POST', conf, null, p);
}

module.exports = {
  Snapshots: {
      List: SnapshotsList,
      BackupSchedule: SnapshotsBackupSchedule,
      BackupUnschedule: SnapshotsBackupUnschedule,
      Restore: SnapshotsRestore
  }
}
