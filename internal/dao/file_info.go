// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"GF_Shop/internal/dao/internal"
)

// internalFileInfoDao is an internal type for wrapping the internal DAO implementation.
type internalFileInfoDao = *internal.FileInfoDao

// fileInfoDao is the data access object for the table file_info.
// You can define custom methods on it to extend its functionality as needed.
type fileInfoDao struct {
	internalFileInfoDao
}

var (
	// FileInfo is a globally accessible object for table file_info operations.
	FileInfo = fileInfoDao{
		internal.NewFileInfoDao(),
	}
)

// Add your custom methods and functionality below.
