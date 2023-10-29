package types

import "sync"

// AccountRolesDictionary map[int32][]string
type AccountRolesDictionary = *sync.Map
