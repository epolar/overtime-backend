package repository

type Repository struct {
	db *DBCli
}

func NewRepository(db *DBCli) *Repository {
	return &Repository{db: db}
}

func (r *Repository) save(v interface{}) error {
	if err := r.db.Save(v).Error; err != nil {
		return err
	}
	return nil
}