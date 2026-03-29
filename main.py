import os
import re
import json
import random
import argparse
from glob import glob

try:
    from transformers.generation.stopping_criteria import StoppingCriteria
except ModuleNotFoundError:
    class StoppingCriteria:  # type: ignore[no-redef]
        pass

LOCAL_MODELS = ["CodeLlama", "DeepSeek-Coder", "Magicoder"]
ARK_MODEL = "Doubao-Seed-Code"
ARK_BASE_URL = "https://ark-cn-beijing.bytedance.net/api/v3"
EXTRA_HEADERS = {'x-is-encrypted': 'true'}
GO_KEYWORDS = [
    "break", "case", "chan", "const", "continue", "default", "defer", "else", "fallthrough", "for",
    "func", "go", "goto", "if", "import", "interface", "map", "package", "range", "return", "select",
    "struct", "switch", "type", "var", "bool", "byte", "complex64", "complex128", "error", "float32",
    "float64", "int", "int8", "int16", "int32", "int64", "rune", "string", "uint", "uint8", "uint16",
    "uint32", "uint64", "uintptr", "true", "false", "iota", "nil", "append", "cap", "close", "complex",
    "copy", "delete", "imag", "len", "make", "new", "panic", "print", "println", "real", "recover",
    "t", "T", "errorf",
]


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
                        if identifier not in self.generated_identifiers and identifier not in GO_KEYWORDS:
                            self.generated_identifiers.append(identifier)
                            return True
                    break
        return False   


def set_seed(seed: int):
    import numpy as np

    random.seed(seed)
    np.random.seed(seed)
    os.environ['PYTHONHASHSEED'] = str(seed)
    try:
        import torch
    except ModuleNotFoundError:
        return

    torch.manual_seed(seed)
    if torch.cuda.is_available():
        torch.cuda.manual_seed(seed)
    if hasattr(torch.backends, "cudnn"):
        torch.backends.cudnn.deterministic = True


def get_row_and_column(input_string, identifier):
    lines = input_string.strip().split('\n')
    row = len(lines) - 1
    last_line = lines[-1] if lines else None
    column = last_line.rfind(identifier) + len(identifier)
    return row, column


def init_gopls(workspace):
    import requests

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
    import requests

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


def build_parser():
    parser = argparse.ArgumentParser()
    parser.add_argument(
        '--model',
        help=(
            'model to use for unit test generation. should be one of '
            f'[{", ".join(LOCAL_MODELS + [ARK_MODEL])}]'
        ),
        required=True,
        type=str,
    )
    parser.add_argument("--seed", type=int, default=42)
    parser.add_argument("--temperature", type=float, default=1.0)
    parser.add_argument(
        "--datasets-root",
        type=str,
        default="./datasets",
        help="dataset root directory; formal aligned runs should point to generated selected-package datasets",
    )
    parser.add_argument(
        "--project",
        type=str,
        default="",
        help="optional project name to restrict dataset traversal",
    )
    parser.add_argument(
        "--dataset-dir",
        action="append",
        default=[],
        help="optional dataset directory basename whitelist; can be passed multiple times",
    )
    parser.add_argument(
        "--max-items",
        type=int,
        default=0,
        help="optional upper bound for dataset items to process, 0 means no limit",
    )
    parser.add_argument(
        "--ark-endpoint-id",
        type=str,
        default=os.environ.get("ARK_ENDPOINT_ID", ""),
        help=f"Ark endpoint id for {ARK_MODEL}",
    )
    parser.add_argument(
        "--ark-api-key-env",
        type=str,
        default="ARK_API_KEY",
        help="environment variable name that stores the Ark API key",
    )
    return parser


def extract_new_identifier(text, generated_identifiers):
    current_identifier = []
    identifiers = list(generated_identifiers)
    for token_str in text:
        if is_go_identifier_part(token_str, current_identifier):
            current_identifier.append(token_str)
            continue
        if not current_identifier:
            continue
        identifier = ''.join(current_identifier)
        current_identifier = []
        if identifier not in identifiers and identifier not in GO_KEYWORDS:
            return identifier
    return None


