package dbtest

import (
	"reflect"

	"github.com/anuchito/dbstore/pb"
)

func IsEquals(want, got *pb.Entity) bool {
	key := want.Key == got.Key
	val := reflect.DeepEqual(want.Value, got.Value)
	return key && val
}
