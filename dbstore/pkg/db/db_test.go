package db

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/anuchito/dbstore/pb"
)

func setupFile(t *testing.T) (*os.File, func()) {
	t.Parallel()

	f, err := ioutil.TempFile("", "dbstore")
	if err != nil {
		t.Fatalf("error creating temp dir: %v", err)
	}

	teardown := func() {
		f.Close()
		os.Remove(f.Name())
	}

	return f, teardown
}

func TestSingleGet(t *testing.T) {
	f, teardown := setupFile(t)
	defer teardown()
	db := New(f)
	key := "foo-key"
	value := "foo-value"
	entity := &pb.Entity{Tombstone: false, Key: key, Value: []byte(value)}
	err := db.Set(entity)
	if err != nil {
		t.Fatalf("error append")
	}
	readEntity, err := db.Get(key)
	if err != nil {
		t.Fatalf("error getting entity %v", err)
	}
	if !reflect.DeepEqual(entity, readEntity) {
		t.Fatalf("expected %v, got %v", entity, readEntity)
	}
}

func TestMultipleGet(t *testing.T) {
	f, teardown := setupFile(t)
	defer teardown()
	db := New(f)
	key := "foo-key"
	value := "foo-value"
	entity := &pb.Entity{Tombstone: false, Key: key, Value: []byte(value)}
	err := db.Set(entity)
	if err != nil {
		t.Fatalf("error setting entity 1")
	}
	key1 := "foo-key-1"
	value1 := "foo-value-1"
	entity1 := &pb.Entity{Tombstone: false, Key: key1, Value: []byte(value1)}
	err = db.Set(entity1)
	if err != nil {
		t.Fatalf("error setting entity 2")
	}
	readEntity, err := db.Get(key)
	if err != nil {
		t.Fatalf("error getting entity %v", err)
	}
	readEntity1, err := db.Get(key1)
	if !reflect.DeepEqual(entity, readEntity) {
		t.Fatalf("expected %v, got %v", entity, readEntity)
	}
	if !reflect.DeepEqual(entity1, readEntity1) {
		t.Fatalf("expected %v, got %v", entity1, readEntity1)
	}
}

func TestSingleDelete(t *testing.T) {
	// prepare
	f, teardown := setupFile(t)
	defer teardown()
	db := New(f)
	key := "foo-key"
	value := "foo-value"
	entity := &pb.Entity{Tombstone: false, Key: key, Value: []byte(value)}
	err := db.Set(entity)
	if err != nil {
		t.Fatalf("error append")
	}
	err = db.Delete(key)
	readEntity, err := db.Get(key)
	if readEntity != nil || err != nil {
		t.Fatalf("readEntity expected nil, got %v", readEntity)
	}
}

func TestSingleRecover(t *testing.T) {
	// prepare
	f, teardown := setupFile(t)
	defer teardown()
	db := New(f)
	key := "foo-key"
	value := "foo-value"
	entity := &pb.Entity{Tombstone: false, Key: key, Value: []byte(value)}
	err := db.Set(entity)
	if err != nil {
		t.Fatalf("error append")
	}

	// clear map
	db.offsetMap = make(map[string]int64)

	err = db.Recover()
	if err != nil {
		t.Fatalf("error recovering %v", err)
	}

	readEntity, err := db.Get(key)
	if err != nil {
		t.Fatalf("error deleting entity %v", err)
	}
	if !reflect.DeepEqual(entity, readEntity) {
		t.Fatalf("error entities not equal after recovering")
	}

}

func TestSingleRecoverWithDelete(t *testing.T) {
	// prepare
	f, teardown := setupFile(t)
	defer teardown()
	db := New(f)
	key := "foo-key"
	value := "foo-value"
	entity := &pb.Entity{Tombstone: false, Key: key, Value: []byte(value)}
	err := db.Set(entity)
	if err != nil {
		t.Fatalf("error SET")
	}

	err = db.Delete(key)
	if err != nil {
		t.Fatalf("error DELETE")
	}

	// clear map
	db.offsetMap = make(map[string]int64)

	err = db.Recover()
	if err != nil {
		t.Fatalf("error recovering %v", err)
	}

	readEntity, err := db.Get(key)
	if readEntity != nil || err != nil {
		t.Fatalf("readEntity expected nil, got %v", readEntity)
	}
}

func TestMultipleRecover(t *testing.T) {
	f, teardown := setupFile(t)
	defer teardown()
	db := New(f)

	// first item
	key := "foo-key"
	value := "foo-value"
	entity := &pb.Entity{Tombstone: false, Key: key, Value: []byte(value)}
	err := db.Set(entity)
	if err != nil {
		t.Fatalf("error setting entity 1")
	}

	// second item
	key1 := "foo-key-1"
	value1 := "foo-value-1"
	entity1 := &pb.Entity{Tombstone: false, Key: key1, Value: []byte(value1)}
	err = db.Set(entity1)
	if err != nil {
		t.Fatalf("error setting entity 2")
	}

	// third item
	key2 := "foo-key-2"
	value2 := "foo-value-2"
	entity2 := &pb.Entity{Tombstone: false, Key: key2, Value: []byte(value2)}
	err = db.Set(entity2)
	if err != nil {
		t.Fatalf("error setting entity 3")
	}

	// act
	// clear map
	db.offsetMap = make(map[string]int64)
	err = db.Recover()
	if err != nil {
		t.Fatalf("error recovering %v", err)
	}
	readEntity, err := db.Get(key)
	if err != nil {
		t.Fatalf("error getting entity %v", err)
	}
	readEntity1, err := db.Get(key1)
	if err != nil {
		t.Fatalf("error getting entity1 %v", err)
	}
	readEntity2, err := db.Get(key2)
	if err != nil {
		t.Fatalf("error getting entity2 %v", err)
	}

	// assert
	if !reflect.DeepEqual(entity, readEntity) {
		t.Fatalf("expected %v, got %v", entity, readEntity)
	}
	if !reflect.DeepEqual(entity1, readEntity1) {
		t.Fatalf("expected %v, got %v", entity1, readEntity1)
	}
	if !reflect.DeepEqual(entity2, readEntity2) {
		t.Fatalf("expected %v, got %v", entity2, readEntity2)
	}
}
