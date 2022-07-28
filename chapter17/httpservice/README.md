* use `sudo` to copy `httpservice.service` to `/etc/systemd/system/` 
* use `sudo` to copy `httpservice` binary file to `/usr/local/bin`
* start service `sudo systemctl start httpservice`
* check service status `sudo systemctl status httpservice`
    ```
    ● https.service - HTTP Server Application
         Loaded: loaded (/etc/systemd/system/https.service; disabled; vendor preset: enabled)
         Active: active (running) since Tue 2022-06-21 11:06:41 AEST; 2min 17s ago
      Main PID: 26593 (https)
      Tasks: 6 (limit: 9294)
      Memory: 2.9M
      CPU: 6ms
      CGroup: /system.slice/https.service
      └─26593 /usr/local/bin/https
    
    Jun 21 11:06:41 nanik systemd[1]: Started HTTP Server Application.
    Jun 21 11:06:41 nanik https[26593]: 2022/06/21 11:06:41 Server starting on port 8111
    ```
