import pathlib
import subprocess
import sys
import unittest
import json
import importlib.util

ROOT = pathlib.Path(__file__).resolve().parents[1]


def load_main_module():
    spec = importlib.util.spec_from_file_location("ratester_main", ROOT / "main.py")
    module = importlib.util.module_from_spec(spec)
    spec.loader.exec_module(module)
    return module


def load_run_test_module():
    spec = importlib.util.spec_from_file_location("ratester_run_test", ROOT / "run_test.py")
    module = importlib.util.module_from_spec(spec)
    spec.loader.exec_module(module)
    return module

class SmokeTests(unittest.TestCase):
    def test_run_test_script_compiles(self):
        result = subprocess.run(
            [sys.executable, '-m', 'py_compile', str(ROOT / 'run_test.py')],
            capture_output=True,
            text=True,
        )
        self.assertEqual(result.returncode, 0, msg=result.stderr)

    def test_main_help_exposes_temperature(self):
        result = subprocess.run(
            [sys.executable, str(ROOT / 'main.py'), '--help'],
            capture_output=True,
            text=True,
        )
        self.assertEqual(result.returncode, 0, msg=result.stderr)
        self.assertIn('--temperature', result.stdout)
        self.assertIn('--project', result.stdout)
        self.assertIn('--max-items', result.stdout)
        self.assertIn('--ark-endpoint-id', result.stdout)
        self.assertIn('--ark-api-key-env', result.stdout)
        self.assertIn('Doubao-Seed-Code', result.stdout)

    def test_run_test_uses_ratester_prefix(self):
        text = (ROOT / 'run_test.py').read_text()
        self.assertIn('RATester_', text)
        self.assertNotIn('IntelliTester_', text)

    def test_provider_settings_uses_absolute_gopls_path(self):
        settings = json.loads((ROOT / 'jsonrpc' / 'provider_settings.json').read_text())
        self.assertEqual(len(settings), 1)
        self.assertTrue(settings[0]['binaryLocation'].startswith('/'))

    def test_extract_prompt_identifiers_skips_keywords_and_function_name(self):
        main = load_main_module()

        identifiers = main.extract_prompt_identifiers(
            "func Request(ctx context.Context, req *Request) {",
            "Request",
        )
        names = [item[0] for item in identifiers]

        self.assertIn("ctx", names)
        self.assertIn("context", names)
        self.assertIn("Context", names)
        self.assertIn("req", names)
        self.assertNotIn("func", names)
        self.assertNotIn("Request", names)

    def test_resolve_goimports_prefers_existing_binary(self):
        run_test = load_run_test_module()

        resolved = run_test.resolve_goimports()

        self.assertTrue(resolved.endswith("goimports"))
        self.assertTrue(pathlib.Path(resolved).exists())

    def test_generated_test_file_maps_back_to_package_dir(self):
        run_test = load_run_test_module()

        package_dir = run_test.package_dir_for_generated_test(
            "./projects/beego/client/cache/RATester_Doubao-Seed-Code_Get_1_test.go",
            "Doubao-Seed-Code",
        )

        self.assertEqual(package_dir, "./projects/beego/client/cache")

    def test_build_ark_messages_enforces_code_only_output(self):
        main = load_main_module()

        messages = main.build_ark_messages("PROMPT")

        self.assertEqual(messages[0]["role"], "system")
        self.assertIn("Go", messages[0]["content"])
        self.assertIn("Do not include markdown fences", messages[0]["content"])
        self.assertIn("Do not explain", messages[0]["content"])
        self.assertEqual(messages[1], {"role": "user", "content": "PROMPT"})

if __name__ == '__main__':
    unittest.main()
