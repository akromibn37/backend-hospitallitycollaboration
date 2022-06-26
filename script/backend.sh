GOOS=linux GOARCH=amd64 go build
scp hospitalityCollaboration root@89.47.167.214:/var/www/backend/hospitalityCollaboration
pm2 delete hospitalityCollaboration
pm2 start hospitalityCollaboration

