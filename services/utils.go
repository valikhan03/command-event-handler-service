package services

import(
	"command-event-handler-service/db"
	"log"
)

func getProductName(pr_type, pr_id int) string {
	var name string
	if pr_type == 3 {
		query := `select title from tb_lots where id=$1`
		err := database.DBConn.Get(&name, query, pr_id)
		if err != nil{
			log.Printf("getProductName error: %v\n", err)
			return ""
		}
	}

	return name
}

