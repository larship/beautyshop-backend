package models

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
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
	Price       float32     `json:"price"`
	Deleted     bool        `json:"deleted"`
}

const tableName = "checkin_list"

func GetBeautyshopCheckInList(beautyshopUuid string, dateFrom string, dateTo string) []CheckInItem {
	sql := fmt.Sprintf(`
		SELECT
			cl.uuid, cl.start_date, cl.end_date, cl.price, cl.deleted, cl.created_date,
			b.uuid, b.name, b.city, b.address,
			w.uuid, w.full_name,
			st.uuid, st.name
		FROM %s cl
		INNER JOIN beautyshops b ON b.uuid = cl.beautyshop_uuid
		INNER JOIN workers w ON w.uuid = cl.worker_uuid
		INNER JOIN service_types st ON st.uuid = cl.service_type_uuid
		WHERE
			cl.beautyshop_uuid = $1 AND
			cl.start_date >= $2 AND
			cl.start_date <= $3
	`, tableName)

	rows, err := database.DB.GetConnection().Query(context.Background(), sql, beautyshopUuid, dateFrom, dateTo)
	if err != nil {
		fmt.Printf("Ошибка получения записей: %v", err)
		return nil
	}

	var checkInItemsList []CheckInItem

	for rows.Next() {
		var checkInItem CheckInItem
		var beautyshop Beautyshop
		var worker Worker
		var serviceType ServiceType

		err = rows.Scan(&checkInItem.Uuid, &checkInItem.StartDate, &checkInItem.EndDate, &checkInItem.Price,
			&checkInItem.Deleted, &checkInItem.CreatedDate,
			&beautyshop.Uuid, &beautyshop.Name, &beautyshop.City, &beautyshop.Address,
			&worker.Uuid, &worker.FullName,
			&serviceType.Uuid, &serviceType.Name)
		if err != nil {
			fmt.Printf("Ошибка получения данных записей: %v", err)
		}

		// @TODO Выпилить POINT и перейти на lat/long
		beautyshop.Coordinates = pgtype.Point{
			Status: pgtype.Null,
		}

		checkInItem.Beautyshop = beautyshop
		checkInItem.Worker = worker
		checkInItem.ServiceType = serviceType

		checkInItemsList = append(checkInItemsList, checkInItem)
	}

	return checkInItemsList
}

// Создать запись в салон красоты.
func CreateCheckIn(beautyshopUuid string, clientUuid string, workerUuid string, serviceTypeUuid string, startTime int64) *CheckInItem {
	sql := fmt.Sprintf(`
		INSERT INTO %s
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, FALSE, $9)
	`, tableName)

	checkInUuid := uuid.New().String()
	checkInStartTime := time.Unix(startTime, 0)
	checkInEndTime := time.Unix(startTime, 0) // TODO Рассчитывать исходя из длительности услуги
	createdTime := time.Now()

	beautyshop := GetBeautyshopByUuid(beautyshopUuid)
	worker := GetWorkerByUuid(workerUuid)
	serviceType := GetWorkerServiceType(workerUuid, serviceTypeUuid)

	if beautyshop == nil || worker == nil || serviceType == nil {
		fmt.Printf("Ошибка при добавлении записи: не смогли найти салон красоты, работника или услугу")
		return nil
	}

	_, err := database.DB.GetConnection().Exec(context.Background(), sql, checkInUuid, beautyshopUuid, clientUuid,
		workerUuid, serviceTypeUuid, checkInStartTime.Format(time.UnixDate), checkInEndTime.Format(time.UnixDate),
		serviceType.Price, createdTime.Format(time.UnixDate))

	if err != nil {
		fmt.Printf("Ошибка при добавлении записи: %v", err)
		return nil
	}

	beautyshop.Workers = nil

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
		Price:       serviceType.Price,
		CreatedDate: createdTime,
	}

	return &checkInItem
}

// Отменить запись
func CancelCheckIn(uuid string) bool {
	sql := fmt.Sprintf(`
		UPDATE %s
		SET
			deleted = TRUE
		WHERE
			uuid = $1
	`, tableName)

	_, err := database.DB.GetConnection().Exec(context.Background(), sql, uuid)

	if err != nil {
		fmt.Printf("Ошибка при отмене записи: %v", err)
		return false
	}

	return true
}
