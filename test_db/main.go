package main

import (
    "database/sql"
    "fmt"
    "log"
    "time"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    // Налаштування підключення до бази даних
    dsn := "root:password@tcp(mysql:3306)/testdb"
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Перевірка підключення
    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }

    // Створення таблиці
    createTableSQL := `CREATE TABLE IF NOT EXISTS users (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(50) NOT NULL,
        created_at DATETIME
    );`
    if _, err := db.Exec(createTableSQL); err != nil {
        log.Fatal(err)
    }

    // Додавання запису
    insertUserSQL := `INSERT INTO users (name, created_at) VALUES (?, ?)`
    _, err = db.Exec(insertUserSQL, "Oleh", time.Now())
    if err != nil {
        log.Fatal(err)
    }

    // Перевірка запису
    rows, err := db.Query("SELECT id, name, created_at FROM users")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next() {
        var id int
        var name string
        var createdAt time.Time
        if err := rows.Scan(&id, &name, &createdAt); err != nil {
            log.Fatal(err)
        }
        fmt.Printf("ID: %d, Name: %s, Created At: %s\n", id, name, createdAt)
    }
}
