package models

import "StudentSystem/database"

// Student Model 结构体
type Student struct {
	UID string `json:"UID"`
	UserName string `json:"UserName"`
	CourseName string `json:"CourseName"`
	Score int `json:"Score"`
}

/*
	Student这个Model的增删改查操作都放在这里
*/
// CreateAStudent 创建Student
func CreateAStudent(Student *Student) (err error){
	err = database.DB.Create(&Student).Error
	return
}

func GetAllStudent() (StudentList []*Student, err error){
	if err = database.DB.Find(&StudentList).Error; err != nil{
		return nil, err
	}
	return
}

func GetAStudent(UID string)(student *Student, err error){
	student = new(Student)
	if err = database.DB.Debug().Where("UID=?", UID).First(student).Error; err!=nil{
		return nil, err
	}
	return
}

func UpdateAStudent(Student *Student)(err error){
	err = database.DB.Save(Student).Error
	return
}

func DeleteAStudent(UID string)(err error){
	err = database.DB.Where("UID=?", UID).Delete(&Student{}).Error
	return
}