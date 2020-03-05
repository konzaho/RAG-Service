# RAG-Service
Service for randomize ArmA escape scenario

API Documentation:

/rag (base URL of the service)  
-
no request header and body required


/rag/generate (generate ArmA escape scenario)
-
Request header:
    "Content-Type" : "application/json"
    
Request template:

	{
	"locations": ["altis", "malden", "tanoa"],
	"times": ["day", "night"],
	"factions": ["CSAT", "NATO"]
	}

Response template:

    {
    "Location":"malden",
    "Time":"night",
    "Faction":"CSAT"
    }


