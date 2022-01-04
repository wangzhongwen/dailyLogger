package dailyLogger

import "os"

func Example()  {

	os.Setenv("debug", "1")
	logger := NewLogger("c:/tmp", "my_")
	NewDelete("c:/tmp",7, logger)
	logger.Printf("test")
	logger.Printf("test222222")
}