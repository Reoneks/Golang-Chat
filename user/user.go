package user

type StatusCode int64

//^ Status codes
const (
	Active   StatusCode = 1
	UnActive StatusCode = 2
	Banned   StatusCode = 3
)

func (s StatusCode) ToInt64() int64 {
	return int64(s)
}

func NewStatusCode(i int64) StatusCode {
	return StatusCode(i)
}

type User struct {
	Id        int64      `json:"id"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Email     string     `json:"email"`
	Password  string     `json:"-"`
	Status    StatusCode `json:"status"`
}
