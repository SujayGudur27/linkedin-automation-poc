# LinkedIn Automation Proof of Concept

## Overview

This project is a **proof of concept** built using **Go** and the **Rod** browser automation library.  
It is created only for **educational and technical evaluation purposes**, as part of an assignment.

Automating LinkedIn violates LinkedIn’s Terms of Service.  
This project is **not intended for real-world usage or production deployment**.

---

## Objective

The objective of this project is to demonstrate:

- Browser automation using Go
- Clean and modular project architecture
- Human-like behavior simulation
- Basic anti-detection techniques
- Safe and responsible automation design

---

## Project Structure

The project is organized into the following modules:

cmd/
└── app/
└── main.go

internal/
├── auth/ - Login and authentication logic
├── browser/ - Browser setup using Rod
├── stealth/ - Human-like delays, typing, scrolling
├── search/ - Profile search logic
├── profile/ - Connection request logic
├── messaging/ - Messaging logic
├── config/ - Configuration loading
└── store/ - Placeholder for persistence

yaml
Copy code

This structure keeps responsibilities separated and makes the code easier to understand and extend.

---

## Features Implemented

### Authentication
- Reads credentials from environment variables
- Allows only test or dummy email accounts
- Blocks real LinkedIn accounts for safety

### Search
- Searches profiles using keywords
- Extracts profile URLs from search results

### Connection Requests
- Navigates to profile pages
- Identifies the Connect button
- Supports adding a short note
- Runs in safe (dry-run) mode by default

### Messaging
- Simulates follow-up messages
- Uses human-like typing delays
- Does not send messages unless automation is enabled

---

## Stealth Techniques

The project includes basic stealth mechanisms such as:

- Random delays between actions
- Human-like typing behavior
- Simple scrolling simulation
- Disabling automation flags in the browser

These techniques demonstrate how automation can mimic normal user behavior.

---

## Safety Measures

To prevent misuse, the project includes multiple safety checks:

- Automation is disabled by default
- Only test email domains are allowed:
  - `@test.com`
  - `@dummy.com`
  - `@example.com`
- Real LinkedIn accounts are blocked
- Actions are logged instead of executed in safe mode

---

## Configuration

### config.yml
```yaml
automation_enabled: false
daily_limit: 10
.env.example
env
Copy code
LINKEDIN_EMAIL=test@example.com
LINKEDIN_PASSWORD=password123
Running the Project
Running the project is optional and not required for evaluation.

bash
Copy code
go run cmd/app/main.go
When run with default settings, the browser opens but no real actions are performed.

Disclaimer
This project is created only to demonstrate technical concepts related to browser automation.
It must not be used on real LinkedIn accounts or in any production environment.

Author
Sujay Gudur

yaml
Copy code

---

## Next steps for you

1. Paste this into `README.md`
2. Save the file
3. Run:
   ```powershell
   git add README.md
   git commit -m "Update README documentation"
   git push

This is the video link of the explanation

https://drive.google.com/file/d/1upe1qVCy_b_ZOXvLeL-s87vBLQOJo9a1/view?usp=sharing
