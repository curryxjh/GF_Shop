// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"GF_Shop/internal/dao/internal"
)

// internalRefundInfoDao is an internal type for wrapping the internal DAO implementation.
type internalRefundInfoDao = *internal.RefundInfoDao

// refundInfoDao is the data access object for the table refund_info.
// You can define custom methods on it to extend its functionality as needed.
type refundInfoDao struct {
	internalRefundInfoDao
}

var (
	// RefundInfo is a globally accessible object for table refund_info operations.
	RefundInfo = refundInfoDao{
		internal.NewRefundInfoDao(),
	}
)

// Add your custom methods and functionality below.
