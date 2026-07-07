# Exercise 2: Data Extraction Template

## Choose an unstructured text (e.g., "John Doe, age 29, lives in Paris and works as a software engineer.").

Alice Smith, a 34-year-old architect from Berlin, recently joined GreenDesign Inc. as a senior consultant.

## Create a template prompt:
"Extract the following fields from the text: Name, Age, Location, Occupation. Return the output in JSON format."

## **Prompt:**
Extract the following fields from the text: Name, Age, Location, Occupation. Return the output in JSON format. Alice Smith, a 34-year-old architect from Berlin, recently joined GreenDesign Inc. as a senior consultant.

## **AI Model:**
```json
{
  "Name": "Alice Smith",
  "Age": 34,
  "Location": "Berlin",
  "Occupation": "Architect"
}
```

## **Prompt:**
Extract the following fields from the text: Name, Age, Location, Occupation. Return the output in JSON format. On Aug 29th, 2024, Optimist Corp. announced in Chicago that its CEO, Brad Doe, 75 years old, would be stepping down after a successful $5 million funding round.

## **AI Model:**
```json
{
  "Name": "Brad Doe",
  "Age": 75,
  "Location": "Chicago",
  "Occupation": "CEO"
}
```

## **Prompt:**
Extract the following fields from the text: Name, Age, Location, Occupation. Return the output in JSON format. Dr. Abraham Musa is the primary author of this paper and a physician specializing in internal medicine.He was 34 years when he authored the paper, in Nigeria
## **AI Model:**
```json
{
  "Name": "Dr. Abraham Musa",
  "Age": 34,
  "Location": "Nigeria",
  "Occupation": "Physician specializing in internal medicine"
}
```

## **Validation Report:**

All inputs have been successfully validated against the defined JSON schema. Each output contains the required fields (Name, Age, Location, Occupation) with correct data types and consistent formatting. No missing keys, structural errors, or type mismatches were detected. The results are consistent across all tested inputs.