package main





import (
        "fmt"
        "os"
        "os/exec"
        "net"
	"time"
	"strings"
	"strconv"
	"runtime"
	"testing"
//      "exec"
)



type Commander interface {
    Ping(host string) (PingResult, error)
    GetSystemInfo() (SystemInfo, error)
}

type PingResult struct {
    Successful bool
    Time       time.Duration
}

type SystemInfo struct {
    Hostname  string
    IPAddress string
}

func (c *commander) Ping(host string) (PingResult, error) {
   var ping_result PingResult
   var command string
   var ping_arguments  string 
   ping_result.Successful = false
   ping_result.Time = -1
   if runtime.GOOS == "windows" {
	ping_arguments =  host
	command = ".\\ping_windows.bat "
        // cmd := exec.Command("ping ", ping_arguments)
   }  else {

       ping_arguments = " -c 5  " + host
       command = "./ping_test.sh"
       // cmd := exec.Command("./ping_test.sh", ping_arguments) // execute command e.g ping -c 4 192.168.1.41 
   }
   cmd := exec.Command(command,ping_arguments)
   out, err := cmd.Output()       // Execute the command and get the output
   if err != nil {
       fmt.Println("Error:", err)
       fmt.Println(string(out))
       ping_result.Successful = false
       return ping_result,err
   } else {
       fmt.Println(string(out)) // Print the output
       // now determine time it took to execute ping 
       var time_keyword string
       if runtime.GOOS == "windows" {
	  time_keyword =" Average = "
       } else {

           time_keyword = "time "
       }	       
       time_keyword_len := len(time_keyword)
       time_found_at := strings.Index(string(out),time_keyword)
       if time_found_at == -1 {
           ping_result.Successful = false
	   return ping_result,err
       } else {
           rest_of_string := (string(out)[time_found_at:])
	   find_time_unit := "ms"
           // mfind_time_unit_index :=  strings.Index(string(out),find_time_unit)
           find_time_unit_index :=  strings.Index(rest_of_string,find_time_unit)
	   fmt.Printf("time_found_at = %d , find_time_unit_index +%d\n",time_found_at,find_time_unit_index)
	   if find_time_unit_index != -1 {
	   	time_in_string := out[(time_found_at + time_keyword_len): (time_found_at + find_time_unit_index)]
		// time_in_number,err_time :=  strconv.Atoi(string(time_in_string))
		time_in_number,err_time :=  strconv.ParseInt(string(time_in_string), 10, 64)
		if err_time != nil {
                    fmt.Println("Error:", err_time)
		    ping_result.Successful = false
		    ping_result.Time = -1
		    return ping_result, err_time
                    
	        }
		fmt.Println("ping time = ",time_in_number);
		ping_result.Time = time.Duration(time_in_number * 1000  * 1000) // ping_result.Time's unit is in microsecond
           }
       }

   }


   ping_result.Successful = true
   return ping_result,err

}
func  GetSystemInfo() (SystemInfo, error) {

	var host_info SystemInfo
	temp_host_name , err := os.Hostname()
	if  err != nil {
            fmt.Println("Error:", err)
            return host_info,err

        }
	host_info.Hostname = temp_host_name
	temp_ip_address , err3 := net.LookupHost( host_info.Hostname) 
	if  err3 != nil {
            fmt.Println("Error:", err3)
            return host_info,err3

        }
	if runtime.GOOS == "windows" {
	    host_info.IPAddress = temp_ip_address[len(temp_ip_address) -1]
	} else {
            host_info.IPAddress = temp_ip_address[1] 
    	}

	return host_info , err3
}
type commander struct{}

func NewCommander() Commander {
    return &commander{}
}

func (c *commander) GetSystemInfo() (SystemInfo, error) {
    /*
    hostname, err := os.Hostname()
    if err != nil {
        return SystemInfo{}, err
    }

    */
    // Get IP address (implement this)
    system_info, err_system_info :=  GetSystemInfo()

    if err_system_info != nil {
        return SystemInfo{},err_system_info
    }
    return system_info,nil
}
/*
func main() {


    my_commander := NewCommander()
    system_info_main,err := my_commander.GetSystemInfo()
    if err != nil {
        fmt.Println("Error:", err)

    } else {
         fmt.Println("system info hostname = ",system_info_main.Hostname)
	 fmt.Println("system info ip address = ",system_info_main.IPAddress)
	 ping_result, _ :=  my_commander.Ping(string(system_info_main.IPAddress))
	 if ping_result.Successful {
              // fmt.Println("success = ip_address = %s, and ping time = %.2f",system_info_main.IPAddress,ping_result.Time)
	      fmt.Println("Success in pinging to ip address = ",system_info_main.IPAddress)
	      fmt.Println(" ping time = ",ping_result.Time)
         } else {
             fmt.Println("Error: in pinging host")
         }
    }


}

*/
func TestGetSystemInfo(t *testing.T) {
    cmdr := NewCommander()
    info, err := cmdr.GetSystemInfo()

    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

    if info.Hostname == "" {
        t.Error("Expected hostname to be non-empty")
    }

    if info.IPAddress == "" {
        t.Error("Expected IP address to be non-empty")
    }
}
