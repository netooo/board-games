module github.com/netooo/board-games/app

go 1.15

replace github.com/netooo/board-games/routing => ./routing

require (
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator v9.31.0+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang-migrate/migrate/v4 v4.14.1
	github.com/google/uuid v1.2.0
	github.com/gorilla/sessions v1.2.1
	github.com/jinzhu/gorm v1.9.16
	github.com/joho/godotenv v1.3.0
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/netooo/board-games/routing v0.0.0-00010101000000-000000000000
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
)
