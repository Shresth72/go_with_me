# Run server
go run . -- users <user_name>

# Connect and send commands using netcat
# Any user that belongs to the same groups in the user specified in the server
 sudo -g <user_name> -- nc -U /tmp/creds.sock
