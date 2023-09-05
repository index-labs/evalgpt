<div align="center">

# Eval GPT

</div>

## What is Eval GPT

The code interpreter tool invokes the LLM to generate Python code. The LLM takes inputs, uses its deep learning
capabilities to generate corresponding Python code. This generated code is then executed by the code interpreter,
returning the execution results.

## Benefits

-

## How it works

1. It utilizes the Language Learning Model (LLM) to transform user inputs into Python code. This is achieved through the
   LLM's deep learning capabilities.
2. The generated Python code is then executed by the code interpreter, which returns the results of the execution.

```bash
$ evalgpt -q "calculate MD5 of the string 'hello, world'" 
The MD5 hash of "hello, world" is: e4d7f1b4ed2e42d15898f4b27b019da4
```

And another example:

```bash
$ evalgpt -q "the title of the link: https://arxiv.org/abs/2302.04761" 
[2302.04761] Toolformer: Language Models Can Teach Themselves to Use Tools
```

## Architecture

![](./architecture.png)


## Quick Start 🚀

### Install `evalgpt`

You can install evalgpt using the following command:

```bash
go install github.com/index-labs/evalgpt@latest
```

You could verify the installation by running the following command:

```bash
evalgpt -h
```

### Configuration

After you install evalgpt command line, before execute it, you must config below options:

**Configure Openai API Key**

```bash
export OPENAI_API_KEY=sk_******
```

also, you can config openai api key by command args, but it's not recommend:

```bash
evalgpt --openai-api-key sk_***** -q <query>

```

**Configure Python Interpreter**

By default, the code interpreter uses the system's Python interpreter. However, you can create a completely new Python
interpreter using Python's virtual environment tools and configure it accordingly.

```bash
python3 -m venv /path/evalgpt/venv
# install third python libraries
/path/evalgpt/venv/bin/pip3 install -r requirements.txt

# config python interpreter
export PYTHON_INTERPRETER=/path/evalgpt/venv/bin/python3
```

or

```bash
evalgpt --python-interpreter /path/evalgpt/venv/bin/python3 -q <query>
```

**Note:**

Before tackling complex tasks, ensure to install necessary Python third-party libraries. This equips your code
interpreter to handle corresponding tasks, boosting efficiency and ensuring smooth operation.

### Usage

**Help**

```bash
> evalgpt -h
NAME:
   evalgpt help - A new cli application

USAGE:
   evalgpt help [global options] command [command options] [arguments...]

DESCRIPTION:
   description

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --openai-api-key value         Openai Api Key, if you use open ai model gpt3 or gpt4, you must set this flag [$OPENAI_API_KEY]
   --model value                  LLM name (default: "gpt-4-0613") [$MODEL]
   --python-interpreter value     python interpreter path (default: "/usr/bin/python3") [$PYTHON_INTERPRETER]
   --verbose, -v                  print verbose log (default: false) [$VERBOSE]
   --query value, -q value        what you want to ask
   --file value [ --file value ]  the path to the file to be parsed and processed, eg. --file /tmp/a.txt --file /tmp/b.txt
   --help, -h                     show help
```

**Note:**

Remember to configure the OpenAI API key and Python interpreter before executing the code interpreter, The following
examples have already been configured with environment variables for the OpenAI API key and the Python interpreter.

**Simple Query**

Get the public IP address of the machine:

```bash
❯ evalgpt -q 'get the public IP of my computer'
Your public IP is: 104.28.240.133
```

Calculate the sha256 hash of a string:

```bash
❯ ./bin/evalgpt -q 'calculate the sha256 of the "hello,world"'
77df263f49123356d28a4a8715d25bf5b980beeeb503cab46ea61ac9f3320eda
```

**Pipeline**

You can user pipeline to input context data and query on it:

```bash
> cat a.csv
date,dau
2023-08-20,1000
2023-08-21,900
2023-08-22,1100
2023-08-23,2000
2023-08-24,1800

> cat a.csv | evalgpt -q 'calculate the average dau'
Average DAU:  1360.0
```

**Interact with files**

convert png file to webp file:

```bash
> ls
a.png

> evalgpt -q 'convert this png file to webp' --file ./a.png
created file: a.webp

> ls
a.png a.webp
```

