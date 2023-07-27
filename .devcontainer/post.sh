echo -e "\nexport PATH=$PATH:/usr/local/go/bin:/home/jordan/node/bin:/home/jordan/w/statyk/output" >> /home/jordan/.bashrc

export PATH=$PATH:/usr/local/go/bin:/home/jordan/node/bin:/home/jordan/w/statyk/output

go install -v github.com/ramya-rao-a/go-outline@v0.0.0-20210608161538-9736a4bde949
go install -v golang.org/x/tools/gopls@latest

npm i -g npm@latest
npm i -g sass

go version
node --version
npm --version
sass --version