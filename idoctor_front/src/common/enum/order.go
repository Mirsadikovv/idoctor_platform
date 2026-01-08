package enum

type Order string

const (
	ASC  Order = "ASC"
	DESC Order = "DESC"
)

func (o Order) IsValid() bool {
	return o == ASC || o == DESC
}

func (o Order) String() string {
	return string(o)
}
