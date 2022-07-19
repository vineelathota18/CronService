package Models

type PermissionType struct{
	Id 		int 	`json:"_id"`
	Name	string	`json:"name"`
}

func (pt *PermissionType) TableName() string {
	return "permission-type"
}
