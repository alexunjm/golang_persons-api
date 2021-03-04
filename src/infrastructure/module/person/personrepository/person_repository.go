package personrepository

import (
	"context"
	"database/sql"
	"fmt"
	"golang_persons-api/src/application/module/person/queries"
	"golang_persons-api/src/domain/module/person"
	"time"

	"github.com/huandu/go-sqlbuilder"
)

// PersonRepository handles database actions
type PersonRepository struct {
	db        *sql.DB
	dbTimeout time.Duration
}

// NewPersonRepository creates a new PersonRepository
func NewPersonRepository(db *sql.DB, dbTimeout time.Duration) *PersonRepository {
	return &PersonRepository{
		db,
		dbTimeout,
	}
}

// Save database action
func (r PersonRepository) Save(ctx context.Context, person person.Person) error {
	fmt.Printf("saving person:\n %+v", person)

	personSQLStruct := sqlbuilder.NewStruct(new(sqlPerson))
	sql, args := personSQLStruct.InsertInto(sqlPersonTable, sqlPerson{
		ID:        person.ID().String(),
		Firstname: person.Firstname().String(),
		Lastname:  person.Lastname().String(),
		Age:       person.Age().Int(),
	}).Build()

	fmt.Println(sql)
	fmt.Println(args)

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	_, err := r.db.ExecContext(ctxTimeout, sql, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist person on database: %v", err)
	}

	return nil
}

// Update database action
func (r PersonRepository) Update(ctx context.Context, person person.Person) error {
	fmt.Printf("updating person :\n %+v", person)

	personSQLStruct := sqlbuilder.NewStruct(new(sqlPerson))
	ub := personSQLStruct.Update(sqlPersonTable, sqlPerson{
		ID:        person.ID().String(),
		Firstname: person.Firstname().String(),
		Lastname:  person.Lastname().String(),
		Age:       person.Age().Int(),
	})
	sql, args := ub.Where(
		ub.E("id", person.ID().String()),
	).Build()

	fmt.Println(sql)
	fmt.Println(args)

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	sqlResult, err := r.db.ExecContext(ctxTimeout, sql, args...)
	fmt.Println(sqlResult)
	if err != nil {
		return fmt.Errorf("error trying to persist person on database: %v", err)
	}

	return nil
}

// Delete database action
func (r PersonRepository) Delete(ctx context.Context, personID person.PersonID) error {
	fmt.Printf("deleting person :\n %+v", personID)

	personSQLStruct := sqlbuilder.NewStruct(new(sqlPerson))
	deleteBuilder := personSQLStruct.DeleteFrom(sqlPersonTable)
	sql, args := deleteBuilder.Where(
		deleteBuilder.E("id", personID.String()),
	).Build()

	fmt.Println(sql)
	fmt.Println(args)

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	sqlResult, err := r.db.ExecContext(ctxTimeout, sql, args...)
	fmt.Println(sqlResult)
	if err != nil {
		return fmt.Errorf("error trying to persist person on database: %v", err)
	}

	return nil
}

// Find database action
func (r PersonRepository) Find(ctx context.Context, personID person.PersonID) (interface{}, error) {
	fmt.Printf("finding person :\n %+v", personID)

	personSQLStruct := sqlbuilder.NewStruct(new(sqlPerson))
	deleteBuilder := personSQLStruct.SelectFrom(sqlPersonTable)
	sql, args := deleteBuilder.Where(
		deleteBuilder.E("id", personID.String()),
	).Build()

	fmt.Println(sql)
	fmt.Println(args)

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	var (
		sqlPersonFound sqlPerson
	)
	row := r.db.QueryRowContext(ctxTimeout, sql, args...)

	// Scan row data to person.
	err := row.Scan(personSQLStruct.Addr(&sqlPersonFound)...)
	if err != nil {
		return queries.FindPersonQueryResponse{}, err
	}
	fmt.Println(sqlPersonFound)

	return queries.FindPersonQueryResponse{
		ID:        sqlPersonFound.ID,
		Firstname: sqlPersonFound.Firstname,
		Lastname:  sqlPersonFound.Lastname,
		Age:       sqlPersonFound.Age,
	}, nil
}

// FindAll database action
func (r PersonRepository) FindAll(ctx context.Context) ([]interface{}, error) {
	fmt.Printf("finding all persons\n")

	personSQLStruct := sqlbuilder.NewStruct(new(sqlPerson))
	deleteBuilder := personSQLStruct.SelectFrom(sqlPersonTable)
	sql, args := deleteBuilder.Build()

	fmt.Println(sql)
	fmt.Println(args)

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	var (
		sqlPersonFound sqlPerson
	)
	rows, err := r.db.QueryContext(ctxTimeout, sql, args...)
	if err != nil {
		return []interface{}{}, err
	}
	defer rows.Close()

	allPersonsFound := make([]interface{}, 0)
	for rows.Next() {

		// Scan row data to person.
		err := rows.Scan(personSQLStruct.Addr(&sqlPersonFound)...)
		if err != nil {
			return []interface{}{}, err
		}
		allPersonsFound = append(allPersonsFound, queries.FindPersonQueryResponse{
			ID:        sqlPersonFound.ID,
			Firstname: sqlPersonFound.Firstname,
			Lastname:  sqlPersonFound.Lastname,
			Age:       sqlPersonFound.Age,
		})
	}

	return allPersonsFound, nil
}
