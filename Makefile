GOOS := darwin
GOARCH := amd64

.PHONY : wechat_public_account pub pub_all

wechat_public_account:
	GOOS=${GOOS} GOARCH=${GOARCH} go build -o  wechant_public_account main.go

pub:
	scp ./wechant_public_account ubuntu@119.29.112.104:~/app/wechant_public_account

pub_all:
	scp ./wechant_public_account ubuntu@119.29.112.104:~/app/wechant_public_account/
	scp -r ./conf ubuntu@119.29.112.104:~/app/wechant_public_account/

