# supervisor-api
supervisor client by golang



## supervisor-client

```go
    import(
      	"github.com/cierdes/supervisor-api/supervisor"
      	"fmt"
    )
    
    func main(){
      	client, err := supervisor.NewSupervisor("192.168.250.178:9001")
		if err != nil {
			fmt.Println(err)
          	 return
		}
		defer client.Close()

		r, err := client.ListMethods()
		if err != nil {
			fmt.Println(err)
          	 return
		}
      
      	 fmt.Println(r)
    }
	
```
