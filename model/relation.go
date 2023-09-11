package model

import (
	"errors"
	"gorm.io/gorm"
	"sync"
)

// 注：Relationship的值受Id的相对位置的影响，故后续传参数需注意

type Relation struct {
	//ID           uint `gorm:"primaryKey"`
	gorm.Model
	UserId1      int64
	UserId2      int64
	Relationship int64 // 0：互不关注  1：只有1关注了2  2：只有2关注了1  3：互相关注
}

type RelationDao struct {
}

var (
	relationDao  *RelationDao
	relationOnce sync.Once
)

//type RelationStatus int64
//
//const (
//	None RelationStatus = iota //用户不存在
//	Concern1
//	Concern2
//	Both //
//)

func RelationInit() error {
	// Migrate the schema
	return db.AutoMigrate(&Relation{})
}

func NewRelationDaoInstance() *RelationDao {
	relationOnce.Do(
		func() {
			relationDao = &RelationDao{}
		})
	return relationDao
}

// relation表中存在UserId1=1 UserId2=2时，则不允许存在UserId1=2 UserId2=1

func (*RelationDao) Create(r *Relation) (UserId1, UserId2 int64, err error) {
	if r.UserId1 == r.UserId2 {
		err = errors.New("The UserIds must be different")
		return -1, -1, err
	}
	err = db.Where("(user_id1 = ? AND user_id2 = ?) OR (user_id2 = ? AND user_id1 = ?)", r.UserId1, r.UserId2, r.UserId1, r.UserId2).First(r).Error
	if err == nil {
		err = errors.New("The relationship has already existed")
		return -1, -1, err
	}

	err = db.Create(r).Error
	if err != nil {
		return -1, -1, err
	}
	return r.UserId1, r.UserId2, nil
}

// 查询两个用户之间的关系，由于两个用户在relation表中的相对位置会影响relationship的值，故设计两个查询函数
// 后续查询两个用户之间的关系时需同时调用两个函数，并根据使用的哪个函数及其返回值来判断两者关系
// 如使用QueryOne2One1返回1，则代表第一个参数代表的用户关注第二个参数代表的用户
// 使用QueryOne2One2返回1，则代表第二个参数代表的用户关注第一个参数代表的用户，可参照relation表辅助理解

// QueryOne2One1 第一个参数是relation表中的UserId1， 第二个参数是UserID2,返回值是两者的Relationship
func (*RelationDao) QueryOne2One1(UserId1, UserId2 int64) int64 {
	if UserId1 == UserId2 {
		return -1
	}
	relation := &Relation{}
	err := db.First(relation, "user_id1 = ? AND user_id2 = ?", UserId1, UserId2).Error

	if err == nil {
		return relation.Relationship
	} else {
		return -1
	}
}

// QueryOne2One2 第一个参数是relation表中的UserId2， 第二个参数是UserID1,返回值是两者的Relationship
func (*RelationDao) QueryOne2One2(UserId2, UserId1 int64) int64 {
	if UserId1 == UserId2 {
		return -1
	}
	relation := &Relation{}
	err := db.First(relation, "user_id1 = ? AND user_id2 = ?", UserId1, UserId2).Error

	if err == nil {
		return relation.Relationship
	} else {
		return -1
	}
}

// QueryOne2Many 关注列表，返回所有关注者的ID
func (*RelationDao) QueryOne2Many(UserId int64) []int64 {
	var relationships []Relation
	var part_relationship []Relation
	db.Where("user_id1 = ? AND relationship = ?", UserId, 1).Find(&part_relationship)
	relationships = append(relationships, part_relationship...)

	db.Where("(user_id1 = ? OR user_id2 = ?) AND relationship = ?", UserId, UserId, 3).First(&part_relationship)
	relationships = append(relationships, part_relationship...)

	db.Where("user_id2 = ? AND relationship = ?", UserId, 2).Find(&part_relationship)
	relationships = append(relationships, part_relationship...)

	users := []int64{}
	for _, subslice := range relationships {
		//fmt.Println(subslice)
		if subslice.UserId1 != UserId {
			users = append(users, subslice.UserId1)
		} else if subslice.UserId2 != UserId {
			users = append(users, subslice.UserId2)
		}
	}
	return users
}

