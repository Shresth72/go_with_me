# Build GO project
go build -o tftp-server main.go server.go

# Run the binary
sudo ./tftp-server &
SERVER_PID=$!
echo "Server started with PID $SERVER_PID"

# Give the server a moment to start up
sleep 1

# Run the tftp commands
echo "get test2.txt" | tftp 127.0.0.1
sleep 1
echo "quit"
sleep 1

# Kill the server process
sudo kill $SERVER_PID

echo "Checking if port 69 is still in use"
if sudo ss -tulnp | grep -q ':69'; then
    echo "Port 69 is still in use. Please manually kill the process using:"
    echo "    sudo kill $SERVER_PID"
else
    echo "Port 69 is free"
fi
