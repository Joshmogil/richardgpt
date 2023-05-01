Stealing richard's code ...
To use OpenAI's API for generating an infinite book, you need to install the `openai` package first:

```sh
pip install openai
```

Here's a Python script that generates text using OpenAI's GPT-3 API:

```python
import openai
import os

# Load your OpenAI API key from an environment variable or secret management service
openai.api_key = os.getenv("OPENAI_API_KEY")

# Function to generate text using GPT-3 API
def generate_text(prompt, model='text-davinci-002', token_limit=2048):
    response = openai.Completion.create(
        engine=model,
        prompt=prompt,
        max_tokens=token_limit,
        n=1,
        stop=None,
        temperature=0.7
    )
    return response.choices[0].text.strip()

# Infinite book generator loop
while True:
    book_segment_prompt = "<Your segment starting prompt here>"
    print(generate_text(book_segment_prompt))
```

Replace `<Your segment starting prompt here>` with the desired input to start generating your infinite book. The API returns a text segment based on the given prompt. The loop continues to generate new text segments until you stop the script.

Replace `text-davinci-002` model as needed. Models supported by OpenAI API are: `text-davinci-002`, `text-curie-002`, `text-babbage-002`, `text-ada-002`.
