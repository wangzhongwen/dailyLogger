
# dailyLogger
daily Logger

按天分割日志


 

#### How to Get:
``` sh
$ go get github.com/wangzhongwen/dailyLogger
```
#### How To Use:
``` go
import "github.com/wangzhongwen/dailyLogger"

logger := NewLogger("c:/tmp", "my_")
NewDelete("c:/tmp",7, logger)
logger.Printf("test")
logger.Printf("test222222")