// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"GF_Shop/internal/dao/internal"
)

// internalPositionInfoDao is an internal type for wrapping the internal DAO implementation.
type internalPositionInfoDao = *internal.PositionInfoDao

// positionInfoDao is the data access object for the table position_info.
// You can define custom methods on it to extend its functionality as needed.
type positionInfoDao struct {
	internalPositionInfoDao
}

var (
	// PositionInfo is a globally accessible object for table position_info operations.
	PositionInfo = positionInfoDao{
		internal.NewPositionInfoDao(),
	}
)

// Add your custom methods and functionality below.
