import pickle
import json
import numpy as np

with open('random_forest_model.pkl', 'rb') as file:
    loaded_model = pickle.load(file)

# JSON object with coordinates to test
json_data_to_test = '{"latitude": 42.4187022, "longitude": -111.3963006}'

# Convert the JSON string to a Python dictionary
data_dict = json.loads(json_data_to_test)

# Extract relevant features (latitude and longitude)
features_to_test = [data_dict['latitude'], data_dict['longitude']]

# Convert the features to a numpy array
features_np = np.array(features_to_test).reshape(1, -1)

# Make predictions using the loaded model
prediction = loaded_model.predict(features_np)

# Print the prediction
print(f"Prediction (Sighting): {prediction}")