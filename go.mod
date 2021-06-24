module github.com/chukwuka-emi

// +heroku goVersion go1.16.4
// +heroku install goVersion go1.16
go 1.16

require (
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/form3tech-oss/jwt-go v3.2.3+incompatible
	github.com/gofiber/fiber/v2 v2.10.0
	github.com/gofiber/jwt/v2 v2.2.1
	github.com/jackc/pgproto3/v2 v2.0.7 // indirect
	github.com/joho/godotenv v1.3.0
	github.com/lib/pq v1.6.0 // indirect
	github.com/sendgrid/rest v2.6.4+incompatible // indirect
	github.com/sendgrid/sendgrid-go v3.10.0+incompatible
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a
	golang.org/x/text v0.3.6 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
	gorm.io/driver/postgres v1.1.0
	gorm.io/gorm v1.21.10
)
