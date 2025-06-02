import os
import subprocess

folders = [
    "test",
]

bash_commands = [
    "node -v",
    "npm -v",
    "npm install",
]

def run_commands_in_folder(folder: str):
    cwd = os.getcwd()
    target_folder = os.path.join(cwd, folder)

    if not os.path.exists(target_folder):
        raise FileNotFoundError(f"Folder {folder} does not exist")
    
    print(f"Running commands in {folder}")
    for command in bash_commands:
        print(f"Running {command}")
        subprocess.run(command, shell=True, check=True, cwd=target_folder)


if __name__ == "__main__":
    for folder in folders:
        run_commands_in_folder(folder)
