
**dbmaker** driver is the odbc-compatible driver for DBMaker database management system. It calls into dmapi54.dll on Windows, and use cgo dynamic linked to libdmapic.so on Linux. NOTE, **dbmaker** doesn't use unixODBC or windows ODBC Driver Manager.

On windows, you should add the path contains dmapi54.dll into the env variable %PATH%, then run your program. e.g. 
```
set PATH=C:\DBMaker\5.4\bin;%PATH%
go test github.com/dbmaker-go/dbmaker
```
If you want to switch to DBMaker 5.5, you should modify zapi_windows.go to load dmapi55.dll
```
  moddmapi54 = syscall.NewLazyDLL("dmapi55.dll")
```

On linux, you can modify cgo linux LDFLAGS/CFLAGS of api_unix.go/zapi_unix.go to swith to DBMaker 5.5. 

The cgo flags for static linking to libdmapic.a:
```
// #cgo linux LDFLAGS: -ldmapic -ldl -lm -L/home/dbmaker/5.4/lib
// #cgo linux CFLAGS: -I/home/dbmaker/5.4/include
// #include "sql.h"
// #include "sqlext.h"
```

For dynamic linking to libdmapic.so:
```
// #cgo linux LDFLAGS: -ldmapic  -L/home/dbmaker/5.4/lib/so
// #cgo linux CFLAGS: -I/home/dbmaker/5.4/include
// #include "sql.h"
// #include "sqlext.h"
```
If using libdmapic.so, env variable LD_LIBRARY_PATH must contain the path in where libdmapic.so resided. e.g.:
```
export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/home/dbmaker/5.4/lib/so
go test github.com/dbmaker-go/dbmaker ## need start DBSAMPLE5 DB before hand.
```

A simple example:
```
import (
    "database/sql"
    _ "github.com/dbmaker-go/dbmaker"
)

func main(){
	db, err := sql.Open("dbmaker","DSN=DB1;UID=SYSADM;PWD=xxx;");
	db.Query(...)
	//...
}
```

For more information about DBMaker, please refer to www.dbmaker.com.tw .
