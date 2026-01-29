from fastapi import FastAPI
from fastapi.concurrency import run_in_threadpool
from tools.firecrawl.main import get_badminton_times

app = FastAPI(title="FastAPI Server for MCP Tools")


@app.get("/health")
def health():
    return {"status" : "ok"}

# use await here for now, then switch to concurrency later when optimizing
# await is used as a easy quick fix
@app.get("/api/badminton-schedule")
async def get_badminton_schedule():
    return await run_in_threadpool(get_badminton_times)