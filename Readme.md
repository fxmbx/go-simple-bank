## golang-migrate/migrate

## syntax

# migrate create [-ext E] [-dir D] [-seq] [-digits N] [-format] NAME

Create a set of timestamped up/down migrations titled NAME, in directory D with extension E.
Use -seq option to generate sequential up/down migrations with N digits.
Use -format option to specify a Go time format string.

# migrate create -ext sql -dir db/migration -seq init_schema

postgres driver

# go get github.com/lib/pq

to check test using read db

# go get github.com/stretchr/testify/

to test using mock db

# go get github.com/golang/mock/mockgen@v1.6.0

mockgen -destination db/mock/store.go github.com/fxmbx/go-simple-bank/db/sqlc Store

gin for routing and some other things

# go get -u github.com/gin-gonic/gin

viper for reading and watching env variables

# go get github.com/spf13/viper

uuid for the tokens

# go get github.com/google/uuid

jwt

# github.com/dgrijalva/jwt-go

paseto

# go get -u github.com/o1egl/paseto
