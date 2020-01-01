package src
import (

	"fmt"

	"runtime"

)




func printStats(mem runtime.MemStats) {

	runtime.ReadMemStats(&mem)

	fmt.Println("mem.Alloc:", mem.Alloc)

	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)

	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)

	fmt.Println("mem.NumGC:", mem.NumGC)

	fmt.Println("-----")

}
func main()  {
	printStats(runtime.MemStats{})
}
