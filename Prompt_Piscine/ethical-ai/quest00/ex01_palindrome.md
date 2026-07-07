# Pseudocode

```py
DEFINE FUNCTION check_palindrome(text)

    SET rev TO the reverse of text

    IF text is equal to rev THEN
        DISPLAY "it is a palindrome"
    ELSE
        DISPLAY "not a palindrome"
```

# Code implementation in Python

```py
def check_palindrome(text):     #declaring the palindrome function
    rev = ''.join(reversed(text))  # reversing the input from the user, and join it back to string with "" separator 
    if text == rev :        # checking to see if reversed input is equal to original input
        print("it is a palindrome")  # if condition is met, then this will be the output
    else:
        print("not a palindrome")  #Output when condition is not met
```
# Testing
```py
check_palindrome("racecar") # it is a palidrome
check_palindrome("hello") # not a palidrome
check_palindrome("A man a plan a canal Panama") # not a palidrome
```

# Using AI to go deeper

What's the time complexity?

What edge cases might I miss?

Are there better approaches?

# Summary of the answers to the questions above by Ai

> ✅ Your solution is correct

> ⏱ Time complexity: O(n) linear time , if input double the number of operations also double

> 💾 Space complexity: O(n) Linear space

> ⚠️ Edge cases: case, spaces, punctuation, empty string

> Yes there are better approaches that are more cleaner and more effiecient.


# Reflection

- I learned how to solve it myself, how to break problems into smaller concepts and I also learned how to read documentation. Implementing the palindrome first without consulting AI helped me to understand string manipulation like reverse string. I can use join() and reversed() functions which I did not know before.

- My understanding has increased exponetially, i am better and more knwoledgeable than before.

- Yes, i can now write similar function e.g, reverse a string without the help of Ai