package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"

	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/spf13/viper"

	_productHttpDeliver "github.com/oryzel/go-product-service/product/delivery/http"
	_productRepo "github.com/oryzel/go-product-service/product/repository"
	_productUcase "github.com/oryzel/go-product-service/product/usecase"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}

	fmt.Printf("Service run on port %v \n", viper.GetString(`server.address`))
}

func main() {

	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)
	if err != nil && viper.GetBool("debug") {
		fmt.Println(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()

	//URL to swagger documentation
	e.Static("/swaggerui", "./swaggerui")

	pr := _productRepo.NewMysqlArticleRepository(dbConn)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	au := _productUcase.NewProductUsecase(pr, timeoutContext)
	_productHttpDeliver.NewProductHandler(e, au)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
