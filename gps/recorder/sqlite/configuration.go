package sqlite

import "time"

type configuration struct {
	Filepath string        `json:"db_filepath"`
	Period   time.Duration `json:"flush_period"`
}
