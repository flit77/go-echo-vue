package handlers

import (
    "database/sql"
    "net/http"
    "strconv"

    "echo-vue/models"

    "github.com/labstack/echo"
)

type H map[string]interface{}

func GetTasks(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.JSON(http.StatusOK, models.GetTasks(db))
    }
}

func PutTask(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        // create new task
        var task models.Task
        // bind income json to created task
        c.Bind(&task)
        // add task within model
        id, err := models.PutTask(db, task.Name)
        // in success case - get json
        if err == nil {
            return c.JSON(http.StatusCreated, H{
                "created": id,
            })
        // errors handling
        } else {
            return err
        }
    }
}

func DeleteTask(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        id, _ := strconv.Atoi(c.Param("id"))
        // utilize a model to delete task
        _, err := models.DeleteTask(db, id)
        // in success - return json
        if err == nil {
            return c.JSON(http.StatusOK, H{
                "deleted": id,
            })
        // errors handling
        } else {
            return err
        }
    }
}