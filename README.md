# Floyd Player 
## Learning GO - Simple M3U List player with mplayer

#### Player
![playlist](screenshots/floyd.png?raw=true)
![mplayer](screenshots/mplayer.png?raw=true)

```
# Clone
git clone https://github.com/talesluna/floyd-player-golang/ && cd floyd-player-golang

# Install go
sudo apt install golang

# Config GOPATH
export GOPATH="../youPath"

# Download packages
go get github.com/golang/gxui

# Build
go build -o floyd

# Run 
go run *.go path/you/list.m3u   # For run dev code
./floyd path/you/list.m3u       # For binary

```
