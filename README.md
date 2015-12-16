# About

Experiments in golang, as a Laravel dev



web router framework: github.com/gin-gonic/gin
orm db driver: github.com/jinzhu/gorm
dotfile: github.com/joho/godotenv
gorazor: razor template engine, github.com/sipin/gorazor
i18n: github.com/nicksnyder/go-i18n

STATUS: experimental


# Views

Run the watcher to generate views on the fly:

    gorazor -watch resources/views views


# TODO

- check out https://github.com/jteeuwen/go-bindata

- jwt: https://github.com/dgrijalva/jwt-go

# TODO cli app
    - something similar to artisan
    - "migrate:seed --refresh", etc


tests:
    - route test!?

- migrations important!
    - should be its own go package, based on gorm
    - come up with some yaml-syntax for declaring migrations, rather than in code (more terse)
    - needs to be able to change column name, table name
    - keep migrations table, if we do manual migrations.. !?
    - hook up manual migration files, like in laravel. solution exists?




# LATER

- faker lib integrated with gorm. i barely use fakes, but they are nice to have... https://github.com/icrowley/fake
- sass compiler i go: https://github.com/c9s/c6

# STATUS

password.Make(): provides helper for bcrypt password hashing
