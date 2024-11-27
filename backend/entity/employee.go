package entity


import (
	
   "time"
   "gorm.io/gorm"
)

type Employee struct {

   gorm.Model

   FirstName string    `json:"first_name"`
   LastName  string    `json:"last_name"`
   Email     string    `json:"email"`
   Age       uint8     `json:"age"`
   Password  string    `json:"-"`
   BirthDay  time.Time `json:"birthday"`

   GenderID  uint      `json:"gender_id"`
   Gender    *Genders  `gorm:"foreignKey: GenderID" json:"gender"`

   RoleID  uint      `json:"roles_id"`
   Role    *Roles  `gorm:"foreignKey: RoleID" json:"role"`

   StatusID uint	 `json:"status_id"`
   Status  *Status `gorm:"foreignKey: StatusID" json:"status"`

}