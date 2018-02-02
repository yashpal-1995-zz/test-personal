package main

import ( 
   "database/sql"
  _"github.com/go-sql-driver/mysql"
  "log"
)

func main(){
    
      db,err :=sql.Open("mysql","user:password@tcp(127.0.0.1:8080)/yash")
 
      // 1st argument- driver name 
     // 2nd argument - driver specific (to access)
  
      if err!=nil{
        log.Fatal(err)
      }
     defer db.Close()
   err=db.Ping() 
// to check if db is available and accessible 
    if err!=nil{
      log.Fatal(err)
   } 
     var (
        id int
        name string
       )
  rows,err:=db.Query("select id,name where id>=?",3)
  //to send the query to db
  if err!=nil{
  log.Fatal(err)
   }
  defer db.Close()
  for rows.Next() {
	err := rows.Scan(&id, &name)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(id, name)
}
err = rows.Err()
if err != nil {
	log.Fatal(err)
}
    
}
