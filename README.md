# mysql_tw
mysql 连接 增删改查

go get github.com/tw-iot/mysql_tw

### 示例
```
package main

import (
	"fmt"
	"github.com/tw-iot/mysql_tw"
	"time"
)

type Data_collect struct {
	Collect_data_id string    `db:"collect_data_id"`
	Collect_time    int64     `db:"collect_time"`
	Device_id       string    `db:"device_id"`
	Create_time     time.Time `db:"create_time"`
}

func main() {
	mysqlInfo := mysql_tw.MysqlInfo{"tcp", "192.168.146.18", 3306,
		"root", "root@890", "iot-data",
		"utf8", "loc=Asia%2FShanghai&parseTime=true",
		100, 20, 100 * time.Second}
	//mysqlInfo := mysql_tw.NewMysqlInfo("192.168.146.18", "root",
	//	"root@890", "iot-data", 3306)
	mysql_tw.MysqlInit(&mysqlInfo)

	data_collect := new(Data_collect)
	row := mysql_tw.MysqlTw.QueryRow("select collect_data_id,collect_time,device_id,create_time from data_collect where collect_data_id = ?", "000127bdce814aa88c12e2c6ab2327ae")
	//row.scan中的字段必须是按照数据库存入字段的顺序，否则报错
	if err := row.Scan(&data_collect.Collect_data_id, &data_collect.Collect_time,
		&data_collect.Device_id, &data_collect.Create_time); err != nil {
		fmt.Printf("scan failed, err:%v", err)
		return
	}
	fmt.Println(data_collect)
}

```
