package models 

import ( 
	"database/sql" 
	_ "github.com/mattn/go-sqlite3"
) 

type Task struct { 
	ID int `json:"id"` 
	Name string `json:"name"` 
} 

type TaskCollection struct { 
	Tasks []Task `json:"items"` 
}

func GetTasks(db *sql.DB) TaskCollection {
    sql := "SELECT * FROM tasks"
    rows, err := db.Query(sql)

    if err != nil {
        panic(err)
    }
    // close handler at the end
    defer rows.Close()

    result := TaskCollection{}
    for rows.Next() {
        task := Task{}
        err2 := rows.Scan(&task.ID, &task.Name)

        if err2 != nil {
            panic(err2)
        }
        result.Tasks = append(result.Tasks, task)
    }
    return result
}

func PutTask(db *sql.DB, name string) (int64, error) {
    sql := "INSERT INTO tasks(name) VALUES(?)"

    stmt, err := db.Prepare(sql)

    if err != nil {
        panic(err)
    }

    defer stmt.Close()

    result, err2 := stmt.Exec(name)

    if err2 != nil {
        panic(err2)
    }

    return result.LastInsertId()
}

func DeleteTask(db *sql.DB, id int) (int64, error) {
    sql := "DELETE FROM tasks WHERE id = ?"

    // выполним SQL запрос
    stmt, err := db.Prepare(sql)
    // выход при ошибке
    if err != nil {
        panic(err)
    }

    // заменим символ '?' в запросе на 'id'
    result, err2 := stmt.Exec(id)
    // выход при ошибке
    if err2 != nil {
        panic(err2)
    }

    return result.RowsAffected()
}
