package swagger

//go:generate rm -rf server
//go:generate mkdir -p server
//go:generate swagger generate server --quiet --target server --name filer-gateway-api --spec swagger.yaml --exclude-main
