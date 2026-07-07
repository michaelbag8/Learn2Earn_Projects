# Modified code based on these creteria
Ignore spaces and punctuation.

Be case-insensitive.

Return the position where the string stops being a palindrome (if not one)

```py
def check_palindrome(text):
    cleaned_text = ""  # start with empty string
    for letter in text:  # loop through each character in text
        if letter.isalnum():  # check if the character is alphanumeric
            cleaned_text += letter.lower()  # convert to lowercase and append

    rev = "".join(reversed(cleaned_text))  # reverse the string and join

    if cleaned_text == rev: # checking to see if there is a match
        return True

    for i in range(len(cleaned_text)):  # loop through each character by index
        if cleaned_text[i] != rev[i]:  # Checks if character at index i is different from the reversed.
            return False, i     # return mismatch index and false

```
# Testing
```py
print(check_palindrome("racecar")) # True
print(check_palindrome("hello"))# False, 0
print(check_palindrome("A man a plan a canal Panama")) # True 
print(check_palindrome("noonpoon")) # False, 3

```
# Asking AI go through my  modified  code by answering these quetions 
1. Did I miss anything? 
2. Can it be more efficient?"


## This is the code genertated by AI: AI Model
```py
def check_palindrome(word):
    left = 0
    right = len(word) - 1

    while left < right:
        # Skip non-alphanumeric characters
        while left < right and not word[left].isalnum():
            left += 1
        while left < right and not word[right].isalnum():
            right -= 1

        # Check characters (case-insensitive)
        if word[left].lower() != word[right].lower():
            return left  # Return the index where it fails

        left += 1
        right -= 1

    return -1  # Return -1 if it is a palindrome

```
> in summary i did miss to **on how the function handle empty string** and yes the code can be modified for more efficiency.

# Reflecting on what AI added that i didn't consider initially

According to AI my modified code was able to handle **case insensitivity** , **punctuation** and **spaces** but failed to handle when the function input is an empty string. This was a scenario I didn't consider.
