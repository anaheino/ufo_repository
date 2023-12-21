import json

with open('../mongo/data.json', 'r') as file:
    ufo_data = json.load(file)

coordinates_data = [{'latitude': entry['latitude'], 'longitude': entry['longitude'], 'sighting': 1} for entry in ufo_data]
# Write the new JSON file with only coordinates
with open('ufo_coordinates.json', 'w') as file:
    json.dump(coordinates_data, file, indent=2)