def extract_prompt_identifiers(text, funcname):
    current_identifier = []
    identifiers = []
    seen = set()
    row = 0
    idx = 0
    for token_str in text:
        if is_go_identifier_part(token_str, current_identifier):
            current_identifier.append(token_str)
        else:
            if current_identifier:
                identifier = ''.join(current_identifier)
                current_identifier = []
                if identifier not in GO_KEYWORDS and identifier != funcname and identifier not in seen:
                    identifiers.append((identifier, row, idx))
                    seen.add(identifier)
            if token_str == "\n":
                row += 1
                idx = -1
        idx += 1
    if current_identifier:
        identifier = ''.join(current_identifier)
        if identifier not in GO_KEYWORDS and identifier != funcname and identifier not in seen:
            identifiers.append((identifier, row, idx))
    return identifiers


def load_local_model(model_name):
    import torch
    from transformers import AutoTokenizer, AutoModelForCausalLM

    if model_name == "CodeLlama":
        tokenizer = AutoTokenizer.from_pretrained('codellama/CodeLlama-7b-Instruct-hf', use_fast=True)
        model = AutoModelForCausalLM.from_pretrained(
            'codellama/CodeLlama-13b-Instruct-hf',
            trust_remote_code=True,
            torch_dtype=torch.bfloat16,
            device_map='auto',
        )
    elif model_name == "DeepSeek-Coder":
        tokenizer = AutoTokenizer.from_pretrained("deepseek-ai/deepseek-coder-6.7b-instruct", use_fast=True)
        model = AutoModelForCausalLM.from_pretrained(
            "deepseek-ai/deepseek-coder-6.7b-instruct",
            trust_remote_code=True,
            torch_dtype=torch.bfloat16,
            device_map='auto',
        )
    else:
        tokenizer = AutoTokenizer.from_pretrained("ise-uiuc/Magicoder-S-DS-6.7B", use_fast=True)
        model = AutoModelForCausalLM.from_pretrained(
            "ise-uiuc/Magicoder-S-DS-6.7B",
            trust_remote_code=True,
            torch_dtype=torch.bfloat16,
            device_map='auto',
        )

    from transformers.generation.stopping_criteria import StoppingCriteria, StoppingCriteriaList

    globals()["StoppingCriteria"] = StoppingCriteria
    globals()["StoppingCriteriaList"] = StoppingCriteriaList
    return tokenizer, model


def create_ark_client(args):
    try:
        from volcenginesdkarkruntime import Ark
    except ModuleNotFoundError as exc:
        raise RuntimeError(
            "Doubao-Seed-Code requires volcengine-python-sdk[ark]. "
            "Install it with: pip install --upgrade \"volcengine-python-sdk[ark]\""
        ) from exc

    api_key = os.environ.get(args.ark_api_key_env)
    if not api_key:
        raise RuntimeError(f"Missing Ark API key in environment variable: {args.ark_api_key_env}")
    if not args.ark_endpoint_id:
        raise RuntimeError("Missing Ark endpoint id. Use --ark-endpoint-id or set ARK_ENDPOINT_ID.")

    return Ark(base_url=ARK_BASE_URL, api_key=api_key)


def build_ark_messages(prompt):
    return [
        {
            "role": "system",
            "content": (
                "You generate Golang unit test snippets. "
                "Continue the current test snippet from the exact cursor position. "
                "Output only raw Go code that should be appended next. "
                "Do not explain. Do not include markdown fences. "
                "Do not repeat the existing snippet prefix."
            ),
        },
        {
            "role": "user",
            "content": prompt,
        },
    ]


def iter_dataset_sources(datasets_root, project, selected_dirs):
    sources = sorted(glob(os.path.join(datasets_root, project, "*")))
    if not selected_dirs:
        return sources
    allowed = set(selected_dirs)
    return [source for source in sources if os.path.basename(source) in allowed]


def generate_local_chunk(tokenizer, model, prompt, stopping_criteria, temperature):
    input_ids = tokenizer.encode(prompt, return_tensors="pt").to(model.device)
    output = model.generate(
        input_ids,
        max_new_tokens=512,
        temperature=temperature,
        eos_token_id=tokenizer.eos_token_id,
        pad_token_id=tokenizer.eos_token_id,
        stopping_criteria=stopping_criteria,
    )
    output = tokenizer.decode(output[0], skip_special_tokens=True)
    return output[len(prompt):]


