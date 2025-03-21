// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"GF_Shop/internal/dao/internal"
)

// internalCollectionInfoDao is an internal type for wrapping the internal DAO implementation.
type internalCollectionInfoDao = *internal.CollectionInfoDao

// collectionInfoDao is the data access object for the table collection_info.
// You can define custom methods on it to extend its functionality as needed.
type collectionInfoDao struct {
	internalCollectionInfoDao
}

var (
	// CollectionInfo is a globally accessible object for table collection_info operations.
	CollectionInfo = collectionInfoDao{
		internal.NewCollectionInfoDao(),
	}
)

// Add your custom methods and functionality below.
