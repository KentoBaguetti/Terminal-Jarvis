# Roadmap

## Idea
Build an AI tool that you can chat to in the terminal that has access to tools and services in your computer, allowing you to do whatever you want with just a prompt.
Since it'll live in the terminal, there will be less distractions that the user may encounter, just a raw assistant to talk to.

Think about it like an AI browser, that isnt a browser, and has access to local systems on your computer.

## Tools/Technologies Used:
- GitHub Copilot-SDK
- Golang
- Python3
- FastAPI
- FireCrawl
- Playwright MCP Server


## Current MCP Tools:
- UBC Badminton Drop-ins Schedule Finder

### commands
Run the CLI Tool:
`go run ./cmd/cli/.`
Starting the Python server:
`python3 -m uvicorn tools.python_server.main:app --host 0.0.0.0 --port 8000 --reload`