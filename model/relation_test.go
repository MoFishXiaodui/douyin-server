package model

import (
	"reflect"
	"testing"
)

func TestRelationCreate(t *testing.T) {
	err := MySQLInit()
	if err != nil {
		t.Errorf("sth wrong happened when init MySQL: %v", err)
	}
	//通过 db 对象执行数据库操作，  db 数据库连接对象
	//并将操作的结果赋值给 user 变量 （单个用户)

	relation0 := &Relation{UserId1: 1, UserId2: 2, Relationship: 1}
	_, _, err = NewRelationDaoInstance().Create(relation0)
	if err != nil {
		t.Errorf("something wrong in the Create, err: %v ", err)
	}

	relation1 := &Relation{UserId1: 2, UserId2: 1, Relationship: 1}
	_, _, err = NewRelationDaoInstance().Create(relation1)
	if err == nil {
		t.Errorf("something wrong in the Create, err: %v ", err)
	}

	_, _, err = NewRelationDaoInstance().Create(relation1)
	if err == nil {
		t.Errorf("something wrong in the Create, err: %v", err)
	}

	relation2 := &Relation{UserId1: 1, UserId2: 1, Relationship: 2}
	_, _, err = NewRelationDaoInstance().Create(relation2)
	if err == nil {
		t.Errorf("something wrong in the Create, err: %v ", err)
	}

	relation3 := &Relation{UserId1: 3, UserId2: 1, Relationship: 2}
	_, _, err = NewRelationDaoInstance().Create(relation3)
	if err != nil {
		t.Errorf("something wrong in the Create, err: %v ", err)
	}

	relation4 := &Relation{UserId1: 1, UserId2: 6, Relationship: 3}
	_, _, err = NewRelationDaoInstance().Create(relation4)
	if err != nil {
		t.Errorf("something wrong in the Create, err: %v ", err)
	}

	relation5 := &Relation{UserId1: 1, UserId2: 8, Relationship: 2}
	_, _, err = NewRelationDaoInstance().Create(relation5)
	if err != nil {
		t.Errorf("something wrong in the Create, err: %v ", err)
	}

	_, _, _ = NewRelationDaoInstance().Create(&Relation{UserId1: 7, UserId2: 1, Relationship: 1})
}

func TestRelationQueryOne2One(t *testing.T) {
	err := MySQLInit()
	if err != nil {
		t.Errorf("sth wrong happened when init MySQL: %v", err)
	}
	real1 := NewRelationDaoInstance().QueryOne2One1(1, 2)
	real2 := NewRelationDaoInstance().QueryOne2One1(2, 2)
	real3 := NewRelationDaoInstance().QueryOne2One1(3, 1)
	real4 := NewRelationDaoInstance().QueryOne2One1(2, 1)
	if real1 != 1 {
		t.Errorf("Expected %v do not match actual %v", 1, real1)
	}
	if real2 != -1 {
		t.Errorf("Expected %v do not match actual %v", -1, real2)
	}
	if real3 != 2 {
		t.Errorf("Expected %v do not match actual %v", 2, real2)
	}
	if real4 != -1 {
		t.Errorf("Expected %v do not match actual %v", -1, real2)
	}

	real5 := NewRelationDaoInstance().QueryOne2One2(2, 1)
	real6 := NewRelationDaoInstance().QueryOne2One2(2, 2)
	real7 := NewRelationDaoInstance().QueryOne2One2(1, 3)
	real8 := NewRelationDaoInstance().QueryOne2One2(1, 2)

	if real5 != 1 {
		t.Errorf("Expected %v do not match actual %v", 1, real1)
	}
	if real6 != -1 {
		t.Errorf("Expected %v do not match actual %v", -1, real2)
	}
	if real7 != 2 {
		t.Errorf("Expected %v do not match actual %v", 2, real2)
	}
	if real8 != -1 {
		t.Errorf("Expected %v do not match actual %v", -1, real2)
	}
}

func TestRelationQueryOne2Many(t *testing.T) {
	err := MySQLInit()
	if err != nil {
		t.Errorf("sth wrong happened when init MySQL: %v", err)
	}
	real1 := NewRelationDaoInstance().QueryOne2Many(1)
	expected := []int64{2, 6, 3}
	if !reflect.DeepEqual(real1, expected) {
		t.Errorf("Expected %v do not match actual %v", expected, real1)
	}

}

func TestRelationQueryMany2One(t *testing.T) {
	err := MySQLInit()
	if err != nil {
		t.Errorf("sth wrong happened when init MySQL: %v", err)
	}
	real1 := NewRelationDaoInstance().QueryMany2one(1)
	expected := []int64{8, 6, 7}
	if !reflect.DeepEqual(real1, expected) {
		t.Errorf("Expected %v do not match actual %v", expected, real1)
	}
}

func TestRelationUpdate(t *testing.T) {
	err := MySQLInit()
	if err != nil {
		t.Errorf("sth wrong happened when init MySQL: %v", err)
	}
	actual1 := NewRelationDaoInstance().Update(1, 1, 1)
	actual2 := NewRelationDaoInstance().Update(1, 2, 3)
	actual3 := NewRelationDaoInstance().Update(1, 6, 2)
	//actual4 := NewRelationDaoInstance().Update(1, 3, 1)
	//actual5 := NewRelationDaoInstance().Update(6, 1, 2)
	//actual6 := NewRelationDaoInstance().Update(7, 1, 1)

	if actual1 != -1 {
		t.Errorf("Expected %v do not match actual %v", -1, actual1)
	}
	if actual2 != 3 {
		t.Errorf("Expected %v do not match actual %v", 3, actual2)
	}
	if actual3 != 2 {
		t.Errorf("Expected %v do not match actual %v", 2, actual3)
	}
	//if actual4 != 0 {
	//	t.Errorf("Expected %v do not match actual %v", 0, actual4)
	//}
	//if actual5 != 2 {
	//	t.Errorf("Expected %v do not match actual %v", 2, actual5)
	//}
	//if actual6 != 0 {
	//	t.Errorf("Expected %v do not match actual %v", 0, actual6)
	//}
}

func TestRelationDelete(t *testing.T) {
	err := MySQLInit()
	if err != nil {
		t.Errorf("sth wrong happened when init MySQL: %v", err)
	}
	actual1 := NewRelationDaoInstance().Delete(2, 1)
	actual2 := NewRelationDaoInstance().Delete(3, 6)
	actual3 := NewRelationDaoInstance().Delete(3, 1)
	actual4 := NewRelationDaoInstance().Delete(3, 3)

	if actual1 != 1 {
		t.Errorf("Expected %v do not match actual %v", 1, actual1)
	}
	if actual2 != -1 {
		t.Errorf("Expected %v do not match actual %v", -1, actual2)
	}
	if actual3 != 1 {
		t.Errorf("Expected %v do not match actual %v", 1, actual3)
	}
	if actual2 != -1 {
		t.Errorf("Expected %v do not match actual %v", -1, actual4)
	}
}
