package models

import (
	"context"
	"fmt"
	"github.com/larship/beautyshop/database"
)

type ServiceType struct {
	Uuid  string  `json:"uuid"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

func GetServiceTypes(beautyshopUuid string) []ServiceType {
	sql := `
		SELECT st.*, wst.price
		FROM service_types st
		INNER JOIN workers_service_types wst ON wst.service_type_uuid = st.uuid
		INNER JOIN beautyshops_workers bw ON bw.worker_uuid = wst.worker_uuid
		WHERE bw.beautyshop_uuid = $1
	`

	rows, err := database.DB.GetConnection().Query(context.Background(), sql, beautyshopUuid)
	if err != nil {
		fmt.Printf("Ошибка получения списка услуг: %v", err)
		return nil
	}

	var serviceTypesList []ServiceType

	for rows.Next() {
		var item ServiceType
		err = rows.Scan(&item.Uuid, &item.Name, &item.Price)
		serviceTypesList = append(serviceTypesList, item)
	}

	return serviceTypesList
}
