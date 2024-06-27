package models

type Auth struct {
	ID     int    `gorm:"primaryKey;type:bigint(20);not null;autoIncrement"`
	ApiKey string `gorm:"type:varchar(32);not null"`
}

func (Auth) TableName() string {
	return "auth"
}

type User struct {
	ID          int         `gorm:"primaryKey;type:bigint(20);not null;autoIncrement"`
	Username    string      `gorm:"type:varchar(64);not null"`
	UserProfile UserProfile `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	UserData    UserData    `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
}

func (User) TableName() string {
	return "user"
}

type UserProfile struct {
	UserID    int    `gorm:"primaryKey;type:bigint(20);not null"`
	FirstName string `gorm:"type:varchar(32);not null"`
	LastName  string `gorm:"type:varchar(64);not null"`
	Phone     string `gorm:"type:varchar(64);not null"`
	Address   string `gorm:"type:varchar(64);not null"`
	City      string `gorm:"type:varchar(64);not null"`
}

func (UserProfile) TableName() string {
	return "user_profile"
}

type UserData struct {
	UserID int    `gorm:"primaryKey;type:bigint(20);not null"`
	School string `gorm:"type:varchar(32);not null"`
}

func (UserData) TableName() string {
	return "user_data"
}

type ProfileAPI struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name" gorm:"column:UserProfile__first_name"`
	LastName  string `json:"last_name" gorm:"column:UserProfile__last_name"`
	City      string `json:"city" gorm:"column:UserProfile__city"`
	School    string `json:"school" gorm:"column:UserData__school"`
}
