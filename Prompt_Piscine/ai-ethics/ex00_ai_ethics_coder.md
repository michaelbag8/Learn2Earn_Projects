# Part A: The Critical Distinction

## Write down your honest answers:
### How have you used AI for coding so far?
    I used AI to help me learn, especially when I am stuck on a problem but I must have tried to solve it by myself before I consult AI.
### Do you ask AI for solutions before trying yourself?
    No, I usually try out the solution myself before I use AI. In most cases it is when I am stuck that I go to AI.
### Can you explain code you've submitted without AI's help?
    Very much, i can explain code that i have submitted without AI’s help
### What would happen if AI was suddenly unavailable during an exam or interview?
    Sincerely i do not have any worries with the unavailability of AI during an Exam or Job test, because I have never used AI during such cases. I use AI to amplify my learning experience, I do not use it for exams or interviews. Hence i will do well in any test or exam without AI.
## Identify your current pattern: Which learner are you now? Learner A: "AI is my answer generator"
I am a Learner B type:  AI is my learning amplifier
    - I tried to solve my problem first.
    - I test my understanding by explaining concepts to my colleagues.
    - I  used AI to get a deeper understanding of concepts.
    - I don’t outsource my thinking to AI

# Write a brief paragraph: Where are you now, and where do you want to be?
Right now, I see myself as a Learner B, someone who uses AI to grow, not to take shortcuts. I try solving problems on my own first. The struggle helps me think more clearly. When I do turn to AI, I am not just looking for answers. I want to understand why something works and how I can make it better. I want to become someone who can confidently break big problems into smaller, manageable pieces and solve them step by step. I want to use AI in a smart way, not to replace my thinking, but to help me go deeper.


# Part B: The Wrong Way vs. The Right Way
## Track B — The Right Way (DO THIS):
### Step 1: Attempt independently
Pseudocode for a palindrome

```py
FUNCTION check_palindrome(text)
    rev ← REVERSE(text)
    
    IF text == rev THEN
        PRINT "it is a palindrome"
    ELSE
        PRINT "not a palindrome"
    END IF
    
END FUNCTION
```

Implement your solution in Python
```py
def check_palindrome(text):     #declaring the palindrome function
    rev = ''.join(reversed(text))  # reversing the input from the user
    if text == rev :        # checking to see if reversed input is equal to original input
        print("it is a palindrome")  # if condition is met, then this will be the output
    else:
        print("not a palindrome")  #Output when condition is not met
```
## Test with examples: *racecar*, *hello*, *A man a plan a canal Panama*
```py
check_palindrome("racecar") # it is a palidrome
check_palindrome("hello") # not a palidrome
check_palindrome("A man a plan a canal Panama") # not a palidrome

```

# Step 2: Strategic AI use After you have a working solution, ask AI:
I wrote a function to check palindromes. Here's my approach: [paste YOUR code]
- What's the time complexity?
- What edge cases am I missing?
- Alternatives and trade-offs?
- How does it perform on very long strings?"

# Then use AI strategically:
Using to check for the following:

    What's the time complexity?

    What edge cases am I missing?

    Alternatives and trade-offs?

    How does it perform on very long strings?"

# 1. Time Complexity
    - Time complexity: O(n) linear time 

    - Space complexity: O(n) linear space


# 2. Edge Cases You're Missing
You handle well: spaces, mixed case, punctuation, numbers.
What's missing:
    - Empty string
    - Single character
    - Unicode/accented characters 
    - None input

# 3. Alternatives & Trade-offs
Two-pointer approach (most efficient)
```py
def check_palindrome(text):
    left, right = 0, len(text) - 1
    while left < right:
        if text[left] != text[right]:
            return False
        left += 1
        right -= 1
    return True
```
# 4. Performance on Very Long Strings

Your code is efficient but has two limitations on very long inputs:
> - No early exit — even if the first and last characters don't match, your code fully builds rev before doing any comparison. On a 10 million character string that fails on character 1, you've wasted the entire reversed() + join operation.
> - Memory usage — you hold text, cleaned_text, and rev in memory simultaneously, roughly 3x your input size. On a 100MB string that's ~300MB of memory.
> - The cleaned_text[::-1] vs ''.join(reversed()) — slice reversal is slightly faster in practice since it's a single C-level operation internally, though both are O(n)

