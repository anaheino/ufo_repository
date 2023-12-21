import pandas as pd
import random
import json

with open('ufo_coordinates.json', 'r') as file:
    ufo_data = pd.DataFrame(json.load(file))

# Set the number of random coordinates you want to generate
num_random_coordinates = 15000000

# Generate random coordinates,
random_coordinates = pd.DataFrame({
    'latitude': [random.uniform(-90.0000000, 90.0000000) for _ in range(num_random_coordinates)],
    'longitude': [random.uniform(-180.0000000, 180.0000000) for _ in range(num_random_coordinates)],
})
random_coordinates['sighting'] = 0.00

# Combine the random coordinates with your existing UFO sightings data
combined_data = pd.concat([ufo_data, random_coordinates], ignore_index=True)
json_data = combined_data.to_json(orient='records')

with open('sighting_dataset.json', 'w') as file:
    file.write(json_data)