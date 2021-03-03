package personrepository

type sqlPerson struct {
	ID        string `db:"id"`
	Firstname string `db:"firstname"`
	Lastname  string `db:"lastname"`
	Age       int    `db:"age"`
}
