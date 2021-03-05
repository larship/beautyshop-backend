package models

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/larship/beautyshop/database"
	"time"
)

type CheckInItem struct {
	Uuid        string      `json:"uuid"`
	Beautyshop  Beautyshop  `json:"beautyshop"`
	Client      Client      `json:"client"`
	Worker      Worker      `json:"worker"`
	ServiceType ServiceType `json:"serviceType"`
	StartDate   time.Time   `json:"startDate"`
	EndDate     time.Time   `json:"endDate"`
	CreatedDate time.Time   `json:"createdDate"`
}

const tableName = "checkin_list"

func GetCheckInList(clientUuid string, from string, to string) []CheckInItem {
	sql := fmt.Sprintf(`
		SELECT
			cl.uuid, cl.start_date, cl.end_date,
			b.name, b.city, b.address,
			w.full_name, w.description
			st.name
		FROM %s cl
		INNER JOIN beautyshops b ON b.uuid = cl.beautyshop_uuid
		INNER JOIN workers w ON w.uuid = cl.worker_uuid
		INNER JOIN service_types st ON st.uuid = cl.service_type_uuid
		WHERE cl.client_uuid = $1
	`, tableName)

	rows, err := database.DB.GetConnection().Query(context.Background(), sql, clientUuid)
	if err != nil {
		fmt.Printf("Ошибка получения элементов расписания: %v", err)
		return nil
	}

	var checkInItem []CheckInItem

	for rows.Next() {
		var item CheckInItem
		// var beautyshop Beautyshop
		// var worker Worker
		// var serviceType ServiceType
		// @TODO Получать остальные структуры
		var tmp string

		err = rows.Scan(&item.Uuid, &item.StartDate, &item.EndDate, &tmp, &tmp, &tmp, &tmp, &tmp)
		if err != nil {
			fmt.Printf("Ошибка получения данных элементов расписания: %v", err)
		}

		checkInItem = append(checkInItem, item)
	}

	return checkInItem
}

// Создать запись в салон красоты.
func CreateCheckIn(beautyshopUuid string, clientUuid string, workerUuid string, serviceTypeUuid string, startTime int64) *CheckInItem {
	sql := fmt.Sprintf(`
		INSERT INTO %s
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`, tableName)

	checkInUuid := uuid.New().String()
	checkInStartTime := time.Unix(startTime, 0)
	checkInEndTime := time.Unix(startTime, 0) // TODO Рассчитывать исходя из длительности услуги
	createdTime := time.Now()

	_, err := database.DB.GetConnection().Exec(context.Background(), sql, checkInUuid, beautyshopUuid, clientUuid,
		workerUuid, serviceTypeUuid, checkInStartTime.Format(time.UnixDate), checkInEndTime.Format(time.UnixDate),
		createdTime.Format(time.UnixDate))

	if err != nil {
		fmt.Printf("Ошибка при добавлении записи: %v", err)
		return nil
	}

	beautyshop := GetBeautyshopByUuid(beautyshopUuid)
	beautyshop.Workers = nil

	worker := GetWorkerByUuid(workerUuid)
	serviceType := GetWorkerServiceType(workerUuid, serviceTypeUuid)

	var checkInItem = CheckInItem{
		Uuid:       checkInUuid,
		Beautyshop: *beautyshop,
		Client: Client{
			Uuid: clientUuid,
		},
		Worker:      *worker,
		ServiceType: *serviceType,
		StartDate:   checkInStartTime,
		EndDate:     checkInEndTime,
		CreatedDate: createdTime,
	}

	return &checkInItem
}
