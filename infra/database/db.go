package database

import (
	"database/sql"
	"fmt"
	"log"

	"toko-belanja-app/entity"
	"toko-belanja-app/infra/config"

	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

func handleDatabaseConnection() {

	appConfig := config.AppConfig()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		appConfig.DbHost,
		appConfig.DbPort,
		appConfig.DbUser,
		appConfig.DbPassword,
		appConfig.DbName,
	)

	db, err = sql.Open(appConfig.DbDialect, dsn)

	if err != nil {
		log.Panicln("error occured while trying to validate database arguments: ", err.Error())
		return
	}

	if err := db.Ping(); err != nil {
		log.Panicln("error occured while trying to connect to database: ", err.Error())
		return
	}
}

func handleRequiredTables() {
	const (
		createTableUsersQuery = `
			CREATE TABLE IF NOT EXISTS
				users
					(
						id SERIAL PRIMARY KEY,
						full_name VARCHAR(100) NOT NULL,
						email VARCHAR(100) NOT NULL,
						password TEXT NOT NULL,
						role VARCHAR(10) NOT NULL DEFAULT 'customer',
						balance INT NOT NULL DEFAULT 0,
						created_at TIMESTAMPTZ DEFAULT now(),
						updated_at TIMESTAMPTZ DEFAULT now(),
						deleted_at TIMESTAMPTZ,
						CONSTRAINT
							unique_email
								UNIQUE(email)
					)
		`

		createTableCategoriesQuery = `
			CREATE TABLE IF NOT EXISTS
				categories
					(
						id SERIAL PRIMARY KEY,
						type VARCHAR(25) NOT NULL,
						sold_product_amount INT DEFAULT 0,
						created_at TIMESTAMPTZ DEFAULT now(),
						updated_at TIMESTAMPTZ DEFAULT now(),
						deleted_at TIMESTAMPTZ
					)
		`

		createTableProductsQuery = `
			CREATE TABLE IF NOT EXISTS
				products
					(
						id SERIAL PRIMARY KEY,
						title VARCHAR(50) NOT NULL,
						price INT NOT NULL,
						stock INT NOT NULL,
						category_id INT NOT NULL,
						created_at TIMESTAMPTZ DEFAULT now(),
						updated_at TIMESTAMPTZ DEFAULT now(),
						deleted_at TIMESTAMPTZ,
						CONSTRAINT
							fk_products_category_id
								FOREIGN KEY (category_id)
									REFERENCES 
										categories(id)
					)
		`

		createTableTransactionHistoriesQuery = `
			CREATE TABLE IF NOT EXISTS
				transaction_histories
					(
						id SERIAL PRIMARY KEY,
						user_id INT NOT NULL,
						product_id INT NOT NULL,
						quantity INT NOT NULL,
						total_price INT NOT NULL,
						created_at TIMESTAMPTZ DEFAULT now(),
						updated_at TIMESTAMPTZ DEFAULT now(),
						deleted_at TIMESTAMPTZ,
						CONSTRAINT
							fk_transaction_history_user_id
								FOREIGN KEY (user_id)
									REFERENCES 
										users(id),
						CONSTRAINT
							fk_transaction_history_product_id
								FOREIGN KEY (product_id)
									REFERENCES 
										products(id)
					)
		`
		createTrigger = `
			CREATE OR REPLACE FUNCTION reduce_balance_on_transaction() RETURNS TRIGGER AS $$
			BEGIN
				UPDATE users
				SET balance = balance - (NEW.quantity * (
					SELECT price FROM products WHERE id = NEW.product_id
				))
				WHERE id = NEW.user_id;
				RETURN NEW;
			END;
			$$ LANGUAGE plpgsql;

			CREATE OR REPLACE TRIGGER transaction_balance_trigger
			AFTER INSERT ON transaction_histories
			FOR EACH ROW
			EXECUTE FUNCTION reduce_balance_on_transaction();


			CREATE OR REPLACE FUNCTION reduce_stock_on_transaction() RETURNS TRIGGER AS $$
			BEGIN
				UPDATE products
				SET stock = stock - NEW.quantity
				WHERE id = NEW.product_id;
				RETURN NEW;
			END;
			$$ LANGUAGE plpgsql;

			CREATE OR REPLACE TRIGGER transaction_stock_trigger
			AFTER INSERT ON transaction_histories
			FOR EACH ROW
			EXECUTE FUNCTION reduce_stock_on_transaction();

			CREATE OR REPLACE FUNCTION increase_sold_amount_on_transaction() RETURNS TRIGGER AS $$
			BEGIN
				UPDATE categories
				SET sold_product_amount = sold_product_amount + NEW.quantity
				WHERE id = (SELECT category_id FROM products WHERE id = NEW.product_id);
				RETURN NEW;
			END;
			$$ LANGUAGE plpgsql;

			CREATE OR REPLACE TRIGGER transaction_category_trigger
			AFTER INSERT ON transaction_histories
			FOR EACH ROW
			EXECUTE FUNCTION increase_sold_amount_on_transaction();
		`

		createAdminQuery = `
			INSERT INTO
				users (
					full_name,
					email,
					password,
					role,
					balance
				)
			VALUES
				($1, $2, $3, 'admin', 0)
			ON CONFLICT(email)
			DO NOTHING
		`
	)

	_, err = db.Exec(createTableUsersQuery)

	if err != nil {
		log.Panic("error while create table users: ", err.Error())
		return
	}

	_, err = db.Exec(createTableCategoriesQuery)

	if err != nil {
		log.Panic("error while create table categories: ", err.Error())
		return
	}

	_, err = db.Exec(createTableProductsQuery)

	if err != nil {
		log.Panic("error while create table products: ", err.Error())
		return
	}

	_, err = db.Exec(createTableTransactionHistoriesQuery)

	if err != nil {
		log.Panic("error while create table transaction_histories: ", err.Error())
		return
	}

	_, err = db.Exec(createTrigger)

	if err != nil {
		log.Panic("error while create OR REPLACE trigger: ", err.Error())
		return
	}

	user := &entity.User{
		FullName: config.AppConfig().AdminFullName,
		Email:    config.AppConfig().AdminEmail,
		Password: config.AppConfig().AdminPassword,
	}

	if user.FullName == "" && user.Email == "" && user.Password == "" {
		log.Panicln("full name, email and password can't be empty please fill .env file ")
		return
	}

	user.HashPassword()

	_, err = db.Exec(createAdminQuery, user.FullName, user.Email, user.Password)

	if err != nil {
		log.Panic("error while create admin: ", err.Error())
		return
	}
}

func InitializeDatabase() {
	handleDatabaseConnection()
	handleRequiredTables()
}

func GetInstanceDatabaseConnection() *sql.DB {
	return db
}
