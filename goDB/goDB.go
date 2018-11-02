package main
import (
 "fmt"
 "database/sql"
 _ "github.com/lib/pq"
)
func main(){
//	connStr:="postgres://pgdb:pgdb@tcp(130.162.69.126:5432)/dmsm_test?sslmode=verify-full"
	connStr := "user=pgdb password=pgdb host=130.162.69.126 dbname=dmsm_test sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("Connection Established")
	}
	var menu_id string
	var display_key string
	var icon_url string
	var seq_no int
	var status string
	defer db.Close()
	rows,err:=db.Query(`SELECT * FROM menu_primary`)
	if err!=nil{
		panic(err)
	}
	for rows.Next(){
		err:=rows.Scan(&menu_id,&display_key,&icon_url,&seq_no,&status)
		if err!=nil{
			panic(err)
		}
		fmt.Println(menu_id,display_key,icon_url,seq_no,status)
	}
	defer rows.Close()
}