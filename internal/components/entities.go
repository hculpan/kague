package components

import (
	"fmt"
	"reflect"
)

var entities []Entity = make([]Entity, 0)

func AddEntity(e Entity) {
	e.SetEntityId(len(entities))
	entities = append(entities, e)
}

func RemoveEntity(id int) error {
	for i := range entities {
		if entities[i].GetEntityId() == id {
			entities = append(entities[:i], entities[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("unable to find entitiy with id '%d'", id)
}

func GetEntityCount() int {
	return len(entities)
}

func GetNextEntity(index int) Entity {
	if index >= 0 && index < len(entities) {
		return entities[index]
	}

	return nil
}

func CheckInterface(obj interface{}, interfaceType interface{}) bool {
	return reflect.TypeOf(obj).Implements(reflect.TypeOf(interfaceType).Elem())
}

func FindNextEntity(startingIndex int, itype interface{}) (Entity, int) {
	for i := startingIndex; i < len(entities); i++ {
		if CheckInterface(entities[i], itype) {
			return entities[i], i
		}
	}

	return nil, -1
}

func GetEntity(id int) (Entity, error) {
	for i := range entities {
		if entities[i].GetEntityId() == id {
			return entities[i], nil
		}
	}

	return nil, fmt.Errorf("unable to find entitiy with id '%d'", id)
}
