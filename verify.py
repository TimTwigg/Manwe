from urllib import request
import json

print("[TEST] /statblock")
res = request.urlopen("http://localhost:8080/statblock?name=Winter+Ghoul&detail_level=2")
assert res.status == 200, f"Error /statblock: {res.status}"
data = json.loads(res.read().decode("utf-8"))
assert data["Name"] == "Winter Ghoul", f"Error /statblock: {data['Name']}"

print("[TEST] /statblock/all")
res = request.urlopen("http://localhost:8080/statblock/all")
assert res.status == 200, f"Error /statblock/all: {res.status}"
data = json.loads(res.read().decode("utf-8"))
assert len(data) > 0, f"Error /statblock/all: {len(data)}"

print("[TEST] /encounter")
res = request.urlopen("http://localhost:8080/encounter?name=Encounter+1")
assert res.status == 200, f"Error /encounter: {res.status}"
data = json.loads(res.read().decode("utf-8"))
assert data["Name"] == "Encounter 1", f"Error /encounter: {data['Name']}"

print("[TEST] /encounter/all")
res = request.urlopen("http://localhost:8080/encounter/all?detail_level=1")
assert res.status == 200, f"Error /encounter/all: {res.status}"
data = json.loads(res.read().decode("utf-8"))
assert len(data) > 0, f"Error /encounter/all: {len(data)}"

print("\nAll tests passed.")