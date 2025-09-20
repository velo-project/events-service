package ports

type UserExistsByIdPort interface {
	Execute(userId int) (bool, error)
}
