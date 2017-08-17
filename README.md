# Installing LIRC
Run lirc.sh; ny default this will install LIRC with IR Tx at GPIO 22 and IR Rx at GPIO 23

# Setting up a remote
To generate a remote configuration file, run the following commands:
```
# Stop lirc to free up /dev/lirc0
sudo /etc/init.d/lirc stop
# Create a new remote control configuration file (using /dev/lirc0) and save the output to ~/lircd.conf
irrecord -d /dev/lirc0 ~/lircd.conf
# Make a backup of the original lircd.conf file
sudo mv /etc/lirc/lircd.conf /etc/lirc/lircd_original.conf
# Copy over your new configuration file
sudo cp ~/lircd.conf /etc/lirc/lircd.conf
# Start up lirc again
sudo /etc/init.d/lirc start
```

Rename the configuration file by editing the 'name' in /etc/lirc/lircd.conf
Restart LIRC: ```sudo /etc/init.d/lirc restart```

# Transmit examples using LIRC
- List all of the commands that LIRC knows for 'your-remote': ```irsend LIST your-remote ""```
- Send the KEY_POWER command once: ```irsend SEND_ONCE your-remote KEY_POWER```

# Installing LIRC Server
#### For quick installation, run the install.sh script. Other wise, read on...


1. Compile app.go for target system. For example, compiling for Raspberry Pi Zero:
```
GOOS=linux GOARCH=arm GOARM=5 go build -v app.go
```
2. SCP app binary to system
3. Install supervisor: ```apt install supervisor```
4. Copy over supervisor.conf to /etc/supervisord/conf.d/
5. Create logs: ```touch /var/log/lirc_server.log```
6. Start supervisor: ```supervisorctl reread && supervisorctl update```
