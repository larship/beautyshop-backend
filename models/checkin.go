package models

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/larship/beautyshop/database"
	"time"
)

type CheckInItem struct {
	Uuid        string
	Beautyshop  Beautyshop
	ClientUuid  string // @TODO Подгружать полностью клиента
	Worker      Worker
	ServiceType ServiceType
	StartDate   time.Time
	EndDate     time.Time
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
func CreateCheckIn(beautyshopUuid string, clientUuid string, workerUuid string, serviceTypeUuid string, startTime int64) bool {
	sql := fmt.Sprintf(`
		INSERT INTO %s
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`, tableName)

	checkInUuid := uuid.New().String()
	startTimeStr := time.Unix(startTime, 0).Format(time.UnixDate)
	endDateStr := time.Unix(startTime, 0).Format(time.UnixDate) // TODO Рассчитывать исходя из длительности услуги

	_, err := database.DB.GetConnection().Exec(context.Background(), sql, checkInUuid, beautyshopUuid, clientUuid,
		workerUuid, serviceTypeUuid, startTimeStr, endDateStr)

	if err != nil {
		fmt.Printf("Ошибка добавления мастера: %v", err)
		return false
	}

	return true
}
