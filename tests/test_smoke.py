import pathlib
import subprocess
import sys
import unittest
import json
import importlib.util
import os
import tempfile

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
        self.assertIn('--datasets-root', result.stdout)
        self.assertIn('--project', result.stdout)
        self.assertIn('--max-items', result.stdout)
        self.assertIn('--dataset-dir', result.stdout)
        self.assertIn('--ark-endpoint-id', result.stdout)
        self.assertIn('--ark-api-key-env', result.stdout)
        self.assertIn('--gopls-endpoint', result.stdout)
        self.assertIn('--results-root', result.stdout)
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

    def test_iter_dataset_sources_uses_custom_root_and_filter(self):
        main = load_main_module()

        root = ROOT / 'tests' / 'tmp_datasets_root'
        source_a = root / 'demo' / 'pkg_a'
        source_b = root / 'demo' / 'pkg_b'
        source_a.mkdir(parents=True, exist_ok=True)
        source_b.mkdir(parents=True, exist_ok=True)

        all_sources = main.iter_dataset_sources(str(root), 'demo', [])
        filtered_sources = main.iter_dataset_sources(str(root), 'demo', ['pkg_b'])

        self.assertEqual([pathlib.Path(p).name for p in all_sources], ['pkg_a', 'pkg_b'])
        self.assertEqual([pathlib.Path(p).name for p in filtered_sources], ['pkg_b'])

    def test_resolve_project_path_returns_absolute_path(self):
        main = load_main_module()

        resolved = main.resolve_project_path("./projects/gin/codec/json/json.go")

        self.assertTrue(pathlib.Path(resolved).is_absolute())
        self.assertTrue(resolved.endswith("/projects/gin/codec/json/json.go"))

    def test_gopls_base_url_uses_env_and_trims_trailing_slash(self):
        main = load_main_module()
        previous = os.environ.get("GOPLS_ENDPOINT")
        try:
            os.environ["GOPLS_ENDPOINT"] = "http://127.0.0.1:24567/"
            self.assertEqual(main.gopls_base_url(), "http://127.0.0.1:24567")
        finally:
            if previous is None:
                os.environ.pop("GOPLS_ENDPOINT", None)
            else:
                os.environ["GOPLS_ENDPOINT"] = previous

    def test_build_metrics_paths_follow_results_layout(self):
        main = load_main_module()

        jsonl_path, summary_path = main.build_metrics_paths("./results", "Doubao-Seed-Code", "gin")

        self.assertTrue(jsonl_path.endswith("results/RATester_Doubao-Seed-Code/metrics/gin.jsonl"))
        self.assertTrue(summary_path.endswith("results/RATester_Doubao-Seed-Code/metrics/gin.summary.json"))

    def test_extract_usage_defaults_missing_usage_to_zero(self):
        main = load_main_module()

        usage = main.extract_usage(object())

        self.assertEqual(
            usage,
            {
                "prompt_tokens": 0,
                "completion_tokens": 0,
                "total_tokens": 0,
            },
        )

    def test_persist_metrics_record_appends_and_updates_summary(self):
        main = load_main_module()
        with tempfile.TemporaryDirectory() as root:
            record = {
                "project": "gin",
                "dataset_dir": "codec_json_MarshalIndent",
                "llm_calls": 2,
                "prompt_tokens": 11,
                "completion_tokens": 7,
                "total_tokens": 18,
                "generation_rounds": 2,
                "latency_ms": 123,
            }

            main.persist_metrics_record(str(root), "Doubao-Seed-Code", "gin", record)

            jsonl_path = pathlib.Path(root) / "RATester_Doubao-Seed-Code" / "metrics" / "gin.jsonl"
            summary_path = pathlib.Path(root) / "RATester_Doubao-Seed-Code" / "metrics" / "gin.summary.json"

            self.assertTrue(jsonl_path.exists())
            self.assertTrue(summary_path.exists())

            jsonl_lines = jsonl_path.read_text().strip().splitlines()
            self.assertEqual(len(jsonl_lines), 1)
            self.assertEqual(json.loads(jsonl_lines[0])["dataset_dir"], "codec_json_MarshalIndent")

            summary = json.loads(summary_path.read_text())
            self.assertEqual(summary["items"], 1)
            self.assertEqual(summary["llm_calls"], 2)
            self.assertEqual(summary["prompt_tokens"], 11)
            self.assertEqual(summary["completion_tokens"], 7)
            self.assertEqual(summary["total_tokens"], 18)

if __name__ == '__main__':
    unittest.main()
