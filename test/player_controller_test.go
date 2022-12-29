package test 

import (
    "testing"
    "context"
    // "fmt"
    "io"
    "strings"
    "strconv"
    "time"
    "net/http"
    "database/sql"
    "encoding/json"
     "net/http/httptest"
    "github.com/go-playground/validator/v10"
    "github.com/stretchr/testify/assert"
    _ "github.com/go-sql-driver/mysql"
    "exercise/webservice/controller"
    "exercise/webservice/model"
    "exercise/webservice/repository"
    "exercise/webservice/service"
    "exercise/webservice/helper"
    "exercise/webservice/router"
    // "exercise/webservice/app/config"
)

func setupTestDB() *sql.DB {
    db, err := sql.Open("mysql", "root@tcp(localhost:3306)/database_name_test")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
} 

func setupRouter(db *sql.DB) http.Handler {
    validate := validator.New()
    playerRepository := repository.NewPlayerRepository()
    playerService := service.NewPlayerService(playerRepository, db, validate)
    playerController := controller.NewPlayerController(playerService)
    router := router.NewRouter(playerController)

    return router
}

func truncatePlayer(db *sql.DB) {
    _, err := db.Exec("TRUNCATE TABLE players")
    helper.PanicIfError(err)
}

func TestCreatePlayerSuccess(t *testing.T) {
    db := setupTestDB()
    truncatePlayer(db)
    router := setupRouter(db)

    requestBody := strings.NewReader(`{"name": "Pomodoro","position": "access","height": 5,"weight": 64,"birth_date": "1996-02-08"}`)

    request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/players", requestBody)
    request.Header.Add("Content-Type", "application/json")

    recorder := httptest.NewRecorder()
    router.ServeHTTP(recorder, request)

    response := recorder.Result()
    assert.Equal(t, 200, response.StatusCode)

    body, _ := io.ReadAll(response.Body)
    var responseBody map[string]interface{}
    json.Unmarshal(body, &responseBody)

    assert.Equal(t, 200, int(responseBody["code"].(float64)))
    assert.Equal(t, "Player created succesfully", responseBody["status"])
    assert.Equal(t, "Pomodoro", responseBody["data"].(map[string]interface{})["name"])
    assert.Equal(t, "access", responseBody["data"].(map[string]interface{})["position"])
    assert.Equal(t, 5, int(responseBody["data"].(map[string]interface{})["height"].(float64)))
    assert.Equal(t, 64, int(responseBody["data"].(map[string]interface{})["weight"].(float64)))
    assert.Equal(t, "1996-02-08", responseBody["data"].(map[string]interface{})["birth_date"])
}

func TestCreatePlayerFailed(t *testing.T) {
    db := setupTestDB()
    truncatePlayer(db)
    router := setupRouter(db)

    requestBody := strings.NewReader(`{"name":"","position": "CMK","height":5,"weight":80,"birth_date":"2002-02-08"}`)

    request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/players", requestBody)
    request.Header.Add("Content-Type", "application/json")

    recorder := httptest.NewRecorder()
    router.ServeHTTP(recorder, request)

    response := recorder.Result()
    assert.Equal(t, 400, response.StatusCode)

    body, _ := io.ReadAll(response.Body)
    var responseBody map[string]interface{}
    json.Unmarshal(body, &responseBody)

    assert.Equal(t, 400, int(responseBody["code"].(float64)))
    assert.Equal(t, "Validation errors", responseBody["status"])
}

func TestUpdatePlayerSuccess(t *testing.T) {
    db := setupTestDB()
    truncatePlayer(db)

    tx, _ := db.Begin()
    playerRepository := repository.NewPlayerRepository()
    data := model.Player{
        Name: "Ngolo Kante",
        Position: "CMF",
        Height: 5.6,
        Weight: 70,
        BirthDate: "1991-03-29",
    }

    player := playerRepository.Create(context.Background(), tx, data)
    tx.Commit()

    router := setupRouter(db)

    requestBody := strings.NewReader(`{"name": "Ngolo Kante","position": "DMF","height": 5.6,"weight": 70,"birth_date": "1991-03-29"}`)
    request := httptest.NewRequest(http.MethodPut, "http://localhost:8080/api/players/"+strconv.Itoa(player.Id), requestBody)
    request.Header.Add("Content-Type", "application/json")

    recorder := httptest.NewRecorder()
    router.ServeHTTP(recorder, request)
    response := recorder.Result()

    assert.Equal(t, 200, response.StatusCode)

    body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

    assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "Player updated successfully", responseBody["status"])
	assert.Equal(t, player.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "DMF", responseBody["data"].(map[string]interface{})["position"])
}

func TestUpdatePlayerFailed(t *testing.T) {
    db := setupTestDB()
    truncatePlayer(db)

    tx, _ := db.Begin()
    playerRepository := repository.NewPlayerRepository()
    data := model.Player{
        Name: "Ngolo Kante",
        Position: "CMF",
        Height: 5.6,
        Weight: 70,
        BirthDate: "1991-03-29",
    }

    player := playerRepository.Create(context.Background(), tx, data)
    tx.Commit()

    router := setupRouter(db)

    requestBody := strings.NewReader(`{"name": "","position": "DMF","height": 5.6,"weight": 70,"birth_date": "1991-03-29"}`)
    request := httptest.NewRequest(http.MethodPut, "http://localhost:8080/api/players/"+strconv.Itoa(player.Id), requestBody)
    request.Header.Add("Content-Type", "application/json")

    recorder := httptest.NewRecorder()
    router.ServeHTTP(recorder, request)
    response := recorder.Result()

    assert.Equal(t, 400, response.StatusCode)

    body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

    assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "Validation errors", responseBody["status"])
}

