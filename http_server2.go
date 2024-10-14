package main

import (
    "fmt"
    "net/http"
    "os"
     "os/exec"
     "net"
     "time"
     "strings"
     "strconv"
     "log"
     "runtime"
     "reflect"
     "encoding/json"
)
type Commander interface {
    Ping(host string,w http.ResponseWriter) (PingResult, error)
    GetSystemInfo(given_hostname string ) (SystemInfo, error)
    // handlePostData(cmdr Commander) http.HandlerFunc
}

type PingResult struct {
    Successful bool
    Time       time.Duration
}

type SystemInfo struct {
    Hostname  string
    IPAddress string
}


type CommandRequest struct {
    Type    string `json:"type"`    // "ping" or "sysinfo"
    Payload string `json:"payload"` // For ping, this is the host
}

type CommandResponse struct {
    Success bool        `json:"success"`
    Data    interface{} `json:"data"`
    Error   string      `json:"error,omitempty"`
}


var all_ip_addresses []string

/*
func (c *commander) Ping(host string) (PingResult, error) {
   var ping_result PingResult
   ping_result.Successful = false
   ping_result.Time = -1

   ping_arguments := " -c 5  " + host
   cmd := exec.Command("./ping_test.sh", ping_arguments) // execute command e.g ping -c 4 192.168.1.41
   out, err := cmd.Output()       // Execute the command and get the output
   if err != nil {
       fmt.Println("Error:", err)
       fmt.Println(string(out))
       ping_result.Successful = false
       return ping_result,err
   } else {
       fmt.Println(string(out)) // Print the output
       // now determine time it took to execute ping
       time_keyword := "time "
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
           fmt.Println("time_found_at = %d , find_time_unit_index +%d\n",time_found_at,find_time_unit_index)
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
*/
func print_all_ip_address(given_ip_address_list []string) {

	for _,ip_address := range given_ip_address_list {
            fmt.Println("ip_address = ",ip_address)
	}



}
func print_all_ip_address_web(given_ip_address_list []string,w http.ResponseWriter)  {
        for _,ip_address := range given_ip_address_list {

            fmt.Fprintln(w,"ip_address = ",ip_address)
        }
}
func  GetSystemInfo(given_hostname string) (SystemInfo, error) {

        var host_info SystemInfo
	dummy_variable := -1
	return_value := &dummy_variable 
        // var err string
// 	var temp_host_name string 
	temp_host_name , err := os.Hostname()
        if  err != nil {
            fmt.Println("Error:", err)
            return host_info,err

        }
	if given_hostname == "" {
            host_info.Hostname = temp_host_name
        } else {
            host_info.Hostname = given_hostname
	}
        temp_ip_address , err3 := net.LookupHost( host_info.Hostname)
        if  err3 != nil {
            fmt.Println("Error:", err3)
            return host_info,err3

        }
	if runtime.GOOS == "windows" {
            print_all_ip_address(temp_ip_address)
	    all_ip_addresses = temp_ip_address
	    checkType( temp_ip_address,return_value )
	    if *return_value == 1 {
		    host_info.IPAddress = temp_ip_address[0]
	    } else if *return_value == 2 {
                host_info.IPAddress = temp_ip_address[len(temp_ip_address) -1]
            }
	} else {
            all_ip_addresses = temp_ip_address 
            print_all_ip_address(temp_ip_address)
	    checkType( temp_ip_address,return_value )
	    if *return_value == 1 {
                host_info.IPAddress = temp_ip_address[0]
	    } else if *return_value == 2 {
		    length_of_slice := len(temp_ip_address)
		if length_of_slice > 1   {
                    host_info.IPAddress = temp_ip_address[1]
	        } else {
		    host_info.IPAddress = temp_ip_address[0]
                }
            }
	}

        return host_info , err3
}
type commander struct{}

func NewCommander() Commander {
    return &commander{}
}



func main() {
    commander := NewCommander()
    server := &http.Server{
        Addr:    ":8080",
        Handler: handleRequests(commander),
    }
    log.Fatal(server.ListenAndServe())
}

