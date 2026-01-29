import os
from dotenv import load_dotenv
from firecrawl import Firecrawl

load_dotenv()

API_KEY = os.getenv("FIRECRAWL_API_KEY")

firecrawl = Firecrawl(api_key=API_KEY)

def get_badminton_times():
    doc = firecrawl.scrape(url="https://recreation.ubc.ca/drop-in/gym-sports/", formats=[{
        "type": "json",
        "prompt": "Extract badminton drop in times from this page after selecting the 3 day option"
        }],
        actions=[
            {"type": "wait", "selector": "button.fc-agendaThreeDay-button"},
            {"type": "click", "selector": "button.fc-agendaThreeDay-button"},
            {"type": "wait", "milliseconds": 1500},
        ])
    return doc.json

if __name__ == "__main__":
    get_badminton_times()

