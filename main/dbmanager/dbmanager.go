package dbmanager

import (
	"context"
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"strings"
)

type DatabaseManager struct {
	sqldb *sql.DB
	db    *bun.DB
}

func (dbm *DatabaseManager) Connect(url string) (err error) {
	dbm.sqldb, err = sql.Open("mysql", url)
	if err != nil {
		return err
	}
	dbm.db = bun.NewDB(dbm.sqldb, mysqldialect.New())
	return nil
}

func (dbm *DatabaseManager) Select(data interface{}, query string, values ...any) (err error) {

	ctx := context.Background()

	if len(query) > 0 {
		if strings.Count(query, "?") != len(values) {
			return errors.New("dbmanager select: arguments count is not equal of the placeholders count in sql query")
		} else {
			err = dbm.db.NewSelect().Model(data).Where(query, values...).Scan(ctx)
			if err != nil {
				return err
			}
		}
	} else {
		err = dbm.db.NewSelect().Model(data).Scan(ctx)
	}
	return nil
}

func (dbm *DatabaseManager) Insert(data interface{}) (err error) {
	ctx := context.Background()

	_, err = dbm.db.NewInsert().Model(data).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (dbm *DatabaseManager) Update(data interface{}, query string, values ...any) (err error) {
	ctx := context.Background()

	if len(query) <= 0 || len(values) <= 0 || strings.Count(query, "?") <= 0 {
		return errors.New("dbmanager update: empty update set is not allowed")

	} else if strings.Count(query, "?") != len(values) {
		return errors.New("dbmanager update: arguments count is not equal of the placeholders count in sql query")

	} else {
		_, err = dbm.db.NewUpdate().Model(data).Set(query, values...).WherePK().Exec(ctx)
	}
	return nil
}

func (dbm *DatabaseManager) Delete(data interface{}) (err error) {
	ctx := context.Background()

	_, err = dbm.db.NewDelete().Model(data).WherePK().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (dbm *DatabaseManager) Close() (err error) {
	err = dbm.db.Close()
	if err != nil {
		return err
	}
	err = dbm.sqldb.Close()
	if err != nil {
		return err
	}
	return nil
}
