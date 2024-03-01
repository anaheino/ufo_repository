from flask import Flask, jsonify, request

import pickle
import numpy as np

app = Flask(__name__)

with open('models/random_forest_model.pkl', 'rb') as file:
    random_forest = pickle.load(file)

@app.route('/api/data', methods=['GET'])
def get_data():
    # Example data
    data = {'message': 'Hello, this is a lightweight REST API without a database!'}
    return jsonify(data)

@app.route('/coordinates/likelihood', methods=['POST'])
def run_coords_through_random_forest():
    print("Running coords through random forest model")
    data_dict = request.get_json()
    if 'latitude' in data_dict and 'longitude' in data_dict:
        features_to_test = [data_dict['latitude'], data_dict['longitude']]
        features_np = np.array(features_to_test).reshape(1, -1)
        probability_estimates = random_forest.predict_proba(features_np)
        probability_sighting = probability_estimates[0, 1]
        response = jsonify({'probability': f'{probability_sighting}'})
        return response
    return "Wrong format, couldn't find latitude and longitude from body"

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
