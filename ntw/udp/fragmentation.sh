# Send pings over Ethernet
# min MTU -> 46bytes 
# max MTU -> 1500bytes
# -M -> prohibit fragmentation
# -s flag -> 1500bytes payload
timeout 2s ping -M do -s 1500 1.1.1.1
# This should fail as Header - 1528 bytes

# Send 28bytes less
timeout 4s ping -M do -s 1472 1.1.1.1
# Should work
