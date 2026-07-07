# Exercise 1: Reducing False or Invented Facts

Write a prompt that produces errors (example: “Give the capitals of five imaginary countries.”).

**Prompt:**  
Give me the square root of –1 as a real number?

**AI Model:**  
There is no real number whose square is –1. Every real number squared yields a non-negative result, so 
$$
(-1 \cdot \sqrt{-1})
$$ 
doesn't exist in the real numbers. Real numbers range from negative infinity to positive infinity, but squaring any number (positive, negative, or zero) always gives 0 or positive—no real solution fits. In complex numbers, it exists, but you specified real numbers.

---

### Rewrite the prompt so it only accepts correct data (example: “List five real countries and their capitals.”).

**Prompt:**  
Give me the square root of 1 as a real number?

**AI Model:**  
The square root of 1 as a real number is simply:

$$
\sqrt{1} = 1
$$

**Reasoning:**  
A square root is a number which, when multiplied by itself, gives the original number. So:

$$
1 \times 1 = 1
$$

Technically, -1 is also a square root of 1, because:

$$
(-1) \times (-1) = 1
$$

But if we specify the principal square root (the positive real number), it is:

$$
\boxed{1}
$$

---

**Compare both results and note the differences:**  
Based on the output, the square root of –1 as a real number cannot exist in the real number system, while the square root of 1 exists entirely in the real number system, with two real solutions.