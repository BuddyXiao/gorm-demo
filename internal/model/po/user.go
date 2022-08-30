package po

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User struct {
	BaseModel
	Username string `gorm:"column:user_name"`
	Password string `gorm:"column:password"`
	RealName string `gorm:"column:real_name"`
	Phone    int    `gorm:"column:phone"`
	Province string `gorm:"column:province"`
	City     string `gorm:"column:city"`
	County   string `gorm:"column:county"`
	UnitId   int    `gorm:"column:unit_id"`
	UnitName string `gorm:"column:unit_name"`
	Deleted  int    `gorm:"column:deleted"`
	// 多对多关系
	Roles []Role `gorm:"many2many:t_user_role"`
}

func (User) TableName() string {
	return "t_user"
}

//func (u *User) AfterFind(tx *gorm.DB) (err error) {
//	fmt.Println("AfterFind 执行了")
//	condition := tx.Statement.BuildCondition("deleted > 0 ")
//	tx.Statement.AddClause(clause.Where{Exprs: condition})
//	return
//}

// BeforeDelete Updating data in same transaction
func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
	fmt.Println("BeforeDelete 执行了")
	fmt.Println("userID:", u.ID)
	// 将delete 改成 update
	fmt.Println("置空后 sql:" + tx.Statement.SQL.String())
	clauses := []clause.Interface{
		clause.Update{},
		clause.Set([]clause.Assignment{{Column: clause.Column{Name: "deleted"}, Value: 100}}),
		clause.Where{Exprs: []clause.Expression{clause.Eq{Column: clause.PrimaryColumn, Value: u.ID}}},
	}
	for _, c := range clauses {
		tx.Statement.AddClause(c)
	}
	tx.Statement.Build("UPDATE", "SET", "WHERE")
	//db, _ := tx.DB()
	//result, err := db.Exec(tx.Statement.SQL.String(), 100, u.ID)
	//_, err = result.RowsAffected()
	fmt.Println(tx.Statement.SQL.String()) // UPDATE `t_user` SET `deleted`=? WHERE `t_user`.`id` = ?
	return
}