func handleRequests(cmdr Commander) http.Handler {
    mux := http.NewServeMux()
    mux.HandleFunc("/execute", handlePostData(cmdr))
    mux.HandleFunc("/execute1", handleCommand(cmdr))
    mux.HandleFunc("/list",handleCommand_listd(cmdr))
    return mux
}
func (c *commander) Ping(host string,w http.ResponseWriter) (PingResult, error) {
   var ping_result PingResult
   var command string
   var ping_arguments  string
   ping_result.Successful = false
   ping_result.Time = -1
   if runtime.GOOS == "windows" {
       ping_arguments =  host
       command = ".\\ping_windows.bat "
   } else {
   	ping_arguments = " -c 5  " + host
	command = "./ping_test.sh"
   }
   cmd := exec.Command(command ,ping_arguments) // execute command e.g ping -c 4 192.168.1.41
   out, err := cmd.Output()       // Execute the command and get the output
   if err != nil {
       fmt.Println("Error:", err)
       fmt.Println(string(out))
       fmt.Fprintln(w,string(out))
       ping_result.Successful = false
       return ping_result,err
   } else {
       var time_keyword string
       fmt.Fprintln(w,string(out)) // Print the output
       // now determine time it took to execute ping
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
                fmt.Println("ping time = ",time_in_number)
                ping_result.Time = time.Duration(time_in_number * 1000  * 1000) // ping_result.Time's unit is in microsecond
           }
       }

   }


   ping_result.Successful = true
   return ping_result,err

}
func (c *commander) GetSystemInfo(hostname string ) (SystemInfo, error) {
    /*
    hostname, err := os.Hostname()
    if err != nil {
        return SystemInfo{}, err
    }

    */
    // Get IP address (implement this)
            system_info, err_system_info :=  GetSystemInfo(hostname)

            if err_system_info != nil {
                return SystemInfo{},err_system_info
            }
            return system_info,nil
         }

func handleCommand_listd(cmdr Commander) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        //   Parse request and execute command
	hostname_web := r.URL.Query().Get("host")
	system_info_list, err := GetSystemInfo(hostname_web)
	if err != nil {
              fmt.Fprintln(w,"error in getting  GetSystemInfo")
        } else {
            fmt.Fprintln(w, "host name = ",system_info_list.Hostname)
	    fmt.Fprintln(w, "ip address = ",system_info_list.IPAddress)
	    print_all_ip_address_web(all_ip_addresses,w)

       }
    }
}
func handleCommand(cmdr Commander) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Parse request and execute command
         hostname_web := r.URL.Query().Get("host") 
	 system_info_main,err := GetSystemInfo(hostname_web)
	 if err != nil {
            fmt.Fprintln(w,"Error:", err)

         } else {
             fmt.Fprintln(w,"system info hostname = ",system_info_main.Hostname)
             fmt.Fprintln(w,"system info ip address = ",system_info_main.IPAddress)
	     print_all_ip_address_web(all_ip_addresses,w)
             ping_result, _ :=  cmdr.Ping(string(system_info_main.IPAddress),w)
             if ping_result.Successful {
                 fmt.Fprintln(w,"Success in pinging to ip address = ",system_info_main.IPAddress)
                 fmt.Fprintln(w," ping time = ",ping_result.Time)
              } else {
                 fmt.Fprintln(w,"Error: in pinging host")
              }
         }


    }
}

func checkType(v interface{} ,value *int) {
	switch reflect.TypeOf(v).Kind() {
	case reflect.String: 
		fmt.Println("Variable is a string")
		*value = 1
		// return "string",nil
	
	case reflect.Slice:
		if reflect.TypeOf(v).Elem().Kind() == reflect.String {
			fmt.Println("Variable is a slice of strings")
			*value = 2
			// return "slice of stgring",nil
		} else {
			fmt.Println("Variable is a slice, but not of strings")
			*value = 3
			// return "slice of unknown","error"
		}
	default:
		fmt.Println("Variable is neither a string nor a slice")
		*value = -1
		// return "neither","error"
	}
}

func handlePostData(cmdr Commander) http.HandlerFunc {
    // return func handlePostData_do(w http.ResponseWriter, r *http.Request) {
    return func (w http.ResponseWriter, r *http.Request) {

        var commandRequest CommandRequest
	var commandResponse CommandResponse
        if r.Method == "POST" { 
            decoder := json.NewDecoder(r.Body)
	    err := decoder.Decode(&commandRequest)
	    if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	    }

	    // Process the data
	    fmt.Println("Received user data:", commandRequest)
	    if commandRequest.Type == "sysinfo" {
                system_info,err := GetSystemInfo(commandRequest.Payload )
		if err != nil {
                    http.Error(w, "error in executing sysinfo command", http.StatusBadRequest)
		    return 
		}
                commandResponse.Success = true
		commandResponse.Data    =  " hostname = " + system_info.Hostname  + " ip_address = " +  system_info.IPAddress
		json.NewEncoder(w).Encode(commandResponse)

	    } else if commandRequest.Type == "ping" {
		    hostname := commandRequest.Payload
		    ping_result,err := cmdr.Ping(hostname,w) 
		    if err != nil {

                         commandResponse.Success = false
			 commandResponse.Error   = "failure in pinging "
			 json.NewEncoder(w).Encode(commandResponse)

		    } else {
                         commandResponse.Success = ping_result.Successful
			 // commandResponse.Data = string(  ping_result.Time)
			 commandResponse.Data =  ping_result.Time
			 commandResponse.Error   = "\nping was successful\n"
			 json.NewEncoder(w).Encode(commandResponse)

		    }
            } else {

                http.Error(w, "invalid command ", http.StatusBadRequest)

            }




    }

   }
}
