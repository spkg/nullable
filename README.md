# Nullable Types

[![GoDoc](https://godoc.org/github.com/spkg/nullable?status.svg)](https://godoc.org/github.com/spkg/nullable)
[![Build Status (Linux)](https://travis-ci.org/spkg/nullable.svg?branch=master)](https://travis-ci.org/spkg/nullable)
[![Build status (Windows)](https://ci.appveyor.com/api/projects/status/txfjx8i49ntan6fm?svg=true)](https://ci.appveyor.com/project/jjeffery/nullable)
[![license](http://img.shields.io/badge/license-MIT-green.svg?style=flat)](https://raw.githubusercontent.com/spkg/nullable/master/LICENSE.md)
[![Coverage](http://gocover.io/_badge/github.com/spkg/nullable)](http://gocover.io/github.com/spkg/nullable)
[![GoReportCard](http://goreportcard.com/badge/spkg/nullable)](http://goreportcard.com/report/spkg/nullable)

The `nullable` package provides a number of types that represent values
that may be null. The standard Go library already includes the following
types in the `database/sql` package for this purpose:

* `NullBool`
* `NullFloat64`
* `NullInt64`
* `NullString`

The types in this package add to this list for convenience. The other
significant difference is that the types in this package all implement
the `json.Marshaler` and `json.Unmarshaler` interfaces, which is
used for serializing to and from JSON.
