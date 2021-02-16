package auth

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"github.com/larship/beautyshop/database"
	"github.com/larship/beautyshop/models"
	"math/rand"
)

const serverSalt = "dg353hy034"

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
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

	sessionIdHash := sha256.Sum256([]byte(RandomString(32)))
	sessionIdHashStr := hex.EncodeToString(sessionIdHash[:])

	saltHash := sha256.Sum256([]byte(RandomString(32)))
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