def generate_ark_chunk(client, args, prompt):
    response = client.chat.completions.create(
        model=args.ark_endpoint_id,
        messages=build_ark_messages(prompt),
        temperature=args.temperature,
        extra_headers=EXTRA_HEADERS,
    )
    return response.choices[0].message.content


def should_stop_generation(test_class, output, stopping_criteria, model_name):
    if "\n}" in test_class:
        return True
    if model_name in LOCAL_MODELS:
        return stopping_criteria[0].output_token_length >= 512
    return output.strip() == ""


def main():
    parser = build_parser()
    args = parser.parse_args()

    set_seed(args.seed)

    tokenizer = None
    model = None
    ark_client = None
    if args.model in LOCAL_MODELS:
        tokenizer, model = load_local_model(args.model)
        from transformers.generation.stopping_criteria import StoppingCriteria, StoppingCriteriaList
        globals()["StoppingCriteria"] = StoppingCriteria
        globals()["StoppingCriteriaList"] = StoppingCriteriaList
    elif args.model == ARK_MODEL:
        ark_client = create_ark_client(args)
    else:
        raise ValueError(f"Unsupported model: {args.model}")

    last_workspace = None
    processed_items = 0
    projects = [args.project] if args.project else sorted(os.listdir(args.datasets_root))
    for project in projects:
        funcnames = dict()
        for source in iter_dataset_sources(args.datasets_root, project, args.dataset_dir):
            for file_name in sorted(os.listdir(source)):
                if args.max_items and processed_items >= args.max_items:
                    return
                with open(source + "/" + file_name, 'r') as file:
                    item = json.load(file)
                    relpath = item["relpath"]

                    stopping_criteria = None
                    if args.model in LOCAL_MODELS:
                        from transformers.generation.stopping_criteria import StoppingCriteriaList
                        stopping_criteria = StoppingCriteriaList([FunctionStoppingCriteria(tokenizer)])

                    # get package
                    with open(relpath, 'r') as focal_file:
                        for line in focal_file.readlines():
                            if line.startswith("package "):
                                package = line.replace("package ", "").strip()
                                break

                    function = item["funcbody"]
                    funcname = item["funcname"]
                    generated_identifiers = [funcname, project]
                    if stopping_criteria is not None:
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
                    
                    row = item["lineno"] - 1
                    context_num = 1

                    source_text = item["focalfunc"] + " {"
                    initial_identifiers = extract_prompt_identifiers(source_text, funcname)
                    for identifier, row_offset, column in initial_identifiers:
                        definition = send_completion_request(row + row_offset, column, relpath)
                        generated_identifiers.append(identifier)
                        if stopping_criteria is not None:
                            stopping_criteria[0].generated_identifiers.append(identifier)
                        if definition.strip() != "" and not definition.startswith("var") and not definition.startswith("package"):
                            prefix += "### CONTEXT_{}\n{}\n\n".format(context_num, definition.strip())
                            context_num += 1

                    with open(target_file_path, 'w') as file:
                        print(target_file_path)
                        test_class = "package {}\nfunc Test{}_{}(t *testing.T) {{\n".format(package, funcname[0].upper() + funcname[1:], funcnames[funcname.lower()])
                        file.write(test_class)
                        generation_round = 0
                        while True:
                            generation_round += 1
                            prompt = prefix + suffix
                            if args.model in LOCAL_MODELS:
                                output = generate_local_chunk(tokenizer, model, prompt, stopping_criteria, args.temperature)
                            else:
                                output = generate_ark_chunk(ark_client, args, prompt)
                            suffix += output
                            test_class += output
                            file.write(output) 
                            file.flush()
                            if should_stop_generation(test_class, output, stopping_criteria, args.model) or generation_round >= 8:
                                break
                            identifier = extract_new_identifier(output, generated_identifiers)
                            if not identifier:
                                break
                            generated_identifiers.append(identifier)
                            if stopping_criteria is not None and identifier not in stopping_criteria[0].generated_identifiers:
                                stopping_criteria[0].generated_identifiers.append(identifier)
                            row, column = get_row_and_column(test_class, identifier)
                            definition = send_completion_request(row, column, target_file_path)
                            if definition != "" and "invalid type" not in definition and not definition.startswith("var") and not definition.startswith("package"):
                                prefix += "### CONTEXT_{}\n{}\n\n".format(context_num, definition.strip())
                                context_num += 1
                    processed_items += 1


if __name__ == "__main__":
    main()
