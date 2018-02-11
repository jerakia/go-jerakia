#!/bin/bash

sudo apt-get update
sudo apt-get install -y git make ruby-bundler ruby-dev build-essential zlib1g-dev libsqlite3-dev

sudo wget -O /usr/local/bin/gimme https://raw.githubusercontent.com/travis-ci/gimme/master/gimme
sudo chmod +x /usr/local/bin/gimme
eval $(/usr/local/bin/gimme 1.9)
/usr/local/bin/gimme 1.9 >> ~/.bashrc

go get github.com/jtopjian/go-jerakia

git clone https://github.com/crayfishx/jerakia
cd jerakia
bundle install
. test/environment.sh
token=$(./bin/jerakia token create myapp | grep myapp)

cat > ~/jrc <<EOF
export JERAKIA_TOKEN="$token"
export JERAKIA_URL="http://localhost:9992/v1"

jcurl() {
  curl -X GET -H 'X-Authentication: \$JERAKIA_TOKEN' \$1
}
EOF

echo "[Unit]
Description = Jerakia

[Service]
ExecStart = /home/ubuntu/files/jerakia.sh
User = ubuntu

[Install]
WantedBy = multi-user.target" | sudo tee -a /etc/systemd/system/jerakia.service

sudo systemctl daemon-reload
chmod +x ~/files/jerakia.sh
sudo service jerakia start
