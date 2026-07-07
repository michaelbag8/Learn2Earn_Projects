# Exercise 3: Chaining Prompts for Agent Flow

Step 1 prompt: Ask for a user’s query (e.g., “What’s the temperature in Paris?”).
User’s  query

**Prompt:**
"What’s the temperature in Lagos?"

**Step 2 prompt:** Simulate a call to the weather API.

**Prompt:**
Simulate a call to the weather API.
- Extract the location from the user query: Lagos
- Simulate fetching data from a weather API:
    API response: { "location": "Lagos", "temperature": "31°C", "condition": "Sunny" }

**Step 3 prompt:** Format the final answer back to the user.

**Step 3 Output:**
"The current temperature in Lagos is 31°C with sunny conditions."

### Document the chained process and outputs.
Simulate a tool-using agent that can answer user queries by chaining multiple prompts:
- Ask for input
- Fetch information using a tool (simulated API)
- Format and return the answer

**Step 1:** Ask for User Query

Action: Prompt the user for a question.

Example User Input:
“What’s the temperature in Lagos?”

Output from Step 1:
Captured query: "What’s the temperature in Lagos?"

**Step 2:** Simulate Tool/API Call

Action: The agent interprets the query and calls the relevant tool. In this exercise, we simulate a weather API.

Process:
Extract location from user query → Lagos
Simulate API call:
```json
{
  "location": "Lagos",
  "temperature": "31°C",
  "condition": "Sunny"
}

Output from Step 2:
Raw API data: { "location": "Lagos", "temperature": "31°C", "condition": "Sunny" }
```
**Step 3:** Format Final Answer

Action: Convert the raw API response into a human-readable answer.

**Formatted Output:**
“The current temperature in Lagos is 31°C with sunny conditions.”
