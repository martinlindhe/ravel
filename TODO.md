# TODO

- check out https://github.com/jteeuwen/go-bindata

- jwt by default: https://github.com/dgrijalva/jwt-go



# TODO cli app
    - something similar to artisan
    - "migrate:seed --refresh", etc


# TODO tests
    - route test!?


# migrations
    - should be its own go package, based on gorm
    - come up with some yaml-syntax for declaring migrations, rather than in code (more terse)
    - needs to be able to change column name, table name
    - keep migrations table, if we do manual migrations.. !?
    - hook up manual migration files, like in laravel. solution exists?

    - TODO: cli that generates a migration file based on Model name (and existing model when cmd is run)





# LATER

- faker lib integrated with gorm. i barely use fakes, but they are nice to have... https://github.com/icrowley/fake
- sass compiler i go: https://github.com/c9s/c6
