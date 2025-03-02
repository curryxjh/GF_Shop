// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"GF_Shop/internal/dao/internal"
)

// internalUserCouponInfoDao is an internal type for wrapping the internal DAO implementation.
type internalUserCouponInfoDao = *internal.UserCouponInfoDao

// userCouponInfoDao is the data access object for the table user_coupon_info.
// You can define custom methods on it to extend its functionality as needed.
type userCouponInfoDao struct {
	internalUserCouponInfoDao
}

var (
	// UserCouponInfo is a globally accessible object for table user_coupon_info operations.
	UserCouponInfo = userCouponInfoDao{
		internal.NewUserCouponInfoDao(),
	}
)

// Add your custom methods and functionality below.
