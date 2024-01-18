<h1 align="center">
<img src="img/InsightNet Scanner.jpeg" alt="InsightNet Scanner" width="200px">
<br>
</h1>
 
 <p align="center">
<a href="#features">Features</a> --
<a href="#Installation">Installation</a> --
<a href="#usage">Usage</a>
</p>

 InsightNet Scanner est un scanneur de ports écrit en Go qui vous permet d'énumérer les ports ouverts en TCP sur un hôte distant.

# Features

 <h1 align="center">
<img src="img/term.png" alt="Linux Terminal" width="700px">
<br>
</h1>

- Scan de port TCP en utilisant des goroutines --> utilisation de paramètres afin d'en tirer le meilleur temps
- Choix du nombre de **workers**
- Choix de la **plage de ports** ou du **port**
- Choix du nombre de **port par plage**

# Installation

 ``` git clone https://github.com/ThomasRAYNAUD/ELP.git ```

L'exécutable .go est dans ./GO/scan.go


# Usage

 ```yaml
Usage:
./naabu [flags]
 INPUT:
-host string[] hosts to scan ports for (comma-separated)
-list, -l string list of hosts to scan ports (file)
-exclude-hosts, -eh string hosts to exclude from the scan (comma-separated)
-exclude-file, -ef string list of hosts to exclude from scan (file)
 PORT:
-port, -p string ports to scan (80,443, 100-200)
-top-ports, -tp string top ports to scan (default 100) [full,100,1000]
-exclude-ports, -ep string ports to exclude from scan (comma-separated)
-ports-file, -pf string list of ports to scan (file)
-port-threshold, -pts int port threshold to skip port scan for the host
-exclude-cdn, -ec skip full port scans for CDN/WAF (only scan for port 80,443)
-display-cdn, -cdn display cdn in use
 RATE-LIMIT:
-c int general internal worker threads (default 25)-rate int packets to send per second (default 1000)
 UPDATE:
-up, -update update naabu to latest version
-duc, -disable-update-check disable automatic naabu update check
 OUTPUT:
-o, -output string file to write output to (optional)
-j, -json write output in JSON lines format
-csv write output in csv format
 CONFIGURATION:
-scan-all-ips, -sa scan all the IP's associated with DNS record
-ip-version, -iv string[] ip version to scan of hostname (4,6) - (default 4)
-scan-type, -s string type of port scan (SYN/CONNECT) (default "s")
-source-ip string source ip and port (x.x.x.x:yyy)
-interface-list, -il list available interfaces and public ip
-interface, -i string network Interface to use for port scan
-nmap invoke nmap scan on targets (nmap must be installed) - Deprecated
-nmap-cli string nmap command to run on found results (example: -nmap-cli
'nmap -sV')
-r string list of custom resolver dns resolution (comma separated or from file)
-proxy string socks5 proxy (ip[:port] / fqdn[:port]
-proxy-auth string socks5 proxy authentication (username:password)
-resume resume scan using resume.cfg
-stream stream mode (disables resume, nmap, verify, retries, shuffling, etc)
-passive display passive open ports using shodan internetdb api
-irt, -input-read-timeout value timeout on input read (default 3m0s)
-no-stdin Disable Stdin processing
 HOST-DISCOVERY:
-sn, -host-discovery Perform Only Host Discovery
-Pn, -skip-host-discovery Skip Host discovery
-ps, -probe-tcp-syn string[] TCP SYN Ping (host discovery needs to be enabled)
-pa, -probe-tcp-ack string[] TCP ACK Ping (host discovery needs to be enabled)
-pe, -probe-icmp-echo ICMP echo request Ping (host discovery needs to be enabled)
-pp, -probe-icmp-timestamp ICMP timestamp request Ping (host discovery needs to be
enabled)
-pm, -probe-icmp-address-mask ICMP address mask request Ping (host discovery needs
to be enabled)
-arp, -arp-ping ARP ping (host discovery needs to be enabled)
-nd, -nd-ping IPv6 Neighbor Discovery (host discovery needs to be enabled)-rev-ptr Reverse PTR lookup for input ips
 OPTIMIZATION:
-retries int number of retries for the port scan (default 3)
-timeout int millisecond to wait before timing out (default 1000)
-warm-up-time int time in seconds between scan phases (default 2)
-ping ping probes for verification of host
-verify validate the ports again with TCP verification
 DEBUG:
-health-check, -hc run diagnostic check up
-debug display debugging information
-verbose, -v display verbose output
-no-color, -nc disable colors in CLI output
-silent display only results in output
-version display version of naabu
-stats display stats of the running scan (deprecated)
-si, -stats-interval int number of seconds to wait between showing a statistics update
(deprecated) (default 5)
-mp, -metrics-port int port to expose nuclei metrics on (default 63636)
```
