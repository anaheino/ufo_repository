import pickle
import json
import numpy as np

with open('random_forest_model.pkl', 'rb') as file:
    loaded_model = pickle.load(file)

json_data_to_test = '{"latitude": 51.52282, "longitude": -1.128462}'
data_dict = json.loads(json_data_to_test)
features_to_test = [data_dict['latitude'], data_dict['longitude']]

# Convert the features to a numpy array
features_np = np.array(features_to_test).reshape(1, -1)
# Get probability estimates using the loaded model
probability_estimates = loaded_model.predict_proba(features_np)
probability_sighting = probability_estimates[0, 1]
# Print the probability
print(f"Probability of Sighting: {probability_sighting}")