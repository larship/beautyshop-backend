package auth

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"github.com/larship/beautyshop/api"
	"github.com/larship/beautyshop/database"
	"github.com/larship/beautyshop/models"
	"math/rand"
)

const serverSalt = "dg353hy034"

func randomString(length int, onlyNums bool) string {
	var letters []rune

	if !onlyNums {
		letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	} else {
		letters = []rune("0123456789")
	}

	s := make([]rune, length)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}

	return string(s)
}

func CheckAuth(clientUuid string, sessionId string, salt string) *models.Client {
	var client models.Client

	sql := `
		SELECT *
		FROM clients c
		WHERE c.uuid = $1
	`

	err := database.DB.GetConnection().QueryRow(context.Background(), sql, clientUuid).Scan(&client.Uuid, &client.FullName,
		&client.Phone, &client.SessionId, &client.SessionPrivateId, &client.Salt)
	if err != nil {
		fmt.Printf("Ошибка получения клиента: %v", err)
		return nil
	}

	serverSaltHash := sha256.Sum256([]byte(serverSalt))
	serverSaltHashStr := hex.EncodeToString(serverSaltHash[:])

	calcPrivateSessionIdHash := sha256.Sum256([]byte(clientUuid + sessionId + salt + serverSaltHashStr))

	if client.SessionPrivateId == hex.EncodeToString(calcPrivateSessionIdHash[:]) {
		return &client
	}

	return nil
}

func CreateUser(fullName string, phone string) *models.Client {
	clientUuid := uuid.New().String()

	sessionIdHash := sha256.Sum256([]byte(randomString(32, false)))
	sessionIdHashStr := hex.EncodeToString(sessionIdHash[:])

	saltHash := sha256.Sum256([]byte(randomString(32, false)))
	saltHashStr := hex.EncodeToString(saltHash[:])

	serverSaltHash := sha256.Sum256([]byte(serverSalt))
	serverSaltHashStr := hex.EncodeToString(serverSaltHash[:])

	sessionPrivateIdHash := sha256.Sum256([]byte(clientUuid + sessionIdHashStr + saltHashStr + serverSaltHashStr))
	sessionPrivateIdHashStr := hex.EncodeToString(sessionPrivateIdHash[:])

	sql := `
		INSERT INTO clients
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := database.DB.GetConnection().Exec(context.Background(), sql, clientUuid, fullName, phone,
		sessionIdHashStr, sessionPrivateIdHashStr, saltHashStr)

	if err != nil {
		fmt.Printf("Ошибка добавления клиента: %v", err)
		return nil
	}

	return &models.Client{
		Uuid:             clientUuid,
		Phone:            phone,
		FullName:         fullName,
		SessionId:        sessionIdHashStr,
		SessionPrivateId: sessionPrivateIdHashStr,
		Salt:             saltHashStr,
	}
}

func getAdminByPhone(phone string) (*models.Client, error) {
	var client models.Client

	sql := `
		SELECT c.*
		FROM clients c
		INNER JOIN beautyshops_admins ba ON ba.client_uuid = c.uuid
		WHERE c.phone = $1
		LIMIT 1
	`

	err := database.DB.GetConnection().QueryRow(context.Background(), sql, phone).Scan(&client.Uuid, &client.FullName,
		&client.Phone, &client.SessionId, &client.SessionPrivateId, &client.Salt)

	if err != nil {
		return nil, err
	}

	return &client, nil
}

func CheckAdminAuth(phone string, code string) *models.Client {
	client, err := getAdminByPhone(phone)

	if err != nil || client == nil {
		fmt.Printf("Ошибка получения клиента: %v", err)
		return nil
	}

	sql := `
		SELECT TRUE
		FROM (
				 SELECT *
				 FROM security_codes
				 WHERE
					phone = $1 AND
					status = 'success' AND
					send_time >= NOW() - INTERVAL '1 HOUR'
				 ORDER BY send_time DESC
				 LIMIT 1
			 ) t
		WHERE code = $2
	`

	codeCorrect := false

	err = database.DB.GetConnection().QueryRow(context.Background(), sql, phone, code).Scan(&codeCorrect)

	if err != nil || !codeCorrect {
		fmt.Printf("Ошибка проверки кода: %v", err)
		return nil
	}

	return client
}

func SendSecurityCode(phone string) bool {
	client, err := getAdminByPhone(phone)

	if err != nil || client == nil {
		fmt.Printf("Ошибка получения клиента: %v", err)
		return false
	}

	code := randomString(5, true)
	err = api.SendSms(phone, code)

	status := "success"
	errorText := ""
	if err != nil {
		status = "error"
		errorText = err.Error()
	}

	codeUuid := uuid.New().String()
	sql := `
		INSERT INTO security_codes
		VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP)
	`

	_, err = database.DB.GetConnection().Exec(context.Background(), sql, codeUuid, phone, code, status, errorText)

	if err != nil {
		fmt.Printf("Ошибка сохранения кода подтверждения: %v", err)
		return false
	}

	if status == "error" {
		return false
	}

	return true
}
