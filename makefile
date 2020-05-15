all:
	go build -ldflags "-X main.version_date=`date +%d_%m_%Y__%H:%M`"
install:
	sudo cp ./rpn /usr/bin/rpn
