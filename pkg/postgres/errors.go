package postgres

import "errors"

// ErrZhukAlreadyExists is used when user is trying to create new zhuk.
var ErrZhukAlreadyExists = errors.New("такой жук уже есть")
