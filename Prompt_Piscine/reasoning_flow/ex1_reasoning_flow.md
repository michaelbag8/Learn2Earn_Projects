# Exercise 1: Logic Puzzle with Explanation

## Write a logic puzzle (e.g., "Alice is older than Bob, Bob is older than Carol. Who is the youngest?"). 

## **Prompt:** Explain your reasoning before giving the final answer: There are two ducks in front of a duck, two ducks behind a duck and a duck in the middle. How many ducks are there?
## **AI Model:**
Let’s carefully reason through this step by step. The riddle is tricky because it plays with overlapping positions rather than being completely separate ducks.
We are told:

1. **Two ducks in front of a duck**
2. **Two ducks behind a duck**
3. **A duck in the middle**

---

### Step 1: Think in terms of positions

Let’s label the ducks as positions in a line from front to back:

* Duck 1 → front
* Duck 2 → middle
* Duck 3 → back

---

### Step 2: Check the statements

1. **Two ducks in front of a duck**

   * If the “middle duck” is Duck 2, then Duck 1 is in front of Duck 2. But we need *two ducks* in front of a duck. So maybe Duck 2 is the middle of the “two in front of a duck”?

   Let’s hold that thought.

2. **Two ducks behind a duck**

   * If Duck 2 is the middle, then Duck 3 is behind Duck 2. Again, two behind?

---

### Step 3: Consider overlapping

The trick is that some ducks serve double duty. Imagine **three ducks in a row**:

1. Duck A
2. Duck B
3. Duck C

* For **Duck B** (middle duck):

  * There’s **one duck in front** (Duck A)
  * There’s **one duck behind** (Duck C)

* For **Duck A** (front duck):

  * There are **two ducks behind** (B and C) ✅

* For **Duck C** (back duck):

  * There are **two ducks in front** (A and B) ✅

* Duck B is literally the **duck in the middle** ✅

---

### Step 4: Count ducks

* Duck A
* Duck B
* Duck C

So that’s **3 ducks in total**.

---

✅ **Answer:**

**3 ducks**

---

## Verify that the model provides step-by-step reasoning followed by the correct answer
The AI model was verified that step by step reasoning was given before it was followed by the  correct answer