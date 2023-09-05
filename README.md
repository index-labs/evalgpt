<div align="center">

# EvalGPT

</div>

## What is EvalGPT

🧩 This project is still in the early stages of development, and we are actively working on it. If you have any questions or suggestions, please submit an issue or PR.

EvalGPT is an code interpreter framework, leveraging the power of large language models such as GPT-4, CodeLlama, and Claude 2. This powerful tool allows users to write tasks, and EvalGPT will assist in writing the code, executing it, and delivering the results.

![](images/architecture.png)

EvalGPT's architecture draws inspiration from [Google's Borg system](https://research.google/pubs/pub43438/). It includes a master node, known as EvalGPT, composed of three components: planning, scheduler, and memory.

When EvalGPT receives a request, it starts planning the task using a Large Language Model (LLM), dividing larger tasks into smaller, manageable ones. For each sub-task, EvalGPT will spawn a new node known as an EvalAgent.

Each EvalAgent is responsible for generating the code based on the assigned small task. Once the code is generated, the EvalAgent initiates a runtime to execute the code, even harnessing external tools when necessary. The results are then collected by the EvalAgent.

EvalAgent nodes can access the memory from the EvalGPT master node, allowing for efficient and effective communication. If an EvalAgent encounters any errors during the process, it reports the error to the EvalGPT master node, which then replans the task to avoid the error.

Finally, the EvalGPT master node collates all results from the EvalAgent nodes and generates the final answer for the request.

## Benefits

1. **Automated Code Writing**: EvalGPT leverages advanced language models to auto-generate code, reducing manual effort and increasing productivity.
2. **Efficient Task Execution**: By breaking down complex tasks into manageable sub-tasks, EvalGPT ensures efficient and parallel execution, speeding up the overall process.
3. **Robust Error Handling**: With its ability to replan tasks in case of errors, EvalGPT ensures reliable operation and accurate results.
4. **Scalability**: EvalGPT is built to handle tasks of varying complexity, making it a scalable solution for a wide range of coding needs.
5. **Resource Optimization**: Inspired by Google Borg's resource management, EvalGPT optimally utilizes computational resources, leading to improved performance.
6. **Extensibility**: With the ability to incorporate external tools into its runtime, EvalGPT is highly adaptable and can be extended to handle a diverse range of tasks.

## Demo

https://github.com/index-labs/evalgpt/assets/7857126/73417c1f-8866-47fb-951a-7fd03c9dbf41

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

### Build it from source code

```bash
git clone https://github.com/index-labs/evalgpt.git

cd evalgpt

go mod tidy && go mod vendor

mkdir -p ./bin

go build -o ./bin/evalgpt ./*.go

./bin/evalgpt -h
```

Then you can find it on bin directory.

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

## Usage

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
❯ evalgpt -q 'calculate the sha256 of the "hello,world"'
77df263f49123356d28a4a8715d25bf5b980beeeb503cab46ea61ac9f3320eda
```

Get the title of a website:

```bash
❯ evalgpt -q "get the title of a website: https://arxiv.org/abs/2302.04761" -v
[2302.04761] Toolformer: Language Models Can Teach Themselves to Use Tools
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

Draw a line graph based on the data from the CSV

```bash
> cat a.csv
date,dau
2023-08-20,1000
2023-08-21,900
2023-08-22,1100
2023-08-23,2000
2023-08-24,1800

> evalgpt -q 'draw a line graph based on the data from the CSV' --file ./a.csv
```

output:

![](images/example_dau.png)

## Architecture Details

### EvalGPT Master Node

The EvalGPT master node serves as the control center of the framework. It houses three critical components: planning, scheduler, and memory.

The planning component leverages large language models to plan tasks based on the user's request. It breaks down complex tasks into smaller, manageable sub-tasks, each of which is handled by an individual EvalAgent node.

The scheduler component is responsible for task distribution. It assigns each sub-task to an EvalAgent node, ensuring efficient utilization of resources and parallel execution of tasks for optimal performance.

The memory component serves as the shared memory space for all EvalAgent nodes. It stores the results of executed tasks and provides a platform for data exchange between different nodes. This shared memory model facilitates complex computations and aids in error handling by allowing for task replanning in case of errors.

In the event of an error during code execution, the master node replans the task to avoid the error, thereby ensuring robust and reliable operation.

Finally, the EvalGPT master node collects the results from all EvalAgent nodes, compiles them, and generates the final answer for the user's request. This centralized control and coordination make the EvalGPT master node a crucial part of the EvalGPT framework.

### EvalAgent Node

EvalAgent nodes are the workhorses of the EvalGPT framework. Spawned by the master node for each sub-task, they're responsible for code generation, execution, and result collection.

The code generation process in an EvalAgent node is guided by the specific task it's assigned. Using the large language model, it produces the necessary code to accomplish the task, ensuring it's suited to the task's requirements and complexity.

Once the code is generated, the EvalAgent node initiates a runtime environment to execute the code. This runtime is flexible, capable of incorporating external tools as needed, and provides a robust platform for code execution.

During execution, the EvalAgent node collects the results and can access the shared memory from the EvalGPT master node. This allows for efficient data exchange and facilitates complex computations requiring significant data manipulation or access to previously computed results.

In case of any errors during code execution, the EvalAgent node reports these back to the EvalGPT master node. The master node then replans the task to avoid the error, ensuring a robust and reliable operation.

In essence, EvalAgent nodes are autonomous units within the EvalGPT framework, capable of generating and executing code, handling errors, and communicating results efficiently.

### Runtime

The runtime of EvalGPT is managed by EvalAgent nodes. Each EvalAgent node generates code for a specific task and initiates a runtime to execute the code. The runtime environment is flexible and can incorporate external tools as necessary, providing a highly adaptable execution context.

The runtime also includes error handling mechanisms. If an EvalAgent node encounters any errors during code execution, it reports these back to the EvalGPT master node. The master node then replans the task to avoid the error, ensuring robust and reliable code execution.

The runtime can interact with the EvalGPT master node's memory, enabling efficient data exchange and facilitating complex computations. This shared memory model allows for the execution of tasks that require significant data manipulation or access to previously computed results.
