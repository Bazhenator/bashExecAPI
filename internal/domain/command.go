package domain

type Command struct {
	ID      int     `db:"id" json:"id"`
	Command string  `db:"command" json:"command"`
	Result  *string `db:"result" json:"result"`
}
