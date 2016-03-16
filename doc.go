/*
Package nullable contains types that represent values that may be null.
The Go standard library already has NullString, NullInt64 and NullFloat64
in the database/sql package. The types in this package add additional
data types.

The types in this package also all implement the json.Marshaler and
json.Unmarshaler interfaces, so they can be serialized to and from JSON.
*/
package nullable
