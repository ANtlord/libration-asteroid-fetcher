// Package datamining provides ...
package datamining

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/lib/pq"
)

type Miner struct {
	User,
	Password,
	Database,
	Host,
	Port string
}

func buildQuery(byIntegers *Integers, planet1, planet2 string, onlyPure bool) *string {
	var sel = "SELECT substring(asteroid.name, 2, length(asteroid.name)-1)::int"
	var from = "FROM libration"
	var joins = []string{
		"JOIN resonance ON libration.resonance_id = resonance.id",
		"JOIN planet AS planet_1 ON planet_1.id = resonance.first_body_id",
		"JOIN planet AS planet_2 ON planet_2.id = resonance.second_body_id",
		"JOIN asteroid ON asteroid.id = resonance.small_body_id",
	}

	var conditions = []string{
		"WHERE planet_1.name='JUPITER' AND planet_2.name='SATURN'",
		fmt.Sprintf("AND planet_1.longitude_coeff = %d", byIntegers.First),
		fmt.Sprintf("AND planet_2.longitude_coeff = %d", byIntegers.Second),
		fmt.Sprintf("AND asteroid.longitude_coeff = %d", byIntegers.Asteroid),
	}
	if onlyPure {
		var lengthStr = "array_length(libration.circulation_breaks, 1)"
		var pureCond = fmt.Sprintf("AND (%s < 2 OR %s IS NULL)", lengthStr, lengthStr)
		conditions = append(conditions, pureCond)
	}

	var order = "ORDER BY substring(asteroid.name, 2, length(asteroid.name)-1)::int;"
	var query = fmt.Sprintf(
		"%s %s %s %s %s", sel, from, strings.Join(joins, " "),
		strings.Join(conditions, " "), order,
	)

	return &query
}

var db *sql.DB = nil

// Do you think, this is mistake? Wrong. It is the lack of time.
func (miner *Miner) getDB() *sql.DB {
	if db == nil {
		var err error
		var connection_address = fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=disable",
			miner.User,
			miner.Password,
			miner.Host,
			miner.Port,
			miner.Database,
		)
		fmt.Println(connection_address)
		db, err = sql.Open("postgres", connection_address)
		if err != nil {
			log.Fatal(err)
		}
	}
	return db
}

// FetchLibrations returns array of numbers of asteroids, that librates in
// pointed Integers.
func (miner *Miner) FetchLibrations(byIntegers *Integers, planet1, planet2 string, onlyPure bool) []string {
	var res = make([]string, 0, 100)
	db = miner.getDB()
	var query = buildQuery(byIntegers, planet1, planet2, onlyPure)
	rows, err := db.Query(*query)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var asteroidNumber string
		rows.Scan(&asteroidNumber)
		res = append(res, asteroidNumber)
	}
	return res
}
