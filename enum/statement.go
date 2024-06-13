package enum

type Statement int

const (
	Select Statement = iota
	Insert
	Delete
	Update
)

func (s Statement) ToString() string {
	return [...]string{"SELECT", "INSERT", "DELETE", "UPDATE"}[s]
}

func (s Statement) ToInt() int {
	return int(s)
}
