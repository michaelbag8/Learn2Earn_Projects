# Exercise 2: Role-Based Mentoring and Feedback

## Ask the model to act as a coding mentor.
### Provide a simple piece of code (e.g., a Python function with a bug).

## **Prompt:** Act as a mentor reviewing this code. Provide corrections and explain improvements.
```py
def string_reversed(str):
    return “”.join(reversed(str))
print(string_reversed("michael", "bag"))
```

Alright,let's go through this carefully. Your code has a few issues, and I’ll break them down, explain why, and show how to improve it.

Your Original Code
```py
def string_reversed(str):
    return “”.join(reversed(str))
print(string_reversed("michael", "bag"))
```
## Issues
**1. Function parameter name:**

Using ***str*** as a parameter shadows Python’s built-in str type. This is risky because it can cause unexpected behavior if you try to use ***str()*** later in the code.

**2. Quotation marks:**
You’re using fancy quotes (“ ”) instead of standard quotes ("). Python will throw a syntax error.

**3. Function call arguments:**
You’re passing two arguments "michael", "bag" to a function that only accepts one. This will raise a TypeError.

Corrected Code
```py
def string_reversed(text):
    return "".join(reversed(text))

print(string_reversed("michael"))
print(string_reversed("bag"))
```

## Explanation of Improvements

| Issue                     | Correction                          | Why it matters                                      |
|---------------------------|-------------------------------------|-----------------------------------------------------|
| Shadowing `str`           | Changed parameter to `text`         | Avoids conflict with Python built-in `str` type.    |
| Wrong quotes              | Replaced `“ ”` with `"`             | Python only recognizes straight quotes for strings. |
| Wrong number of arguments | Call function once per string       | Matches function definition (1 argument).           |
| Readability               | Renamed parameter descriptively     | Makes the code easier to understand.                |


### Optional Enhancement
If you want to reverse multiple strings at once, you could allow your function to accept a list or multiple arguments:
```py
def string_reversed(*texts):
    return [text[::-1] for text in texts]

print(string_reversed("michael", "bag"))
# Output: ['leahcim', 'gab']
```
Using *texts lets you pass multiple strings.
text[::-1] is a Pythonic way to reverse a string.

# Review the response and check if the role influences feedback style.(Reflection)

The role actually influence the feedback, the LLM i used acted as a coding mentor and reviewed my code pointing out that errors and how the code can be improved for efficiency