// QueryMany2one 粉丝列表,返回所有粉丝的ID
func (*RelationDao) QueryMany2one(UserId int64) []int64 {
	var relationships []Relation
	var part_relationship []Relation
	db.Where("user_id1 = ? AND relationship = ?", UserId, 2).Find(&part_relationship)
	relationships = append(relationships, part_relationship...)

	db.Where("(user_id1 = ? OR user_id2 = ?) AND relationship = ?", UserId, UserId, 3).First(&part_relationship)
	relationships = append(relationships, part_relationship...)

	db.Where("user_id2 = ? AND relationship = ?", UserId, 1).Find(&part_relationship)
	relationships = append(relationships, part_relationship...)

	users := []int64{}
	for _, subslice := range relationships {
		if subslice.UserId1 != UserId {
			users = append(users, subslice.UserId1)
		} else if subslice.UserId2 != UserId {
			users = append(users, subslice.UserId2)
		}
	}
	return users
}

//	func (*UserDao) QuerywithId(id uint) UserStatus {
//		user := &User{}
//		err := db.First(user, "id = ?", id).Error
//		if err != nil {
//			return Inexistence
//		} else {
//			return Existence
//		}
//	}

func (*RelationDao) Update(UserId1, UserId2, r int64) int64 {
	if UserId1 == UserId2 {
		return -1
	}
	if r < 0 || r > 3 {
		return -1
	}
	relation := &Relation{}
	expect1 := NewRelationDaoInstance().QueryOne2One1(UserId1, UserId2)
	if expect1 != -1 {
		db.First(relation, "user_id1 = ? AND user_id2 = ?", UserId1, UserId2)
		db.Model(relation).Where("user_id1 = ? AND user_id2 = ?", UserId1, UserId2).Update("Relationship", r)
		db.Save(relation)
		return relation.Relationship
	} else {
		return -1
	}

	//	if r == 1 {
	//		if relation.Relationship == 0 {
	//			relation.Relationship = 1
	//		} else if relation.Relationship == 1 {
	//			relation.Relationship = 0
	//		} else if relation.Relationship == 2 {
	//			relation.Relationship = 3
	//		} else {
	//			relation.Relationship = 2
	//		}
	//	} else {
	//		if relation.Relationship == 0 {
	//			relation.Relationship = 2
	//		} else if relation.Relationship == 1 {
	//			relation.Relationship = 3
	//		} else if relation.Relationship == 2 {
	//			relation.Relationship = 0
	//		} else {
	//			relation.Relationship = 1
	//		}
	//	}
	//	db.Save(relation)
	//	return relation.Relationship
	//}
	//
	//expect2 := NewRelationDaoInstance().QueryOne2One2(UserId2, UserId1)
	//if expect2 != -1 {
	//	db.First(relation, "user_id1 = ? AND user_id2 = ?", UserId1, UserId2)
	//	if r == 1 {
	//		if relation.Relationship == 0 {
	//			relation.Relationship = 2
	//		} else if relation.Relationship == 1 {
	//			relation.Relationship = 3
	//		} else if relation.Relationship == 2 {
	//			relation.Relationship = 0
	//		} else {
	//			relation.Relationship = 1
	//		}
	//	} else {
	//		if relation.Relationship == 0 {
	//			relation.Relationship = 1
	//		} else if relation.Relationship == 1 {
	//			relation.Relationship = 0
	//		} else if relation.Relationship == 2 {
	//			relation.Relationship = 3
	//		} else {
	//			relation.Relationship = 2
	//		}
	//	}
	//}

}

func (*RelationDao) Delete(UserId1, UserId2 uint) int64 {
	if UserId1 == UserId2 {
		return -1
	}
	//expect := NewUserDaoIstance().QuerywithId(id)
	relation := &Relation{}
	//db.Where("ID = ?", id).Delete(user)
	err := db.First(relation, "(user_id1 = ? AND user_id2 = ?) OR (user_id1 = ? AND user_id2 = ?)", UserId1, UserId2, UserId2, UserId1).Delete(relation).Error
	if err != nil {
		return -1
	}
	return 1
}
