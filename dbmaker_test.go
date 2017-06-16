// test connecting to dbsample5 and select * from sysinfo;

package dbmaker

import (
  "database/sql"
  "fmt"
  "testing"
)

func TestGetSysinfo(t *testing.T) {
  if db, err := sql.Open("dbmaker", "DSN=DBSAMPLE5; UID=SYSADM; PWD=;"); err != nil {
    t.Error(err)
  } else {
    if rows, err := db.Query("select * from sysinfo;"); err != nil {
      t.Error(err)
    } else {
      cols, err := rows.Columns()
      if err != nil {
        t.Error(err)
      }
      for _, colName := range cols {
        fmt.Printf("%s, ", colName)
      }
      fmt.Printf("\n------------------------------------------------------\n")
      var id,info,value string
      for rows.Next() {
        if err := rows.Scan(&id,&info,&value); err != nil {
          t.Error(err)
        }
        fmt.Printf("%s,%-20s,%s\n",id,info,value)
      }
      rows.Close()
    }
    
    db.Close()
  }
}
