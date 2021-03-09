package models

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/larship/beautyshop/database"
)

type Worker struct {
	Uuid        string        `json:"uuid"`
	FullName    string        `json:"fullName"`
	Description string        `json:"description"`
	Services    []ServiceType `json:"services"`
}

func GetWorkerByUuid(workerUuid string) *Worker {
	sql := `
		SELECT w.*
		FROM workers w
		WHERE w.uuid = $1
	`

	var worker Worker

	err := database.DB.GetConnection().QueryRow(context.Background(), sql, workerUuid).Scan(&worker.Uuid,
		&worker.FullName, &worker.Description)

	if err != nil {
		fmt.Printf("Ошибка получения мастера: %v", err)
		return nil
	}

	return &worker
}

func GetWorkers(beautyshopUuid string) []*Worker {
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

	var workerList []*Worker

	for rows.Next() {
		var item Worker
		err = rows.Scan(&item.Uuid, &item.FullName, &item.Description)
		workerList = append(workerList, &item)
	}

	workerUuids := []string{}
	for _, worker := range workerList {
		workerUuids = append(workerUuids, worker.Uuid)
	}

	sql = `
		SELECT wst.worker_uuid, st.uuid, st.name, wst.price, wst.duration
		FROM workers_service_types wst
		INNER JOIN service_types st ON st.uuid = wst.service_type_uuid
		WHERE wst.worker_uuid = ANY($1)
	`
	rows, err = database.DB.GetConnection().Query(context.Background(), sql, workerUuids)

	if err != nil {
		fmt.Printf("Ошибка получения услуг мастеров: %v", err)
		return nil
	}

	for rows.Next() {
		var item ServiceType
		var workerUuid string
		err = rows.Scan(&workerUuid, &item.Uuid, &item.Name, &item.Price, &item.Duration)

		for _, worker := range workerList {
			if workerUuid == worker.Uuid {
				worker.Services = append(worker.Services, item)
			}
		}
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
		fmt.Printf("Ошибка добавления мастера: %v", err)
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
