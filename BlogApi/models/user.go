package models

import (
	"github.com/astaxie/beego/orm"
)

var (
	UserList []*Users
)

func init() {
	// UserList = make(map[string]*User)
	// u := User{"user_11111", "tes", "12345", Profile{"male", 20, "jakarta", "tes@gmail.com"}}
	// UserList["user_11111"] = &u
	orm.RegisterModel(new(Users))
}

type Users struct {
	Id   int    `orm:"pk;column(id)" json:"Id"`
	Name string `orm:"column(name)" json:"Name"`
}

// func AddUser(u User) string {
// 	u.Id = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
// 	UserList[u.Id] = &u
// 	return u.Id
// }

// func GetUser(uid string) (u *User, err error) {
// 	if u, ok := UserList[uid]; ok {
// 		return u, nil
// 	}
// 	return nil, errors.New("User not exists")
// }

func GetAllUsers() []*Users {

	o := orm.NewOrm()
	// o.Raw("SELECT id, category_id,user_id,title,slug,body,post_image FROM Posts",).QueryRows(&PostList)
	o.QueryTable("users").All((&UserList))
	return UserList
}

// func UpdateUser(uid string, uu *User) (a *User, err error) {
// 	if u, ok := UserList[uid]; ok {
// 		if uu.Username != "" {
// 			u.Username = uu.Username
// 		}
// 		if uu.Password != "" {
// 			u.Password = uu.Password
// 		}
// 		if uu.Profile.Age != 0 {
// 			u.Profile.Age = uu.Profile.Age
// 		}
// 		if uu.Profile.Address != "" {
// 			u.Profile.Address = uu.Profile.Address
// 		}
// 		if uu.Profile.Gender != "" {
// 			u.Profile.Gender = uu.Profile.Gender
// 		}
// 		if uu.Profile.Email != "" {
// 			u.Profile.Email = uu.Profile.Email
// 		}
// 		return u, nil
// 	}
// 	return nil, errors.New("User Not Exist")
// }

// func Login(username, password string) bool {
// 	for _, u := range UserList {
// 		if u.Name == username && u.Name == password {
// 			return true
// 		}
// 	}
// 	return false
// }

// func DeleteUser(uid string) {
// 	delete(UserList, uid)
// }
