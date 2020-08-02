GOOS := darwin
GOARCH := amd64

.PHONY : wechat_public_account pub pub_all

wechat_public_account:
	GOOS=${GOOS} GOARCH=${GOARCH} go build -o  wechat_public_account main.go

pub:
	scp ./wechat_public_account ubuntu@119.29.112.104:~/app/wechat_public_account

pub_all:
	scp ./wechat_public_account ubuntu@119.29.112.104:~/app/wechat_public_account/
	scp -r ./conf ubuntu@119.29.112.104:~/app/wechat_public_account/

