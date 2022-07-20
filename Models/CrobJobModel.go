package Models

import (
	"time"
)

type CronJob struct {
	Id 				int 		`json:"_id"`
	Name			string		`json:"name"`
	Expression		string		`json:"expression"`
	URL				string		`json:"url"`
	Timezone		string		`json:"timezone"`
	HttpMethod		string		`json:"http_method"`
	HttpHeader		string		`json:"http_header"`
	PostData		string		`json:"post_data"`
	RetryCount		int			`json:"retry_count"`
	UserId			int			`json:"user_id"`
	Status			int 		`json:"status"`
	CreatedAt		time.Time	`json:"created_at"`
	UpdatedAT		time.Time	`json:"updated_at"`

}

func (cj *CronJob) TableName() string {
	return "cronjob"
}
