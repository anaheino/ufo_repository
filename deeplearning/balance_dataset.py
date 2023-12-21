import pandas as pd
import random
import json

# Assuming your existing dataset is in a DataFrame called 'ufo_data'
# Your DataFrame should have columns like 'latitude', 'longitude', and 'likelihood'
with open('ufo_coordinates.json', 'r') as file:
    ufo_data = pd.DataFrame(json.load(file))

# Extract the range of latitude and longitude from your existing dataset
min_lat = ufo_data['latitude'].min()
max_lat = ufo_data['latitude'].max()
min_lon = ufo_data['longitude'].min()
max_lon = ufo_data['longitude'].max()

# Set the number of random coordinates you want to generate
num_random_coordinates = 15000000  # Adjust this based on your needs

# Generate random coordinates
random_coordinates = pd.DataFrame({
    'latitude': [random.uniform(min_lat, max_lat) for _ in range(num_random_coordinates)],
    'longitude': [random.uniform(min_lon, max_lon) for _ in range(num_random_coordinates)],
})
random_coordinates['sighting'] = 0

# Combine the random coordinates with your existing UFO sightings data
combined_data = pd.concat([ufo_data, random_coordinates], ignore_index=True)
json_data = combined_data.to_json(orient='records')

with open('sighting_dataset.json', 'w') as file:
    file.write(json_data)