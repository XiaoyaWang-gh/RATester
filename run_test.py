import os
import re
import shutil
import argparse
import subprocess
from tqdm import tqdm
from glob import glob


def find_file(filename, start_dir):
    for root, dirs, files in os.walk(start_dir):
        if filename in files:
            return os.path.join(root, filename)
    return None


def run_import(model, project):
    files = glob(f"./projects/{project}/**/IntelliTester_{model}_*_test.go", recursive=True)
    files = sorted(files)
    packages = set()

    for file_path in tqdm(files):
        packages.add(file_path.split(f"{model}_")[0])
        subprocess.run(["goimports", "-w", file_path], capture_output=True, text=True)
    return packages


def run_test(model, project, packages):
    pattern = rf"IntelliTester_{model}.*_test\.go"
    for package in packages:
        while True:
            result = subprocess.run(["go", "test", "-timeout", "30s"], cwd=package, capture_output=True, text=True, errors='ignore')
            matches = re.findall(pattern, result.stderr)

            if "go get" in result.stderr:
                go_get_pattern = re.compile(r'go get (\S+)')
                match = go_get_pattern.search(result.stderr)
                if match:
                    go_get_command = match.group(0)
                result = subprocess.run(
                    go_get_command.split(),
                    cwd=f"./projects/{project}",
                    capture_output=True,
                    text=True,
                    errors='ignore'
                )
                continue

            if len(matches) == 0:
                break

            for file_name in matches:
                source_path = os.path.join(package, file_name)
                if not os.path.exists(source_path):
                    source_path = find_file(file_name, f"./projects/{project}")
                if source_path == None: continue
                model_dir = source_path.split(f"/IntelliTester_{model}_")[0].replace(f"./projects", f"./error/IntelliTester_{model}")
                model_path = os.path.join(model_dir, file_name)
                if not os.path.exists(model_path):
                    os.makedirs(model_dir, exist_ok=True)
                    shutil.move(source_path, model_path)
                    print(f"Found compilation error in {source_path}, move to {model_path}")


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Process some parameters.")
    parser.add_argument('--model', help='model to use for unit test generation. should be one of [CodeLlama, DeepSeek-Coder, Magicoder]', required=True, type=str)
    parser.add_argument('--project', type=str, help='project to be tested. should be one of [beego, echo, fiber, frp, gin, hugo, nps, traefik]', required=True, type=str)
    args = parser.parse_args()
    model = args.model
    project = args.project
    packages = run_import(model, project)
    run_test(model, project, packages)