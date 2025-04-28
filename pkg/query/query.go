package query

import (
	"time"

	"github.com/pravandkatyare/reduct-go/pkg/client"
)

// Package query provides functionality to create and execute queries on a database.
// Query is a struct that represents a query to be executed on the database.
type Query struct {
	*client.Client
	EntryName string
	Start     time.Time
	End       time.Time
	When      string
}

// NewQuery creates a new Query instance with the given parameters.
func (q *Query) NewEntry(entryName string, start, end time.Time, when string) {
	q.EntryName = entryName
	q.Start = start
	q.End = end
	q.When = when

	// make http request to delete the entry
}

// RemoveEntry is for removing an entry in the database. It takes entryName, start time, end time and when as parameters.
func (q *Query) RemoveEntry(entryName string, start, end time.Time, when string) {
	q.EntryName = entryName
	q.Start = start
	q.End = end
	q.When = when

	// make http request to delete the entry

}

// DeleteEntry is for deleting an entry in the database. It takes entryName, start time, end time and when as parameters.
func (q *Query) DeleteEntry(entryName string, start, end time.Time, when string) {
	q.EntryName = entryName
	q.Start = start
	q.End = end
	q.When = when

	// make http request to delete the entry
}

// UpdateEntry is for updating an entry in the database. It takes entryName, start time, end time and when as parameters.
func (q *Query) UpdateEntry(entryName string, start, end time.Time, when string) {
	q.EntryName = entryName
	q.Start = start
	q.End = end
	q.When = when

	// make http request to delete the entry
}

// ComplexQuery is for queries which are not handled by above methods. It is a generic method to handle complex queries.
// It takes entryName, start time, end time and when as parameters. (parameters will change, this is for representation purpose only)
func (q *Query) ComplexQuery(entryName string, start, end time.Time, when string) {
	q.EntryName = entryName
	q.Start = start
	q.End = end
	q.When = when

	// make http request to delete the entry
}
