fuser -k 3019/tcp
fuser -k 3008/tcp
fuser -k 8000/tcp

cd linktreeResultar

#npm run dev
pm2 start npm -- name "linktree-clone-rm" -- start

cd ..

cd Resultarmind.com.br
#npm run dev
pm2 start npm -- name "saas-theme" -- start

cd ..


cd backend
export PATH=$PATH:/usr/local/go/bin

#echo "stoping any service on port 8000"

#fuser -k 8000/tcp
echo "we are now going to start the Serve side"
echo "Go is ready"
#go run cmd/main.go admin admin
go build cmd/main.go
go run cmd/main.go 
#./main
