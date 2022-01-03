package database

import "github.com/NexClipper/sudory-prototype-r1/pkg/model"

func (d *DBManipulator) CreateClient(m *model.Client) (int64, error) {
	tx := d.session()
	tx.Begin()

	cnt, err := tx.Insert(m)

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	return cnt, err
}
