package storage

type Storage interface {
	StorageUser      // StorageUser - crud interface for working with users
	StorageDatapoint // StorageDatapoint - all commands related to datapoints
}
