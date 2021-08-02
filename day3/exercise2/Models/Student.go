package Models
import (
	"fmt"
	Config2 "freshers-bootcamp/day3/exercise2/Config"
	_ "github.com/go-sql-driver/mysql"
)
//GetAllUsers Fetch all user data
func GetAllStudents(student *[]Student) (err error) {
	if err = Config2.DB.Find(student).Error; err != nil {
		return err
	}
	return nil
}
//CreateUser ... Insert New data
func CreateStudent(student *Student) (err error) {
	if err = Config2.DB.Create(student).Error; err != nil {
		return err
	}
	return nil
}
//GetUserByID ... Fetch only one user by Id
func GetStudentByID(student *Student, id string) (err error) {
	if err = Config2.DB.Where("id = ?", id).First(student).Error; err != nil {
		return err
	}
	return nil
}
//UpdateUser ... Update user
func UpdateStudent(student *Student, id string) (err error) {
	fmt.Println(student)
	Config2.DB.Save(student)
	return nil
}
//DeleteUser ... Delete user
func DeleteStudent(student *Student, id string) (err error) {
	Config2.DB.Where("id = ?", id).Delete(student)
	return nil
}