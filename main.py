import os
import re
import json
import torch
import random
import requests
import argparse
import argparse
import numpy as np
from glob import glob
from transformers import AutoTokenizer, AutoModelForCausalLM
from transformers.generation.stopping_criteria import StoppingCriteria, StoppingCriteriaList


def is_go_identifier_part(s, current_identifier):
    pattern_start = re.compile(r'^[a-zA-Z_][a-zA-Z0-9_]*$')
    pattern_inner = re.compile(r'^[a-zA-Z_0-9][a-zA-Z_0-9]*$')
    if current_identifier == [] and pattern_start.match(s) or current_identifier != [] and pattern_inner.match(s):
        return True
    else:
        return False
    

class FunctionStoppingCriteria(StoppingCriteria):
    def __init__(self, tokenizer):
        self.tokenizer = tokenizer
        self.output_token_length = 0
        self.current_token_str = None
        self.previous_token_str = None
        self.generated_identifiers = []
        self.current_identifier = []

    def __call__(self, input_ids, scores, **kwargs):
        # Get the last two tokens
        output = self.tokenizer.decode(input_ids[0])
        self.current_token_str = output[-1]
        self.previous_token_str = output[-2]
        self.output_token_length += 1
        # Check if the last token is '}' and the previous token is '\n' or output token length exceeds 512
        if (self.current_token_str == '}' and self.previous_token_str == '\n') or self.output_token_length >= 512:
            return True
            
        if not is_go_identifier_part(self.current_token_str, self.current_identifier):
            for token_str in output[:-1][::-1]:
                if is_go_identifier_part(token_str, self.current_identifier):
                    self.current_identifier.append(token_str)
                else:
                    if self.current_identifier:
                        identifier = ''.join(self.current_identifier[::-1])
                        self.current_identifier = []
                        if identifier not in self.generated_identifiers and identifier not in ["break", "case", "chan", "const", "continue", "default", "defer", "else", "fallthrough", "for", "func", "go", "goto", "if", "import", "interface", "map", "package", "range", "return", "select", "struct", "switch", "type", "var", "bool", "byte", "complex64", "complex128", "error", "float32", "float64", "int", "int8", "int16", "int32", "int64", "rune", "string", "uint", "uint8", "uint16", "uint32", "uint64", "uintptr", "true", "false", "iota", "nil", "append", "cap", "close", "complex", "copy", "delete", "imag", "len", "make", "new", "panic", "print", "println", "real", "recover", "t", "T", "errorf"]:
                            self.generated_identifiers.append(identifier)
                            return True
                    break
        return False   


def set_seed(seed: int):
    random.seed(seed)
    np.random.seed(seed)
    os.environ['PYHTONHASHSEED'] = str(seed)
    torch.manual_seed(seed)
    torch.cuda.manual_seed(seed)
    torch.backends.cudnn.deterministic = True


def get_row_and_column(input_string, identifier):
    lines = input_string.strip().split('\n')
    row = len(lines) - 1
    last_line = lines[-1] if lines else None
    column = last_line.rfind(identifier) + len(identifier)
    return row, column


def init_gopls(workspace):
    url = "http://localhost:2000/init"
    data = {
        "workspace": workspace,
    }

    headers = {
        "Content-Type": "application/json"
    }

    response = requests.post(url, data=json.dumps(data), headers=headers)

    if response.status_code == 200:
        return
    else:
        raise Exception(f"Error: {response.status_code} - {response.text}")


