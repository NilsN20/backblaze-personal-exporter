# Backblaze Personal Exporter
Open metrics exporter for monitoring various Metrics from the Backblaze Personal Program.

This is **not** for Backblaze B2
# Installation
## Windows
Download backblaze-personal-exporter.exe from the latest [release](https://github.com/NilsN20/backblaze-personal-exporter/releases/latest) <br />
Run the Program in a Terminal. 
``` 
C:\>backblaze-personal-exporter.exe
``` 
## Linux
Download backblaze-personal-exporter from the latest [release](https://github.com/NilsN20/backblaze-personal-exporter/releases/latest) <br />
Run the Program. <br>For Linux you'll have to specify the Path of the Backblaze Programdata folder. This will likely be in your Wine prefix under C:\ProgramData\Backblaze\
``` 
$./backblaze-personal-exporter --path /home/wine/prefix/drive_c/ProgramData/Backblaze
```
# Arguments
| Argument         | Default                   | Description                                                                                                                                    | Format  |
|------------------|---------------------------|------------------------------------------------------------------------------------------------------------------------------------------------|---------|
| --port           | 8090                      | Port the exporter listens on                                                                                                                   | Number  |
| --updateInterval | 5                         | How often the exporter updates data from the Backblaze files. A higher value results in faster metric refreshes at the cost of disk operations | Seconds |
| --backblazeData  | C:\ProgramData\Backblaze\ | Directory of the Backblaze programdata files                                                                                                   | Path    |


# Metrics
| Metric             | Description                                                                         | Format         |
|--------------------|-------------------------------------------------------------------------------------|----------------|
| transferStatus     | Weather or not it is currently transferring. 1 = Transferring, 0 = Not Transferring | Boolean (1/0)  |
| remainingFilesNum  | Number of remaining files which need to be backed up                                | Number         |
| totalFilesNum      | Number of total files marked for backup                                             | Number         |
| remainingBytes     | Number of remaining bytes which need to be backed up                                | Bytes          |
| totalBytes         | Number of total bytes marked for backup                                             | Bytes          |
| lastBackupComplete | Last time data was fully backed up                                                  | Unix Timestamp |
