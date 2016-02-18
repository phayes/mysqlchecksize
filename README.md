# mysqlchecksize
A small utility for checking disk space used by mysql databases. Shows database sizes in KB. 

###Example
```bash
$ mysqlchecksize
bk_gsw_tizra	9336
cacti	1776
dev2_highwire_jnl_template	8848
$ mysqlchecksize cacti
1776
```

###Installation

For this utility to work, the executable must be owned by mysql and the SUID sticky bit must be set

```bash
$ wget "https://phayes.github.io/bin/current/mysqlchecksize/linux/mysqlchecksize.tar.gz"
$ tar -xf mysqlchecksize.tar.gz             # Extract mysqlchecksize from tarball
$ sudo cp mysqlchecksize /usr/bin           # Copy the mysqlchecksize executable to somewhere in your $PATH
$ sudo chown mysql /usr/bin/mysqlchecksize  # MySQL user must own the executable
$ sudo chmod u+s /usr/bin/mysqlchecksize    # Set the sticky bit so the file executes as the mysql user
```