**Recommendation for long strings:** use the two-pointer approach since it avoids creating any extra copies and can short-circuit early (stops as soon as a mismatch is found, rather than always processing the full string).

# Step 3: Reflection
## What did you learn by struggling first?
By struggling first, indirectly train me to be independent, I learned more about my code, why it did not work the first time, learned how to debug and learned how to explain my code with confidence.
## How is your understanding different than if you'd just asked for the solution?
My understanding has increased exponetially, i am better and more knwoledgeable than before. I learned how to solve it myself, how to break problems into smaller concepts and I also learned how to read documentation. Implementing the palindrome first without consulting AI helped me to understand string manipulation like reverse string. I can use join() and reversed() functions which I did not know before

## Can you now implement similar functions (reverse a string, find duplicates) without AI?
Yes, I can now implement similar functions without the help of AI.

## What mental model did you build?
The mental model I built is this: I take my time to truly understand the problem. Work through it step by step. Think about the edge cases that could cause issues. And once it works, ask myself if it’s the best and most efficient way to do it

# Part C: Testing Your Understanding
Without using AI, complete these variations:
    Ignore spaces and punctuation
    Make it case-insensitive
    Return the position where the string stops being a palindrome (if not a palindrome)

```py
def check_palindrome(text):
    cleaned_text ="".join(letter.lower() for letter in text if letter.isalnum()) # cleaning the input

    rev = ''.join(reversed(cleaned_text)) # reversing the string
    
    if cleaned_text == rev: # checking to see if there is a match
        return True

    for i in range(len(cleaned_text)):  # loop through each character by index
        if cleaned_text[i] != rev[i]:  # Checks if character at index i is different from the reversed.
            return False, i 
```
# Part D: The Fairness Contract

## I will use AI when:
    -  Have tried to solve the problem for at least 15 to 20 minutes.
    -  To get deeper understanding why my solution works and why it doesn’t.
    -  To examine other alternative after my working solution.
    -  To amplify my learning.

## I will NOT use AI when:
    - I am still struggling with fundamentals concepts.
    - I am writing an exam or assessment.
    - I am doing code review.
    - Instructions clearly says do not use AI.

## I know I'm using AI fairly when:
    - I can write similar code without AI's help.
    - I can explain the logic of my code without AI.
    - I am more confident in my own understanding.
    - I don't over rely on AI.


Sign and date this contract.

![my_signature](https://i.ibb.co/ycRvT6gM/mike-sign.jpg)

20th February,2026


# Part E: Real-World Scenario Analysis
### Interview: "Explain how you'd implement a caching system." If you always relied on AI, can you answer?
    If i had always relied on AI to design systems, i wouldn't be able to answer the question.

### Production bug at 2 AM: AI is unavailable. Can you debug code you don't fully understand?
    No, i can't debug the code because it was generated by AI and i do not understand the code.

### New tech with little documentation: If you never learned to read docs and experiment, what happens?
    It will be very difficult for me to use the software product, and i might likely fall back to using AI as shortcut for my inability to learn how to read documentation

## Write a paragraph: How does using AI fairly now prepare you for these scenarios?
   Using AI fairly today prepares me for the real world challenges because I use it responsibly, I don't rely on it, rather use it to amplify my learning. Therefore in the real world i will be more equipped to handle any challenges.


# Part F: Building Irreplaceable Skills

| Skill                   | Description                              | Rating | Improvement Plan      |
| :----------------       | :---------                               | ----:  | -----------------     |
| Problem Decomposition   | Breaking down problems logically         | 4/5    |Learn critical thinking|
| System Thinking         | Understanding how components interact    | 3/5    | Seek to understand    |
| Critical Evaluation     | Knowing when code is wrong or inefficient| 4/5    | More Code review      |
| Debugging Mindset       | Investigating unexpected behavior        | 3/5    | Study more            |
| Conceptual Understanding| Knowing WHY, not just HOW                | 2/5    | Ask WHY,all the time  |



# My lowest-rated skill 3 specific actions i will take this week to improve it without relying on AI.

My lowest rated skill is *Conceptual Understanding*
1. I will spend more time to read documentation and to understand why my code works
2. I will seek help from my peers
3. I will write more code without using AI
