package models

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/larship/beautyshop/database"
)

type Worker struct {
	Uuid        string `json:"id"`
	FullName    string `json:"full-name"`
	Description string `json:"description"`
}

func GetWorkers(beautyshopUuid string) []Worker {
	sql := `
		SELECT h.*
		FROM workers h
		INNER JOIN beautyshops_workers bh ON bh.worker_uuid = h.uuid
		WHERE bh.beautyshop_uuid = $1
	`

	rows, err := database.DB.GetConnection().Query(context.Background(), sql, beautyshopUuid)
	if err != nil {
		fmt.Printf("Ошибка получения списка мастеров: %v", err)
		return nil
	}

	var workerList []Worker

	for rows.Next() {
		var item Worker
		err = rows.Scan(&item.Uuid, &item.FullName, &item.Description)
		workerList = append(workerList, item)
	}

	return workerList
}

func AddWorker(beautyshopUuid string, workerFullName string, workerDesc string) bool {
	// @TODO Также добавлять список услуг мастера!

	var sql string
	var err error
	beautyshop := GetBeautyshopByUuid(beautyshopUuid)

	if beautyshop == nil {
		fmt.Printf("При добавлении мастера не смогли найти салон красоты с uuid = %s", beautyshopUuid)
		return false
	}

	workerUuid := uuid.New().String()

	sql = `
		INSERT INTO workers
		VALUES ($1, $2, $3)
	`

	_, err = database.DB.GetConnection().Exec(context.Background(), sql, workerUuid, workerFullName, workerDesc)

	if err != nil {
		fmt.Printf("Ошибка добавления мастеа: %v", err)
		return false
	}

	sql = `
		INSERT INTO beautyshops_workers
		VALUES ($1, $2)
	`

	_, err = database.DB.GetConnection().Exec(context.Background(), sql, beautyshopUuid, workerUuid)

	if err != nil {
		fmt.Printf("Ошибка добавления связи салон красоты - мастер: %v", err)
		return false
	}

	return true
}
