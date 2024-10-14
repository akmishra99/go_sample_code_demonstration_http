this is sample implementation of requirements at https://gist.github.com/greenido/c725d654b2d4acbf90c46c1bda96a950 

there is windows application installation file go_sample_code_demonstration_http.msi

it is installation file for windows 11 64 bit (x86_64)


in the installation directory , there are following files

10/13/2024  06:07 PM    <DIR>          .
10/13/2024  05:33 PM    <DIR>          ..
10/13/2024  05:54 PM    <DIR>          docs
10/13/2024  05:34 PM         3,119,893 GetSystemInfo
10/13/2024  05:35 PM         3,444,224 GetSystemInfo.exe
10/13/2024  05:34 PM             4,687 GetSystemInfo.go
10/13/2024  05:34 PM             5,074 GetSystemInfo_test.go
10/13/2024  05:34 PM         7,555,222 http_server2
10/13/2024  06:07 PM         8,104,960 http_server2.exe
10/13/2024  06:07 PM            12,427 http_server2.go
10/13/2024  05:34 PM                34 ping_test.sh
10/13/2024  05:34 PM                 7 ping_windows.bat
10/13/2024  05:55 PM    <DIR>          tests
               9 File(s)     22,246,528 bytes
               4 Dir(s)  1,309,888,794,624 bytes free


docs directory contains this README.md this file 

tests directory contains Linux bash scripts to test using curl from Linux hosts ( I have tested this test bash (shell scripts) program using ubuntu Ubuntu 24.04 LTS)

GetSystemInfo.exe and
http_server2.exe

are windows 11 executable files

and

GetSystemInfo

and

http_server2

are Linux (Ubuntu 24.04 LTS) executables

GetSystemInfo.go

is golang (go) file which implements finding ip address from given hostname, if host name is not given then hostname of host is used on which this program is running, there is 

 GetSystemInfo_test.go

is golang (go) test file to test functionality implemented by GetSystemInfo.go

( here is result of invoking following command 

 C:\Users\akmis\go_sample_code>go test -v  GetSystemInfo_test.go
=== RUN   TestGetSystemInfo
--- PASS: TestGetSystemInfo (0.03s)
PASS
ok      command-line-arguments  0.193s   


)

to run GetSystemInfo.exe use 


(
  
C:\Users\akmis\go_sample_code>GetSystemInfo.exe
system info hostname =  windows11
system info ip address =  10.0.0.139

C:\Users\akmis\go_sample_code>ping 10.0.0.139

Pinging 10.0.0.139 with 32 bytes of data:
Reply from 10.0.0.139: bytes=32 time<1ms TTL=128
Reply from 10.0.0.139: bytes=32 time<1ms TTL=128
Reply from 10.0.0.139: bytes=32 time<1ms TTL=128
Reply from 10.0.0.139: bytes=32 time<1ms TTL=128

Ping statistics for 10.0.0.139:
    Packets: Sent = 4, Received = 4, Lost = 0 (0% loss),
Approximate round trip times in milli-seconds:
    Minimum = 0ms, Maximum = 0ms, Average = 0ms

time_found_at = 469 , find_time_unit_index +12
ping time =  0
Success in pinging to ip address =  10.0.0.139
 ping time =  0s

C:\Users\akmis\go_sample_code>


)

above is sample execution of program GetSystemInfo.exe on windows 11 




to run on Linux host use 
(




akmishra@akmishra-A520I-AC2:~/scratch_oct_13_2024$ ./GetSystemInfo
system info hostname =  akmishra-A520I-AC2
system info ip address =  192.168.1.41
PING 192.168.1.41 (192.168.1.41) 56(84) bytes of data.
64 bytes from 192.168.1.41: icmp_seq=1 ttl=64 time=0.027 ms
64 bytes from 192.168.1.41: icmp_seq=2 ttl=64 time=0.038 ms
64 bytes from 192.168.1.41: icmp_seq=3 ttl=64 time=0.039 ms
64 bytes from 192.168.1.41: icmp_seq=4 ttl=64 time=0.041 ms
64 bytes from 192.168.1.41: icmp_seq=5 ttl=64 time=0.040 ms

--- 192.168.1.41 ping statistics ---
5 packets transmitted, 5 received, 0% packet loss, time 4119ms
rtt min/avg/max/mdev = 0.027/0.037/0.041/0.005 ms

time_found_at = 444 , find_time_unit_index +9
ping time =  4119
Success in pinging to ip address =  192.168.1.41
 ping time =  4.119s
akmishra@akmishra-A520I-AC2:~/scratch_oct_13_2024$
)
above is sample execution of program GetSystemInfo on (Ubuntu 24.04 LTS)



http_server2.go implements basic http server which runs on windows 11 and ubuntu Linux and accepts two commands using JSON as transport between Linux test machine (using curl) commands are sysinfo and ping to run http_server2.exe   use 

