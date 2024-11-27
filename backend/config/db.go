package config


import (

   "fmt"

   "time"

   "github.com/TANTA1-KW/SE67/entity"

   "gorm.io/driver/sqlite"

   "gorm.io/gorm"

)


var db *gorm.DB


func DB() *gorm.DB {

   return db

}


func ConnectionDB() {

   database, err := gorm.Open(sqlite.Open("sa.db?cache=shared"), &gorm.Config{})

   if err != nil {

       panic("failed to connect database")

   }

   fmt.Println("connected database")

   db = database

}


func SetupDatabase() {


   db.AutoMigrate(

       &entity.Users{},
       &entity.Genders{},

       &entity.Employee{},
       &entity.Roles{},
       &entity.Status{},
   )


   GenderMale := entity.Genders{Gender: "Male"}
   GenderFemale := entity.Genders{Gender: "Female"}
   db.FirstOrCreate(&GenderMale, &entity.Genders{Gender: "Male"})
   db.FirstOrCreate(&GenderFemale, &entity.Genders{Gender: "Female"})


   hashedPassword01, _ := HashPassword("123456")
   BirthDay01, _ := time.Parse("2006-01-02", "1988-11-12")

   User := &entity.Users{
       FirstName: "Software",
       LastName:  "Engineer",
       Email:     "se@gmail.com",
       Age:       80,
       Password:  hashedPassword01,
       BirthDay:  BirthDay01,
       GenderID:  1,
   }
   db.FirstOrCreate(User, &entity.Users{
       Email: "se@gmail.com",
   })

   Captain := entity.Roles{Role: "Captain"}
   Crew := entity.Roles{Role: "Crew"}
   Mechanics := entity.Roles{Role: "Mechanics"}
   Engineer := entity.Roles{Role: "Engineer"}
   Receptionist := entity.Roles{Role: "Receptionist"}
   Housekeeper := entity.Roles{Role: "Housekeeper"}
   Chef := entity.Roles{Role: "Chef"}
   Waiter := entity.Roles{Role: "Waiter"}
   Doctor := entity.Roles{Role: "Doctor"}
   Nurse := entity.Roles{Role: "Nurse"}
   Security := entity.Roles{Role: "Security"}

   db.FirstOrCreate(&Captain, &entity.Roles{Role: "Captain"})
   db.FirstOrCreate(&Crew, &entity.Roles{Role: "Crew"})
   db.FirstOrCreate(&Mechanics, &entity.Roles{Role: "Mechanics"})
   db.FirstOrCreate(&Engineer, &entity.Roles{Role: "Engineer"})
   db.FirstOrCreate(&Receptionist, &entity.Roles{Role: "Receptionist"})
   db.FirstOrCreate(&Housekeeper, &entity.Roles{Role: "Housekeeper"})
   db.FirstOrCreate(&Chef, &entity.Roles{Role: "Chef"})
   db.FirstOrCreate(&Waiter, &entity.Roles{Role: "Waiter"})
   db.FirstOrCreate(&Doctor, &entity.Roles{Role: "Doctor"})
   db.FirstOrCreate(&Nurse, &entity.Roles{Role: "Nurse"})
   db.FirstOrCreate(&Security, &entity.Roles{Role: "Security"})

   Working := entity.Status{Status: "Working"}
   Sick := entity.Status{Status: "Sick"}
   OnLeave := entity.Status{Status: "On Leave"}

   db.FirstOrCreate(&Working, &entity.Status{Status: "Working"})
   db.FirstOrCreate(&Sick, &entity.Status{Status: "Sick"})
   db.FirstOrCreate(&OnLeave, &entity.Status{Status: "On Leave"})

   hashedPassword02, _ := HashPassword("123456")
   BirthDay02, _ := time.Parse("2006-01-02", "2000-12-12")

   Employee  := &entity.Employee{
       FirstName: "S",
       LastName: "E",
       Email: "67@gmail.com",
       Age: 67,
       Password:  hashedPassword02,
       BirthDay:  BirthDay02,
       GenderID:  1,
       RoleID:  1,
       StatusID: 1,
   }
   db.FirstOrCreate(Employee, &entity.Employee{
    Email: "67@gmail.com",
})
}