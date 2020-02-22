package models

type Object struct {
	ObjectID   int `gorm:"primary_key;"`
	ObjectName string
}

func (object *Object) TableName() string {
	// custom table name, this is default
	return "statuses.objects"
}

// func (object *Object) Create() map[string]interface{} {

// 	if err := database.DB.Create(object).Error; err != nil {
// 		errmsg := err.Error()
// 		resp := u.Message(false, errmsg)
// 		return resp
// 	}

// 	resp := u.Message(true, "success")
// 	resp["object"] = object
// 	return resp
// }
