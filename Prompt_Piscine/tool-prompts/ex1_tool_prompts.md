# Exercise 1: Structured JSON Output

### Write a prompt that asks for unstructured information (e.g., “List three cities with their populations.”).

**Prompt:** 
List the top three most populous countries with their population

**AI Model:**
Here are the top three most populous countries in the world with their estimated populations (2026 figures):

1. India – approximately 1,476,625,576 people (most populous country) (Worldometer)
2. China – approximately 1,412,914,089 people (second largest) (Worldometer)
3. United States – approximately 349,035,494 people (third largest) (Worldometer)
These estimates are based on the most recent global population data available. (Worldometer)


***Rewrite the prompt to enforce structured JSON output:***

Example: “Return the following in JSON format: [{city: string, population: number}].”

**Prompt:** 
Return the following in JSON format: [{city: string, population: number}]. List the top three most populous countries with their population.

**AI Model:**
```json
[
{"city": "India", "population": 1430000000},
{"city": "China", "population": 1410000000},
{"city": "United States", "population": 340000000}
]
```
## Test with multiple inputs and verify the JSON stays valid.

**Prompt:** 
List the top three footballers with countries and their population.

**AI Model:**
Here are three of the top footballers in the world, along with their countries and the approximate population of those countries:

1. Lionel Messi – Argentina
Population of Argentina: ~46 million
2. Cristiano Ronaldo – Portugal
Population of Portugal: ~10 million
3. Kylian Mbappé – France
Population of France: ~68 million
Note: Population figures are approximate recent estimates.

**Prompt:**
Return the following in JSON format: [{city: string, population: number}]. List the top three footballers with countries and their population.

**AI Model:**
```json
[
{"city": "Argentina", "population": 46000000},
{"city": "Portugal", "population": 10300000},
{"city": "France", "population": 68000000}
]
```
**Prompt:**
Return the following in JSON format: [{city: string, population: number}]. List the three poorest cities with their population.

**AI Model:**
```json
[
  {"city": "Kinshasa", "population": 17000000},
  {"city": "Lagos", "population": 21000000},
  {"city": "Kabul", "population": 4600000}
]
```

**Prompt:**
Return the following in JSON format: [{city: string, population: number}]. List the three richest cities with their population.

**AI Model:**
```json
[
  {"city": "Tokyo", "population": 37400000},
  {"city": "New York City", "population": 8400000},
  {"city": "London", "population": 9500000}
]
```
### Comparison
The difference between the two prompts is mainly about output structure and control:
“Return the following in JSON format: [{city: string, population: number}]. List the three richest cities with their population.”
- Clearly specifies the output format (JSON).
- Reduces complexity in how the answer should be presented.
- Useful for APIs, data processing, or programming tasks.

“List the three richest cities with their population.”
- Does not specify a format.
- The response was bullet points.
- More flexible but less controlled.
- Better suited for general reading

The first prompt is structured and constrained, guiding both content and format, while the second is open-ended, focusing only on content without formatting requirements.


