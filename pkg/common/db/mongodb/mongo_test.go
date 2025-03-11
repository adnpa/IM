package mongodb

import "testing"

type TestObj struct {
	Id   int64
	Name string
	Age  int
}

func TestInsert(t *testing.T) {
	obj := &TestObj{1, "jjj", 18}
	err := Insert("ttt", obj)
	t.Log(err)
}

func TestGet(t *testing.T) {
	tar := &TestObj{}
	obj, err := GetById("msg", 1)
	t.Log(err)
	r, _ := obj.Raw()
	t.Log(r.Validate())
	t.Log(r.String())
	obj.Decode(&tar)
	t.Log(tar.Id)
	t.Log(tar.Name)
}
