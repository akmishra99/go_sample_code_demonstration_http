#!/bin/bash
#
curl -X POST -H "Content-Type: application/json" -d '{"type": "sysinfo", "Payload": "windows11"}' http://localhost:8080/execute
