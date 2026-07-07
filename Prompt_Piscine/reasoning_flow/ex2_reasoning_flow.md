# Exercise 2: Step-by-Step Reasoning Prompt

## Choose a math or coding problem (e.g., "Calculate the factorial of 5").
### My Math problem:
Find the number of permutations and combinations if n = 12 and r = 2.

## **Prompt:** Let's think step by step before answering.Find the number of permutations and combinations if n = 12 and r = 2."

## **AI Model:**
## Permutations and Combinations for n = 12, r = 2

We are asked to find the **number of permutations** and **combinations** when n = 12  and r = 2 .



## Step 1: Formulas

1. **Permutation** P(n, r) — order matters:  


$$
nP_r = \frac{n!}{(n-r)!}
$$




2. **Combination** C(n, r)  — order does **not** matter:


$$
nC_r = \frac{n!}{r!(n-r)!}
$$




## Step 2: Permutations

### Permutation Calculation

$$
P(12, 2) = \frac{12!}{(12-2)!} = \frac{12!}{10!}
$$

Since 

$$
12! = 12 \times 11 \times 10!
$$

we have:

$$
P(12, 2) = \frac{12 \times 11 \times 10!}{10!} = 12 \times 11 = 132
$$

✅ **Number of permutations = 132**



## Step 3: Combinations
### Combination Calculation

$$
C(12, 2) = \frac{12!}{2! \cdot (12-2)!} = \frac{12!}{2! \cdot 10!}
$$

$$
C(12, 2) = \frac{12 \times 11 \times 10!}{2 \times 1 \times 10!} = \frac{12 \times 11}{2} = 66
$$
✅ **Number of combinations = 66**

---

## Step 4: Summary

| Type            | Formula                          | Result |
|-----------------|--------------------------------- |--------|
| Permutations    | P(12, 2) = 12.11                 | 132    |
| Combinations    | C(12, 2) = 12.11/2               | 66     |

# Without reasoning instruction

## **Prompt:**"Find the number of permutations and combinations if n = 12 and r = 2."

### 📌 Given:

* ( n = 12 )
* ( r = 2 )

* **Permutations:** **132**
* **Combinations:** **66**

# Compare outputs with and without the reasoning instruction.

Without the reasoning instruction, the output is just the final numbers, permutations = 132 and combinations = 66. It is fast and short but doesn’t show the formulas or steps, making it harder to verify or learn from.
With the reasoning instruction, the output breaks the problem into steps, shows the formulas, and explains how the numbers are calculated. This is longer but easier to verify, reduces errors, and has higher educational value.The reasoning adds clarity and reliability at the cost of brevity.