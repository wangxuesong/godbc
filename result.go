// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package odbc

import (
	"database/sql/driver"
	"errors"
)

type Result struct {
	c        *Conn
	rowCount int64
}

func (r *Result) LastInsertId() (int64, error) {
	// TODO(brainman): implement (*Result).LastInsertId
	// return 0, errors.New("not implemented")
	s, err := r.c.Prepare("select dbinfo('serial8') from sysmaster:sysdual")
	if err != nil {
		return 0, err
	}
	defer s.Close()
	rows, err := s.Query(nil)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	dest := make([]driver.Value, 1)
	err = rows.Next(dest)
	if err != nil {
		return 0, err
	}
	if dest[0] == nil {
		return -1, errors.New("There is no generated identity value")
	}
	lastInsertId := dest[0].(int64)
	return lastInsertId, nil
}

func (r *Result) RowsAffected() (int64, error) {
	return r.rowCount, nil
}
