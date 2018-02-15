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
token=$(./bin/jerakia token create myapp --quiet)

cat > ~/jrc <<EOF
export JERAKIA_TOKEN="$token"
export JERAKIA_URL="http://localhost:9992/v1"

jcurl() {
  curl -X GET -H "X-Authentication: \$JERAKIA_TOKEN" \$1
}
EOF

cat > ~/jerakia.sh <<EOF                                      
#!/bin/bash                                                   
                                                              
cd /home/ubuntu/jerakia
export RUBYLIB=\${PWD}/lib
export JERAKIA_CONFIG=./test/fixtures/etc/jerakia/jerakia.yaml
export PATH=\${PATH}:\${PWD}/bin
./bin/jerakia server
EOF
chmod +x ~/jerakia.sh

echo "[Unit]
Description = Jerakia

[Service]
ExecStart = /home/ubuntu/jerakia.sh
User = ubuntu

[Install]
WantedBy = multi-user.target" | sudo tee -a /etc/systemd/system/jerakia.service

sudo systemctl daemon-reload
sudo service jerakia start
