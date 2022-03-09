#/usr/bin/bash
go build ./request.go
./request >output.json
echo "data recived in 'output.json'"
