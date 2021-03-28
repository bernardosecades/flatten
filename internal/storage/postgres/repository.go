package postgres

import (
	"database/sql"
	"fmt"
	"time"

	flatten "github.com/bernardosecades/flatten/internal"
	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

type postgresRepository struct {
	SQL *sql.DB
}

func NewPostgresRepository(host string, port string, user string, pass string, dbname string) flatten.Repository {

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return &postgresRepository{SQL: db}
}

func (r *postgresRepository) SaveHistory(req []byte, res []byte, depth int) error {

	u := uuid.Must(uuid.NewV4(), nil)
	_, err := r.SQL.Exec("INSERT INTO history (id, request, response, depth, created_at) VALUES ($1, $2::bytea, $3::bytea, $4, $5)", u.String(), req, res, depth, time.Now())

	return err
}

func (r *postgresRepository) GetHistoryByLimit(l uint) ([]flatten.History, error) {

	var histories []flatten.History

	rows, err := r.SQL.Query("SELECT request, response, depth, created_at FROM history ORDER BY created_at DESC LIMIT $1", l)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		history := flatten.History{}
		err := rows.Scan(&history.Request, &history.Response, &history.Depth, &history.CreatedAt)
		if err != nil {
			return nil, err
		}
		histories = append(histories, history)
	}

	return histories, nil
}
