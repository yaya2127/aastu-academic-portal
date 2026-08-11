import subprocess
import os

repo_dir = r'c:\yared-projects\aastu-academic-portal'
os.chdir(repo_dir)

def run_cmd(cmd):
    result = subprocess.run(cmd, shell=True, capture_output=True, text=True)
    print(">", cmd)
    if result.stdout:
        print(result.stdout.strip())

run_cmd('git config user.name "yaya2127"')
run_cmd('git config user.email "kinetibebyared@gmail.com"')

run_cmd("git add .gitattributes services/gpa_calculator.go")
run_cmd('git commit -m "feat(linguist): add .gitattributes to prioritize Go language statistics on GitHub" --allow-empty')

print("\nCommitted .gitattributes!")
