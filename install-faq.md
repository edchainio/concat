**I got the following error: ./setup.sh: line 9: go: command not found** 

Please ensure you have go installed. If you have not installed go, please follow directions to install go at: 
https://golang.org/doc/install

**I already installed go and still got the above error**

Add go to your path environment. Instructions from the fo install guide:
Add /usr/local/go/bin to the PATH environment variable. You can do this by adding this line to your /etc/profile 
(for a system-wide installation) or $HOME/.profile:

export PATH=$PATH:/usr/local/go/bin

**When I run the setup and install commands, I get the following error: gx not found**

Please install gx by running:

go get -u github.com/whyrusleeping/gx (source: https://github.com/whyrusleeping/gx) 

Also install gx-go: 

go get -u github.com/whyrusleeping/gx-go

The packages are installed in your $GOPATH/bin folder. If this folder is not on your path. you must add it to your path for the gx and gx-go commands to work. 
