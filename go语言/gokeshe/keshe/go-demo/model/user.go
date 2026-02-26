package model

// User
type User struct {
	ID       int    `gorm:"primaryKey；AUTO_INCREMENT"` //设置为主键
	Username string `gorm:"not null;unique;size:255"`  //唯一，不为空
	Password string `gorm:"not null"`
	Mobile   string `gorm:"unique;not null;"`
}
