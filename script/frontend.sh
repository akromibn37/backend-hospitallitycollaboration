cd frontend
npm run build
pm2 delete static-page-server-3000
pm2 serve build/ 3000 --spa