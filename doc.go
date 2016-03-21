/*
Package nullable contains types that represent values that may be null.
The Go standard library already has NullBool, NullFloat64, NullInt64
and NullString in the database/sql package. The types in this package
add additional data types.

The types in this package also all implement the json.Marshaler and
json.Unmarshaler interfaces, so they can be serialized to and from JSON.
These types also have convience methods for converting to and from pointers.
*/
package nullable
