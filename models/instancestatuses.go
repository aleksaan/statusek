package models

import (
	"time"

	rc "github.com/aleksaan/statusek/returncodes"
	"gorm.io/gorm"
)

type InstanceStatus struct {
	Status
	IsSet   bool
	SetAt   time.Time
	EventID uint
}

func GetInstanceStatuses(db *gorm.DB, statuses *[]InstanceStatus, instanceToken string) rc.ReturnCode {

	query := `select
		s.id,
		s.status_name,
		s.status_desc,
		s.status_type,
		case 
			when e.id is null
				then false
			else true
		end as is_set,
		e.created_at  as set_at,
		e.id as event_id
	from instances i
		inner join statuses s
			on i.object_id = s.object_id
		left join events e
			on s.id = e.status_id and i.id = e.instance_id
	where 
		i.instance_token = ?`

	db.Raw(query, instanceToken).Scan(&statuses)

	return rc.SUCCESS
}
