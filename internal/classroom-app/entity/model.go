package entity

import "database/sql"

type Models struct {
	Classrooms ClassroomModel
	Tasks      TaskModel
	Users      UserModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Classrooms: ClassroomModel{
			DB: db,
		},
		Tasks: TaskModel{
			DB: db,
		},
		Users: UserModel{
			DB: db,
		},
	}
}
