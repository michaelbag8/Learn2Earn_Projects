# Exercise 3: Multi-Step Question Answering

## Select a multi-step question (e.g., "What is 25% of the sum of 40 and 60?").

**My multi-step question:** After collecting eggs from his chickens, Dale puts the eggs into cartons to sell. Dale fills ‍15 cartons and has ‍7 eggs left over. Each carton holds ‍12 eggs. How many eggs did Dale collect?

### Create a prompt that asks the model to:
- Restate the problem
- Break it into smaller steps
- Solve each step before giving the final answer

### **Prompt:** Restate the problem, break it into smaller steps and solve each step before giving the final answer : After collecting eggs from his chickens, Dale puts the eggs into cartons to sell. Dale fills ‍15 cartons and has ‍7 eggs left over. Each carton holds ‍12 eggs. How many eggs did Dale collect?

**AI Model:**
Let's carefully solve this step by step.

---

### **Step 1: Restate the problem**

We are asked:

> Dale fills 15 cartons with eggs and has 7 eggs left over. Each carton holds 12 eggs. How many eggs did he collect in total?

We need to **find the total number of eggs** Dale collected.

---

### **Step 2: Break it into smaller steps**

1. **Calculate the number of eggs in the full cartons**  
   Each carton holds 12 eggs, and Dale filled 15 cartons.

   \[
   \text{Eggs in cartons} = 15 \times 12
   \]

2. **Add the leftover eggs**  
   Dale has 7 eggs left over.

   \[
   \text{Total eggs} = (\text{Eggs in cartons}) + (\text{Leftover eggs})
   \]

---

### **Step 3: Solve each step**

1. Eggs in cartons:

\[
15 \times 12 = 180
\]

2. Total eggs:

\[
180 + 7 = 187
\]

---

### **✅ Step 4: Final Answer**

Dale collected **187 eggs**.

---

## Review the reasoning process and check the accuracy of the final solution

The reasoning is sound, following a clear process of restating the problem, breaking it into smaller steps, solving, and arriving at the final answer; the arithmetic is correct, making the final answer of 187 eggs accurate.

