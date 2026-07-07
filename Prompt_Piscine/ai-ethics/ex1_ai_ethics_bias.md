# Understanding AI Bias

# Part A: Build Your Foundation (no AI first)

## Find or simulate a short example of model bias (e.g., gender, nationality, profession).

A real example of gender bias in AI systems that I found was published in a book by Hyzyk and Misanchuk (2024). is shown below.

![ChatGPT](https://i.ibb.co/2bbBxbc/Bias-07.png)

Despite the scenario not specifying gender, ChatGPT’s bias leads it to assume the doctor is male and dismiss the possibility that the nurse could be male.


# Explain in your own words why that bias may exist.
This gender bias exists because of the data that was used to train the model and how language reflects society. It exists because of what ChatGPT or LLMs learn from textbook, websites , articles etc


## Propose one mitigation (data balancing, prompt phrasing, model choice).
To mitigate this bias AI should be trained with the right data, and prompt phrasing should be gender specific

# Part B: Strategic AI Use

## Ask AI to deepen your understanding: 

"I found this bias example: [your example]. I think it exists because: [your explanation]. My proposed mitigation: [your strategy]. What am I missing? What underlying causes haven't I considered?"

You've got the right idea, but here's what you're missing:

**Underlying causes beyond training data:**
- **Ambiguity resolution**: LLMs resolve unclear pronouns using statistical patterns, not reasoning
- **Language structure itself**: Even balanced data contains embedded societal biases in how we talk about professions
- **Intersectionality**: Race, class, and age biases compound gender bias
- **Who chooses the "right" data**: Lack of diverse AI teams means blind spots in what gets fixed

**Better mitigations than "use right data":**
- Post-training techniques (RLHF, constitutional AI)
- Flag ambiguity instead of assuming
- System prompts that prevent stereotyping
- Continuous monitoring and user feedback
- Diverse development teams

**Key insight**: Your solution puts the burden on users to phrase prompts carefully. Better approach: make the model handle ambiguity properly rather than defaulting to stereotypes.

The tricky part: Perfect debiasing is impossible. Models need to balance reflecting reality with avoiding harmful assumptions.

# Then explore edge cases:
> "What happens when this bias appears in healthcare, hiring, or criminal justice? Here are my predictions: ***unequal treatment under the law, in hiring one gender or ethnicity, getting hired over the other, treatment delayed for some gender or ethnicity.**
What are the real-world implications?

# AI Reponses
**Bias in healthcare, hiring, and criminal justice leads to unequal treatment, delayed or denied opportunities, unfair outcomes, and long-term systemic inequality, eroding trust, efficiency, and social mobility.**

# Part C: Critical Reflection

### What % did you complete before using AI?
i completed about 80% before using AI

### Did AI replace your thinking or amplify it?
AI didn’t replace my thinking, it rather amplify it

### Could you explain this to someone else without AI?
Yes, i can confidently explain it without AI

### What did you contribute that AI couldn't?
What I really contributed was my own thinking first. Before I ever asked AI anything, I took time to understand the example, form my own opinion, and come up with my own explanation. I used my personal judgment to analyze the bias and think about possible solutions.




# Reference:
Hyzyk, J., & Misanchuk, M. (2024). ChatGPT in STEM teaching: An introduction to using LLM-based tools in higher ed. eCampusOntario. https://ecampusontario.pressbooks.pub/llmtoolsforstemteachinginhighered/