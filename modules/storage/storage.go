package storage

type Storage interface {
	StorageUser        // StorageUser - crud interface for working with users
	StorageDatapoint   // StorageDatapoint - all commands related to datapoints
	StoragePermissions // StoragePermissions - all commands related to permissions (roles / endpoints)
	StorageDevices     // StorageDevices - all commands related to devices
	StorageEndpoints   // StorageEndpoints - all commands related to endpoints
	StorageRoles       // StorageRoles - all commands related to roles
}
