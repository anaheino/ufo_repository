from sklearn.ensemble import RandomForestClassifier
from sklearn.model_selection import train_test_split
from sklearn.metrics import accuracy_score
import pickle
import pandas as pd
import json

with open('sighting_dataset.json', 'r') as file:
    ufo_data = pd.DataFrame(json.load(file))
X = ufo_data[['latitude', 'longitude']]
y = ufo_data['sighting']

X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.2, random_state=42)

# Initialize the Random Forest classifier
rf_model = RandomForestClassifier(n_estimators=100, random_state=42)

# Train the model
rf_model.fit(X_train, y_train)

# Make predictions on the test set
predictions = rf_model.predict(X_test)

# Evaluate the model
accuracy = accuracy_score(y_test, predictions)
print(f"Accuracy: {accuracy}")

with open('random_forest_model.pkl', 'wb') as file:
    pickle.dump(rf_model, file)
