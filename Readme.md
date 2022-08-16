## golang-migrate/migrate

## syntax

# migrate create [-ext E] [-dir D] [-seq] [-digits N] [-format] NAME

Create a set of timestamped up/down migrations titled NAME, in directory D with extension E.
Use -seq option to generate sequential up/down migrations with N digits.
Use -format option to specify a Go time format string.

# migrate create -ext sql -dir db/migration -seq init_schema
