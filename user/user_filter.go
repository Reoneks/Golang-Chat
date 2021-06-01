package user

type UserFilter struct {
	FirstName *string
	LastName  *string
	Email     *string
	Status    []int64
}
