package data

type Car struct {
	tableName struct{} `pg:"car"`
	ID        int      `pg:"id,pk" json:"-"`
	RegNum    string   `pg:"reg_num" json:"regNum"`
	Mark      string   `pg:"mark" json:"mark"`
	Model     string   `pg:"model" json:"model"`
	Year      int      `pg:"year" json:"year"`
	Owner     *Owner   `pg:"rel:has-one" json:"owner"`
	OwnerId   int      `pg:"owner_id" json:"-"`
}

type Owner struct {
	tableName  struct{} `pg:"owner"`
	ID         int      `pg:"id,pk" json:"-"`
	Name       string   `pg:"name" json:"name"`
	Surname    string   `pg:"surname" json:"surname"`
	Patronymic string   `pg:"patronymic" json:"patronymic"`
}
