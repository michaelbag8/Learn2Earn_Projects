## Exercise 2: Simulated API Call (Weather Example)
Objective: Practice designing prompts that mimic function calls.

Create a prompt: “Simulate a weather API response for Paris today.”

**Prompt:**
Simulate a weather API response for Abuja today.

**AI Model:**
```json
{
  "location": {
    "city": "Abuja",
    "country": "Nigeria",
    "latitude": 9.0765,
    "longitude": 7.3986,
    "timezone": "Africa/Lagos"
  },
  "current": {
    "timestamp": "2026-02-23T12:00:00+01:00",
    "temperature_c": 32.5,
    "temperature_f": 90.5,
    "feels_like_c": 34.0,
    "feels_like_f": 93.2,
    "humidity": 28,
    "pressure_hPa": 1012,
    "wind_kph": 15,
    "wind_mph": 9.3,
    "wind_degree": 220,
    "wind_direction": "SW",
    "uv_index": 7,
    "visibility_km": 10,
    "condition": {
      "text": "Partly cloudy",
      "icon": "https://cdn.weatherapi.com/weather/64x64/day/116.png",
      "code": 1003
    },
    "precip_mm": 0.0,
    "precip_in": 0.0
  },
  "forecast": {
    "date": "2026-02-23",
    "astro": {
      "sunrise": "06:45 AM",
      "sunset": "06:15 PM",
      "moonrise": "03:20 AM",
      "moonset": "02:45 PM",
      "moon_phase": "Waning Crescent",
      "moon_illumination": "28%"
    },
    "day": {
      "maxtemp_c": 34.0,
      "maxtemp_f": 93.2,
      "mintemp_c": 22.0,
      "mintemp_f": 71.6,
      "avgtemp_c": 28.0,
      "avgtemp_f": 82.4,
      "maxwind_kph": 18,
      "maxwind_mph": 11.2,
      "totalprecip_mm": 0.0,
      "totalprecip_in": 0.0,
      "avgvis_km": 10,
      "avghumidity": 30,
      "condition": {
        "text": "Sunny intervals",
        "icon": "https://cdn.weatherapi.com/weather/64x64/day/113.png",
        "code": 1000
      },
      "uv_index": 7
    }
  }
}
```

Define an expected JSON schema, e.g.:
{ "location": "Paris", "temperature_celsius": 18, "condition": "Cloudy" }

**Prompt:**
Simulate a weather API response for Abuja today. In this format { "location": "Paris", "temperature_celsius": 18, "condition": "Cloudy" }

**AI Model:**
```json
{
  "location": "Abuja",
  "temperature_celsius": 32,
  "condition": "Partly Cloudy"
}
```
## Test the model with different cities and compare outputs
**Prompt:**
Simulate a weather API response for Kaduna today. In this format { "location": "Paris", "temperature_celsius": 18, "condition": "Cloudy" }

**AI Model:**
```json
{
  "location": "Kaduna",
  "temperature_celsius": 30,
  "condition": "Sunny"
}
```
**Prompt:**
Simulate a weather API response for Maiduguri today. In this format { "location": "Paris", "temperature_celsius": 18, "condition": "Cloudy" }
**AI Model:**
```json
{
  "location": "Maiduguri",
  "temperature_celsius": 35,
  "condition": "Hot and Sunny"
}
```
## JSON Validation
JSON validation ensures that AI outputs are structured, consistent, and machine-readable. First, i checked the syntax to confirm arrays, objects, quotes, and commas are correct. Then verify that all required fields exist and that each field has the proper data type. i test multiple inputs to ensure the format is consistent, and optionally automate validation using a JSON parser.

## Output Comparison

**Prompt 1:** Open-ended, detailed, full API-style JSON (includes wind, humidity, forecast). Good for realistic simulations but verbose. It is detailed and flexible

**Prompt 2:** Specifies format, minimal fields (location, temperature_celsius, condition). It is concise and predictable
