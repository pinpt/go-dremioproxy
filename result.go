package dremioproxy

import (
	"bufio"
	"database/sql/driver"
	"encoding/json"
	"io"
)

type result struct {
	reader  io.ReadCloser
	scanner *bufio.Reader
	columns []string
	offset  int
	data    map[string]interface{}
}

func newResult(r io.ReadCloser) (driver.Rows, error) {
	theresult := &result{
		reader:  r,
		scanner: bufio.NewReader(r),
	}
	if err := theresult.readNext(); err != nil {
		return nil, err
	}
	colnames := make([]string, 0)
	for k := range theresult.data {
		colnames = append(colnames, k)
	}
	theresult.columns = colnames
	return theresult, nil
}

func (r *result) readNext() error {
	buf, err := r.scanner.ReadBytes('\n')
	if err != nil {
		return err
	}
	data := make(map[string]interface{})
	if err := json.Unmarshal(buf, &data); err != nil {
		return err
	}
	r.data = data
	return nil
}

// Columns returns the names of the columns. The number of
// columns of the result is inferred from the length of the
// slice. If a particular column name isn't known, an empty
// string should be returned for that entry.
func (r *result) Columns() []string {
	return r.columns
}

// Close closes the rows iterator.
func (r *result) Close() error {
	return r.reader.Close()
}

// Next is called to populate the next row of data into
// the provided slice. The provided slice will be the same
// size as the Columns() are wide.
//
// Next should return io.EOF when there are no more rows.
//
// The dest should not be written to outside of Next. Care
// should be taken when closing Rows not to modify
// a buffer held in dest.
func (r *result) Next(dest []driver.Value) error {
	if r.data == nil {
		if err := r.readNext(); err != nil {
			return err
		}
	}
	for i := 0; i < len(r.columns); i++ {
		key := r.columns[i]
		val := r.data[key]
		dest[i] = val
	}
	r.data = nil
	r.offset++
	return nil
}

var _ driver.Rows = (*result)(nil)
