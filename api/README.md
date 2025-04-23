# USEFUL PACKAGES
- [air](https://github.com/air-verse/air) : Used to live releoad the application
- [gin-swagger](https://github.com/swaggo/gin-swagger) swagger doc generations
    - [declarative comment docs](https://github.com/swaggo/swag/blob/master/README.md#how-to-use-it-with-gin)
- [Gocloack](https://github.com/Nerzal/gocloak) : Golang's keycloak client
- [Gin cors](github.com/gin-contrib/cors) : Gin middleware/handler to enable CORS support.
- [Validator](https://github.com/go-playground/validator) : For struct validation
# Generate api doc: 
`swag init -g ./cmd/main.go -o api/docs`
