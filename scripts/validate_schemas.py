# Validate the schemas in the catalog

from jsonschema import validate, ValidationError
from pathlib import Path
import json
from enum import Enum

ASSET_PATH: Path = Path("./../assets")

class Colors:
    HEADER = '\033[95m'
    OKBLUE = '\033[94m'
    OKCYAN = '\033[96m'
    OKGREEN = '\033[92m'
    WARNING = '\033[93m'
    FAIL = '\033[91m'
    ENDC = '\033[0m'
    BOLD = '\033[1m'
    UNDERLINE = '\033[4m'

class LogType(Enum):
    INFO = 0
    WARNING = 1
    ERROR = 2
    GOOD = 3

def read_schema(schemaName: str) -> dict:
    """Reads a schema from the schema directory and returns it as a dictionary."""
    schemaPath: Path = ASSET_PATH / "schemas" / schemaName
    with schemaPath.open() as schemaFile:
        return json.load(schemaFile)

def log(message: str, logtype: LogType):
    """Logs a message to the console with a color based on the log type."""
    color = ""
    intro = ""
    if logtype == LogType.INFO:
        color = Colors.OKBLUE
        intro = "[INFO]"
    elif logtype == LogType.WARNING:
        color = Colors.WARNING
        intro = "[WARNING]"
    elif logtype == LogType.ERROR:
        color = Colors.FAIL
        intro = "[ERROR]"
    elif logtype == LogType.GOOD:
        color = Colors.OKGREEN
        intro = "[SUCCEED]"
    print(color + intro + Colors.ENDC + " " + message)

def dump(message: str, filename: str = "dump.txt"):
    """Dumps a message to a file."""
    p = Path(filename)
    num: int = 1
    while p.exists():
        p = Path(f"{filename.split(".")[0]}_{num}.txt")
        num += 1
    with p.open("w") as f:
        f.write(message)

def validate_stat_blocks():
    """Validates all stat blocks in the stat_blocks directory."""
    schema = read_schema("stat_block.schema.json")
    statBlocksDir = ASSET_PATH / "stat_blocks"
    statBlocks = [json.loads(f.read_text()) for f in statBlocksDir.iterdir()]
    for block in statBlocks:
        try:
            validate(block, schema)
            log(f"Validated {block["Name"]}", LogType.GOOD)
        except ValidationError as e:
            log(f"Error validating {block["Name"]}: {e.message}", LogType.ERROR)
            dump(str(e))
    
if __name__ == "__main__":
    validate_stat_blocks()