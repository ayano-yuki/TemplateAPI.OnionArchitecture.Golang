import os
import re
import sys

from pathlib import Path

def replace_in_file(path: Path, pattern: re.Pattern, replacement: str):
    text = path.read_text(encoding="utf-8")
    new_text = pattern.sub(replacement, text)
    if text != new_text:
        path.write_text(new_text, encoding="utf-8")
        print(f"Updated: {path}")

def main():
    old = input("Old module name (e.g., github.com/olduser/oldproject): ").strip()
    new = input("New module name (e.g., github.com/newuser/newproject): ").strip()

    print(f"\nRenaming module: {old} -> {new}\n")

    # 1. go.mod の書き換え
    mod_path = Path("go.mod")
    if not mod_path.exists():
        print("Error: go.mod not found")
        sys.exit(1)

    replace_in_file(mod_path, re.compile(re.escape(old)), new)

    # 2. .goファイルの import のみ対象にして書き換え
    pattern = re.compile(rf'(?<=["(]){re.escape(old)}(?=/)')  # import "API/x/y" → "TEST/x/y"

    print("Replacing import paths in .go files...\n")
    for file in Path(".").rglob("*.go"):
        replace_in_file(file, pattern, new)

    # 3. go mod tidy を実行
    print("\nRunning `go mod tidy` ...\n")
    os.system("go mod tidy")

    print("\nDone.")

if __name__ == "__main__":
    main()
