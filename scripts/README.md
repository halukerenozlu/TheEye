# Scripts — TheEye

This repo uses small PowerShell scripts to avoid long Docker Compose commands on Windows.

## Prerequisites
- Docker Desktop installed and running
- Run scripts from repo root: `C:\dev\TheEye`

If PowerShell blocks script execution:
```powershell
Set-ExecutionPolicy -Scope CurrentUser RemoteSigned
```

## Commands

### Start stack (no rebuild)
Starts services in the background. Uses existing images/containers.
```powershell
.\scripts\up.ps1
```

### Start stack (rebuild images)
Rebuilds images and starts services in the background. Use after Dockerfile/code changes.
```powershell
.\scripts\up-build.ps1
```

### Stop (do NOT delete containers)
Stops running containers but keeps them (fast resume later).
```powershell
.\scripts\stop.ps1
```

### Start (resume previously created containers)
Starts previously created containers (after `stop`).
```powershell
.\scripts\start.ps1
```

### Restart (no delete)
```powershell
.\scripts\restart.ps1
```

### Show logs
All services:
```powershell
.\scripts\logs.ps1
```

Single service (example: api):
```powershell
.\scripts\logs.ps1 api
```

### Health check
Checks the API health endpoint.
```powershell
.\scripts\health.ps1
```

## Destructive commands

### Down (delete containers + network, keep DB data)
Stops and removes containers and networks, but keeps named volumes (Postgres data remains).
```powershell
.\scripts\down.ps1
```

### Purge (DELETE EVERYTHING, including DB data)
⚠️ This removes volumes; Postgres data will be wiped.
```powershell
.\scripts\purge.ps1
```
