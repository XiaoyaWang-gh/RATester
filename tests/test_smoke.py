import pathlib
import subprocess
import sys
import unittest

ROOT = pathlib.Path(__file__).resolve().parents[1]

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

    def test_run_test_uses_ratester_prefix(self):
        text = (ROOT / 'run_test.py').read_text()
        self.assertIn('RATester_', text)
        self.assertNotIn('IntelliTester_', text)

if __name__ == '__main__':
    unittest.main()
