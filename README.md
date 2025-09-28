#
For Swagger Configuration

##
* go install github.com/swaggo/swag/cmd/swag@latest
* export PATH=$PATH:$(go env GOPATH)/bin (For Bash)
* $env:PATH += ";$(go env GOPATH)\bin" (For Powershell)
* Check for swagger version using swag --version command.
* Put swagger comments on main.go file and the controllers and models as mentioned in the swagger docs.
* Run command swag init -g ./cmd/main/main.go (swag init -g {path to main.go}).
* Ensure main.go has the following two things imported 1. path to docs package (eg _"com.shreyash/go_inventory/v2/docs") and the import for swagger dependency (httpSwagger "github.com/swaggo/http-swagger").
* Register the dependency for swagger in the main.go file with r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) or r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)



