import json

with open('../mongo/data.json', 'r') as file:
    ufo_data = json.load(file)

coordinates_data = [{'latitude': entry['latitude'], 'longitude': entry['longitude'], 'sighting': 1.00} for entry in ufo_data]
with open('older_random_forest/ufo_coordinates.json', 'w') as file:
    json.dump(coordinates_data, file, indent=2)