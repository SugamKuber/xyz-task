package dal

import (
	"context"
	db "server/internals/init"
)

func SaveCarData(image []byte, carType, color, make, model, caption string) error {
	dbConn := db.GetDB()
	_, err := dbConn.Exec(context.Background(), `
        INSERT INTO cars (image, type, color, make, model, caption) 
        VALUES ($1, $2, $3, $4, $5, $6)`,
		image, carType, color, make, model, caption)
	return err
}

func GetCarData(id int) (map[string]interface{}, error) {
	dbConn := db.GetDB()
	row := dbConn.QueryRow(context.Background(), `
        SELECT type, color, make, model, caption 
        FROM cars WHERE id=$1`, id)

	var carType, color, make, model, caption string
	err := row.Scan(&carType, &color, &make, &model, &caption)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"type":    carType,
		"color":   color,
		"make":    make,
		"model":   model,
		"caption": caption,
	}, nil
}
