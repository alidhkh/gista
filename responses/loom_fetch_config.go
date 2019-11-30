package responses

import "github.com/alidhkh/gista/models"

type LoomFetchConfig struct {
	Response
	SystemControl models.SystemControl `json:"system_control"`
	TraceControl  models.TraceControl  `json:"trace_control"`
	Id            int                  `json:"id"`
}