(
   
PS C:\Users\akmis\go_sample_code> .\http_server2


and on Linux (ubuntu Ubuntu 24.04 LTS)

run 

 curl -X POST -H "Content-Type: application/json" -d '{"type": "sysinfo", "Payload": "windows11"}' http://192.168.1.40:8080/execute
{"success":true,"data":" hostname = windows11 ip_address = 10.0.0.139"}



or 

akmishra@akmishra-A520I-AC2:~/scratch_oct_13_2024$ ../scratch_oct_13_2024_test/sysinfo_windows11.sh
{"success":true,"data":" hostname = windows11 ip_address = 10.0.0.139"}


in above command sysinfo was dispatched using curl to windows 11 (ip address 192.168.1.40) to get ip address from host name , windows 11 has more then one network interface 


here is output from http_server2 on windows command terminal 



PS C:\Users\akmis\go_sample_code> .\http_server2
Received user data: {sysinfo windows11}
ip_address =  fe80::b7a9:ac98:bc66:bfdb%vEthernet (WSL (Hyper-V firewall))
ip_address =  fe80::ef9:80b6:46f8:d0e4%Ethernet 3
ip_address =  fe80::4b82:2001:cb6b:ad0c%Ethernet 2
ip_address =  fe80::5930:5ba7:3f7d:93ff%Ethernet
ip_address =  2601:644:937e:ad50:b076:a211:f719:3f00
ip_address =  2601:644:937e:ad50::5ec7
ip_address =  2601:644:937e:ad50:4edb:6bf8:c957:66d5
ip_address =  172.27.208.1
ip_address =  192.168.1.40
ip_address =  192.168.56.1
ip_address =  10.0.0.139
Variable is a slice of strings


in above last ip address of host for which sysinfo command was executed is  returned on this case it is "10.0.0.139", this is by design


for invalid command  use 

akmishra@akmishra-A520I-AC2:~/scratch_oct_13_2024$ curl -X POST -H "Content-Type: application/json" -d '{"type": "invalid", "Payload": "windows11"}' http://192.168.1.40:8080/execute
invalid command
akmishra@akmishra-A520I-AC2:~/scratch_oct_13_2024$


and on http_server2 command window output is
PS C:\Users\akmis\go_sample_code> .\http_server2

Received user data: {invalid windows11}


)

now description and example of http_server2 handling ping command 

to test pinging to host akmishra-A520I-AC2 use


akmishra@akmishra-A520I-AC2:~/scratch_oct_13_2024$ curl -X POST -H "Content-Type: application/json" -d '{"type": "ping", "Payload": "akmishra-A520I-AC2"}' http://192.168.1.40:8080/execute

here is output of above command 


C:\Users\akmis\go_sample_code>ping akmishra-A520I-AC2

Pinging akmishra-A520I-AC2.local [2601:644:937e:ad50:7656:3cff:feb9:b471] with 32 bytes of data:
Reply from 2601:644:937e:ad50:7656:3cff:feb9:b471: time<1ms
Reply from 2601:644:937e:ad50:7656:3cff:feb9:b471: time=1ms
Reply from 2601:644:937e:ad50:7656:3cff:feb9:b471: time<1ms
Reply from 2601:644:937e:ad50:7656:3cff:feb9:b471: time<1ms

Ping statistics for 2601:644:937e:ad50:7656:3cff:feb9:b471:
    Packets: Sent = 4, Received = 4, Lost = 0 (0% loss),
Approximate round trip times in milli-seconds:
    Minimum = 0ms, Maximum = 1ms, Average = 0ms

{"success":true,"data":0,"error":"\nping was successful\n"}
akmishra@akmishra-A520I-AC2:~/scratch_oct_13_2024$


here is output of http_server2 

on terminal window


PS C:\Users\akmis\go_sample_code> .\http_server2
Received user data: {ping akmishra-A520I-AC2}
time_found_at = 608 , find_time_unit_index +12
ping time =  0

above output prints result of JSON encoded result in which success = true indicates that ping command was successful and 0 indicates that it ping time is 0 seconds,error indicates that  ping was successful


for invalid hostname or host which down here is example 


akmishra@akmishra-A520I-AC2:~/scratch_oct_13_2024$ curl -X POST -H "Content-Type: application/json" -d '{"type": "ping", "Payload": "akmishra-A520I-AC"}' http://192.168.1.40:8080/execute

C:\Users\akmis\go_sample_code>ping akmishra-A520I-AC
Ping request could not find host akmishra-A520I-AC. Please check the name and try again.

{"success":false,"data":null,"error":"failure in pinging "}


and here is output of http_server2 on command window 

PS C:\Users\akmis\go_sample_code> .\http_server2

Received user data: {ping akmishra-A520I-AC}
Error: exit status 1

C:\Users\akmis\go_sample_code>ping akmishra-A520I-AC
Ping request could not find host akmishra-A520I-AC. Please check the name and try again.
