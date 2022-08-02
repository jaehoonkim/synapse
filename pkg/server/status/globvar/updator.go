package globvar

import (
	"database/sql"
	"time"

	// "github.com/NexClipper/sudory/pkg/server/control/vault"

	"github.com/NexClipper/logger"
	"github.com/NexClipper/sudory/pkg/server/database/vanilla"
	"github.com/NexClipper/sudory/pkg/server/macro"
	"github.com/NexClipper/sudory/pkg/server/macro/logs"
	globvarv2 "github.com/NexClipper/sudory/pkg/server/model/global_variables/v2"
	"github.com/NexClipper/sudory/pkg/server/status/state"
	"github.com/pkg/errors"
)

type GlobalVariantUpdate struct {
	*vanilla.SqlDbEx
	offset time.Time //updated column
}

func NewGlobalVariablesUpdate(db *sql.DB) *GlobalVariantUpdate {
	return &GlobalVariantUpdate{SqlDbEx: vanilla.NewSqlDbEx(db, state.ENV__CONTROL_TRANSACTION_TIMEOUT__())}
}

// Update
//  Update = read db -> global_variables
func (worker *GlobalVariantUpdate) Update() (err error) {
	records := make([]globvarv2.GlobalVariables, 0, state.ENV__INIT_SLICE_CAPACITY__())
	globvar := globvarv2.GlobalVariables{}
	globvar.Updated = *vanilla.NewNullTime(worker.offset)

	globvar_cond := vanilla.GreaterThan("updated", globvar.Updated)

	err = vanilla.Stmt.Select(globvar.TableName(), globvar.ColumnNames(), globvar_cond.Parse(), nil, nil).
		QueryRows(worker)(func(scan vanilla.Scanner, _ int) (err error) {
		err = globvar.Scan(scan)
		if err != nil {
			return errors.Wrapf(err, "failed to scan")
		}
		records = append(records, globvar)
		return
	})
	if err != nil {
		return
	}

	for i := range records {
		record := &records[i]
		gv, err := ParseKey(record.Name)

		switch err {
		case nil:
			if err := storeManager.Call(gv, record.Value.String); err != nil {
				return errors.Wrapf(err, "store global_variables%v",
					logs.KVL(
						"key", record.Name,
						"value", record.Value.String,
					))
			}
		default:
			logger.Warningf("%w: parse record name to key%v", err, logs.KVL(
				"key", record.Name,
			))
		}
	}

	//update offset
	worker.offset = time.Now()

	return
}

// WhiteListCheck
//  리스트 체크
func (worker *GlobalVariantUpdate) WhiteListCheck() (err error) {
	records := make([]globvarv2.GlobalVariables, 0, state.ENV__INIT_SLICE_CAPACITY__())

	globvar := globvarv2.GlobalVariables{}
	globvar.Updated = *vanilla.NewNullTime(worker.offset)
	globvar_cond := vanilla.IsNull("deleted")

	err = vanilla.Stmt.Select(globvar.TableName(), globvar.ColumnNames(), globvar_cond.Parse(), nil, nil).
		QueryRows(worker)(func(scan vanilla.Scanner, _ int) (err error) {
		err = globvar.Scan(scan)
		if err != nil {
			return errors.Wrapf(err, "failed to scan")
		}
		records = append(records, globvar)
		return
	})
	if err != nil {
		return
	}

	count := 0
	push, pop := macro.StringBuilder()
	for _, key := range KeyNames() {
		var found bool = false
	LOOP:
		for i := range records {
			if key == records[i].Name {
				found = true
				break LOOP
			}
		}
		if !found {
			count++
			push(key)
		}
	}
	if 0 < count {
		return errors.Errorf("not exists global_variables keys=['%s']", pop("', '"))
	}

	return nil
}

func (worker *GlobalVariantUpdate) Merge() (err error) {
	records := make([]globvarv2.GlobalVariables, 0, state.ENV__INIT_SLICE_CAPACITY__())

	globvar := globvarv2.GlobalVariables{}
	globvar.Updated = *vanilla.NewNullTime(worker.offset)

	globvar_cond := vanilla.IsNull("deleted")

	err = vanilla.Stmt.Select(globvar.TableName(), globvar.ColumnNames(), globvar_cond.Parse(), nil, nil).
		QueryRows(worker)(func(scan vanilla.Scanner, _ int) (err error) {
		err = globvar.Scan(scan)
		if err != nil {
			return errors.Wrapf(err, "failed to scan")
		}
		records = append(records, globvar)
		return
	})
	if err != nil {
		return
	}

	for _, key := range KeyNames() {
		var found bool = false
	LOOP:
		for i := range records {
			if key == records[i].Name {
				found = true
				break LOOP
			}
		}
		if !found {
			globvar_key, err := ParseKey(key)
			if err != nil {
				return errors.Wrapf(err, "ParseGlobVar%v",
					logs.KVL(
						"key", key,
					))
			}

			globvar, updated_columns, ok := GetDefaultGlobalVariable(globvar_key, time.Now())
			if !ok {
				return errors.Errorf("not found global_variables%v",
					logs.KVL(
						"key", key,
					))
			}

			stmt, err := vanilla.Stmt.InsertOrUpdate(globvar.TableName(), globvar.ColumnNames(), updated_columns, globvar.Values())
			if err != nil {
				return errors.Wrapf(err, "failed to make create or update statement")
			}

			_, err = stmt.Exec(worker)
			if err != nil {
				return errors.Wrapf(err, "failed to create or update global_variables")
			}
		}
	}

	return nil
}

// func foreach_environment(elems []envv1.Environment, fn func(elem envv1.Environment) bool) {
// 	for n := range elems {
// 		ok := fn(elems[n])
// 		if !ok {
// 			return
// 		}
// 	}
// }
