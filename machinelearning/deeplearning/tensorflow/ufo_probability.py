import tensorflow as tf
from tensorflow.keras import layers, models

# Define the neural network model
def create_model(input_size, hidden_size, output_size):
    model = models.Sequential()
    model.add(layers.Input(shape=(input_size,)))
    model.add(layers.Dense(hidden_size, activation='relu'))
    model.add(layers.Dense(output_size, activation='softmax'))
    return model

def init_model(input_size, hidden_size, output_size, learning_rate):
    model = create_model(input_size, hidden_size, output_size)
    model.compile(optimizer=tf.keras.optimizers.Adam(learning_rate=learning_rate),
                  loss='sparse_categorical_crossentropy',  # Use 'sparse_categorical_crossentropy' for integer labels
                  metrics=['accuracy'])
    return model

def set_hyper_parameters():
   input_size = 2  # Assuming 2D coordinates
   hidden_size = 64
   output_size = 3  # Assuming we want 3 percentage outputs
   learning_rate = 0.001
   return init_model(input_size, hidden_size, output_size, learning_rate)

def create_network():
    coordinates = tf.constant([[0.5, 0.5], [1.0, 1.0], [0.1, 0.1]], dtype=tf.float32)
    labels = tf.constant([0, 1, 2], dtype=tf.int32)  # Dummy labels
    # Training the model
    model = set_hyper_parameters()
    model.fit(coordinates, labels, epochs=100, batch_size=1)

    # Example prediction
    sample_coordinates = tf.constant([[0.2, 0.2]], dtype=tf.float32)
    prediction = model.predict(sample_coordinates)
    print("Predicted percentages:", prediction)

if __name__ == "__main__":
    create_network()