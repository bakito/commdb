package types

// Command is our command table structure.
type Command struct {
	ID       int64    `json:"id"` // auto-increment by-default by xorm
	Command  string   `xorm:"TEXT" json:"command"`
	Keywords []string `xorm:"TEXT" json:"keywords"`
}
