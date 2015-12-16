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

- få gorazor att funka. eller snarare fatta hur man ska göra!

- kolla in https://github.com/jteeuwen/go-bindata

- views: blade liknande. gorazor https://github.com/sipin/gorazor

CLI APP:
    - liknande artisan
    - "migrate:seed --refresh", etc


tests:
    - route test!?

- migrations important!
    - borde vara ett eget lib baserat på gorm
    - eget textformat med migrations?
    - kunna byta namn på kolumn, tabell
    - keep migrations table, if we do manual migrations.. !?
    - hook up manual migration files, like in laravel. solution exists? proper migration stuff, google some?!


# ASSETS
    - sass compiler i go: https://github.com/c9s/c6



# SENARE

- faker lib, använder jag knappt ... https://github.com/icrowley/fake

# STATUS

password.Make(): provides helper for bcrypt password hashing
