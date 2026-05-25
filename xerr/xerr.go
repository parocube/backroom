package xerr

import "errors"

var ErrOptimisticLockConflict = errors.New("optimistic lock conflict")
var ErrInvalidUUID = errors.New("invalid uuid")
