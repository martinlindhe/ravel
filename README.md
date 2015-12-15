# About

ubique.se source repository

web router framework: github.com/gin-gonic/gin
orm db driver: github.com/jinzhu/gorm
dotfile: github.com/joho/godotenv


STATUS: private


# Views

Run the watcher to generate views on the fly:

    gorazor -watch tpl views


# TODO

- use elixir for js, sass .... !!!

- views: blade liknande. gorazor https://github.com/sipin/gorazor


tests:
    - route test!?

- migrations important!
    - borde vara ett eget lib baserat på gorm
    - eget textformat med migrations?
    - kunna byta namn på kolumn, tabell
    - keep migrations table, if we do manual migrations.. !?
    - hook up manual migration files, like in laravel. solution exists? proper migration stuff, google some?!


# SENARE

- faker lib, använder jag knappt ... https://github.com/icrowley/fake

# STATUS

password.Make(): provides helper for bcrypt password hashing
