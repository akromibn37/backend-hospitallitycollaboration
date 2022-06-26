#install necessary software
sudo apt-get update
sudo apt-get install nginx
sudo apt-get install git
sudo apt install nano
sudo apt install vim
sudo ufw allow 'Nginx HTTP'

#install mysql and set user
sudo apt install mysql-server
sudo systemctl start mysql.service
sudo mysql
ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password by 'mynewpassword';
sudo mysql_secure_installation
mysql -u root -p
CREATE USER 'hospitality'@'localhost' IDENTIFIED BY 'Hospitality@0891777740';
CREATE USER 'rck'@'localhost' IDENTIFIED WITH mysql_native_password BY 'RCK@0891777740';
GRANT ALL PRIVILEGES ON *.* TO 'hospitality'@'localhost' WITH GRANT OPTION;
GRANT ALL PRIVILEGES ON *.* TO 'rck'@'localhost' WITH GRANT OPTION;
FLUSH PRIVILEGES;
exit

#install software for frontend react
sudo apt-get update
sudo apt-get install software-properties-common
sudo add-apt-repository universe
sudo add-apt-repository ppa:certbot/certbot
sudo apt-get update
sudo apt-get install certbot python3-certbot-nginx
sudo apt-get install npm
npm install -g serve
npm install -g pm2

#clone frontend repo
cd /var/www
git clone https://github.com/akromibn37/hospitallitycollaboration.git
mv hospitallitycollaboration frontend
npm install

#fix node and npm error
sudo rm -rf /usr/local/bin/npm /usr/local/share/man/man1/node* ~/.npm
sudo rm -rf /usr/local/lib/node*
sudo rm -rf /usr/local/bin/node*
sudo rm -rf /usr/local/include/node*
sudo apt-get purge nodejs npm
sudo apt autoremove
curl -sL https://deb.nodesource.com/setup_18.x -o /tmp/nodesource_setup.sh
sudo bash /tmp/nodesource_setup.sh
sudo apt-get install -y nodejs
node -v
npm -v
cd /var/www/frontend
npm install
npx browserslist@latest --update-db
export NODE_OPTIONS=--openssl-legacy-provider
npm run build
npm install -g serve
npm install -g pm2
pm2 serve build/ 3000 --spa
cd /etc/nginx/sites-available
vi default
# edit site as document and save
nginx -t
systemctl restart nginx
# change ip address to domain name

#add ssl encrypt
cd
sudo apt install snapd 
sudo snap install --classic certbot
cd /
certbot --nginx -d {domain_name} -d www.{domain_name}
systemctl restart nginx


###### CREATE DATABASE FOR PROJECT ######
mysql -u hospitality -p
CREATE DATABASE hospitality;
use hospitality;

## CREATE TABLE
## INITIAL DATA IN TABLE

###### INITIAL BACKEND SCRIPT ######
GOOS=linux GOARCH=amd64 go build
scp hospitalityCollaboration root@89.47.167.214:/var/www/backend/hospitalityCollaboration
scp conf/app_server.ini root@89.47.167.214:/var/www/backend/conf/app.ini