def send_completion_request(row, column, file_path):
    url = "http://localhost:2000/definition"
    data = {
        "row": row,
        "column": column,
        "path": file_path
    }

    headers = {
        "Content-Type": "application/json"
    }

    response = requests.post(url, data=json.dumps(data), headers=headers)

    if response.status_code == 200:
        data = json.loads(response.json())
        return data['contents']['value'].replace('```go\n', '').strip('```').replace('\n```', '\n').replace('\n\n', '\n')
    else:
        raise Exception(f"Error: {response.status_code} - {response.text}")


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument('--model', help='model to use for unit test generation. should be one of [CodeLlama, DeepSeek-Coder, Magicoder]', required=True, type=str)
    parser.add_argument("--seed", type=int, default=42)
    parser.add_argument("--temperature", type=float, default=1.0)
    args = parser.parse_args()  

    set_seed(args.seed)

    if args.model == "CodeLlama":
        tokenizer = AutoTokenizer.from_pretrained('codellama/CodeLlama-7b-Instruct-hf', use_fast=True)
        model = AutoModelForCausalLM.from_pretrained('codellama/CodeLlama-13b-Instruct-hf', trust_remote_code=True, torch_dtype=torch.bfloat16, device_map='auto')
    elif args.model == "DeepSeek-Coder":
        tokenizer = AutoTokenizer.from_pretrained("deepseek-ai/deepseek-coder-6.7b-instruct", use_fast=True)
        model = AutoModelForCausalLM.from_pretrained("deepseek-ai/deepseek-coder-6.7b-instruct", trust_remote_code=True, torch_dtype=torch.bfloat16, device_map='auto')
    else:
        tokenizer = AutoTokenizer.from_pretrained("ise-uiuc/Magicoder-S-DS-6.7B", use_fast=True)
        model = AutoModelForCausalLM.from_pretrained("ise-uiuc/Magicoder-S-DS-6.7B", trust_remote_code=True, torch_dtype=torch.bfloat16, device_map='auto')

    last_workspace = None
    for project in os.listdir("./datasets"):
        funcnames = dict()
        for source in glob(f"./datasets/{project}/*"):
            for file_name in os.listdir(source):
                with open(source + "/" + file_name, 'r') as file:
                    item = json.load(file)
                    relpath = item["relpath"]

                    stopping_criteria = StoppingCriteriaList([FunctionStoppingCriteria(tokenizer)]) 

                    # get package
                    with open(relpath, 'r') as focal_file:
                        for line in focal_file.readlines():
                            if line.startswith("package "):
                                package = line.replace("package ", "").strip()
                                break

                    function = item["funcbody"]
                    funcname = item["funcname"]
                    stopping_criteria[0].generated_identifiers.append(funcname)
                    stopping_criteria[0].generated_identifiers.append(project)

                    prefix = "I will give you a Golang method or function, please generate a Golang unit test.\n\n"
                    suffix = "### METHOD (FUNCTION) UNDER TEST\n{}\n\n### GENERATED TEST SNIPPET\nfunc Test{}(t *testing.T) {{\n".format(function, funcname)
                    if funcname.lower() not in funcnames.keys():
                        funcnames[funcname.lower()] = 0
                    funcnames[funcname.lower()] += 1
                    workspace = relpath.rsplit('/', 1)[0]
                    
                    target_file_path = workspace + "/" + f"RATester_{args.model}_{funcname}_{funcnames[funcname.lower()]}_test.go"

                    if last_workspace != workspace:
                        last_workspace = workspace
                        init_gopls(workspace)
                    
                    current_identifier = []
                    idx = 0
                    row = item["lineno"] - 1
                    context_num = 1

                    for token_str in (item["focalfunc"] + " {"):
                        if is_go_identifier_part(token_str):
                            current_identifier.append(token_str)
                        else:
                            if current_identifier:
                                identifier = ''.join(current_identifier)
                                current_identifier = []
                                if identifier not in ["break", "case", "chan", "const", "continue", "default", "defer", "else", "fallthrough", "for", "func", "go", "goto", "if", "import", "interface", "map", "package", "range", "return", "select", "struct", "switch", "type", "var", "bool", "byte", "complex64", "complex128", "error", "float32", "float64", "int", "int8", "int16", "int32", "int64", "rune", "string", "uint", "uint8", "uint16", "uint32", "uint64", "uintptr", "true", "false", "iota", "nil", "append", "cap", "close", "complex", "copy", "delete", "imag", "len", "make", "new", "panic", "print", "println", "real", "recover"] and identifier != funcname and identifier not in stopping_criteria[0].generated_identifiers:
                                    column = idx
                                    definition = send_completion_request(row, column, relpath)
                                    stopping_criteria[0].generated_identifiers.append(identifier)
                                    if definition.strip() != "" and not definition.startswith("var") and not definition.startswith("package"):
                                        prefix += "### CONTEXT_{}\n{}\n\n".format(context_num, definition.strip())
                                        context_num += 1
                            if token_str == "\n":
                                row += 1
                                idx = -1
                        idx += 1

                    with open(target_file_path, 'w') as file:
                        print(target_file_path)
                        test_class = "package {}\nfunc Test{}_{}(t *testing.T) {{\n".format(package, funcname[0].upper() + funcname[1:], funcnames[funcname.lower()])
                        file.write(test_class)
                        while True:
                            prompt = prefix + suffix
                            input = tokenizer.encode(prompt, return_tensors="pt").to(model.device)
                            output = model.generate(
                                input,
                                max_new_tokens=512,
                                temperature=args.temperature,
                                eos_token_id=tokenizer.eos_token_id,
                                pad_token_id=tokenizer.eos_token_id,
                                stopping_criteria=stopping_criteria
                            )
                            output = tokenizer.decode(output[0], skip_special_tokens=True)
                            output = output[len(prompt): ]
                            suffix += output
                            test_class += output
                            file.write(output) 
                            file.flush()
                            if "\n}" in test_class or stopping_criteria[0].output_token_length >= 512:
                                break
                            identifier = stopping_criteria[0].generated_identifiers[-1]
                            row, column = get_row_and_column(test_class, identifier)
                            definition = send_completion_request(row, column, target_file_path)
                            if definition != "" and "invalid type" not in definition and not definition.startswith("var") and not definition.startswith("package"):
                                prefix += "### CONTEXT_{}\n{}\n\n".format(context_num, definition.strip())
                                context_num += 1


if __name__ == "__main__":
    main()