func TestGetPlayerSuccess(t *testing.T) {
    db := setupTestDB()
    truncatePlayer(db)

    tx, _ := db.Begin()
    playerRepository := repository.NewPlayerRepository()
    data := model.Player{
        Name: "Ngolo Kante",
        Position: "CMF",
        Height: 5.6,
        Weight: 70,
        BirthDate: "1991-03-29",
    }

    player := playerRepository.Create(context.Background(), tx, data)
    tx.Commit()

    router := setupRouter(db)

    request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/players/"+strconv.Itoa(player.Id), nil)
    request.Header.Add("Content-Type", "application/json")

    recorder := httptest.NewRecorder()
    router.ServeHTTP(recorder, request)
    response := recorder.Result()

    assert.Equal(t, 200, response.StatusCode)

    body, _ := io.ReadAll(response.Body)
    var responseBody map[string]interface{}
    json.Unmarshal(body, &responseBody)

    assert.Equal(t, 200, int(responseBody["code"].(float64)))
    assert.Equal(t, "Player found", responseBody["status"])
    assert.Equal(t, player.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
    assert.Equal(t, "Ngolo Kante", responseBody["data"].(map[string]interface{})["name"])
    assert.Equal(t, "CMF", responseBody["data"].(map[string]interface{})["position"])
    assert.Equal(t, float64(5.6), responseBody["data"].(map[string]interface{})["height"])
    assert.Equal(t, float64(70), responseBody["data"].(map[string]interface{})["weight"])
    assert.Equal(t, "1991-03-29", responseBody["data"].(map[string]interface{})["birth_date"])

    // fmt.Printf("Response body: %v\n", responseBody)
} 

func TestGetPlayerFailed(t *testing.T) {
    db := setupTestDB()
    truncatePlayer(db)

    router := setupRouter(db)

    request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/players/1", nil)
    request.Header.Add("Content-Type", "application/json")

    recorder := httptest.NewRecorder()
    router.ServeHTTP(recorder, request)
    response := recorder.Result()

    assert.Equal(t, 404, response.StatusCode)

    body, _ := io.ReadAll(response.Body)
    var responseBody map[string]interface{}
    json.Unmarshal(body, &responseBody)

    assert.Equal(t, 404, int(responseBody["code"].(float64)))
    assert.Equal(t, "Not Found", responseBody["status"])
}

func TestDeletePlayerSuccess(t *testing.T) {
    db := setupTestDB()
    truncatePlayer(db)

    tx, _ := db.Begin()
    playerRepository := repository.NewPlayerRepository()
    data := model.Player{
        Name: "Ngolo Kante",
        Position: "CMF",
        Height: 5.6,
        Weight: 70,
        BirthDate: "1991-03-29",
    }
    player := playerRepository.Create(context.Background(), tx, data)
    tx.Commit()

    router := setupRouter(db)

    request := httptest.NewRequest(http.MethodDelete, "http://localhost:8080/api/players/"+strconv.Itoa(player.Id), nil)
	request.Header.Add("Content-Type", "application/json")

    recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

    assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "Player deleted successfully", responseBody["status"])
}

func TestDeletePlayerFailed(t *testing.T) {
    db := setupTestDB()
    truncatePlayer(db)

    router := setupRouter(db)

    request := httptest.NewRequest(http.MethodDelete, "http://localhost:8080/api/players/4", nil)
    request.Header.Add("Content-Type", "application/json")

    recorder := httptest.NewRecorder()
    router.ServeHTTP(recorder, request)
    response := recorder.Result()

    assert.Equal(t, 404, response.StatusCode)

    body, _ := io.ReadAll(response.Body)
    var responseBody map[string]interface{}
    json.Unmarshal(body, &responseBody)

    assert.Equal(t, 404, int(responseBody["code"].(float64)))
    assert.Equal(t, "Not Found", responseBody["status"])
}

func TestListPlayerSuccess(t *testing.T) {
    db := setupTestDB()
    truncatePlayer(db)

    tx, _ := db.Begin()
    playerRepository := repository.NewPlayerRepository()
    data1 := model.Player{
        Name: "Ngolo Kante",
        Position: "CMF",
        Height: 5.6,
        Weight: 70,
        BirthDate: "1991-03-29",
    }
    
    playerRepository.Create(context.Background(), tx, data1)

    data2 := model.Player{
        Name: "Cristiano Ronaldo",
        Position: "ST",
        Height: 5.9,
        Weight: 80,
        BirthDate: "1985-02-05",
    }

    playerRepository.Create(context.Background(), tx, data2)
    tx.Commit()

    router := setupRouter(db)
    request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/players", nil)

    recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])

    players := responseBody["data"].([]interface{})

    playerResponse1 := players[0].(map[string]interface{})
    playerResponse2 := players[1].(map[string]interface{})

    assert.Equal(t, "Ngolo Kante", playerResponse1["name"])
    assert.Equal(t, "CMF", playerResponse1["position"])
    assert.Equal(t, float64(5.6), playerResponse1["height"])
    assert.Equal(t, float64(70), playerResponse1["weight"])
    assert.Equal(t, "1991-03-29", playerResponse1["birth_date"])

    assert.Equal(t, "Cristiano Ronaldo", playerResponse2["name"])
    assert.Equal(t, "ST", playerResponse2["position"])
    assert.Equal(t, float64(5.9), playerResponse2["height"])
    assert.Equal(t, float64(80), playerResponse2["weight"])
    assert.Equal(t, "1985-02-05", playerResponse2["birth_date"])

}