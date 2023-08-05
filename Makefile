test:
	go test
bin: test
	cd cmd; go build -o ../../bin/icinga-telegram-notification icinga-telegram-notification.go
upload:
	scp ../bin/icinga-telegram-notification root@$YOURSERVERNAME:/$YOURLOCATION
