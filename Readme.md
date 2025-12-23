# Golang CLI Reminder ⏰

A lightweight **cross-platform CLI reminder tool** written in Go.  
It allows you to set reminders using **natural language time expressions** and shows a native desktop notification when the time is reached.

---

## ✨ Features

-  Natural language time parsing  
-  Native OS notifications (Windows / macOS / Linux)  
-  Non-blocking execution (runs in background)  
-  Simple CLI interface  
-  Single binary usage  

---

## 🚀 Usage

```
golang-cli-reminder <time> <message>
```

## How It Works

Parses human-readable time using the when library

Validates that the given time is in the future

Spawns a background process using an environment marker

Sleeps until the reminder time

Displays a native system notification using beeep

The CLI exits immediately while the reminder continues running independently.

