import pandas as pd
import random
import json

with open('ufo_coordinates.json', 'r') as file:
    ufo_data = pd.DataFrame(json.load(file))

num_random_coordinates = 15000000

ufo_data[['latitude', 'longitude']] = ufo_data[['latitude', 'longitude']].apply(
    lambda x: x.apply(lambda y: float(int(y * 10000)) / 10000)
)
print(ufo_data.iloc[0])
random_coordinates = pd.DataFrame({
    'latitude': [round(random.uniform(-90.0000, 90.0000), 4) for _ in range(num_random_coordinates)],
    'longitude': [round(random.uniform(-180.0000, 180.0000), 4) for _ in range(num_random_coordinates)],
})

merged_df = pd.merge(ufo_data, random_coordinates, on=['latitude', 'longitude'], how='outer', indicator=True)
filtered_df = merged_df[merged_df['_merge'] == 'right_only'].drop(columns=['_merge'])
filtered_df['sighting'] = [0.0] * len(filtered_df)
print(filtered_df.iloc[0])
combined_data = pd.concat([ufo_data, filtered_df], ignore_index=True)
json_data = combined_data.to_json(orient='records')
print(len(json_data))
with open('sighting_dataset.json', 'w') as file:
    file.write(json_data)